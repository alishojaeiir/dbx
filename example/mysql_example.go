package main

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/dbx"
	"log"
)

func mysql_example() {
	config := dbx.NewConfig(
		dbx.WithDriver("mysql"),
		dbx.WithHost("127.0.0.1"),
		dbx.WithPort(3306),
		dbx.WithUsername("root"),
		dbx.WithPassword(""),
		dbx.WithDBName("db"),
		dbx.WithSSLMode("disable"),
		dbx.WithMaxIdleConns(15),
		dbx.WithMaxOpenConns(100),
		dbx.WithConnMaxLifetime(5),
		dbx.WithRetryAttempts(3),
	)
	db, err := dbx.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Executor().Close()

	row, err := db.Executor().QueryRowContext(context.Background(), "SELECT VERSION()")
	if err != nil {
		log.Fatalf("Failed to query: %v", err)
	}
	var version string
	if err := row.Scan(&version); err != nil {
		log.Fatalf("Failed to scan: %v", err)
	}
	fmt.Println("MySQL version:", version)
}
