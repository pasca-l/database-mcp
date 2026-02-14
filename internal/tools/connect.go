package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pasca-l/database-mcp/internal/db"
)

type ConnectInput struct {
	DSN string `json:"dsn" jsonschema_description:"PostgreSQL connection string (e.g. postgres://user:password@localhost:5432/dbname)"`
}

type ConnectOutput struct {
	Success bool `json:"success" jsonschema_description:"Whether the connection was successful"`
}

func NewConnectTool(conn *db.Connection) *McpTool[ConnectInput, ConnectOutput] {
	return &McpTool[ConnectInput, ConnectOutput]{
		Tool: &mcp.Tool{
			Name:        "connect",
			Description: "Connect to a PostgreSQL database using a connection string (e.g., postgres://user:password@localhost:5432/dbname)",
		},
		Handler: buildConnectToolHandler(conn),
	}
}

func buildConnectToolHandler(conn *db.Connection) func(ctx context.Context, req *mcp.CallToolRequest, input ConnectInput) (*mcp.CallToolResult, ConnectOutput, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input ConnectInput) (*mcp.CallToolResult, ConnectOutput, error) {
		err := conn.Connect(ctx, input.DSN)
		if err != nil {
			return nil, ConnectOutput{}, err
		}

		return nil, ConnectOutput{
			Success: true,
		}, nil
	}
}
