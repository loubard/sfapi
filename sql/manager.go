package sql

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/loubard/sfapi/models"
)

// GetByID a payment by a payment id
func GetByID(db *gorm.DB, pID string) (*models.Payment, error) {
	p := &models.Payment{}

	result := db.Where("payments.payment = ?", pID).First(p)

	if result.RecordNotFound() {
		return p, errors.New("Record not found")
	}
	return p, nil
}
