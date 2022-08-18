CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SET TIMEZONE = 'Etc/GMT-7';

CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

create table marital_statuses (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	fill_spouse_name BOOLEAN NOT NULL,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER marital_statuses BEFORE UPDATE ON marital_statuses FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO marital_statuses ("id", "name", "mapping_name", "fill_spouse_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', 't', now());

create table religions (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER religions BEFORE UPDATE ON religions FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO religions ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table education_levels (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER education_levels BEFORE UPDATE ON education_levels FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO education_levels ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table residence_ownerships (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER residence_ownerships BEFORE UPDATE ON residence_ownerships FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO residence_ownerships ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table occupations (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER occupations BEFORE UPDATE ON occupations FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO occupations ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table line_of_businesses (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER line_of_businesses BEFORE UPDATE ON line_of_businesses FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO line_of_businesses ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table incomes (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	min_value NUMERIC(20, 3) NOT NULL,
	max_value NUMERIC(20, 3) NOT NULL,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER incomes BEFORE UPDATE ON incomes FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO incomes ("id", "name", "mapping_name", "min_value", "max_value", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 0.000, 0.000, 't', now());

create table investment_purposes (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER investment_purposes BEFORE UPDATE ON investment_purposes FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO investment_purposes ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table genders (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER genders BEFORE UPDATE ON genders FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO genders ("id", "name", "mapping_name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table provinces (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	code TEXT NOT NULL CHECK (char_length(code) <= 128),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER provinces BEFORE UPDATE ON provinces FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO provinces ("id", "code", "name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table cities (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	province_id uuid NOT NULL REFERENCES provinces(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 128),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	mapping_name TEXT CHECK (char_length(mapping_name) <= 100),
	type INT NOT NULL,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER cities BEFORE UPDATE ON cities FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO cities ("id", "province_id", "code", "name", "mapping_name", "type", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', '', '', 0, 't', now());

create table districts (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	city_id uuid NOT NULL REFERENCES cities(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 128),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER districts BEFORE UPDATE ON districts FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO districts ("id", "city_id", "code", "name", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', '', 't', now());

create table sub_districts (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	district_id uuid NOT NULL REFERENCES districts(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 128),
	name TEXT NOT NULL CHECK (char_length(name) <= 128),
	postal_code TEXT NOT NULL CHECK (char_length(postal_code) <= 32),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER sub_districts BEFORE UPDATE ON sub_districts FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO sub_districts ("id", "district_id", "code", "name", "postal_code", "status", "deleted_at")
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', '', '', 't', now());

create table account_openings (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	user_id uuid NOT NULL REFERENCES users(id),
	email TEXT NOT NULL CHECK (char_length(email) <= 128) DEFAULT '',
	email_valid_at TEXT NOT NULL CHECK (char_length(email_valid_at) <= 64) DEFAULT '',
	name TEXT NOT NULL CHECK (char_length(name) <= 128) DEFAULT '',
	marital_status_id uuid NOT NULL REFERENCES marital_statuses(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	gender_id uuid NOT NULL REFERENCES genders(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	birth_place TEXT NOT NULL CHECK (char_length(birth_place) <= 128) DEFAULT '',
	birth_place_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	birth_date TEXT NOT NULL CHECK (char_length(birth_date) <= 128) DEFAULT '',
	mother_name TEXT NOT NULL CHECK (char_length(mother_name) <= 128) DEFAULT '',
	phone TEXT NOT NULL CHECK (char_length(phone) <= 128) DEFAULT '',
	phone_valid_at TEXT NOT NULL CHECK (char_length(phone_valid_at) <= 64) DEFAULT '',
	other_phone TEXT NOT NULL CHECK (char_length(other_phone) <= 128) DEFAULT '',
	citizenship TEXT NOT NULL CHECK (char_length(citizenship) <= 128) DEFAULT '',
	religion_id uuid NOT NULL REFERENCES religions(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	education_level_id uuid NOT NULL REFERENCES education_levels(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	residence_ownership_id uuid NOT NULL REFERENCES residence_ownerships(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	residence_since TEXT NOT NULL CHECK (char_length(residence_since) <= 128) DEFAULT '',
	npwp_number TEXT NOT NULL CHECK (char_length(npwp_number) <= 128) DEFAULT '',
	reason_no_npwp TEXT NOT NULL CHECK (char_length(reason_no_npwp) <= 512) DEFAULT '',
	fax_number TEXT NOT NULL CHECK (char_length(fax_number) <= 128) DEFAULT '',
	mailing_address_type INTEGER NOT NULL CHECK (char_length(mailing_address_type) <= 16) DEFAULT 0,
	-- identity_type_id uuid NOT NULL REFERENCES identities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	identity_number TEXT NOT NULL CHECK (char_length(identity_number) <= 64) DEFAULT '',
	-- identity_image_id uuid NOT NULL REFERENCES files(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	identity_published_at TEXT NOT NULL CHECK (char_length(identity_published_at) <= 64) DEFAULT '',
	identity_expired_at TEXT NOT NULL CHECK (char_length(identity_expired_at) <= 64) DEFAULT '',
	referral_id uuid NOT NULL REFERENCES users(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	address TEXT NOT NULL CHECK (char_length(address) <= 512) DEFAULT '',
	city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	address_different_with_identity BOOLEAN NOT NULL DEFAULT 'false',
	domicile_address TEXT NOT NULL CHECK (char_length(domicile_address) <= 512) DEFAULT '',
	domicile_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	bank_name TEXT NOT NULL CHECK (char_length(bank_name) <= 128) DEFAULT '',
	bank_branch TEXT NOT NULL CHECK (char_length(bank_branch) <= 128) DEFAULT '',
	account_name TEXT NOT NULL CHECK (char_length(account_name) <= 128) DEFAULT '',
	account_number TEXT NOT NULL CHECK (char_length(account_number) <= 128) DEFAULT '',
	-- bank_rdn uuid NOT NULL REFERENCES banks(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	is_agree BOOLEAN NOT NULL DEFAULT 'true',
	company_name TEXT NOT NULL CHECK (char_length(company_name) <= 128) DEFAULT '',
	company_address TEXT NOT NULL CHECK (char_length(company_address) <= 512) DEFAULT '',
	company_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	company_phone TEXT NOT NULL CHECK (char_length(company_phone) <= 128) DEFAULT '',
	occupation_id uuid NOT NULL REFERENCES occupations(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	position TEXT NOT NULL CHECK (char_length(position) <= 128) DEFAULT '',
	line_of_business_id uuid NOT NULL REFERENCES line_of_businesses(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	length_of_work INTEGER NOT NULL CHECK (char_length(length_of_work) <= 16) DEFAULT 0,
	income_id uuid NOT NULL REFERENCES incomes(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	other_income TEXT NOT NULL CHECK (char_length(other_income) <= 128) DEFAULT '',
	-- source_of_income_id uuid NOT NULL REFERENCES source_of_incomes(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	investment_purpose_id uuid NOT NULL REFERENCES investment_purposes(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	-- asset_id uuid NOT NULL REFERENCES assets(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	total_asset TEXT NOT NULL CHECK (char_length(total_asset) <= 128) DEFAULT '',
	company_email TEXT NOT NULL CHECK (char_length(company_email) <= 128) DEFAULT '',
	company_fax_number TEXT NOT NULL CHECK (char_length(company_fax_number) <= 128) DEFAULT '',
	emergency_name TEXT NOT NULL CHECK (char_length(emergency_name) <= 128) DEFAULT '',
	emergency_relation TEXT NOT NULL CHECK (char_length(emergency_relation) <= 128) DEFAULT '',
	emergency_phone TEXT NOT NULL CHECK (char_length(emergency_phone) <= 128) DEFAULT '',
	emergency_other_phone TEXT NOT NULL CHECK (char_length(emergency_other_phone) <= 128) DEFAULT '',
	emergency_address TEXT NOT NULL CHECK (char_length(emergency_address) <= 512) DEFAULT '',
	emergency_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_name TEXT NOT NULL CHECK (char_length(pr_name) <= 128) DEFAULT '',
	pr_birth_place TEXT NOT NULL CHECK (char_length(pr_birth_place) <= 128) DEFAULT '',
	pr_birth_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_birth_date TEXT NOT NULL CHECK (char_length(pr_birth_date) <= 128) DEFAULT '',
	pr_phone TEXT NOT NULL CHECK (char_length(pr_phone) <= 128) DEFAULT '',
	pr_other_phone TEXT NOT NULL CHECK (char_length(pr_other_phone) <= 128) DEFAULT '',
	-- pr_identity_type_id uuid NOT NULL REFERENCES identities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_identity_number TEXT NOT NULL CHECK (char_length(pr_identity_number) <= 128) DEFAULT '',
	-- pr_identity_image_id uuid NOT NULL REFERENCES files(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_identity_expired_at TEXT NOT NULL CHECK (char_length(pr_identity_expired_at) <= 128) DEFAULT '',
	pr_citizenship TEXT NOT NULL CHECK (char_length(pr_citizenship) <= 128) DEFAULT '',
	pr_gender_id uuid NOT NULL REFERENCES genders(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_address_different_with_identity BOOLEAN NOT NULL DEFAULT 'false',
	pr_address TEXT NOT NULL CHECK (char_length(pr_address) <= 512) DEFAULT '',
	pr_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_occupation_id uuid NOT NULL REFERENCES occupations(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_company_name TEXT NOT NULL CHECK (char_length(pr_company_name) <= 128) DEFAULT '',
	pr_line_of_business_id uuid NOT NULL REFERENCES line_of_businesses(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_position TEXT NOT NULL CHECK (char_length(pr_position) <= 128) DEFAULT '',
	pr_incomes_id uuid NOT NULL REFERENCES incomes(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_company_phone TEXT NOT NULL CHECK (char_length(pr_company_phone) <= 128) DEFAULT '',
	pr_company_fax TEXT NOT NULL CHECK (char_length(pr_company_fax) <= 128) DEFAULT '',
	pr_length_of_work INTEGER NOT NULL CHECK (char_length(pr_length_of_work) <= 16) DEFAULT 0,
	-- pr_source_of_income_id uuid NOT NULL REFERENCES source_of_incomes(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	pr_company_address TEXT NOT NULL CHECK (char_length(pr_company_address) <= 512) DEFAULT '',
	pr_company_city_id uuid NOT NULL REFERENCES cities(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	-- selfie_image_id uuid NOT NULL REFERENCES files(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	-- selfie_image_id uuid NOT NULL REFERENCES files(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	-- selfie_video_id uuid NOT NULL REFERENCES files(id) DEFAULT '00000000-0000-0000-0000-000000000000',
	status TEXT NOT NULL CHECK (char_length(status) <= 128),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
