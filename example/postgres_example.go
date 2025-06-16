package example

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/dbx"
	"log"
)

func postgres_example() {
	cfg := dbx.NewConfig(
		dbx.WithDriver("postgres"),
		dbx.WithHost("127.0.0.1"),
		dbx.WithPort(5432),
		dbx.WithUsername("root"),
		dbx.WithPassword("password"),
		dbx.WithDBName("db"),
		dbx.WithSSLMode("disable"),
		dbx.WithMaxIdleConns(15),
		dbx.WithMaxOpenConns(100),
		dbx.WithConnMaxLifetime(5),
		dbx.WithRetryAttempts(3),
	)

	conn, err := dbx.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Executor().Close()

	row, err := conn.Executor().QueryRowContext(context.Background(), "SELECT version()")
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	var version string
	if err := row.Scan(&version); err != nil {
		log.Fatalf("failed to scan row: %v", err)
	}

	fmt.Println("PostgreSQL version:", version)
}
