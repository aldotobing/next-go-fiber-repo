-- +migrate Up
CREATE TABLE IF NOT EXISTS "countries" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	country_code TEXT NOT NULL CHECK (char_length(country_code) <= 10),
  	name TEXT NOT NULL CHECK (char_length(name) <= 100),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE
);
INSERT INTO countries (country_code, name, status, created_at, updated_at) VALUES ('62', 'Indonesia', true, current_timestamp, current_timestamp);

-- +migrate Down
DROP TABLE IF EXISTS "countries";