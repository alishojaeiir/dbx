package dbx

// Database represents a database connection.
type Database struct {
	DB      DBExecutor
	Dialect string
}

func (d *Database) Executor() *QueryExecutor {
	return &QueryExecutor{DB: d.DB}
}
