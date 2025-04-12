package dbx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL
	_ "github.com/lib/pq"              // PostgreSQL
	_ "github.com/mattn/go-sqlite3"    // SQLite
)

var DSNRegistry = make(map[string]DSNBuilder)

func RegisterDSNBuilder(dsn string, builder DSNBuilder) {
	DSNRegistry[dsn] = builder
}

func getDSNBuilder(dsn string) (DSNBuilder, error) {
	builder, exists := DSNRegistry[dsn]
	if !exists {
		return nil, fmt.Errorf("database driver %s is not registered", dsn)
	}
	return builder, nil
}

func init() {
	RegisterDSNBuilder("postgres", &PostgresDSNBuilder{})
	RegisterDSNBuilder("mysql", &MySQLDSNBuilder{})
	RegisterDSNBuilder("sqlite3", &SQLiteDSNBuilder{})
}

// MySQLDSNBuilder builds DSN for MySQL.
type MySQLDSNBuilder struct{}

// BuildDSN constructs the DSN string for MySQL.
func (m *MySQLDSNBuilder) BuildDSN(config Config) string {
	tlsValue := "false"
	switch config.SSLMode {
	case "disable", "false":
		tlsValue = "false"
	case "require", "true":
		tlsValue = "true"
	case "verify-ca", "verify-full", "skip-verify":
		tlsValue = "skip-verify"
	default:
		tlsValue = config.SSLMode
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, tlsValue)
}

// PostgresDSNBuilder builds DSN for PostgreSQL.
type PostgresDSNBuilder struct{}

// BuildDSN constructs the DSN string for PostgreSQL.
func (p *PostgresDSNBuilder) BuildDSN(config Config) string {
	tlsValue := "disable"
	switch config.SSLMode {
	case "disable", "false":
		tlsValue = "disable"
	case "require", "true":
		tlsValue = "require"
	case "verify-ca":
		tlsValue = "verify-ca"
	case "verify-full", "skip-verify":
		tlsValue = "verify-full"
	default:
		tlsValue = config.SSLMode
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, tlsValue)
}

// SQLiteDSNBuilder builds DSN for SQLite.
type SQLiteDSNBuilder struct{}

// BuildDSN constructs the DSN string for SQLite.
func (s *SQLiteDSNBuilder) BuildDSN(config Config) string {
	return config.DBName
}
