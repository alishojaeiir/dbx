package dbx_test

import (
	"context"
	"database/sql"
	"github.com/alishojaeiir/dbx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupSQLite(t *testing.T) *dbx.Database {
	config := dbx.Config{
		Driver:       "sqlite3",
		DBName:       ":memory:",
		MaxIdleConns: 1,
		MaxOpenConns: 1,
	}
	db, err := dbx.Connect(config)
	assert.NoError(t, err, "connection should succeed")
	assert.NotNil(t, db, "database instance should not be nil")

	result, err := db.Executor().ExecContext(context.Background(), "CREATE TABLE users (id INTEGER, name TEXT)")
	assert.NoError(t, err, "create table should succeed")
	assert.NotNil(t, result, "result should not be nil")

	return db
}

func TestConnect(t *testing.T) {
	config := dbx.Config{
		Driver:       "sqlite3",
		DBName:       ":memory:",
		MaxIdleConns: 1,
		MaxOpenConns: 1,
	}
	db, err := dbx.Connect(config)
	assert.NoError(t, err, "connection should succeed")
	assert.NotNil(t, db, "database instance should not be nil")
	defer db.Executor().Close()
}

func TestExecContext(t *testing.T) {
	db := setupSQLite(t)
	if db == nil {
		t.Fatal("database connection failed")
	}
	defer db.Executor().Close()

	result, err := db.Executor().ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 1, "Alice")
	assert.NoError(t, err, "insert should succeed")
	assert.NotNil(t, result, "result should not be nil")

	rowsAffected, err := result.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowsAffected, "one row should be affected")
}

func TestQueryRowContext(t *testing.T) {
	db := setupSQLite(t)
	if db == nil {
		t.Fatal("database connection failed")
	}
	defer db.Executor().Close()

	_, err := db.Executor().ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 1, "Alice")
	assert.NoError(t, err, "insert should succeed")

	row, err := db.Executor().QueryRowContext(context.Background(), "SELECT name FROM users WHERE id = ?", 1)
	assert.NoError(t, err, "query should succeed")
	assert.NotNil(t, row, "row should not be nil")

	var name string
	err = row.Scan(&name)
	assert.NoError(t, err, "scan should succeed")
	assert.Equal(t, "Alice", name, "name should match inserted value")
}

func TestQueryContext(t *testing.T) {
	db := setupSQLite(t)
	if db == nil {
		t.Fatal("database connection failed")
	}
	defer db.Executor().Close()

	_, err := db.Executor().ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 1, "Alice")
	assert.NoError(t, err)
	_, err = db.Executor().ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 2, "Bob")
	assert.NoError(t, err)

	rows, err := db.Executor().QueryContext(context.Background(), "SELECT id, name FROM users")
	assert.NoError(t, err)
	defer rows.Close()

	var users []struct {
		ID   int
		Name string
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		assert.NoError(t, err)
		users = append(users, struct {
			ID   int
			Name string
		}{ID: id, Name: name})
	}
	assert.NoError(t, rows.Err(), "rows should have no errors")
	assert.Equal(t, 2, len(users), "should return two users")
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "Bob", users[1].Name)
}

func TestBeginTx(t *testing.T) {
	db := setupSQLite(t)
	if db == nil {
		t.Fatal("database connection failed")
	}
	defer db.Executor().Close()

	tx, err := db.Executor().BeginTx(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, tx, "transaction should not be nil")

	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 1, "Alice")
	assert.NoError(t, err)

	err = tx.Commit()
	assert.NoError(t, err)

	row, err := db.Executor().QueryRowContext(context.Background(), "SELECT name FROM users WHERE id = ?", 1)
	assert.NoError(t, err)
	var name string
	err = row.Scan(&name)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", name)

	tx, err = db.Executor().BeginTx(context.Background(), nil)
	assert.NoError(t, err)
	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (id, name) VALUES (?, ?)", 2, "Bob")
	assert.NoError(t, err)
	err = tx.Rollback()
	assert.NoError(t, err)

	row, err = db.Executor().QueryRowContext(context.Background(), "SELECT name FROM users WHERE id = ?", 2)
	assert.NoError(t, err)
	err = row.Scan(&name)
	assert.Equal(t, sql.ErrNoRows, err, "should return no rows after rollback")
}
