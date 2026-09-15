package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// allTransactionTypes is sent when get_transactions is called without types.
var allTransactionTypes = []string{
	"TRADE", "RECEIVE_AND_DELIVER", "DIVIDEND_OR_INTEREST", "ACH_RECEIPT", "ACH_DISBURSEMENT",
	"CASH_RECEIPT", "CASH_DISBURSEMENT", "ELECTRONIC_FUND", "WIRE_OUT", "WIRE_IN", "JOURNAL",
	"MEMORANDUM", "MARGIN_CALL", "MONEY_MARKET", "SMA_ADJUSTMENT",
}

type accountsIn struct {
	Fields string `json:"fields,omitempty" jsonschema:"Set to positions to include positions"`
}

type accountIn struct {
	AccountHash string `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	Fields      string `json:"fields,omitempty" jsonschema:"Set to positions to include positions"`
}

type ordersIn struct {
	AccountHash     string `json:"accountHash,omitempty" jsonschema:"Encrypted account hash; omit to list orders for all accounts"`
	FromEnteredTime string `json:"fromEnteredTime,omitempty" jsonschema:"Start time, ISO-8601 like 2024-01-31T00:00:00.000Z; default 30 days ago"`
	ToEnteredTime   string `json:"toEnteredTime,omitempty" jsonschema:"End time, ISO-8601 like 2024-01-31T23:59:59.000Z; default now"`
	MaxResults      *int   `json:"maxResults,omitempty" jsonschema:"Max orders to return (Schwab default 3000)"`
	Status          string `json:"status,omitempty" jsonschema:"One of AWAITING_PARENT_ORDER, AWAITING_CONDITION, AWAITING_STOP_CONDITION, AWAITING_MANUAL_REVIEW, ACCEPTED, AWAITING_UR_OUT, PENDING_ACTIVATION, QUEUED, WORKING, REJECTED, PENDING_CANCEL, CANCELED, PENDING_REPLACE, REPLACED, FILLED, EXPIRED, NEW, AWAITING_RELEASE_TIME, PENDING_ACKNOWLEDGEMENT, PENDING_RECALL, UNKNOWN"`
}

type orderIn struct {
	AccountHash string `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	OrderID     string `json:"orderId" jsonschema:"Order ID"`
}

type transactionsIn struct {
	AccountHash string   `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	StartDate   string   `json:"startDate,omitempty" jsonschema:"Start time, ISO-8601 like 2024-01-31T00:00:00.000Z; default 30 days ago"`
	EndDate     string   `json:"endDate,omitempty" jsonschema:"End time, ISO-8601 like 2024-01-31T23:59:59.000Z; default now"`
	Symbol      string   `json:"symbol,omitempty" jsonschema:"Only transactions for this symbol"`
	Types       []string `json:"types,omitempty" jsonschema:"Any of TRADE, RECEIVE_AND_DELIVER, DIVIDEND_OR_INTEREST, ACH_RECEIPT, ACH_DISBURSEMENT, CASH_RECEIPT, CASH_DISBURSEMENT, ELECTRONIC_FUND, WIRE_OUT, WIRE_IN, JOURNAL, MEMORANDUM, MARGIN_CALL, MONEY_MARKET, SMA_ADJUSTMENT; default all"`
}

type transactionIn struct {
	AccountHash   string `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	TransactionID string `json:"transactionId" jsonschema:"Transaction ID"`
}

type orderBodyIn struct {
	AccountHash string         `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	Order       map[string]any `json:"order" jsonschema:"Schwab order JSON object, e.g. {orderType: LIMIT, session: NORMAL, duration: DAY, price: 100.00, orderStrategyType: SINGLE, orderLegCollection: [{instruction: BUY, quantity: 1, instrument: {symbol: AAPL, assetType: EQUITY}}]}"`
}

type replaceIn struct {
	AccountHash string         `json:"accountHash" jsonschema:"Encrypted account hash (hashValue from get_account_numbers)"`
	OrderID     string         `json:"orderId" jsonschema:"ID of the order to replace"`
	Order       map[string]any `json:"order" jsonschema:"Complete new Schwab order JSON object (same shape as place_order)"`
}

// registerAccountTools registers account, order, transaction and preference tools.
// Order-changing tools are registered only when allowTrading is set.
func registerAccountTools(s *mcp.Server, c *Client, allowTrading bool) {
	// orDefault fills an empty ISO time with now minus ago, in Schwab's millisecond UTC layout.
	orDefault := func(v string, ago time.Duration) string {
		if v != "" {
			return v
		}
		return time.Now().UTC().Add(-ago).Format("2006-01-02T15:04:05.000Z")
	}
	acct := func(hash string) string { return "/trader/v1/accounts/" + url.PathEscape(hash) }
	const month = 30 * 24 * time.Hour

	add(s, &mcp.Tool{Name: "get_account_numbers", Annotations: readOnly(),
		Description: "List plain account numbers and their encrypted hashValue. Other account tools take the hashValue."},
		func(ctx context.Context, _ struct{}) ([]byte, error) {
			return c.get(ctx, "/trader/v1/accounts/accountNumbers", nil)
		})

	add(s, &mcp.Tool{Name: "get_accounts", Annotations: readOnly(),
		Description: "Get balances (and optionally positions) for all linked accounts."},
		func(ctx context.Context, in accountsIn) ([]byte, error) {
			return c.get(ctx, "/trader/v1/accounts", toQuery(in))
		})

	add(s, &mcp.Tool{Name: "get_account", Annotations: readOnly(),
		Description: "Get balances (and optionally positions) for one account."},
		func(ctx context.Context, in accountIn) ([]byte, error) {
			return c.get(ctx, acct(in.AccountHash), toQuery(in, "accountHash"))
		})

	add(s, &mcp.Tool{Name: "get_orders", Annotations: readOnly(),
		Description: "List orders for one account, or all accounts when accountHash is omitted. Defaults to the last 30 days."},
		func(ctx context.Context, in ordersIn) ([]byte, error) {
			in.FromEnteredTime = orDefault(in.FromEnteredTime, month)
			in.ToEnteredTime = orDefault(in.ToEnteredTime, 0)
			p := "/trader/v1/orders"
			if in.AccountHash != "" {
				p = acct(in.AccountHash) + "/orders"
			}
			return c.get(ctx, p, toQuery(in, "accountHash"))
		})

	add(s, &mcp.Tool{Name: "get_order", Annotations: readOnly(),
		Description: "Get one order by ID."},
		func(ctx context.Context, in orderIn) ([]byte, error) {
			return c.get(ctx, acct(in.AccountHash)+"/orders/"+url.PathEscape(in.OrderID), nil)
		})

	add(s, &mcp.Tool{Name: "get_transactions", Annotations: readOnly(),
		Description: "List transactions for an account. Defaults to the last 30 days and all transaction types."},
		func(ctx context.Context, in transactionsIn) ([]byte, error) {
			in.StartDate = orDefault(in.StartDate, month)
			in.EndDate = orDefault(in.EndDate, 0)
			if len(in.Types) == 0 {
				in.Types = allTransactionTypes
			}
			return c.get(ctx, acct(in.AccountHash)+"/transactions", toQuery(in, "accountHash"))
		})

	add(s, &mcp.Tool{Name: "get_transaction", Annotations: readOnly(),
		Description: "Get one transaction by ID."},
		func(ctx context.Context, in transactionIn) ([]byte, error) {
			return c.get(ctx, acct(in.AccountHash)+"/transactions/"+url.PathEscape(in.TransactionID), nil)
		})

	add(s, &mcp.Tool{Name: "get_user_preference", Annotations: readOnly(),
		Description: "Get user preferences, including account nicknames and streamer info."},
		func(ctx context.Context, _ struct{}) ([]byte, error) {
			return c.get(ctx, "/trader/v1/userPreference", nil)
		})

	add(s, &mcp.Tool{Name: "preview_order", Annotations: readOnly(),
		Description: "Validate an order and show estimated cost, commissions and warnings without placing it."},
		func(ctx context.Context, in orderBodyIn) ([]byte, error) {
			b, _, err := c.do(ctx, http.MethodPost, acct(in.AccountHash)+"/previewOrder", nil, in.Order)
			return b, err
		})

	if !allowTrading {
		return
	}

	// placed reports the new order ID, or warns not to resubmit when Schwab omitted it.
	placed := func(h http.Header) []byte {
		if id, ok := orderIDFromLocation(h); ok {
			return []byte("order accepted, orderId: " + id)
		}
		return []byte("order accepted (HTTP 201) but ID unknown — check get_orders; do NOT resubmit")
	}

	add(s, &mcp.Tool{Name: "place_order", Annotations: destructive(),
		Description: "Place a live order. Call preview_order with the same order first and confirm with the user. Returns the orderId."},
		func(ctx context.Context, in orderBodyIn) ([]byte, error) {
			_, h, err := c.do(ctx, http.MethodPost, acct(in.AccountHash)+"/orders", nil, in.Order)
			if err != nil {
				return nil, err
			}
			return placed(h), nil
		})

	add(s, &mcp.Tool{Name: "replace_order", Annotations: destructive(),
		Description: "Replace a working order with a new one (the old order is canceled). Call preview_order with the new order first and confirm with the user. Returns the new orderId."},
		func(ctx context.Context, in replaceIn) ([]byte, error) {
			_, h, err := c.do(ctx, http.MethodPut, acct(in.AccountHash)+"/orders/"+url.PathEscape(in.OrderID), nil, in.Order)
			if err != nil {
				return nil, err
			}
			return placed(h), nil
		})

	add(s, &mcp.Tool{Name: "cancel_order", Annotations: destructive(),
		Description: "Cancel a working order. Confirm with the user first; check the result with get_order."},
		func(ctx context.Context, in orderIn) ([]byte, error) {
			b, _, err := c.do(ctx, http.MethodDelete, acct(in.AccountHash)+"/orders/"+url.PathEscape(in.OrderID), nil, nil)
			return b, err
		})
}
