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
	RetryInterval   int    //Interval between retry attempts (seconds) (default: 1).
}

// ConfigOption defines a function that modifies a Config.
type ConfigOption func(*Config)

// NewConfig creates a Config with default values, modified by the provided options.
// Default values:
//   - Driver: "sqlite3"
//   - DBName: ":memory:"
//   - Host: "127.0.0.1"
//   - SSLMode: "disable"
//   - MaxIdleConns: 20
//   - MaxOpenConns: 50
//   - ConnMaxLifetime: 600 seconds (10 minutes)
//   - RetryAttempts: 3
//   - RetryInterval: 1 second
func NewConfig(opts ...ConfigOption) Config {
	cfg := Config{
		Driver:          "sqlite3",  // default SQLite
		DBName:          ":memory:", // in-memory database
		Host:            "127.0.0.1",
		SSLMode:         "disable",
		MaxIdleConns:    20,
		MaxOpenConns:    50,
		ConnMaxLifetime: 600,
		RetryAttempts:   3,
		RetryInterval:   1,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func WithDriver(driver string) ConfigOption {
	return func(c *Config) { c.Driver = driver }
}

func WithHost(host string) ConfigOption {
	return func(c *Config) { c.Host = host }
}

func WithPort(port int) ConfigOption {
	return func(c *Config) { c.Port = port }
}

func WithUsername(username string) ConfigOption {
	return func(c *Config) { c.Username = username }
}

func WithPassword(password string) ConfigOption {
	return func(c *Config) { c.Password = password }
}

func WithDBName(dbName string) ConfigOption {
	return func(c *Config) { c.DBName = dbName }
}

func WithSSLMode(sslMode string) ConfigOption {
	return func(c *Config) { c.SSLMode = sslMode }
}

func WithMaxIdleConns(maxIdleConns int) ConfigOption {
	return func(c *Config) { c.MaxIdleConns = maxIdleConns }
}

func WithMaxOpenConns(maxOpenConns int) ConfigOption {
	return func(c *Config) { c.MaxOpenConns = maxOpenConns }
}

func WithConnMaxLifetime(seconds int) ConfigOption {
	return func(c *Config) { c.ConnMaxLifetime = seconds }
}

func WithRetryAttempts(attempts int) ConfigOption {
	return func(c *Config) { c.RetryAttempts = attempts }
}

func WithRetryInterval(seconds int) ConfigOption {
	return func(c *Config) { c.RetryInterval = seconds }
}
