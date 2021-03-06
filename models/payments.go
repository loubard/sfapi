package models

// FetchResponse holds one payment
type FetchResponse struct {
	Data *Payment `json:"data"`
}

// ListResponse holds one payment
type ListResponse struct {
	Data *[]Payment `json:"data"`
}

// BeneficiaryParty holds information about amount and currency for receiver
type BeneficiaryParty struct {
	ID                int    `json:"-"`
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// SenderCharge holds information about amount and currency for sender
type SenderCharge struct {
	ID       int    `json:"-"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// ChargesInformation holding information about receiver and sender
type ChargesInformation struct {
	ID                      int             `json:"-"`
	BearerCode              string          `json:"bearer_code"`
	ReceiverChargesAmount   string          `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string          `json:"receiver_charges_currency"`
	SenderCharges           []*SenderCharge `json:"sender_charges"`
}

// DebtorParty account and bank information
type DebtorParty struct {
	ID                int    `json:"-"`
	AccountName       string `json:"account_name" db:"account_name"`
	AccountNumber     string `json:"account_number" db:"account_number"`
	AccountNumberCode string `json:"account_number_code" db:"account_number_code"`
	Address           string `json:"address" db:"address"`
	BankID            string `json:"bank_id" db:"bank_id"`
	BankIDCode        string `json:"bank_id_code" db:"bank_id_code"`
	Name              string `json:"name" db:"name"`
}

// Fx information
type Fx struct {
	ID                int    `json:"-"`
	ContractReference string `json:"contract_reference" db:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate" db:"exchange_rate"`
	OriginalAmount    string `json:"original_amount" db:"original_amount"`
	OriginalCurrency  string `json:"original_currency" db:"original_currency"`
}

// SponsorParty information
type SponsorParty struct {
	ID            int    `json:"-"`
	AccountNumber string `json:"account_number" db:"account_number"`
	BankID        string `json:"bank_id" db:"bank_id"`
	BankIDCode    string `json:"bank_id_code" db:"bank_id_code"`
}

// Attributes holds the detail about the payment
type Attributes struct {
	ID                   int    `json:"-"`
	Amount               string `json:"amount"`
	Currency             string `json:"currency"`
	EndToEndReference    string `json:"end_to_end_reference"`
	NumericReference     string `json:"numeric_reference"`
	Payment              string `json:"payment_id"`
	PaymentPurpose       string `json:"payment_purpose"`
	PaymentScheme        string `json:"payment_scheme"`
	PaymentType          string `json:"payment_type"`
	ProcessingDate       string `json:"processing_date"`
	Reference            string `json:"reference"`
	SchemePaymentSubType string `json:"scheme_payment_sub_type"`
	SchemePaymentType    string `json:"scheme_payment_type"`

	BeneficiaryParty     *BeneficiaryParty   `json:"beneficiary_party"`
	BeneficiaryPartyID   uint                `json:"-"`
	ChargesInformation   *ChargesInformation `json:"charges_information"`
	ChargesInformationID uint                `json:"-"`
	DebtorParty          *DebtorParty        `json:"debtor_party"`
	DebtorPartyID        uint                `json:"-"`
	Fx                   *Fx                 `json:"fx"`
	FxID                 uint                `json:"-"`
	SponsorParty         *SponsorParty       `json:"sponsor_party"`
	SponsorPartyID       uint                `json:"-"`
}

// Payment metadata
type Payment struct {
	ID           int    `json:"-"`
	Organisation string `json:"organisation_id"`
	Type         string `json:"type"`
	Version      int    `json:"version"`
	Payment      string `json:"payment_id"`

	Attributes   *Attributes `json:"attributes"`
	AttributesID uint        `json:"-"`
}
