# dbx
A lightweight Go package for simplified database connections and query execution.

## Features
- Automatic DSN generation based on driver (Postgres, MySQL, SQLite).
- Built-in connection retry mechanism.
- Context-aware query execution with prepared statements.

## Installation
```bash
go get github.com/alishojaeiir/dbx
```
Make sure you have the appropriate database drivers installed:

* PostgreSQL: go get github.com/lib/pq
* MySQL: go get github.com/go-sql-driver/mysql
* SQLite: go get github.com/mattn/go-sqlite3

## Usage
Here's a simple example demonstrating how to connect to a SQLite database, execute a query, and retrieve the result:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/alishojaeiir/dbx"
)

func main() {
	// Configure the connection
	config := dbx.Config{
		Driver: "sqlite3",
		DBName: ":memory:",
	}

	// Connect to the database
	db, err := dbx.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Executor().Close() // Close the connection when done

	// Run a simple query
	rows, err := db.Executor().QueryContext(context.Background(), "SELECT sqlite_version()")
	if err != nil {
		log.Fatalf("Failed to query: %v", err)
	}
	defer rows.Close() // Close the rows after use

	// Fetch the result
	if rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			log.Fatalf("Failed to scan: %v", err)
		}
		fmt.Println("SQLite version:", version)
	}
}
```

## Supported Drivers
* PostgreSQL: Use "postgres" as the driver name.
* MySQL: Use "mysql" as the driver name.
* SQLite: Use "sqlite3" as the driver name.

## Configuration Options
The Config struct allows you to customize the database connection:

* Driver: The database driver name (e.g., "postgres", "mysql", "sqlite3").
* Host: The database host (default: "127.0.0.1" for network-based drivers).
* Port: The database port (e.g., 5432 for PostgreSQL).
* Username: Database user.
* Password: Database password.
* DBName: Name of the database.
* SSLMode: SSL mode (e.g., "disable", "require").
* MaxIdleConns: Maximum number of idle connections (default: 10).
* MaxOpenConns: Maximum number of open connections (default: 50).
* ConnMaxLifetime: Maximum lifetime of a connection in seconds (default: 300).
* RetryAttempts: Number of retry attempts for failed connections (default: 3).

## Contributing
Feel free to submit issues or pull requests to the [GitHub repository](https://github.com/alishojaeiir/dbx).

