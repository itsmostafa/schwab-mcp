package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"slices"
	"strconv"
	"strings"
)

// maxBody caps response size so an oversized reply cannot exhaust memory.
const maxBody = 64 << 20

// Client calls the Schwab API with a bearer token from Auth.
type Client struct {
	BaseURL string
	Auth    *Auth
	HTTP    *http.Client
}

// do sends a request with an optional JSON body and returns the response body and headers.
// Non-2xx responses become errors carrying the status and Schwab's error body.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any) ([]byte, http.Header, error) {
	token, err := c.Auth.Token(ctx)
	if err != nil {
		return nil, nil, err
	}
	u := c.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	var rd io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		rd = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, rd)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	mutation := method != http.MethodGet
	resp, err := c.HTTP.Do(req)
	if err != nil {
		if mutation {
			return nil, nil, fmt.Errorf("%s %s: outcome unknown (%w); check get_orders before retrying", method, path, err)
		}
		return nil, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxBody+1))
	if err == nil && len(data) > maxBody {
		err = fmt.Errorf("response exceeds %d MB; narrow the request", maxBody>>20)
	}
	if err != nil {
		// 201 means Schwab accepted the order; callers only need its Location header.
		if mutation && resp.StatusCode == http.StatusCreated {
			return nil, resp.Header, nil
		}
		err = fmt.Errorf("%s %s: HTTP %d: reading body: %w", method, path, resp.StatusCode, err)
		if mutation {
			err = fmt.Errorf("%w; outcome unknown, check get_orders before retrying", err)
		}
		return nil, resp.Header, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err := fmt.Errorf("%s %s: HTTP %d: %s", method, path, resp.StatusCode, data)
		if mutation && resp.StatusCode >= 500 {
			err = fmt.Errorf("%w; outcome unknown, check get_orders before retrying", err)
		}
		return nil, resp.Header, err
	}
	return data, resp.Header, nil
}

// get sends a GET request and returns the response body.
func (c *Client) get(ctx context.Context, path string, q url.Values) ([]byte, error) {
	b, _, err := c.do(ctx, http.MethodGet, path, q, nil)
	return b, err
}

// toQuery turns a tool input struct into query params using its json tags.
// Slices are comma-joined, null values dropped, and keys in skip (path params) omitted.
func toQuery(in any, skip ...string) url.Values {
	// Marshal/decode cannot fail for plain tool input structs.
	b, _ := json.Marshal(in)
	var m map[string]any
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	_ = dec.Decode(&m)

	q := url.Values{}
	for k, v := range m {
		if v == nil || slices.Contains(skip, k) {
			continue
		}
		if list, ok := v.([]any); ok {
			parts := make([]string, len(list))
			for i, e := range list {
				parts[i] = fmt.Sprint(e)
			}
			q.Set(k, strings.Join(parts, ","))
			continue
		}
		q.Set(k, fmt.Sprint(v))
	}
	return q
}

// orderIDFromLocation returns the numeric order ID at the end of a Location header.
func orderIDFromLocation(h http.Header) (string, bool) {
	u, err := url.Parse(h.Get("Location"))
	if err != nil {
		return "", false
	}
	id := path.Base(u.Path)
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return "", false
	}
	return id, true
}
