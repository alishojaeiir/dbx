package example

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/dbx"
	"log"
)

func slite_example() {
	config := dbx.NewConfig()
	db, err := dbx.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Executor().Close()

	_, err = db.Executor().ExecContext(context.Background(), "CREATE TABLE users (id INTEGER, name TEXT)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	row, err := db.Executor().QueryRowContext(context.Background(), "SELECT sqlite_version()")
	if err != nil {
		log.Fatalf("Failed to query: %v", err)
	}
	var version string
	if err := row.Scan(&version); err != nil {
		log.Fatalf("Failed to scan: %v", err)
	}
	fmt.Println("SQLite version:", version)
}
