package sql

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/loubard/sfapi/models"
)

// GetByID a payment by a payment id
func GetByID(db *gorm.DB, pID string) (*models.Payment, error) {
	p := &models.Payment{
		Attributes: &models.Attributes{
			BeneficiaryParty: &models.BeneficiaryParty{},
			ChargesInformation: &models.ChargesInformation{
				SenderCharges: []*models.SenderCharge{},
			},
			DebtorParty:  &models.DebtorParty{},
			Fx:           &models.Fx{},
			SponsorParty: &models.SponsorParty{},
		},
	}

	result := db.Joins("INNER JOIN attributes ON attributes.id=payments.id").
		Joins("INNER JOIN beneficiary_parties ON beneficiary_parties.id=attributes.beneficiary_party_id").
		Where("payments.payment = ?", pID).
		First(p)

	if result.RecordNotFound() {
		return p, errors.New("Record not found")
	}
	return p, nil
}
