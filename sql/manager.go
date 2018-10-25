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

// GetAll payments
func GetAll(db *gorm.DB) *[]models.Payment {
	var pp []models.Payment
	db.Find(&pp)
	return &pp
}

// Delete a payment by id
func Delete(db *gorm.DB, pID string) error {
	result := db.Delete(models.Payment{}, "payments.payment = ?", pID)
	if result.Error != nil {
		return errors.New("Error deleting record")
	}
	return nil
}

// Create a payment resource
func Create(db *gorm.DB, p *models.Payment) error {
	result := db.Create(p)
	if result.Error != nil {
		return errors.New("Error creating record")
	}
	return nil
}
