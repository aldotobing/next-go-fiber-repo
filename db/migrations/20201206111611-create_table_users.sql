-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "code" TEXT NOT NULL CHECK (char_length(code) <= 50),
    "email" TEXT NOT NULL CHECK (char_length(email) <= 100),
    "name" TEXT NOT NULL CHECK (char_length(name) <= 100),
    "profile_image_id" uuid REFERENCES files(id),
    "gender" TEXT NOT NULL CHECK (char_length(gender) <= 50),
    "phone" TEXT NOT NULL CHECK (char_length(phone) <= 100),
    "city_id" uuid REFERENCES cities(id),
    "address" TEXT NOT NULL CHECK (char_length(address) <= 500),
    "status" TEXT NOT NULL CHECK (char_length(status) <= 50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "users";