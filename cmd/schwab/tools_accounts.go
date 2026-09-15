package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
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

// cappedOrderTypes can never fill worse than their limit price. Other types
// (MARKET, STOP, TRAILING_STOP, ...) fill at whatever the market is when they execute.
var cappedOrderTypes = []string{"LIMIT", "STOP_LIMIT", "TRAILING_STOP_LIMIT", "LIMIT_ON_CLOSE", "NET_DEBIT", "NET_CREDIT", "NET_ZERO"}

// quoteCheckTimeout bounds the pre-order quote request.
const quoteCheckTimeout = 5 * time.Second

// liveQuote is the part of a /marketdata/v1/quotes entry the order check reads.
type liveQuote struct {
	Realtime *bool `json:"realtime"`
	Quote    struct {
		BidPrice  float64 `json:"bidPrice"`
		AskPrice  float64 `json:"askPrice"`
		LastPrice float64 `json:"lastPrice"`
		Mark      float64 `json:"mark"`
	} `json:"quote"`
}

// orderShape is the part of a Schwab order the order check reads.
type orderShape struct {
	OrderType string `json:"orderType"`
	// json.Number accepts 100.5 and "100.5" and rejects NaN/Inf strings.
	Price              json.Number `json:"price"`
	StopPrice          json.Number `json:"stopPrice"`
	StopType           string      `json:"stopType"`
	OrderLegCollection []struct {
		Instruction string `json:"instruction"`
		Instrument  struct {
			Symbol string `json:"symbol"`
		} `json:"instrument"`
	} `json:"orderLegCollection"`
	ChildOrderStrategies []orderShape `json:"childOrderStrategies"`
}

// checkOrder re-quotes an order's symbols just before it is sent, because the agent
// built the order from prices that may be minutes old. Anywhere in the order tree it rejects
// uncapped order types (unless cfg.AllowMarketOrders) and legs without a real-time quote.
// For orders that can execute now it also rejects legs with no bid/ask on the side they trade
// against, and single-leg LIMIT (or already-triggered STOP_LIMIT) orders that cross the live
// bid/ask by more than cfg.MaxPriceDeviationBps. Resting limits below the ask (buys)
// or above the bid (sells) always pass. Children of a TRIGGER order wait for their parent
// to fill, so they get no bid/ask or price band check.
func checkOrder(ctx context.Context, c *Client, order map[string]any, cfg Config) error {
	b, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("order not sent: %w", err)
	}
	var top orderShape
	if err := json.Unmarshal(b, &top); err != nil {
		return fmt.Errorf("order not sent: cannot read order: %w", err)
	}
	// An OCO wrapper has no orderType of its own; its children are the real orders.
	orders := []orderShape{top}
	if top.OrderType == "" && len(top.ChildOrderStrategies) > 0 {
		orders = top.ChildOrderStrategies
	}
	// all is every order in the tree; wrappers (no orderType, only children) are skipped.
	var all []orderShape
	var walk func(o orderShape)
	walk = func(o orderShape) {
		if o.OrderType != "" || len(o.ChildOrderStrategies) == 0 {
			all = append(all, o)
		}
		for _, ch := range o.ChildOrderStrategies {
			walk(ch)
		}
	}
	walk(top)

	var symbols []string
	for _, o := range all {
		if !cfg.AllowMarketOrders && !slices.Contains(cappedOrderTypes, o.OrderType) {
			return fmt.Errorf("order not sent: orderType %q can fill at any price; use a LIMIT-type order, "+
				"or set SCHWAB_ALLOW_MARKET_ORDERS=true", o.OrderType)
		}
		for _, l := range o.OrderLegCollection {
			if !slices.Contains(symbols, l.Instrument.Symbol) {
				symbols = append(symbols, l.Instrument.Symbol)
			}
		}
	}
	if len(symbols) == 0 || slices.Contains(symbols, "") {
		return errors.New("order not sent: every order leg needs instrument.symbol")
	}

	// A slow quote is stale by the time the order goes out.
	// ponytail: if the token hits its 60s refresh point during this window, the order request
	// still waits on one OAuth round trip after the check; refresh with a wider skew here if that matters.
	qctx, cancel := context.WithTimeout(ctx, quoteCheckTimeout)
	defer cancel()
	raw, err := c.get(qctx, "/marketdata/v1/quotes", url.Values{"symbols": {strings.Join(symbols, ",")}})
	if err != nil {
		return fmt.Errorf("order not sent: live quote check failed: %w", err)
	}
	var quotes map[string]liveQuote
	if err := json.Unmarshal(raw, &quotes); err != nil {
		return fmt.Errorf("order not sent: live quote check failed: %w", err)
	}

	for _, sym := range symbols {
		if q, ok := quotes[sym]; !ok || q.Realtime == nil || !*q.Realtime {
			return fmt.Errorf("order not sent: no real-time quote for %s (live bid %g, ask %g, last %g)",
				sym, q.Quote.BidPrice, q.Quote.AskPrice, q.Quote.LastPrice)
		}
	}

	tol := float64(cfg.MaxPriceDeviationBps) / 1e4
	for _, o := range orders {
		for _, l := range o.OrderLegCollection {
			sym, buy := l.Instrument.Symbol, strings.HasPrefix(l.Instruction, "BUY")
			q := quotes[sym]
			live := fmt.Sprintf("(live %s bid %g, ask %g, last %g)", sym, q.Quote.BidPrice, q.Quote.AskPrice, q.Quote.LastPrice)
			switch {
			case buy && q.Quote.AskPrice <= 0:
				return fmt.Errorf("order not sent: no live ask for %s %s", sym, live)
			case !buy && q.Quote.BidPrice <= 0:
				return fmt.Errorf("order not sent: no live bid for %s %s", sym, live)
			}
			// Spreads have a net price no single quote maps to, so only single legs get the band check.
			if len(o.OrderLegCollection) != 1 || o.OrderType != "LIMIT" && o.OrderType != "STOP_LIMIT" {
				continue
			}
			ref := q.Quote.BidPrice
			if buy {
				ref = q.Quote.AskPrice
			}
			price, perr := o.Price.Float64()
			stop, serr := o.StopPrice.Float64()
			if perr != nil || o.OrderType == "STOP_LIMIT" && serr != nil {
				return fmt.Errorf("order not sent: cannot read price/stopPrice of %s order", o.OrderType)
			}
			// A STOP_LIMIT the market has not reached rests until Schwab sees the live price hit its stop.
			if o.OrderType == "STOP_LIMIT" {
				// Schwab triggers on the price stopType names. STANDARD (the default) is the last trade
				// for equities but bid/ask-based elsewhere, so it counts as reached when either is.
				var basis []float64
				switch o.StopType {
				case "", "STANDARD":
					basis = []float64{q.Quote.LastPrice, ref}
				case "LAST":
					basis = []float64{q.Quote.LastPrice}
				case "BID":
					basis = []float64{q.Quote.BidPrice}
				case "ASK":
					basis = []float64{q.Quote.AskPrice}
				case "MARK":
					basis = []float64{q.Quote.Mark}
				default:
					return fmt.Errorf("order not sent: unsupported stopType %q", o.StopType)
				}
				if !slices.ContainsFunc(basis, func(p float64) bool { return buy && p >= stop || !buy && p <= stop }) {
					continue
				}
			}
			if buy && price > ref*(1+tol) || !buy && price < ref*(1-tol) {
				return fmt.Errorf("order not sent: %s %s limit %g crosses the live market by more than %d bps %s; "+
					"prices moved since the order was built, re-check quotes and rebuild it",
					l.Instruction, sym, price, cfg.MaxPriceDeviationBps, live)
			}
		}
	}
	return nil
}

// registerAccountTools registers account, order, transaction and preference tools.
// Order-changing tools are registered only when cfg.AllowTrading is set.
func registerAccountTools(s *mcp.Server, c *Client, cfg Config) {
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

	if !cfg.AllowTrading {
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
		Description: "Place a live order. Call preview_order with the same order first and confirm with the user. " +
			"The server re-quotes first and rejects stale LIMIT prices, non-real-time quotes and MARKET/STOP orders. Returns the orderId."},
		func(ctx context.Context, in orderBodyIn) ([]byte, error) {
			if err := checkOrder(ctx, c, in.Order, cfg); err != nil {
				return nil, err
			}
			_, h, err := c.do(ctx, http.MethodPost, acct(in.AccountHash)+"/orders", nil, in.Order)
			if err != nil {
				return nil, err
			}
			return placed(h), nil
		})

	add(s, &mcp.Tool{Name: "replace_order", Annotations: destructive(),
		Description: "Replace a working order with a new one (the old order is canceled). Call preview_order with the new order first and confirm with the user. " +
			"The server re-quotes first, as for place_order. Returns the new orderId."},
		func(ctx context.Context, in replaceIn) ([]byte, error) {
			if err := checkOrder(ctx, c, in.Order, cfg); err != nil {
				return nil, err
			}
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
