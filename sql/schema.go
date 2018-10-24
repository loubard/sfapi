package sql

var schema = `
CREATE TABLE debtor_parties (
	id INTEGER PRIMARY KEY,
	account_name VARCHAR(255),
	account_number VARCHAR(255),
	account_number_code VARCHAR(255),
	address VARCHAR(255),
	bank_id VARCHAR(255),
	bank_id_code VARCHAR(255),
	name VARCHAR(255)
);

CREATE TABLE fx(
	id INTEGER PRIMARY KEY,
	contract_reference VARCHAR(255),
	exchange_rate VARCHAR(255),
	original_amount VARCHAR(255),
	original_currency VARCHAR(255)
);

CREATE TABLE charges_information (
	id INTEGER PRIMARY KEY,
	bearer_code VARCHAR(255),
	receiver_charges_amount VARCHAR(255),
	receiver_charges_currency VARCHAR(255)
);

CREATE TABLE sender_charges (
	id INTEGER PRIMARY KEY,
	amount VARCHAR(255),
	currency VARCHAR(255),

	charges_information_id INTEGER,
	FOREIGN KEY(charges_information_id) REFERENCES charges_information(id)
);

CREATE TABLE beneficiary_parties (
	id INTEGER PRIMARY KEY,
	account_name VARCHAR(255),
	account_number VARCHAR(255),
	account_number_code VARCHAR(255),
	account_type INTEGER,
	address VARCHAR(255),
	bank_id VARCHAR(255),
	bank_id_code VARCHAR(255),
	name VARCHAR(255)
);

CREATE TABLE sponsor_parties(
	id INTEGER PRIMARY KEY,
	account_number VARCHAR(255),
	bank_id VARCHAR(255),
	bank_id_code VARCHAR(255)
);

CREATE TABLE attributes (
	id INTEGER PRIMARY KEY,
	amount VARCHAR(255),
	currency VARCHAR(255),
	end_to_end_reference VARCHAR(255),
	numeric_reference VARCHAR(255),
	payment_id VARCHAR(255),
	payment_purpose VARCHAR(255),
	payment_scheme VARCHAR(255),
	payment_type VARCHAR(255),
	processing_date VARCHAR(255),
	reference VARCHAR(255),
	scheme_payment_sub_type VARCHAR(255),
	scheme_payment_type VARCHAR(255),

	beneficiary_party_id INTEGER,
	charges_information_id INTEGER,
	debtor_party_id INTEGER,
	fx_id INTEGER,
	sponsor_party_id INTEGER,
	FOREIGN KEY(beneficiary_party_id) REFERENCES beneficiary_parties(id),
	FOREIGN KEY(charges_information_id) REFERENCES charges_information(id),
	FOREIGN KEY(debtor_party_id) REFERENCES debtor_parties(id),
	FOREIGN KEY(fx_id) REFERENCES fx(id),
	FOREIGN KEY(sponsor_party_id) REFERENCES sponsor_parties(id)
);

CREATE TABLE payments (
	id INTEGER PRIMARY KEY,
	payment_id VARCHAR(255),
	organisation_id VARCHAR(255),
	type VARCHAR(255),
	version INTEGER,

	attribute_id INTEGER,
	FOREIGN KEY(attribute_id) REFERENCES attributes(id)
);
`
