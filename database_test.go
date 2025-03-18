package dbx_test

import (
	"context"
	"github.com/alishojaeiir/dbx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	config := dbx.Config{
		Driver: "sqlite3",
		DBName: ":memory:",
	}
	db, err := dbx.Connect(config)
	assert.NoError(t, err)
	assert.NotNil(t, db)
	defer db.Executor().Close()
}

func TestExecContext(t *testing.T) {
	config := dbx.Config{Driver: "sqlite3", DBName: ":memory:"}
	db, _ := dbx.Connect(config)
	defer db.Executor().Close()

	_, err := db.Executor().ExecContext(context.Background(), "CREATE TABLE test (id INT)")
	assert.NoError(t, err)
}
