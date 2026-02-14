package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pasca-l/database-mcp/db"
)

type WriteQueryInput struct {
	Query string `json:"query" jsonschema_description:"SQL query to execute (INSERT, UPDATE, DELETE, etc.). Use $1, $2, etc. for parameters to prevent SQL injection."`
	Args  []any  `json:"args,omitempty" jsonschema_description:"Parameters for the query"`
}

type WriteQueryOutput struct {
	RowsAffected int64 `json:"rows_affected" jsonschema_description:"Number of rows affected by the query"`
}

func NewWriteQueryTool(conn *db.Connection) *McpTool[WriteQueryInput, WriteQueryOutput] {
	return &McpTool[WriteQueryInput, WriteQueryOutput]{
		Tool: &mcp.Tool{
			Name:        "write_query",
			Description: "Execute a write query (INSERT, UPDATE, DELETE) to modify data in the database. Use with caution as this will change database contents.",
		},
		Handler: buildWriteQueryToolHandler(conn),
	}
}

func buildWriteQueryToolHandler(conn *db.Connection) func(ctx context.Context, req *mcp.CallToolRequest, input WriteQueryInput) (*mcp.CallToolResult, WriteQueryOutput, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input WriteQueryInput) (*mcp.CallToolResult, WriteQueryOutput, error) {
		rowsAffected, err := conn.WriteQuery(ctx, input.Query, input.Args)
		if err != nil {
			return nil, WriteQueryOutput{}, fmt.Errorf("error executing query: %w", err)
		}

		return nil, WriteQueryOutput{
			RowsAffected: rowsAffected,
		}, nil
	}
}
