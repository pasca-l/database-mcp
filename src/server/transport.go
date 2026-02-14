package server

import "github.com/modelcontextprotocol/go-sdk/mcp"

func NewTransport() mcp.Transport {
	return &mcp.StdioTransport{}
}
