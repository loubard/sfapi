package sql

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/loubard/sfapi/models"
	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	Seed(db)
	var c int
	db.Model(&models.Payment{}).Count(&c)
	assert.Equal(t, 1, c)
}
