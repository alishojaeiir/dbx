package example

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/dbx"
	"log"
)

func main() {
	cfg := dbx.Config{
		Driver:          "mysql",
		Host:            "127.0.0.1",
		Port:            3306,
		Username:        "root",
		Password:        "",
		DBName:          "db",
		SSLMode:         "false",
		MaxIdleConns:    15,
		MaxOpenConns:    100,
		ConnMaxLifetime: 5,
		RetryAttempts:   3,
	}

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

	fmt.Println("mysql version:", version)
}
