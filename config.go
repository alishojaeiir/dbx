package dbx

// Config holds the configuration for establishing a database connection.
type Config struct {
	Driver          string // Driver is the database driver name (e.g., "mysql", "postgres", "sqlite3").
	Host            string // Host is the database host address (default: "127.0.0.1" for network drivers).
	Port            int    // Port is the database port (e.g., 3306 for MySQL, 5432 for PostgreSQL).
	Username        string // Username is the database user.
	Password        string // Password is the database user's password.
	DBName          string // DBName is the name of the database to connect to.
	SSLMode         string // SSLMode specifies the SSL/TLS mode (e.g., "disable", "require").
	MaxIdleConns    int    // MaxIdleConns sets the maximum number of idle connections (default: 10).
	MaxOpenConns    int    // MaxOpenConns sets the maximum number of open connections (default: 50).
	ConnMaxLifetime int    // ConnMaxLifetime sets the maximum lifetime of a connection in seconds (default: 300).
	RetryAttempts   int    // RetryAttempts specifies the number of retry attempts for connection (default: 3).
}
