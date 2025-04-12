package dbx

// Database represents a database connection with a specific dialect.
type Database struct {
	DB      DBExecutor
	Dialect string
}

// Executor returns a QueryExecutor for executing database queries.
func (d *Database) Executor() *QueryExecutor {
	return &QueryExecutor{DB: d.DB}
}
