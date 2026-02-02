package database

import (
	"context"
	"database/sql"
	"time"
)

// Config for CockroachDB connection
type Config struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// DB wraps *sql.DB for CockroachDB
type DB struct {
	*sql.DB
}

// NewCockroach creates a CockroachDB connection. Caller must import a Postgres-compatible driver (e.g. github.com/lib/pq or github.com/jackc/pgx).
func NewCockroach(ctx context.Context, cfg Config, driverName string) (*DB, error) {
	db, err := sql.Open(driverName, cfg.DSN)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return &DB{DB: db}, nil
}

// Close closes the database connection
func (d *DB) Close() error {
	return d.DB.Close()
}
