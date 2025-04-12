package dbx

import (
	"database/sql"
	"fmt"
	"time"
)

// Connect establishes a connection to the database with retry support.
// It uses the provided Config to build a DSN and attempts to connect, retrying up to RetryAttempts times.
// On success, it configures the connection pool with MaxIdleConns, MaxOpenConns, and ConnMaxLifetime.
func Connect(config Config) (*Database, error) {
	dsnBuilder, err := getDSNBuilder(config.Driver)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnsupportedDriver, err)
	}
	dsn := dsnBuilder.BuildDSN(config)
	var conn *sql.DB
	retries := max(1, config.RetryAttempts)
	for i := 0; i < retries; i++ {
		conn, err = sql.Open(config.Driver, dsn)
		if err == nil {
			if err = conn.Ping(); err == nil {
				break
			}
			conn.Close()
		}
		if i < retries-1 {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}
	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	fmt.Println("Database connection established successfully")

	return &Database{DB: conn, Dialect: config.Driver}, nil
}

// max returns the greater of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
