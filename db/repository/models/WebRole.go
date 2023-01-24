package models

// WeebRole ...
type WeebRole struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Header *string `json:"header"`
}

// WeebRoleParameter ...
type WeebRoleParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WeebRoleOrderBy ...
	WeebRoleOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// WeebRoleOrderByrByString ...
	WeebRoleOrderByrByString = []string{
		"def._name",
	}

	// WeebRoleSelectStatement ...
	WeebRoleSelectStatement = `SELECT def.id, def._name, def._header
	FROM role def
	`

	// WeebRoleWhereStatement ...
	WeebRoleWhereStatement = `WHERE def._name IS not NULL and def.deleted_at is null and def.is_mysm = 1 `
)
