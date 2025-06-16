package dbx

// Config holds the configuration for establishing a database connection.
type Config struct {
	Driver          string `koanf:"driver"` //postgres, mysql, sqlite3
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	Username        string `koanf:"username"`
	Password        string `koanf:"password"`
	DBName          string `koanf:"db_name"`
	SSLMode         string `koanf:"ssl_mode"`
	MaxIdleConns    int    `koanf:"max_idle_conns"`
	MaxOpenConns    int    `koanf:"max_open_conns"`
	ConnMaxLifetime int    `koanf:"conn_max_lifetime"`
	RetryAttempts   int    `koanf:"retry_attempts"`
	RetryInterval   int    `koanf:"retry_interval"`
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
