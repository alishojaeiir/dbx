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
}
