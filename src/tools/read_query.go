package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pasca-l/database-mcp/db"
	"github.com/pasca-l/database-mcp/utils"
)

type ReadQueryInput struct {
	Query string `json:"query" jsonschema_description:"SQL SELECT query to execute (use $1, $2, etc. for parameters)"`
	Args  []any  `json:"args,omitempty" jsonschema_description:"Parameters for the query"`
}

type ReadQueryOutput struct {
	Columns  []string `json:"columns" jsonschema_description:"Column names"`
	Rows     [][]any  `json:"rows" jsonschema_description:"Query result rows (limited to 100)"`
	RowCount int      `json:"row_count" jsonschema_description:"Number of rows returned (max 100)"`
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
		if err := utils.ValidateReadQuery(input.Query); err != nil {
			return nil, ReadQueryOutput{}, err
		}

		columns, rows, err := conn.ReadQuery(ctx, input.Query, input.Args)
		if err != nil {
			return nil, ReadQueryOutput{}, fmt.Errorf("error executing query: %w", err)
		}

		return nil, ReadQueryOutput{
			Columns:  columns,
			Rows:     rows,
			RowCount: len(rows),
		}, nil
	}
}
