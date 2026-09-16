package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestTransactionsChunking(t *testing.T) {
	var mu sync.Mutex
	var windows [][2]string
	cs := session(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		windows = append(windows, [2]string{r.URL.Query().Get("startDate"), r.URL.Query().Get("endDate")})
		n := len(windows)
		mu.Unlock()
		fmt.Fprintf(w, `[{"activityId":%d,"description":"AT&T INC","transferItems":[
			{"instrument":{"symbol":"AAPL"},"amount":10,"cost":-1000},
			{"feeType":"COMMISSION","amount":0.0,"cost":0.0},
			{"feeType":"SEC_FEE","amount":0.01,"cost":-0.01}]}]`, n)
	}), Config{})
	// reqs returns the windows requested so far and resets the list.
	reqs := func() [][2]string {
		mu.Lock()
		defer mu.Unlock()
		w := windows
		windows = nil
		return w
	}

	got := call(t, cs, "get_transactions", map[string]any{"accountHash": "H1",
		"startDate": "2023-01-01T00:00:00-05:00", "endDate": "2025-07-01T00:00:00-05:00"})
	var rows []struct {
		ActivityID    int              `json:"activityId"`
		TransferItems []map[string]any `json:"transferItems"`
	}
	if err := json.Unmarshal([]byte(got), &rows); err != nil {
		t.Fatalf("result %q: %v", got, err)
	}

	reqd := reqs()
	if len(reqd) != 3 {
		t.Fatalf("requests = %q, want 3", reqd)
	}
	const layout = "2006-01-02T15:04:05.000Z"
	if reqd[0][1] != "2025-07-01T05:00:00.000Z" || reqd[2][0] != "2023-01-01T05:00:00.000Z" {
		t.Errorf("outer bounds = %q", reqd)
	}
	for i, w := range reqd {
		s, serr := time.Parse(layout, w[0])
		e, eerr := time.Parse(layout, w[1])
		if serr != nil || eerr != nil || !s.Before(e) || e.Sub(s) > 364*24*time.Hour {
			t.Errorf("window %d = %q", i, w)
		}
		if i > 0 {
			if prev, _ := time.Parse(layout, reqd[i-1][0]); !e.Add(time.Millisecond).Equal(prev) {
				t.Errorf("window %d end %s not 1ms before window %d start %s", i, w[1], i-1, reqd[i-1][0])
			}
		}
	}

	if len(rows) != 3 {
		t.Fatalf("rows = %s", got)
	}
	for i, r := range rows {
		if r.ActivityID != i+1 {
			t.Errorf("row %d activityId %d, want newest window first", i, r.ActivityID)
		}
		if len(r.TransferItems) != 2 || r.TransferItems[0]["feeType"] != nil || r.TransferItems[1]["feeType"] != "SEC_FEE" {
			t.Errorf("row %d transferItems = %v", i, r.TransferItems)
		}
	}

	// The default range is one window and still goes through the parsed path that drops zero fees.
	text := call(t, cs, "get_transactions", map[string]any{"accountHash": "H1"})
	if n := len(reqs()); n != 1 || strings.Contains(text, "COMMISSION") || !strings.Contains(text, `"AT&T INC"`) {
		t.Errorf("default range made %d requests, result %s", n, text)
	}

	// Unparseable dates and ranges over 20 years fail before any Schwab call.
	for _, dates := range [][2]string{{"2023-01-01", "2024-01-01T00:00:00Z"}, {"1900-01-01T00:00:00Z", "2025-01-01T00:00:00Z"}} {
		if text, isErr := callResult(t, cs, "get_transactions", map[string]any{"accountHash": "H1",
			"startDate": dates[0], "endDate": dates[1]}); !isErr || len(reqs()) != 0 {
			t.Errorf("dates %q: isErr %v, text %q", dates, isErr, text)
		}
	}

	// go-sdk rejects a string for the array-typed types field before the handler runs.
	if text, isErr := callResult(t, cs, "get_transactions", map[string]any{"accountHash": "H1", "types": `["TRADE"]`}); !isErr || !strings.Contains(text, "validating") {
		t.Errorf("string types: isErr %v, text %q", isErr, text)
	}
}
