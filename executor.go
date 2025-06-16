package dbx

import (
	"context"
	"database/sql"
	"fmt"
)

// QueryExecutor provides methods to execute database queries with prepared statements.
type QueryExecutor struct {
	DB DBExecutor
}

// QueryRowContext executes a query that returns at most one row.
func (e *QueryExecutor) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	stmt, err := e.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	return stmt.QueryRowContext(ctx, args...), nil
}

// ExecContext executes a query that modifies the database with the given context.
// It prepares the statement, executes it with the provided arguments, and closes the statement.
func (e *QueryExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := e.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	return stmt.ExecContext(ctx, args...)
}

// QueryContext executes a query that returns multiple rows.
func (e *QueryExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := e.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	return stmt.QueryContext(ctx, args...)
}

func (e *QueryExecutor) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return e.DB.PrepareContext(ctx, query)
}

func (e *QueryExecutor) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return e.DB.BeginTx(ctx, opts)
}

func (e *QueryExecutor) Close() error {
	return e.DB.Close()
}
