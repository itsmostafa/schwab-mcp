<p align="center">
  <img src="assets/img/schwab-logo.png" alt="Schwab MCP logo" width="180">
</p>

<h1 align="center">Schwab MCP</h1>

<p align="center">
  <strong>Let your AI agent read your Schwab accounts, pull market data, and place trades.</strong><br>
  A single, ultra-fast binary MCP server for the Schwab Trader API. Works with Claude Desktop, Claude Code, and Codex.
</p>

<p align="center">
  <a href="https://github.com/itsmostafa/schwab-mcp/releases/latest"><img src="https://img.shields.io/github/v/release/itsmostafa/schwab-mcp?sort=semver" alt="Latest release"></a>
  <a href="go.mod"><img src="https://img.shields.io/github/go-mod/go-version/itsmostafa/schwab-mcp" alt="Go version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/itsmostafa/schwab-mcp" alt="License: MIT"></a>
</p>

```text
┌──────────────────┐   stdio (MCP)   ┌──────────┐   HTTPS + OAuth   ┌───────────────────┐
│  Claude / Codex  │ ──────────────▶ │  schwab  │ ────────────────▶ │ Schwab Trader API │
└──────────────────┘                 └──────────┘                   └───────────────────┘
                                          │
                                   token.json (auto-refreshed)
```

> Unofficial project. Not affiliated with or endorsed by Charles Schwab & Co., Inc.

## Why this exists

**The problem:** asking an agent "how is my portfolio doing?" or "what does the SPY option chain look like for next Friday?" means exporting CSVs, copy-pasting quotes, or writing your own OAuth client against the Schwab API. Schwab's OAuth flow alone (HTTPS callback, 30-minute access tokens, 7-day refresh tokens) is enough to stall a weekend project.

**The solution:** `schwab` handles the login flow, stores and refreshes tokens for you, and exposes accounts, orders, transactions, and market data as MCP tools. Install one binary, log in once a week, and your agent can answer questions against live data. Trading stays off unless you explicitly turn it on.

## Quickstart

**1. Install** (macOS and Linux, amd64 and arm64):

```sh
curl -fsSL https://raw.githubusercontent.com/itsmostafa/schwab-mcp/main/install.sh | sh
```

**2. Log in** with your Schwab app credentials ([how to get them](#1-create-a-schwab-app)):

```sh
SCHWAB_APP_KEY=... SCHWAB_APP_SECRET=... schwab login
```

**3. Add it to your agent** (Claude Code shown; [other clients below](#5-add-to-your-mcp-client)):

```sh
claude mcp add schwab -e SCHWAB_APP_KEY=... -e SCHWAB_APP_SECRET=... -- ~/.local/bin/schwab
```

Then ask: *"Show my account balances and today's top movers in the S&P 500."*

## What you get

- **Zero-dependency install.** One static Go binary with checksum-verified downloads and `schwab update` to upgrade in place.
- **Login without the OAuth headache.** `schwab login` runs the HTTPS callback server for you and saves the token. Expired access tokens refresh automatically on the next request.
- **Log in from a desktop app, no terminal needed.** Ask the agent to call the `login` tool, approve in the browser, and retry.
- **Read-only by default.** Order placement tools do not exist until you set `SCHWAB_ALLOW_TRADING=true`. `preview_order` always works, so you can check cost and warnings with no risk.
- **Run it in several clients at once.** Multiple agents share one token file safely on macOS and Linux.
- **Full market data coverage.** Quotes, option chains and expirations, OHLCV history, movers, market hours, and instrument search.
- **Agent-friendly responses.** Tools are annotated read-only or destructive, and descriptions steer the agent to keep responses small and to preview before trading.

## Setup

### 1. Create a Schwab app

1. Sign in at <https://developer.schwab.com> and create an app with the **Accounts and Trading Production** and **Market Data Production** APIs.
2. Set the callback URL to exactly `https://127.0.0.1:8182/callback`.
3. Wait until the app status is **Ready For Use**, then copy the App Key and Secret.

### 2. Install

```sh
curl -fsSL https://raw.githubusercontent.com/itsmostafa/schwab-mcp/main/install.sh | sh
```

This installs the latest release to `~/.local/bin/schwab`. Set `SCHWAB_INSTALL_DIR` to install somewhere else. Run `schwab update` to upgrade in place.

Or build from source with Go:

```sh
go install github.com/itsmostafa/schwab-mcp/cmd/schwab@latest
```

This installs `schwab` in `$(go env GOPATH)/bin`.

### 3. Configure

| Variable | Required | Default |
|---|---|---|
| `SCHWAB_APP_KEY` | yes | |
| `SCHWAB_APP_SECRET` | yes | |
| `SCHWAB_CALLBACK_URL` | no | `https://127.0.0.1:8182/callback` (must match the app registration exactly) |
| `SCHWAB_TOKEN_FILE` | no | `<user config dir>/schwab-mcp/token.json` |
| `SCHWAB_ALLOW_TRADING` | no | unset. Set to `true` to enable `place_order`, `replace_order`, `cancel_order` |
| `SCHWAB_ALLOW_MARKET_ORDERS` | no | unset. Set to `true` to allow order types without a price cap (`MARKET`, `STOP`, `TRAILING_STOP`, ...) |
| `SCHWAB_MAX_PRICE_DEVIATION_BPS` | no | `50`. How far, in basis points, a `LIMIT` price may cross the live bid/ask |

### 4. Log in

```sh
SCHWAB_APP_KEY=... SCHWAB_APP_SECRET=... schwab login
```

Open the printed URL and approve access. Schwab then redirects to `https://127.0.0.1:8182/callback`, which is served by `schwab` with a self-signed certificate, so the browser shows a certificate warning. Click through it (for example "Advanced" then "Proceed"). The token is saved to `SCHWAB_TOKEN_FILE`.

Schwab refresh tokens last 7 days, so log in again about once a week. From a desktop client with no terminal, ask the agent to call the `login` tool: it returns the URL, and you retry the request after approving.

### 5. Add to your MCP client

Use the absolute path to the binary (`~/.local/bin/schwab` from the install script, `~/go/bin/schwab` from `go install`); desktop apps often do not have these directories on `PATH`.

Claude Code:

```sh
claude mcp add schwab -e SCHWAB_APP_KEY=your-app-key -e SCHWAB_APP_SECRET=your-app-secret -- /Users/you/.local/bin/schwab
```

Claude Desktop (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "schwab": {
      "command": "/Users/you/.local/bin/schwab",
      "env": {
        "SCHWAB_APP_KEY": "your-app-key",
        "SCHWAB_APP_SECRET": "your-app-secret"
      }
    }
  }
}
```

Codex (`~/.codex/config.toml`):

```toml
[mcp_servers.schwab]
command = "/Users/you/.local/bin/schwab"

[mcp_servers.schwab.env]
SCHWAB_APP_KEY = "your-app-key"
SCHWAB_APP_SECRET = "your-app-secret"
```

Several clients can run the server at once; they share the token file safely on macOS and Linux (there is no file lock on Windows).

## Trading

> [!WARNING]
> With `SCHWAB_ALLOW_TRADING=true`, an agent can place, replace, and cancel **live orders with real money**. Tool descriptions tell the agent to call `preview_order` and confirm with you first, but nothing enforces that. Leave it unset unless you want this, and review every order request before approving the tool call.

Without the flag, only read-only tools and `preview_order` are available.

An agent builds orders from prices it fetched earlier, and markets move in the seconds it spends reasoning. So `place_order` and `replace_order` fetch a fresh quote for every leg right before sending, and reject the order without sending it when:

- the order type has no price cap (`MARKET`, `STOP`, `TRAILING_STOP`, ...), unless `SCHWAB_ALLOW_MARKET_ORDERS=true`
- a leg's quote is not real-time, or has no bid (sells) or ask (buys)
- a single-leg `LIMIT` buy (or a `STOP_LIMIT` whose stop the market already passed, judged by its `stopType`; `STANDARD` counts as passed when either the last trade or the bid/ask has) is priced more than `SCHWAB_MAX_PRICE_DEVIATION_BPS` above the live ask, or a sell that far below the live bid

Resting limits (a buy below the ask, a sell above the bid) always pass. Multi-leg prices and the children of `TRIGGER` orders are not band-checked. The rejection includes the live bid, ask and last so the agent can rebuild the order. `get_quotes` and `get_option_chain` start with a warning line when Schwab serves delayed data.

## Tools

| Area | Tools |
|---|---|
| Auth | `login` |
| Accounts | `get_account_numbers`, `get_accounts`, `get_account`, `get_user_preference` |
| Orders | `get_orders`, `get_order`, `preview_order`, and with trading enabled `place_order`, `replace_order`, `cancel_order` |
| Transactions | `get_transactions`, `get_transaction` |
| Market data | `get_quotes`, `get_option_chain`, `get_option_expirations`, `get_price_history`, `get_movers`, `get_market_hours`, `search_instruments`, `get_instrument` |

Account tools take the encrypted `hashValue` from `get_account_numbers`, not the plain account number.

## Contributing

Bug reports and pull requests are welcome. Open an [issue](https://github.com/itsmostafa/schwab-mcp/issues) to report a problem or propose a tool, and see the [`Taskfile.yml`](Taskfile.yml) for build, test, and lint commands.

If this saves you time, a ⭐ on the repo helps others find it.

## License

[MIT](LICENSE)
