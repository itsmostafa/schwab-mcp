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
}

func (f *fakeSchwab) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	f.reqs = append(f.reqs, r.Method+" "+r.URL.RequestURI())
	loc := f.location
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
	w.Write([]byte(`{}`))
}

// session connects an in-memory MCP client to a server backed by fake.
func session(t *testing.T, fake http.Handler, allowTrading bool) *mcp.ClientSession {
	srv := httptest.NewServer(fake)
	t.Cleanup(srv.Close)
	a := NewAuth(Config{TokenFile: filepath.Join(t.TempDir(), "token.json")})
	a.tok = &token{AccessToken: "tok", ExpiresAt: time.Now().Add(time.Hour)}
	s := mcp.NewServer(&mcp.Implementation{Name: "schwab", Version: "test"}, nil)
	registerTools(s, &Client{BaseURL: srv.URL, Auth: a, HTTP: srv.Client()}, allowTrading)

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

func call(t *testing.T, cs *mcp.ClientSession, name string, args map[string]any) string {
	res, err := cs.CallTool(t.Context(), &mcp.CallToolParams{Name: name, Arguments: args})
	if err != nil {
		t.Fatal(err)
	}
	text := res.Content[0].(*mcp.TextContent).Text
	if res.IsError {
		t.Fatalf("%s: tool error: %s", name, text)
	}
	return text
}

func TestTradingToolsGated(t *testing.T) {
	names := toolNames(t, session(t, &fakeSchwab{}, false))
	for _, n := range []string{"place_order", "replace_order", "cancel_order"} {
		if slices.Contains(names, n) {
			t.Errorf("%s registered without SCHWAB_ALLOW_TRADING", n)
		}
	}
	if !slices.Contains(names, "preview_order") || !slices.Contains(names, "get_quotes") {
		t.Errorf("read-only tools missing: %v", names)
	}
	if names := toolNames(t, session(t, &fakeSchwab{}, true)); !slices.Contains(names, "place_order") {
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
	fake := &fakeSchwab{location: "https://api.schwabapi.com/trader/v1/accounts/H1/orders/12345"}
	cs := session(t, fake, true)
	order := map[string]any{"orderType": "MARKET"}

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
		"POST /trader/v1/accounts/H1/orders",
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
