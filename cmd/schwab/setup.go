package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// goRunDir matches the temp directory `go run` builds into, deleted on exit.
var goRunDir = regexp.MustCompile(`/go-build\d+/`)

// runMCPSetup registers this binary as the "schwab" MCP server with Claude Code
// and Codex, via their own CLIs, baking in the SCHWAB_* variables from the
// current environment: clients launch the server without the user's shell env.
func runMCPSetup(ctx context.Context) error {
	if _, err := loadConfig(); err != nil {
		return err
	}
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("finding current executable: %w", err)
	}
	if goRunDir.MatchString(exe) {
		return fmt.Errorf("refusing to configure %s: `go run` binaries are deleted on exit; build or install schwab first", exe)
	}

	var env []string
	for _, kv := range os.Environ() {
		if strings.HasPrefix(kv, "SCHWAB_") {
			env = append(env, kv)
		}
	}

	var errs []error
	for _, c := range setupCommands(exe, env) {
		if _, err := exec.LookPath(c.cli); err != nil {
			fmt.Printf("%s: skipped, `%s` not on PATH (see README to configure by hand)\n", c.name, c.cli)
			continue
		}
		var prev []byte
		if c.reset != nil {
			prev = claudeUserEntry()
			// Fails when there is no entry yet; nothing to report either way.
			exec.CommandContext(ctx, c.cli, c.reset...).Run()
		}
		cmd := exec.CommandContext(ctx, c.cli, c.add...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", c.name, err))
			if prev != nil {
				// Not ctx: an interrupted add must still put the old entry back.
				if err := exec.Command(c.cli, "mcp", "add-json", "schwab", string(prev), "-s", "user").Run(); err != nil {
					errs = append(errs, fmt.Errorf("%s: restoring previous entry: %w", c.name, err))
				}
			}
		}
	}
	return errors.Join(errs...)
}

// claudeUserEntry returns Claude Code's current user-scope schwab entry, or nil
// if there is none or the config cannot be read.
func claudeUserEntry() []byte {
	dir := os.Getenv("CLAUDE_CONFIG_DIR")
	if dir == "" {
		dir, _ = os.UserHomeDir()
	}
	b, err := os.ReadFile(filepath.Join(dir, ".claude.json"))
	if err != nil {
		return nil
	}
	var cfg struct {
		MCPServers map[string]json.RawMessage `json:"mcpServers"`
	}
	if json.Unmarshal(b, &cfg) != nil {
		return nil
	}
	return cfg.MCPServers["schwab"]
}

type setupCommand struct {
	name, cli  string
	reset, add []string
}

func setupCommands(exe string, env []string) []setupCommand {
	claude := []string{"mcp", "add", "schwab", "-s", "user"}
	codex := []string{"mcp", "add", "schwab"}
	for _, kv := range env {
		// One flag per pair: claude's -e is variadic and would swallow the name.
		claude = append(claude, "-e", kv)
		codex = append(codex, "--env", kv)
	}
	return []setupCommand{
		// `claude mcp add` refuses an existing name; `codex mcp add` overwrites.
		{"Claude Code", "claude", []string{"mcp", "remove", "schwab", "-s", "user"}, append(claude, "--", exe)},
		{"Codex", "codex", nil, append(codex, "--", exe)},
	}
}
