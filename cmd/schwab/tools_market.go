package main

import (
	"context"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type quotesIn struct {
	Symbols    []string `json:"symbols" jsonschema:"symbols to quote, e.g. AAPL, $SPX, /ES, or an option symbol like AAPL  240517C00190000"`
	Fields     []string `json:"fields,omitempty" jsonschema:"data subsets: quote, fundamental, extended, reference, regular; default all"`
	Indicative *bool    `json:"indicative,omitempty" jsonschema:"also return indicative quotes for ETF symbols"`
}

type chainIn struct {
	Symbol                 string   `json:"symbol" jsonschema:"underlying symbol, e.g. AAPL or $SPX"`
	ContractType           string   `json:"contractType,omitempty" jsonschema:"CALL, PUT, or ALL"`
	StrikeCount            *int     `json:"strikeCount,omitempty" jsonschema:"number of strikes above and below the at-the-money price"`
	IncludeUnderlyingQuote *bool    `json:"includeUnderlyingQuote,omitempty" jsonschema:"include the underlying quote"`
	Strategy               string   `json:"strategy,omitempty" jsonschema:"SINGLE, ANALYTICAL, COVERED, VERTICAL, CALENDAR, STRANGLE, STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, or ROLL; default SINGLE"`
	Interval               *float64 `json:"interval,omitempty" jsonschema:"strike interval for spread strategy chains"`
	Strike                 *float64 `json:"strike,omitempty" jsonschema:"only return this strike price"`
	Range                  string   `json:"range,omitempty" jsonschema:"ITM, NTM, OTM, SAK, SBK, SNK, or ALL"`
	FromDate               string   `json:"fromDate,omitempty" jsonschema:"earliest expiration, YYYY-MM-DD"`
	ToDate                 string   `json:"toDate,omitempty" jsonschema:"latest expiration, YYYY-MM-DD"`
	Volatility             *float64 `json:"volatility,omitempty" jsonschema:"volatility for ANALYTICAL strategy calculations"`
	UnderlyingPrice        *float64 `json:"underlyingPrice,omitempty" jsonschema:"underlying price for ANALYTICAL strategy calculations"`
	InterestRate           *float64 `json:"interestRate,omitempty" jsonschema:"interest rate for ANALYTICAL strategy calculations"`
	DaysToExpiration       *int     `json:"daysToExpiration,omitempty" jsonschema:"days to expiration for ANALYTICAL strategy calculations"`
	ExpMonth               string   `json:"expMonth,omitempty" jsonschema:"expiration month: JAN through DEC, or ALL"`
	OptionType             string   `json:"optionType,omitempty" jsonschema:"option type, e.g. S (standard) or NS (non-standard); default all"`
	Entitlement            string   `json:"entitlement,omitempty" jsonschema:"retail client entitlement: PN (paying non-pro), NP (non-paying), or PP (paying pro)"`
}

type symbolIn struct {
	Symbol string `json:"symbol" jsonschema:"underlying symbol, e.g. AAPL"`
}

type priceHistoryIn struct {
	Symbol                string `json:"symbol" jsonschema:"symbol, e.g. AAPL"`
	PeriodType            string `json:"periodType,omitempty" jsonschema:"day, month, year, or ytd"`
	Period                *int   `json:"period,omitempty" jsonschema:"number of periodType units to return"`
	FrequencyType         string `json:"frequencyType,omitempty" jsonschema:"minute, daily, weekly, or monthly"`
	Frequency             *int   `json:"frequency,omitempty" jsonschema:"frequencyType units per candle, e.g. 1, 5, 10, 15, 30 for minute"`
	StartDate             *int64 `json:"startDate,omitempty" jsonschema:"start time in epoch milliseconds"`
	EndDate               *int64 `json:"endDate,omitempty" jsonschema:"end time in epoch milliseconds"`
	NeedExtendedHoursData *bool  `json:"needExtendedHoursData,omitempty" jsonschema:"include extended hours candles"`
	NeedPreviousClose     *bool  `json:"needPreviousClose,omitempty" jsonschema:"include previous close price and date"`
}

type moversIn struct {
	Symbol    string `json:"symbol" jsonschema:"index: $DJI, $COMPX, $SPX, NYSE, NASDAQ, OTCBB, INDEX_ALL, EQUITY_ALL, OPTION_ALL, OPTION_PUT, or OPTION_CALL"`
	Sort      string `json:"sort,omitempty" jsonschema:"VOLUME, TRADES, PERCENT_CHANGE_UP, or PERCENT_CHANGE_DOWN"`
	Frequency *int   `json:"frequency,omitempty" jsonschema:"minimum percent change: 0, 1, 5, 10, 30, or 60"`
}

type marketHoursIn struct {
	Markets []string `json:"markets" jsonschema:"markets: equity, option, bond, future, forex"`
	Date    string   `json:"date,omitempty" jsonschema:"date YYYY-MM-DD; default today"`
}

type instrumentsIn struct {
	Symbol     string `json:"symbol" jsonschema:"symbol, regex, or description text to search for"`
	Projection string `json:"projection" jsonschema:"symbol-search, symbol-regex, desc-search, desc-regex, search, or fundamental"`
}

type instrumentIn struct {
	Cusip string `json:"cusip" jsonschema:"CUSIP of the instrument"`
}

func registerMarketTools(s *mcp.Server, c *Client) {
	add(s, &mcp.Tool{
		Name:        "get_quotes",
		Description: "Get quotes for one or more symbols (equities, ETFs, indexes, futures, options).",
		Annotations: readOnly(),
	}, func(ctx context.Context, in quotesIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/quotes", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "get_option_chain",
		Description: "Get the option chain for an underlying symbol. Use strikeCount, range, or fromDate/toDate to keep the response small.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in chainIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/chains", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "get_option_expirations",
		Description: "Get the option expiration dates for an underlying symbol.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in symbolIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/expirationchain", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "get_price_history",
		Description: "Get OHLCV price history candles for a symbol, by period or by startDate/endDate in epoch milliseconds.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in priceHistoryIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/pricehistory", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "get_movers",
		Description: "Get the top 10 movers for an index or market segment.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in moversIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/movers/"+url.PathEscape(in.Symbol), toQuery(in, "symbol"))
	})

	add(s, &mcp.Tool{
		Name:        "get_market_hours",
		Description: "Get market hours for one or more markets on a date.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in marketHoursIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/markets", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "search_instruments",
		Description: "Search instruments by symbol or description; projection fundamental returns fundamentals for a symbol.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in instrumentsIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/instruments", toQuery(in))
	})

	add(s, &mcp.Tool{
		Name:        "get_instrument",
		Description: "Get an instrument by CUSIP.",
		Annotations: readOnly(),
	}, func(ctx context.Context, in instrumentIn) ([]byte, error) {
		return c.get(ctx, "/marketdata/v1/instruments/"+url.PathEscape(in.Cusip), nil)
	})
}
