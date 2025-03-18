package dbx

import (
	"context"
	"database/sql"
)

// DSNBuilder defines the interface for building DSN strings.
type DSNBuilder interface {
	BuildDSN(config Config) string
}

// DBExecutor defines the essential database operations used by repositories.
// It abstracts common query and transaction methods to facilitate mocking, testing, and flexibility.
type DBExecutor interface {
	// QueryContext executes a query that returns multiple rows, typically a SELECT statement.
	// It accepts a context for timeout and cancellation control.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext executes a query expected to return a single row.
	// Useful for queries like SELECT with LIMIT 1, COUNT, or primary key lookups.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row

	// ExecContext executes a query without returning any rows, typically used for
	// INSERT, UPDATE, DELETE, or DDL statements.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// PrepareContext creates a prepared statement for repeated execution.
	// It improves performance and helps prevent SQL injection.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	// BeginTx starts a new database transaction with the provided context and options.
	// It allows for atomic execution of multiple statements.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	Close() error
}
