package db

import (
	"context"
	"fmt"
)

func (c *Connection) ReadQuery(ctx context.Context, query string, maxRows int) ([]string, [][]any, error) {
	pool, err := c.GetPool()
	if err != nil {
		return nil, nil, err
	}

	if maxRows <= 0 {
		maxRows = 100
	}

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}

	results := make([][]any, 0, maxRows)
	rowCount := 0

	for rows.Next() {
		if rowCount >= maxRows {
			break
		}

		values, err := rows.Values()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		results = append(results, values)
		rowCount++
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return columns, results, nil
}

func (c *Connection) WriteQuery(ctx context.Context, query string) (int64, error) {
	pool, err := c.GetPool()
	if err != nil {
		return 0, err
	}

	commandTag, err := pool.Exec(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return commandTag.RowsAffected(), nil
}
