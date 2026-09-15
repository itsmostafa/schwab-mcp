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
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
)

// version is set by release builds via -ldflags "-X main.version=...".
var version = "dev"

func init() {
	// Without -X, use the version the go tool stamps: the tag for
	// `go install ...@vX.Y.Z`, a pseudo-version for a local `go build`.
	if info, ok := debug.ReadBuildInfo(); ok && version == "dev" && info.Main.Version != "" && info.Main.Version != "(devel)" {
		version = info.Main.Version
	}
}

// Config holds settings read from SCHWAB_* environment variables.
type Config struct {
	AppKey, AppSecret, CallbackURL, TokenFile string
	AllowTrading                              bool
	// AllowMarketOrders permits order types without a price cap (MARKET, STOP, ...).
	AllowMarketOrders bool
	// MaxPriceDeviationBps is how far a LIMIT order may cross the live bid/ask.
	MaxPriceDeviationBps int
}

func loadConfig() (Config, error) {
	cfg := Config{
		AppKey:               os.Getenv("SCHWAB_APP_KEY"),
		AppSecret:            os.Getenv("SCHWAB_APP_SECRET"),
		CallbackURL:          os.Getenv("SCHWAB_CALLBACK_URL"),
		TokenFile:            os.Getenv("SCHWAB_TOKEN_FILE"),
		AllowTrading:         os.Getenv("SCHWAB_ALLOW_TRADING") == "true",
		AllowMarketOrders:    os.Getenv("SCHWAB_ALLOW_MARKET_ORDERS") == "true",
		MaxPriceDeviationBps: 50,
	}
	if cfg.AppKey == "" || cfg.AppSecret == "" {
		return cfg, errors.New("SCHWAB_APP_KEY and SCHWAB_APP_SECRET must be set")
	}
	if v := os.Getenv("SCHWAB_MAX_PRICE_DEVIATION_BPS"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			return cfg, fmt.Errorf("SCHWAB_MAX_PRICE_DEVIATION_BPS must be a non-negative integer, got %q", v)
		}
		cfg.MaxPriceDeviationBps = n
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := newRootCmd().ExecuteContext(ctx)
	stop()
	if err != nil {
		fmt.Fprintln(os.Stderr, "schwab:", err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	// Wraps ctx-only funcs as RunE.
	run := func(f func(context.Context) error) func(*cobra.Command, []string) error {
		return func(cmd *cobra.Command, _ []string) error { return f(cmd.Context()) }
	}
	root := &cobra.Command{
		Use:     "schwab",
		Short:   "MCP stdio server for the Schwab Trader API (serves when run without a command)",
		Version: version,
		// Stdout carries the MCP protocol; main reports errors on stderr.
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          run(serve),
	}
	root.SetVersionTemplate("{{.Version}}\n")

	mcpCmd := &cobra.Command{Use: "mcp", Short: "Manage MCP client configuration"}
	mcpCmd.AddCommand(&cobra.Command{
		Use:   "setup",
		Short: "Register this binary with Claude Code and Codex",
		Args:  cobra.NoArgs,
		RunE:  run(runMCPSetup),
	})
	root.AddCommand(
		&cobra.Command{Use: "serve", Short: "Run the MCP server over stdio", Args: cobra.NoArgs, RunE: run(serve)},
		&cobra.Command{Use: "login", Short: "Authorize with Schwab and save the token", Args: cobra.NoArgs, RunE: run(login)},
		&cobra.Command{Use: "update", Short: "Replace this binary with the latest release", Args: cobra.NoArgs, RunE: run(runUpdate)},
		&cobra.Command{Use: "version", Short: "Print the version", Args: cobra.NoArgs, Run: func(*cobra.Command, []string) {
			fmt.Println(version)
		}},
		mcpCmd,
	)
	return root
}

func serve(ctx context.Context) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	s := mcp.NewServer(&mcp.Implementation{Name: "schwab", Version: version}, nil)
	c := &Client{
		BaseURL: "https://api.schwabapi.com",
		Auth:    NewAuth(cfg),
		HTTP:    &http.Client{Timeout: 30 * time.Second},
	}
	registerTools(s, c, cfg)
	return s.Run(ctx, &mcp.StdioTransport{})
}

func login(ctx context.Context) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	authURL, done, err := NewAuth(cfg).Login(ctx)
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
}
