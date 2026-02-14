package db

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNoConnection = errors.New("no database connection established")
)

type Connection struct {
	pool *pgxpool.Pool
	mu   sync.RWMutex
	dsn  string
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(ctx context.Context, dsn string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pool != nil {
		c.pool.Close()
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	c.pool = pool
	c.dsn = dsn

	return nil
}

func (c *Connection) GetPool() (*pgxpool.Pool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.pool == nil {
		return nil, ErrNoConnection
	}

	return c.pool, nil
}

func (c *Connection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.pool != nil {
		c.pool.Close()
		c.pool = nil
		c.dsn = ""
	}

	return nil
}
