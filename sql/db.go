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
		&models.DebtorParty{},
		&models.Fx{},
		&models.SponsorParty{},
	)
	db.Create(buildPayment("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"))
	db.Create(buildPayment("7eb8277a-6c91-45e9-8a03-a27f82aca350"))
}

func buildPayment(pID string) *models.Payment {
	// Insert initial data
	sp := &models.SponsorParty{AccountNumber: "56781234", BankID: "123123", BankIDCode: "GBDSC"}
	fx := &models.Fx{
		ContractReference: "FX123",
		ExchangeRate:      "2.00000",
		OriginalAmount:    "200.42",
		OriginalCurrency:  "USD",
	}
	dp := &models.DebtorParty{
		AccountName:       "EJ Brown Black",
		AccountNumber:     "GB29XABC10161234567801",
		AccountNumberCode: "IBAN",
		Address:           "10 Debtor Crescent Sourcetown NE1",
		BankID:            "203301",
		BankIDCode:        "GBDSC",
		Name:              "Emelia Jane Brown",
	}
	sc1 := &models.SenderCharge{Amount: "5.00", Currency: "GBP"}
	sc2 := &models.SenderCharge{Amount: "10.00", Currency: "USD"}
	ci := &models.ChargesInformation{
		BearerCode:              "SHAR",
		ReceiverChargesAmount:   "1.00",
		ReceiverChargesCurrency: "USD",

		SenderCharges: []*models.SenderCharge{sc1, sc2},
	}
	bp := &models.BeneficiaryParty{
		AccountName:       "W Owens",
		AccountNumber:     "31926819",
		AccountNumberCode: "BBAN",
		AccountType:       0,
		Address:           "1 The Beneficiary Localtown SE2",
		BankID:            "403000",
		BankIDCode:        "GBDSC",
		Name:              "Wilfred Jeremiah Owens",
	}
	at := &models.Attributes{
		Amount:               "100.21",
		Currency:             "GBP",
		EndToEndReference:    "Wil piano Jan",
		NumericReference:     "1002001",
		Payment:              "123456789012345678",
		PaymentPurpose:       "Paying for goods/services",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",

		BeneficiaryParty:   bp,
		ChargesInformation: ci,
		DebtorParty:        dp,
		Fx:                 fx,
		SponsorParty:       sp,
	}
	p := &models.Payment{
		Organisation: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Type:         "Payment",
		Payment:      pID,
		Version:      0,

		Attributes: at,
	}

	return p
}
