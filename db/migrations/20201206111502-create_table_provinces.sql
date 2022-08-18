-- +migrate Up
CREATE TABLE IF NOT EXISTS "provinces" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	country_id uuid NOT NULL REFERENCES countries(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 50),
	name TEXT NOT NULL CHECK (char_length(name) <= 50),
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE 
);

-- +migrate Down
DROP TABLE IF EXISTS "files";