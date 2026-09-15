package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// add registers a tool whose handler returns raw text (usually Schwab JSON).
// A returned error becomes a tool result with IsError set.
func add[In any](s *mcp.Server, t *mcp.Tool, fn func(ctx context.Context, in In) ([]byte, error)) {
	mcp.AddTool(s, t, func(ctx context.Context, _ *mcp.CallToolRequest, in In) (*mcp.CallToolResult, any, error) {
		b, err := fn(ctx, in)
		if err != nil {
			return nil, nil, err
		}
		if len(b) == 0 {
			b = []byte("OK")
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: string(b)}}}, nil, nil
	})
}

// readOnly marks a tool that does not change anything.
func readOnly() *mcp.ToolAnnotations {
	return &mcp.ToolAnnotations{ReadOnlyHint: true}
}

// destructive marks a tool that changes live orders.
func destructive() *mcp.ToolAnnotations {
	return &mcp.ToolAnnotations{DestructiveHint: new(true)}
}

func registerTools(s *mcp.Server, c *Client, allowTrading bool) {
	add(s, &mcp.Tool{
		Name: "login",
		Description: "Start Schwab OAuth login. Returns a URL the user must open in a browser to approve access; " +
			"then retry the failed request. Use when a tool reports not logged in.",
		Annotations: &mcp.ToolAnnotations{DestructiveHint: new(false)},
	}, func(ctx context.Context, _ struct{}) ([]byte, error) {
		authURL, err := c.Auth.StartLogin()
		if err != nil {
			return nil, err
		}
		return []byte("Open this URL in a browser and approve access (accept the self-signed certificate warning " +
			"on the redirect), then retry the previous request:\n" + authURL), nil
	})
	registerAccountTools(s, c, allowTrading)
	registerMarketTools(s, c)
}
