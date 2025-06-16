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

## Quick Start
### Default SQLite Connection
If no options are provided, `NewConfig` creates an in-memory SQLite database:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/alishojaeiir/dbx"
)

func main() {
	config := dbx.NewConfig() // Uses SQLite with :memory:
	db, err := dbx.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Executor().Close()

	_, err = db.Executor().ExecContext(context.Background(), "CREATE TABLE users (id INTEGER, name TEXT)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully")
}
```

Make sure you have the appropriate database drivers installed:

* PostgreSQL: go get github.com/lib/pq
* MySQL: go get github.com/go-sql-driver/mysql
* SQLite: go get github.com/mattn/go-sqlite3

## Usage
Here's a simple example demonstrating how to connect to a SQLite database, execute a query, and retrieve the result:

### mysql
```go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/alishojaeiir/dbx"
)

func main() {
	config := dbx.NewConfig(
		dbx.WithDriver("mysql"),
		dbx.WithHost("127.0.0.1"),
		dbx.WithPort(3306),
		dbx.WithUsername("root"),
		dbx.WithDBName("db"),
		dbx.WithSSLMode("disable"),
		dbx.WithMaxIdleConns(15),
		dbx.WithMaxOpenConns(100),
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
```

## Supported Drivers
* PostgreSQL: Use "postgres" as the driver name.
* MySQL: Use "mysql" as the driver name.
* SQLite: Use "sqlite3" as the driver name.

## Configuration Options
The dbx.Config struct is configured using NewConfig with the following options:

| Option                  | Description                                     | Default Value |
|-------------------------|-------------------------------------------------|---------------|
| `WithDriver`            | Database driver name (e.g., "mysql", "sqlite3") | "sqlite3"     |
| `WithHost`              | Database host address                           | "127.0.0.1"   |
| `WithPort`              | Database port                                   | -             |
| `WithUsername`          | Database username                               | -             |
| `WithPassword`          | Database password                               | -             |
| `WithDBName`            | Database name                                   | ":memory:"    |
| `WithSSLMode`           | SSL/TLS mode (e.g., "disable", "require")       | "disable"     |
| `WithMaxIdleConns`      | Maximum number of idle connections              | 20            |
| `WithMaxOpenConns`      | Maximum number of open connections              | 50            |
| `WithConnMaxLifetime`   | Maximum lifetime of a connection (seconds)      | 600           |
| `WithRetryAttempts`     | Number of connection retry attempts             | 3             |
| `WithRetryInterval`     | Interval between retry attempts (seconds)       | 1             |

**Note**: For non-SQLite drivers (e.g., MySQL, PostgreSQL), WithDBName is required. For production, configure WithSSLMode("require") or WithSSLMode("verify-full").

## Contributing
Feel free to submit issues or pull requests to the [GitHub repository](https://github.com/alishojaeiir/dbx).

