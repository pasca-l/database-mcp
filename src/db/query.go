package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	MaxExtractRows = 100
	QueryTimeout   = 30 * time.Second
)

func (c *Connection) ReadQuery(ctx context.Context, query string, args []any) ([]string, [][]any, error) {
	queryCtx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	pool, err := c.GetPool()
	if err != nil {
		return nil, nil, err
	}

	pgxRows, err := pool.Query(queryCtx, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer pgxRows.Close()

	columns := extractColumns(pgxRows)
	rows, err := extractRows(pgxRows)
	if err != nil {
		return nil, nil, err
	}
	return columns, rows, nil
}

func (c *Connection) WriteQuery(ctx context.Context, query string, args []any) (int64, error) {
	queryCtx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	pool, err := c.GetPool()
	if err != nil {
		return 0, err
	}

	commandTag, err := pool.Exec(queryCtx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	return commandTag.RowsAffected(), nil
}

func extractColumns(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()

	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}
	return columns
}

func extractRows(rows pgx.Rows) ([][]any, error) {
	results := make([][]any, 0, MaxExtractRows)
	rowCount := 0

	for rows.Next() {
		if rowCount >= MaxExtractRows {
			break
		}

		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		results = append(results, values)
		rowCount++
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return results, nil
}
