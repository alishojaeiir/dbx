package dbx

import (
	"database/sql"
	"fmt"
	"time"
)

// Connect establishes a connection to the database with retry support.
func Connect(config Config) (*Database, error) {

	if config.Driver == "" {
		return nil, fmt.Errorf("%w: driver is required", ErrInvalidConfig)
	}
	if config.DBName == "" && config.Driver != "sqlite3" {
		return nil, fmt.Errorf("%w: DBName is required for non-SQLite drivers", ErrInvalidConfig)
	}

	dsnBuilder, err := getDSNBuilder(config.Driver)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnsupportedDriver, err)
	}
	dsn := dsnBuilder.BuildDSN(config)
	var conn *sql.DB
	retries := max(1, config.RetryAttempts)
	interval := config.RetryInterval
	if interval == 0 {
		interval = 1
	}
	for i := 0; i < retries; i++ {
		conn, err = sql.Open(config.Driver, dsn)
		if err == nil {
			if err = conn.Ping(); err == nil {
				break
			}
			conn.Close()
		}
		if i < retries-1 {
			time.Sleep(time.Second * time.Duration(interval))
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
