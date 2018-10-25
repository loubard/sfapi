package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/loubard/sfapi/models"
)

// Seed create the schema and insert data
func Seed(db *gorm.DB) {
	// Create table
	db.AutoMigrate(
		&models.Payment{},
		&models.Attributes{},
		&models.BeneficiaryParty{},
		&models.ChargesInformation{},
		&models.SenderCharge{},
		&models.DebtorParty{},
		&models.Fx{},
		&models.SponsorParty{},
	)

	at := &models.Attributes{
		Amount:               "100.21",
		Currency:             "GBP",
		EndToEndReference:    "Wil piano Jan",
		NumericReference:     "1002001",
		PaymentID:            "123456789012345678",
		PaymentPurpose:       "Paying for goods/services",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",
	}
	p := &models.Payment{
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Type:           "Payment",
		PaymentID:      "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		Version:        0,

		Attributes: at,
	}
	db.Create(p)
}
