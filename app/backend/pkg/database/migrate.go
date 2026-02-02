package database

import (
	"context"
	"database/sql"
)

// Migrator runs database migrations (e.g. Goose)
type Migrator interface {
	Up(ctx context.Context, db *sql.DB) error
	Down(ctx context.Context, db *sql.DB) error
}

// GooseMigrator placeholder for Goose-based migrations
type GooseMigrator struct {
	Dialect string
}

// Up runs migrations up
func (g *GooseMigrator) Up(ctx context.Context, db *sql.DB) error {
	// Implement with pressly/goose or similar
	return nil
}

// Down runs migrations down
func (g *GooseMigrator) Down(ctx context.Context, db *sql.DB) error {
	return nil
}
