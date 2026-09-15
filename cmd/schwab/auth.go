package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var errNotLoggedIn = errors.New(`not logged in: call the login tool or run "schwab login"`)

const (
	refreshLifetime = 7 * 24 * time.Hour
	expirySkew      = 60 * time.Second
	loginTimeout    = 5 * time.Minute
)

const authorizeURL = "https://api.schwabapi.com/v1/oauth/authorize"

// tokenHTTP is used for token endpoint calls.
var tokenHTTP = &http.Client{Timeout: 30 * time.Second}

// token is the on-disk token file format.
type token struct {
	AccessToken      string    `json:"access_token"`
	RefreshToken     string    `json:"refresh_token"`
	ExpiresAt        time.Time `json:"expires_at"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}

func (t *token) fresh() bool {
	return t != nil && t.AccessToken != "" && time.Until(t.ExpiresAt) > expirySkew
}

// Auth manages Schwab OAuth tokens.
type Auth struct {
	Cfg      Config
	TokenURL string

	mu  sync.Mutex // guards tok and token file access in this process
	tok *token

	loginMu  sync.Mutex // guards background login state
	loginURL string     // set while a background login is running
	loginErr error      // last background login failure, returned once
}

func NewAuth(cfg Config) *Auth {
	return &Auth{Cfg: cfg, TokenURL: "https://api.schwabapi.com/v1/oauth/token"}
}

// Token returns a valid access token, reloading from disk or refreshing as needed.
func (a *Auth) Token(ctx context.Context) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.tok.fresh() {
		return a.tok.AccessToken, nil
	}

	unlock, err := a.lockFile()
	if err != nil {
		return "", err
	}
	defer unlock()

	// Another process may have refreshed or logged in since we last read.
	t, err := a.load()
	if errors.Is(err, fs.ErrNotExist) {
		return "", errNotLoggedIn
	}
	if err != nil {
		return "", err
	}
	a.tok = t
	if t.fresh() {
		return t.AccessToken, nil
	}
	if t.RefreshToken == "" || !time.Now().Before(t.RefreshExpiresAt) {
		return "", errNotLoggedIn
	}

	form := url.Values{"grant_type": {"refresh_token"}, "refresh_token": {t.RefreshToken}}
	nt, status, err := a.exchange(ctx, form)
	if status == http.StatusBadRequest || status == http.StatusUnauthorized {
		return "", fmt.Errorf("%w (refresh rejected: %v)", errNotLoggedIn, err)
	}
	if err != nil {
		return "", err
	}
	if nt.RefreshToken == "" {
		nt.RefreshToken = t.RefreshToken
	}
	// Schwab refresh tokens expire 7 days after login; refreshing does not extend that.
	nt.RefreshExpiresAt = t.RefreshExpiresAt
	if err := a.save(nt); err != nil {
		return "", err
	}
	a.tok = nt
	return nt.AccessToken, nil
}

// Login starts the OAuth callback listener and returns the URL to open.
// done receives nil on success, or the failure; it is sent exactly once.
func (a *Auth) Login(ctx context.Context) (authURL string, done <-chan error, err error) {
	cb, err := url.Parse(a.Cfg.CallbackURL)
	if err != nil {
		return "", nil, fmt.Errorf("SCHWAB_CALLBACK_URL: %w", err)
	}
	host := cb.Hostname()
	if cb.Scheme != "https" || (host != "127.0.0.1" && host != "localhost") {
		return "", nil, fmt.Errorf("SCHWAB_CALLBACK_URL must be https on 127.0.0.1 or localhost, got %q", a.Cfg.CallbackURL)
	}
	port := cb.Port()
	if port == "" {
		port = "443"
	}
	path := cb.Path
	if path == "" {
		path = "/"
	}

	cert, err := selfSignedCert()
	if err != nil {
		return "", nil, err
	}
	ln, err := tls.Listen("tcp", net.JoinHostPort(host, port), &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		return "", nil, fmt.Errorf("login listener: %w", err)
	}

	state := rand.Text()
	authURL = authorizeURL + "?" + url.Values{
		"client_id":     {a.Cfg.AppKey},
		"redirect_uri":  {a.Cfg.CallbackURL},
		"response_type": {"code"},
		"state":         {state},
	}.Encode()

	ctx, cancel := context.WithTimeout(ctx, loginTimeout)
	var (
		hmu     sync.Mutex // serializes callbacks so only one exchange runs
		ok      bool
		lastErr error
	)
	success := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		q := r.URL.Query()
		if q.Get("state") != state {
			http.Error(w, "invalid state", http.StatusBadRequest)
			return
		}
		hmu.Lock()
		defer hmu.Unlock()
		if ok {
			fmt.Fprintln(w, "Already logged in. You can close this tab.")
			return
		}
		code := q.Get("code")
		if code == "" {
			http.Error(w, "missing code: "+q.Get("error"), http.StatusBadRequest)
			return
		}
		if err := a.finishLogin(ctx, code); err != nil {
			lastErr = err
			http.Error(w, "login failed: "+err.Error(), http.StatusBadGateway)
			return
		}
		ok = true
		fmt.Fprintln(w, "Logged in to Schwab. You can close this tab.")
		close(success)
	})
	srv := &http.Server{Handler: mux, ReadHeaderTimeout: 10 * time.Second}
	go srv.Serve(ln)

	ch := make(chan error, 1)
	go func() {
		defer cancel()
		var err error
		select {
		case <-success:
		case <-ctx.Done():
			err = fmt.Errorf("login not completed: %w", ctx.Err())
			hmu.Lock()
			if lastErr != nil {
				err = fmt.Errorf("%w (last error: %v)", err, lastErr)
			}
			hmu.Unlock()
		}
		sctx, scancel := context.WithTimeout(context.Background(), 2*time.Second)
		if srv.Shutdown(sctx) != nil {
			srv.Close()
		}
		scancel()
		ch <- err
	}()

	openBrowser(authURL)
	return authURL, ch, nil
}

// StartLogin runs Login in the background for the MCP tool.
// A running login returns its URL; a failed one returns its error once.
func (a *Auth) StartLogin() (authURL string, err error) {
	a.loginMu.Lock()
	defer a.loginMu.Unlock()
	if a.loginURL != "" {
		return a.loginURL, nil
	}
	if err := a.loginErr; err != nil {
		a.loginErr = nil
		return "", fmt.Errorf("previous login failed: %w; call login again", err)
	}
	authURL, done, err := a.Login(context.Background())
	if err != nil {
		return "", err
	}
	a.loginURL = authURL
	go func() {
		err := <-done
		if err != nil {
			fmt.Fprintln(os.Stderr, "schwab: login:", err)
		}
		a.loginMu.Lock()
		a.loginURL, a.loginErr = "", err
		a.loginMu.Unlock()
	}()
	return authURL, nil
}

// finishLogin exchanges the authorization code and saves the token.
func (a *Auth) finishLogin(ctx context.Context, code string) error {
	form := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {a.Cfg.CallbackURL},
	}
	t, _, err := a.exchange(ctx, form)
	if err != nil {
		return err
	}
	t.RefreshExpiresAt = time.Now().Add(refreshLifetime)

	a.mu.Lock()
	defer a.mu.Unlock()
	unlock, err := a.lockFile()
	if err != nil {
		return err
	}
	defer unlock()
	if err := a.save(t); err != nil {
		return err
	}
	a.tok = t
	return nil
}

// exchange posts form to the token endpoint with Basic auth.
// status is the HTTP status when a response was received, else 0.
func (a *Auth) exchange(ctx context.Context, form url.Values) (*token, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, 0, err
	}
	req.SetBasicAuth(a.Cfg.AppKey, a.Cfg.AppSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := tokenHTTP.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("token request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("token response: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, resp.StatusCode, fmt.Errorf("token endpoint: HTTP %d: %s", resp.StatusCode, body)
	}
	var r struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, resp.StatusCode, fmt.Errorf("token response: %w", err)
	}
	if r.AccessToken == "" {
		return nil, resp.StatusCode, errors.New("token response: missing access_token")
	}
	return &token{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(r.ExpiresIn) * time.Second),
	}, resp.StatusCode, nil
}

func (a *Auth) load() (*token, error) {
	b, err := os.ReadFile(a.Cfg.TokenFile)
	if err != nil {
		return nil, err
	}
	var t token
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, fmt.Errorf("token file %s: %w", a.Cfg.TokenFile, err)
	}
	return &t, nil
}

// save writes the token atomically with 0600 permissions.
func (a *Auth) save(t *token) error {
	dir := filepath.Dir(a.Cfg.TokenFile)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.CreateTemp(dir, ".token-*.json") // CreateTemp uses mode 0600
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Rename(f.Name(), a.Cfg.TokenFile)
	}
	if err != nil {
		os.Remove(f.Name())
		return fmt.Errorf("save token: %w", err)
	}
	return nil
}

// lockFile takes the cross-process lock that guards the token file.
func (a *Auth) lockFile() (unlock func(), err error) {
	if err := os.MkdirAll(filepath.Dir(a.Cfg.TokenFile), 0o700); err != nil {
		return nil, err
	}
	return lockPath(a.Cfg.TokenFile + ".lock")
}

// selfSignedCert makes an in-memory cert for the loopback callback listener.
func selfSignedCert() (tls.Certificate, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, err
	}
	now := time.Now()
	tmpl := &x509.Certificate{
		Subject:     pkix.Name{CommonName: "schwab-mcp callback"},
		NotBefore:   now.Add(-time.Hour),
		NotAfter:    now.Add(24 * time.Hour),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost"},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}, nil
}

// openBrowser tries to open u; failures are ignored since the URL is also returned.
// The child gets no stdout, so it cannot corrupt the MCP stdio stream. Tests may replace it.
var openBrowser = func(u string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", u)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", u)
	default:
		cmd = exec.Command("xdg-open", u)
	}
	if cmd.Start() == nil {
		go cmd.Wait()
	}
}
