package models

import "database/sql"

// Setting ...
type Setting struct {
	ID        string         `json:"id"`
	Code      sql.NullString `json:"code"`
	Details   sql.NullString `json:"details"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

var (
	// SettingOrderBy ...
	SettingOrderBy = []string{"def.id", "def.code", "def.created_at", "def.updated_at"}
	// SettingOrderByrByString ...
	SettingOrderByrByString = []string{
		"def.code",
	}

	// SettingSelectStatement ...
	SettingSelectStatement = `SELECT def.id, def.code, def.details, def.created_at, def.updated_at, def.deleted_at
	FROM settings def`

	// SettingWhereStatement ...
	SettingWhereStatement = `WHERE def.deleted_at IS NULL`
)
