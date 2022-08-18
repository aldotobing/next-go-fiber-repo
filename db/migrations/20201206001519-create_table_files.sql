-- +migrate Up
CREATE TABLE IF NOT EXISTS "files" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	type TEXT NOT NULL CHECK (char_length(type) <= 50),
	url TEXT NOT NULL CHECK (char_length(url) <= 255),
	user_upload TEXT NOT NULL CHECK (char_length(user_upload) <= 100),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE 
);

-- +migrate Down
DROP TABLE IF EXISTS "files";