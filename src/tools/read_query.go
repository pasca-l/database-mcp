package tools

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pasca-l/database-mcp/db"
)

var (
	ErrNotSelectQuery = errors.New("read_query only supports SELECT statements")
)

type ReadQueryInput struct {
	Query   string `json:"query" jsonschema_description:"SQL SELECT query to execute"`
	MaxRows int    `json:"max_rows" jsonschema_description:"Maximum number of rows to return (default 100)"`
}

type ReadQueryOutput struct {
	Columns  []string `json:"columns" jsonschema_description:"Column names"`
	Rows     [][]any  `json:"rows" jsonschema_description:"Query result rows"`
	RowCount int      `json:"row_count" jsonschema_description:"Number of rows returned"`
}

func NewReadQueryTool(conn *db.Connection) *McpTool[ReadQueryInput, ReadQueryOutput] {
	return &McpTool[ReadQueryInput, ReadQueryOutput]{
		Tool: &mcp.Tool{
			Name:        "read_query",
			Description: "Execute a SELECT query to read data from the database. This tool is read-only and will not modify any data.",
		},
		Handler: buildReadQueryToolHandler(conn),
	}
}

func buildReadQueryToolHandler(conn *db.Connection) func(ctx context.Context, req *mcp.CallToolRequest, input ReadQueryInput) (*mcp.CallToolResult, ReadQueryOutput, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input ReadQueryInput) (*mcp.CallToolResult, ReadQueryOutput, error) {
		// validate if query is a SELECT query
		trimmed := strings.TrimSpace(strings.ToUpper(input.Query))
		if !strings.HasPrefix(trimmed, "SELECT") {
			return nil, ReadQueryOutput{}, fmt.Errorf("%w, got: %s", ErrNotSelectQuery, input.Query)
		}

		columns, rows, err := conn.ReadQuery(ctx, input.Query, input.MaxRows)
		if err != nil {
			return nil, ReadQueryOutput{}, fmt.Errorf("error executing query: %w", err)
		}

		if rows == nil {
			rows = [][]any{}
		}

		return nil, ReadQueryOutput{
			Columns:  columns,
			Rows:     rows,
			RowCount: len(rows),
		}, nil
	}
}
