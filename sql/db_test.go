package sql

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	Seed(db)
	var c int
	err = db.Get(&c, "SELECT count(*) FROM payments")
	assert.Nil(t, err)
	assert.Equal(t, 2, c)
}
