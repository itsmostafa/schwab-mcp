package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestRefresh(t *testing.T) {
	var gotAuth, gotForm, gotType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		gotAuth = u + ":" + p
		gotType = r.Header.Get("Content-Type")
		b, _ := io.ReadAll(r.Body)
		gotForm = string(b)
		w.Write([]byte(`{"access_token":"new","refresh_token":"r2","expires_in":1800}`))
	}))
	defer srv.Close()

	cfg := Config{AppKey: "key", AppSecret: "sec", TokenFile: filepath.Join(t.TempDir(), "d", "token.json")}
	a := NewAuth(cfg)
	a.TokenURL = srv.URL
	if _, err := a.Token(t.Context()); !errors.Is(err, errNotLoggedIn) {
		t.Fatalf("no token file: got %v", err)
	}

	rexp := time.Now().Add(time.Hour).Round(time.Second)
	if err := a.save(&token{AccessToken: "old", RefreshToken: "r1", ExpiresAt: time.Now(), RefreshExpiresAt: rexp}); err != nil {
		t.Fatal(err)
	}
	tok, err := a.Token(t.Context())
	if err != nil || tok != "new" {
		t.Fatalf("Token = %q, %v", tok, err)
	}
	if gotAuth != "key:sec" || gotType != "application/x-www-form-urlencoded" ||
		gotForm != "grant_type=refresh_token&refresh_token=r1" {
		t.Fatalf("refresh request: auth %q type %q form %q", gotAuth, gotType, gotForm)
	}
	fi, err := os.Stat(cfg.TokenFile)
	if err != nil || fi.Mode().Perm() != 0o600 {
		t.Fatalf("token file mode: %v %v", fi, err)
	}
	var ft token
	b, _ := os.ReadFile(cfg.TokenFile)
	if err := json.Unmarshal(b, &ft); err != nil {
		t.Fatal(err)
	}
	if ft.AccessToken != "new" || ft.RefreshToken != "r2" || !ft.RefreshExpiresAt.Equal(rexp) {
		t.Fatalf("saved token: %s", b)
	}
}

func TestLoginRejectsWrongState(t *testing.T) {
	openBrowser = func(string) {}
	a := NewAuth(Config{AppKey: "key", AppSecret: "sec", CallbackURL: "https://127.0.0.1:18183/cb",
		TokenFile: filepath.Join(t.TempDir(), "token.json")})
	ctx, cancel := context.WithCancel(t.Context())
	authURL, done, err := a.Login(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(authURL, "state=") || !strings.Contains(authURL, "client_id=key") {
		t.Fatalf("authURL %s", authURL)
	}
	hc := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	resp, err := hc.Get("https://127.0.0.1:18183/cb?code=abc&state=wrong")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("wrong state: HTTP %d", resp.StatusCode)
	}
	cancel()
	if err := <-done; err == nil {
		t.Fatal("canceled login reported success")
	}
	if _, err := a.load(); err == nil {
		t.Fatal("token saved despite wrong state")
	}
}

func TestToQuery(t *testing.T) {
	in := priceHistoryIn{Symbol: "AAPL", StartDate: new(int64(1700000000123)), NeedExtendedHoursData: new(false)}
	q := toQuery(in, "symbol")
	if q.Get("startDate") != "1700000000123" || q.Get("needExtendedHoursData") != "false" || q.Has("symbol") || q.Has("period") {
		t.Fatalf("toQuery = %v", q)
	}
	if q := toQuery(quotesIn{Symbols: []string{"AAPL", "MSFT"}}); q.Get("symbols") != "AAPL,MSFT" {
		t.Fatalf("slice join = %v", q)
	}
}

// fakeSchwab records request paths and serves place-order responses.
type fakeSchwab struct {
	mu       sync.Mutex
	reqs     []string
	location string
	quote    string // body served for GET /marketdata/v1/quotes
}

func (f *fakeSchwab) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	f.reqs = append(f.reqs, r.Method+" "+r.URL.RequestURI())
	loc, quote := f.location, f.quote
	f.mu.Unlock()
	if r.Header.Get("Authorization") != "Bearer tok" {
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/orders") {
		if loc != "" {
			w.Header().Set("Location", loc)
			// Promise a body that never arrives: a 201 must still yield the order ID.
			w.Header().Set("Content-Length", "5")
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
	if r.URL.Path == "/marketdata/v1/quotes" && quote != "" {
		w.Write([]byte(quote))
		return
	}
	w.Write([]byte(`{}`))
}

// session connects an in-memory MCP client to a server backed by fake.
func session(t *testing.T, fake http.Handler, cfg Config) *mcp.ClientSession {
	srv := httptest.NewServer(fake)
	t.Cleanup(srv.Close)
	a := NewAuth(Config{TokenFile: filepath.Join(t.TempDir(), "token.json")})
	a.tok = &token{AccessToken: "tok", ExpiresAt: time.Now().Add(time.Hour)}
	s := mcp.NewServer(&mcp.Implementation{Name: "schwab", Version: "test"}, nil)
	registerTools(s, &Client{BaseURL: srv.URL, Auth: a, HTTP: srv.Client()}, cfg)

	st, ct := mcp.NewInMemoryTransports()
	if _, err := s.Connect(t.Context(), st, nil); err != nil {
		t.Fatal(err)
	}
	cs, err := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "test"}, nil).Connect(t.Context(), ct, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { cs.Close() })
	return cs
}

func toolNames(t *testing.T, cs *mcp.ClientSession) []string {
	res, err := cs.ListTools(t.Context(), nil)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, tool := range res.Tools {
		names = append(names, tool.Name)
	}
	return names
}

// callResult calls a tool and returns its text and whether it is a tool error.
func callResult(t *testing.T, cs *mcp.ClientSession, name string, args map[string]any) (string, bool) {
	res, err := cs.CallTool(t.Context(), &mcp.CallToolParams{Name: name, Arguments: args})
	if err != nil {
		t.Fatal(err)
	}
	return res.Content[0].(*mcp.TextContent).Text, res.IsError
}

func call(t *testing.T, cs *mcp.ClientSession, name string, args map[string]any) string {
	text, isErr := callResult(t, cs, name, args)
	if isErr {
		t.Fatalf("%s: tool error: %s", name, text)
	}
	return text
}

func TestTradingToolsGated(t *testing.T) {
	names := toolNames(t, session(t, &fakeSchwab{}, Config{}))
	for _, n := range []string{"place_order", "replace_order", "cancel_order"} {
		if slices.Contains(names, n) {
			t.Errorf("%s registered without SCHWAB_ALLOW_TRADING", n)
		}
	}
	if !slices.Contains(names, "preview_order") || !slices.Contains(names, "get_quotes") {
		t.Errorf("read-only tools missing: %v", names)
	}
	if names := toolNames(t, session(t, &fakeSchwab{}, Config{AllowTrading: true})); !slices.Contains(names, "place_order") {
		t.Errorf("place_order missing with trading enabled: %v", names)
	}
}

func TestOutcomeUnknownOnlyForMutations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()
	a := NewAuth(Config{TokenFile: filepath.Join(t.TempDir(), "token.json")})
	a.tok = &token{AccessToken: "tok", ExpiresAt: time.Now().Add(time.Hour)}
	c := &Client{BaseURL: srv.URL, Auth: a, HTTP: srv.Client()}

	for path, want := range map[string]bool{"/trader/v1/accounts/H1/previewOrder": false, "/trader/v1/accounts/H1/orders": true} {
		_, _, err := c.do(t.Context(), http.MethodPost, path, nil, map[string]any{})
		if err == nil || strings.Contains(err.Error(), "outcome unknown") != want {
			t.Errorf("POST %s: err %v, want outcome unknown = %v", path, err, want)
		}
	}
}

func TestToolRequests(t *testing.T) {
	fake := &fakeSchwab{
		location: "https://api.schwabapi.com/trader/v1/accounts/H1/orders/12345",
		quote:    `{"AAPL":{"realtime":true,"quote":{"bidPrice":99.9,"askPrice":100,"lastPrice":100}}}`,
	}
	cs := session(t, fake, Config{AllowTrading: true, MaxPriceDeviationBps: 50})
	order := limitOrder("BUY", "AAPL", 100)

	if got := call(t, cs, "place_order", map[string]any{"accountHash": "H1", "order": order}); got != "order accepted, orderId: 12345" {
		t.Errorf("place_order = %q", got)
	}
	fake.mu.Lock()
	fake.location = ""
	fake.mu.Unlock()
	if got := call(t, cs, "place_order", map[string]any{"accountHash": "H1", "order": order}); !strings.Contains(got, "ID unknown") {
		t.Errorf("place_order without Location = %q", got)
	}
	call(t, cs, "get_movers", map[string]any{"symbol": "$SPX", "sort": "VOLUME"})
	call(t, cs, "get_order", map[string]any{"accountHash": "H1", "orderId": "7"})
	call(t, cs, "get_quotes", map[string]any{"symbols": []string{"AAPL", "MSFT"}, "indicative": false})

	want := []string{
		"GET /marketdata/v1/quotes?symbols=AAPL",
		"POST /trader/v1/accounts/H1/orders",
		"GET /marketdata/v1/quotes?symbols=AAPL",
		"POST /trader/v1/accounts/H1/orders",
		"GET /marketdata/v1/movers/$SPX?sort=VOLUME",
		"GET /trader/v1/accounts/H1/orders/7",
		"GET /marketdata/v1/quotes?indicative=false&symbols=AAPL%2CMSFT",
	}
	fake.mu.Lock()
	defer fake.mu.Unlock()
	if !slices.Equal(fake.reqs, want) {
		t.Errorf("requests:\n got %q\nwant %q", fake.reqs, want)
	}
}

func limitOrder(instruction, symbol string, price any) map[string]any {
	return map[string]any{"orderType": "LIMIT", "price": price, "orderLegCollection": []any{
		map[string]any{"instruction": instruction, "quantity": 1, "instrument": map[string]any{"symbol": symbol, "assetType": "EQUITY"}},
	}}
}

// An order reaches Schwab only if it passes the live quote check.
func TestOrderPriceCheck(t *testing.T) {
	fake := &fakeSchwab{
		location: "https://api.schwabapi.com/trader/v1/accounts/H1/orders/1",
		quote: `{"AAPL":{"realtime":true,"quote":{"bidPrice":99.9,"askPrice":100,"lastPrice":100}},
			"SLOW":{"realtime":false,"quote":{"bidPrice":10,"askPrice":10.1}},
			"DEAD":{"realtime":true,"quote":{"bidPrice":0,"askPrice":0}},
			"GAP":{"realtime":true,"quote":{"bidPrice":100.9,"askPrice":101,"lastPrice":99,"mark":100.95}}}`,
	}
	strict := session(t, fake, Config{AllowTrading: true, MaxPriceDeviationBps: 50})
	loose := session(t, fake, Config{AllowTrading: true, MaxPriceDeviationBps: 50, AllowMarketOrders: true})
	market := func(orderType string) map[string]any {
		o := limitOrder("BUY", "AAPL", nil)
		o["orderType"] = orderType
		delete(o, "price")
		return o
	}
	stopLimit := func(instruction, symbol, stopType string, stop, limit float64) map[string]any {
		o := limitOrder(instruction, symbol, limit)
		o["orderType"], o["stopPrice"], o["stopType"] = "STOP_LIMIT", stop, stopType
		return o
	}
	// bracket is a BUY LIMIT AAPL 100 that triggers an OCO of children.
	bracket := func(children ...map[string]any) map[string]any {
		oco := map[string]any{"orderStrategyType": "OCO", "childOrderStrategies": children}
		o := limitOrder("BUY", "AAPL", 100)
		o["orderStrategyType"], o["childOrderStrategies"] = "TRIGGER", []any{oco}
		return o
	}

	for _, tc := range []struct {
		name    string
		cs      *mcp.ClientSession
		order   map[string]any
		wantErr string // empty means the order must be sent
	}{
		{"buy within band", strict, limitOrder("BUY", "AAPL", 100.4), ""},
		{"resting buy below ask", strict, limitOrder("BUY", "AAPL", 90), ""},
		{"resting sell above bid", strict, limitOrder("SELL", "AAPL", 120), ""},
		{"buy crosses ask", strict, limitOrder("BUY", "AAPL", 101), "crosses the live market"},
		{"sell crosses bid", strict, limitOrder("SELL", "AAPL", "98.00"), "crosses the live market"},
		{"unreadable price", strict, limitOrder("BUY", "AAPL", "abc"), "cannot read order"},
		{"NaN price", strict, limitOrder("BUY", "AAPL", "NaN"), "cannot read order"},
		{"missing price", strict, limitOrder("BUY", "AAPL", nil), "cannot read price"},
		{"triggered stop-limit crosses ask", strict, stopLimit("BUY", "AAPL", "", 99, 150), "crosses the live market"},
		{"untriggered stop-limit rests", strict, stopLimit("BUY", "AAPL", "", 105, 150), ""},
		{"triggered sell stop-limit within band", strict, stopLimit("SELL", "AAPL", "STANDARD", 101, 99.5), ""},
		// GAP: last 99 is below bid/ask, so the trigger basis decides whether the stop is reached.
		{"LAST stop not reached", strict, stopLimit("BUY", "GAP", "LAST", 100, 110), ""},
		{"ASK stop reached", strict, stopLimit("BUY", "GAP", "ASK", 100, 110), "crosses the live market"},
		{"STANDARD reached on ask", strict, stopLimit("BUY", "GAP", "STANDARD", 100, 110), "crosses the live market"},
		{"MARK stop not reached", strict, stopLimit("SELL", "GAP", "MARK", 100, 90), ""},
		{"unknown stopType", strict, stopLimit("BUY", "GAP", "BOGUS", 100, 110), "unsupported stopType"},
		{"market blocked", strict, market("MARKET"), "can fill at any price"},
		{"stop blocked", strict, market("STOP"), "can fill at any price"},
		{"market allowed", loose, market("MARKET"), ""},
		{"OCO hides market child", strict, map[string]any{"orderStrategyType": "OCO",
			"childOrderStrategies": []any{limitOrder("SELL", "AAPL", 120), market("MARKET")}}, "can fill at any price"},
		{"bracket hides stop child", strict, bracket(limitOrder("SELL", "AAPL", 120), market("STOP")), "can fill at any price"},
		{"bracket stop child allowed", loose, bracket(limitOrder("SELL", "AAPL", 120), market("STOP")), ""},
		{"bracket stop-limit child", strict, bracket(limitOrder("SELL", "AAPL", 120), stopLimit("SELL", "AAPL", "", 95, 94)), ""},
		{"bracket child delayed quote", strict, bracket(limitOrder("SELL", "SLOW", 11)), "no real-time quote"},
		{"bracket child with no bid rests", strict, bracket(limitOrder("SELL", "DEAD", 1)), ""},
		{"delayed quote", strict, limitOrder("BUY", "SLOW", 10), "no real-time quote"},
		{"no ask", strict, limitOrder("BUY", "DEAD", 1), "no live ask"},
		{"unknown symbol", strict, limitOrder("BUY", "NOPE", 1), "no real-time quote"},
	} {
		fake.mu.Lock()
		n := len(fake.reqs)
		fake.mu.Unlock()
		text, isErr := callResult(t, tc.cs, "place_order", map[string]any{"accountHash": "H1", "order": tc.order})
		fake.mu.Lock()
		sent := slices.Contains(fake.reqs[n:], "POST /trader/v1/accounts/H1/orders")
		fake.mu.Unlock()
		if tc.wantErr == "" && (isErr || !sent) || tc.wantErr != "" && (!isErr || sent || !strings.Contains(text, tc.wantErr)) {
			t.Errorf("%s: sent=%v isErr=%v text=%q, want error %q", tc.name, sent, isErr, text, tc.wantErr)
		}
	}

	text, isErr := callResult(t, strict, "replace_order", map[string]any{"accountHash": "H1", "orderId": "7", "order": market("MARKET")})
	fake.mu.Lock()
	putSent := slices.Contains(fake.reqs, "PUT /trader/v1/accounts/H1/orders/7")
	fake.mu.Unlock()
	if !isErr || putSent || !strings.Contains(text, "can fill at any price") {
		t.Errorf("replace_order MARKET: sent=%v isErr=%v text=%q", putSent, isErr, text)
	}

	if got := call(t, strict, "get_quotes", map[string]any{"symbols": []string{"AAPL", "SLOW"}}); !strings.HasPrefix(got, "WARNING: delayed, not real-time quotes for SLOW;") {
		t.Errorf("get_quotes delayed warning = %q", got)
	}
}

// A tar member named "schwab" that is a symlink (or any other non-regular entry)
// must not be extracted and installed over the running binary.
func TestExtractBinaryRejectsNonRegularMember(t *testing.T) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)
	if err := tw.WriteHeader(&tar.Header{
		Name:     "schwab",
		Typeflag: tar.TypeSymlink,
		Linkname: "/etc/passwd",
		Mode:     0o777,
	}); err != nil {
		t.Fatal(err)
	}
	for _, closer := range []func() error{tw.Close, gw.Close} {
		if err := closer(); err != nil {
			t.Fatal(err)
		}
	}

	if err := extractBinaryFromTar(buf.Bytes(), "schwab", io.Discard); err == nil || !strings.Contains(err.Error(), "not a regular file") {
		t.Fatalf("expected non-regular member to be rejected, got %v", err)
	}
}

func TestSetupCommands(t *testing.T) {
	cmds := setupCommands("/bin/schwab", []string{"SCHWAB_APP_KEY=k", "SCHWAB_APP_SECRET=s"})
	want := [][]string{
		{"mcp", "remove", "schwab", "-s", "user"},
		{"mcp", "add", "schwab", "-s", "user", "-e", "SCHWAB_APP_KEY=k", "-e", "SCHWAB_APP_SECRET=s", "--", "/bin/schwab", "mcp"},
		nil,
		{"mcp", "add", "schwab", "--env", "SCHWAB_APP_KEY=k", "--env", "SCHWAB_APP_SECRET=s", "--", "/bin/schwab", "mcp"},
	}
	got := [][]string{cmds[0].reset, cmds[0].add, cmds[1].reset, cmds[1].add}
	if !slices.EqualFunc(got, want, slices.Equal) {
		t.Fatalf("got %q\nwant %q", got, want)
	}
}

func TestSetupClaudeDesktop(t *testing.T) {
	path := filepath.Join(t.TempDir(), "claude_desktop_config.json")
	seed := `{"mcpServers":{"lumi":{"command":"/bin/lumi"},"schwab":{"command":"/old"}},"preferences":{"sidebarMode":"chat"}}`
	if err := os.WriteFile(path, []byte(seed), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := setupClaudeDesktop(path, "/bin/schwab", []string{"SCHWAB_ALLOW_TRADING=true"}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(path)
	var got struct {
		MCPServers map[string]struct {
			Command string            `json:"command"`
			Args    []string          `json:"args"`
			Env     map[string]string `json:"env"`
		} `json:"mcpServers"`
		Preferences map[string]string `json:"preferences"`
	}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	s := got.MCPServers["schwab"]
	if s.Command != "/bin/schwab" || !slices.Equal(s.Args, []string{"mcp"}) || s.Env["SCHWAB_ALLOW_TRADING"] != "true" {
		t.Fatalf("schwab entry = %+v", s)
	}
	if got.MCPServers["lumi"].Command != "/bin/lumi" || got.Preferences["sidebarMode"] != "chat" {
		t.Fatalf("other keys lost: %s", b)
	}

	for _, seed := range []string{`null`, `{"mcpServers":null}`} {
		if err := os.WriteFile(path, []byte(seed), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := setupClaudeDesktop(path, "/bin/schwab", nil); err != nil {
			t.Fatalf("seed %s: %v", seed, err)
		}
	}
}

func TestLoadCredentialsFromFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("SCHWAB_APP_KEY", "env-key")
	t.Setenv("SCHWAB_APP_SECRET", "")
	// Stdin may be a terminal when the test binary runs directly; a pipe must not prompt.
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdin := os.Stdin
	os.Stdin = r
	t.Cleanup(func() { os.Stdin = stdin; r.Close(); w.Close() })

	// No file and stdin is not a terminal: error, no prompt.
	if _, err := loadConfig(); err == nil || !strings.Contains(err.Error(), ".config/schwab/config") {
		t.Fatalf("missing credentials: err = %v", err)
	}

	path := filepath.Join(home, ".config", "schwab", "config")
	os.MkdirAll(filepath.Dir(path), 0o700)
	os.WriteFile(path, []byte("SCHWAB_APP_KEY=file-key\nSCHWAB_APP_SECRET = file-secret \n"), 0o600)
	cfg, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	// The environment wins over the file.
	if cfg.AppKey != "env-key" || cfg.AppSecret != "file-secret" {
		t.Fatalf("got key %q secret %q", cfg.AppKey, cfg.AppSecret)
	}
}
