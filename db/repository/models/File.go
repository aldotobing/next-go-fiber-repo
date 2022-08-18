package models

import "database/sql"

// File ....
type File struct {
	ID         string         `db:"id"`
	Type       sql.NullString `db:"type"`
	URL        sql.NullString `db:"url"`
	UserUpload sql.NullString `db:"user_upload"`
	IsUsed     sql.NullBool   `db:"is_used"`
	CreatedAt  string         `db:"created_at"`
	UpdatedAt  string         `db:"updated_at"`
	DeletedAt  sql.NullString `db:"deleted_at"`
}

var (
	// FileTypeDefault ...
	FileTypeDefault = "default"

	// FileWhitelist ...
	FileWhitelist = []string{FileTypeDefault}

	// FileMultipleUploadWhitelist ...
	FileMultipleUploadWhitelist = []string{}

	// FileSelectString ...
	FileSelectString = `SELECT f."id", f."type", f."url", f."user_upload", f."created_at", f."updated_at",
	f."deleted_at" FROM "files" f`

	// UnassignedQueryString ...
	UnassignedQueryString = ` `
)
