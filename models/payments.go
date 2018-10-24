package models

// FetchResponse holds one payment
type FetchResponse struct {
	Data *Payment `json:"data"`
}

// BeneficiaryParty holds information about amount and currency for receiver
type BeneficiaryParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// SenderCharges holds information about amount and currency for sender
type SenderCharges struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// ChargesInformation holding information about receiver and sender
type ChargesInformation struct {
	BearerCode              string          `json:"bearer_code"`
	ReceiverChargesAmount   string          `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string          `json:"receiver_charges_currency"`
	SenderCharges           []SenderCharges `json:"sender_charges"`
}

// DebtorParty account and bank information
type DebtorParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// Fx information
type Fx struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// SponsorParty information
type SponsorParty struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
}

// Attributes holds the detail about the payment
type Attributes struct {
	Amount               string              `json:"amount"`
	BeneficiaryParty     *BeneficiaryParty   `json:"beneficiary_party"`
	ChargesInformation   *ChargesInformation `json:"charges_information"`
	Currency             string              `json:"currency"`
	DebtorParty          *DebtorParty        `json:"debtor_party"`
	EndToEndReference    string              `json:"end_to_end_reference"`
	Fx                   *Fx                 `json:"fx"`
	NumericReference     string              `json:"numeric_reference"`
	PaymentID            string              `json:"payment_id"`
	PaymentPurpose       string              `json:"payment_purpose"`
	PaymentScheme        string              `json:"payment_scheme"`
	PaymentType          string              `json:"payment_type"`
	ProcessingDate       string              `json:"processing_date"`
	Reference            string              `json:"reference"`
	SchemePaymentSubType string              `json:"scheme_payment_sub_type"`
	SchemePaymentType    string              `json:"scheme_payment_type"`
	SponsorParty         *SponsorParty       `json:"sponsor_party"`
}

// Payment metadata
type Payment struct {
	Attributes     *Attributes `json:"attributes"`
	ID             string      `json:"payment_id"`
	OrganisationID string      `json:"organisation_id"`
	Type           string      `json:"type"`
	Version        int         `json:"version"`
}
