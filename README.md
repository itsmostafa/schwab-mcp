# Schwab - MCP

An MCP stdio server that lets AI agents (Claude Desktop, Claude Code, Codex) use the Schwab Trader API: accounts, orders, transactions, and market data.

## 1. Create a Schwab app

1. Sign in at <https://developer.schwab.com> and create an app with the **Accounts and Trading Production** and **Market Data Production** APIs.
2. Set the callback URL to exactly `https://127.0.0.1:8182/callback`.
3. Wait until the app status is **Ready For Use**, then copy the App Key and Secret.

## 2. Install

```sh
go install github.com/itsmostafa/schwab-mcp/cmd/schwab@latest
```

This installs a binary named `schwab` in `$(go env GOPATH)/bin`.

## 3. Configure

| Variable | Required | Default |
|---|---|---|
| `SCHWAB_APP_KEY` | yes | |
| `SCHWAB_APP_SECRET` | yes | |
| `SCHWAB_CALLBACK_URL` | no | `https://127.0.0.1:8182/callback` (must match the app registration exactly) |
| `SCHWAB_TOKEN_FILE` | no | `<user config dir>/schwab-mcp/token.json` |
| `SCHWAB_ALLOW_TRADING` | no | unset. Set to `true` to enable `place_order`, `replace_order`, `cancel_order` |

## 4. Log in

```sh
SCHWAB_APP_KEY=... SCHWAB_APP_SECRET=... schwab login
```

Open the printed URL and approve access. Schwab then redirects to `https://127.0.0.1:8182/callback`, which is served by `schwab` with a self-signed certificate, so the browser shows a certificate warning. Click through it (for example "Advanced" then "Proceed"). The token is saved to `SCHWAB_TOKEN_FILE`.

Schwab refresh tokens last 7 days, so log in again about once a week. From a desktop client with no terminal, ask the agent to call the `login` tool: it returns the URL, and you retry the request after approving.

## 5. Add to your MCP client

Use the absolute path to the binary; desktop apps often do not have `~/go/bin` on `PATH`.

Claude Desktop (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "schwab": {
      "command": "/Users/you/go/bin/schwab",
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
command = "/Users/you/go/bin/schwab"

[mcp_servers.schwab.env]
SCHWAB_APP_KEY = "your-app-key"
SCHWAB_APP_SECRET = "your-app-secret"
```

Several clients can run the server at once; they share the token file safely on macOS and Linux (there is no file lock on Windows).

## Trading

> **Warning:** with `SCHWAB_ALLOW_TRADING=true`, an agent can place, replace, and cancel **live orders with real money**. Tool descriptions tell the agent to call `preview_order` and confirm with you first, but nothing enforces that. Leave it unset unless you want this, and review every order request before approving the tool call.

Without the flag, only read-only tools and `preview_order` are available.

## Tools

- Auth: `login`
- Accounts: `get_account_numbers`, `get_accounts`, `get_account`, `get_user_preference`
- Orders: `get_orders`, `get_order`, `preview_order`, and with trading enabled `place_order`, `replace_order`, `cancel_order`
- Transactions: `get_transactions`, `get_transaction`
- Market data: `get_quotes`, `get_option_chain`, `get_option_expirations`, `get_price_history`, `get_movers`, `get_market_hours`, `search_instruments`, `get_instrument`

Account tools take the encrypted `hashValue` from `get_account_numbers`, not the plain account number.
