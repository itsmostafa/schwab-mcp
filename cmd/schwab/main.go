// Command schwab is an MCP stdio server for the Schwab Trader API.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Config holds settings read from SCHWAB_* environment variables.
type Config struct {
	AppKey, AppSecret, CallbackURL, TokenFile string
	AllowTrading                              bool
}

func loadConfig() (Config, error) {
	cfg := Config{
		AppKey:       os.Getenv("SCHWAB_APP_KEY"),
		AppSecret:    os.Getenv("SCHWAB_APP_SECRET"),
		CallbackURL:  os.Getenv("SCHWAB_CALLBACK_URL"),
		TokenFile:    os.Getenv("SCHWAB_TOKEN_FILE"),
		AllowTrading: os.Getenv("SCHWAB_ALLOW_TRADING") == "true",
	}
	if cfg.AppKey == "" || cfg.AppSecret == "" {
		return cfg, errors.New("SCHWAB_APP_KEY and SCHWAB_APP_SECRET must be set")
	}
	if cfg.CallbackURL == "" {
		cfg.CallbackURL = "https://127.0.0.1:8182/callback"
	}
	if cfg.TokenFile == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return cfg, fmt.Errorf("set SCHWAB_TOKEN_FILE: %w", err)
		}
		cfg.TokenFile = filepath.Join(dir, "schwab-mcp", "token.json")
	}
	return cfg, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "schwab:", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	auth := NewAuth(cfg)

	cmd := "serve"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "serve":
		s := mcp.NewServer(&mcp.Implementation{Name: "schwab", Version: "0.1.0"}, nil)
		c := &Client{
			BaseURL: "https://api.schwabapi.com",
			Auth:    auth,
			HTTP:    &http.Client{Timeout: 30 * time.Second},
		}
		registerTools(s, c, cfg.AllowTrading)
		return s.Run(ctx, &mcp.StdioTransport{})
	case "login":
		authURL, done, err := auth.Login(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Open this URL and approve access (accept the self-signed certificate warning):")
		fmt.Println(authURL)
		if err := <-done; err != nil {
			return err
		}
		fmt.Println("Logged in. Token saved to", cfg.TokenFile)
		return nil
	default:
		return fmt.Errorf("unknown command %q (usage: schwab [serve|login])", cmd)
	}
}
