-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_credentials"
(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "user_id" uuid NOT NULL REFERENCES users(id),
    "type" TEXT NOT NULL CHECK (char_length(type) <= 50),
    "email" TEXT NOT NULL CHECK (char_length(email) <= 100),
    "password" TEXT NOT NULL CHECK (char_length(password) <= 500),
    registration_details JSONB CHECK (char_length(registration_details::TEXT) <= 1000),
    "status" TEXT NOT NULL CHECK (char_length(status) <= 50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "user_credentials";