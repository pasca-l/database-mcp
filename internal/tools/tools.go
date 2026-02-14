package tools

import "github.com/modelcontextprotocol/go-sdk/mcp"

type McpTool[In, Out any] struct {
	Tool    *mcp.Tool
	Handler mcp.ToolHandlerFor[In, Out]
}
