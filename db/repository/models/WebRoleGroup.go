package models

// WebRoleGroup ...
type WebRoleGroup struct {
	ID         *string                 `json:"id"`
	Name       *string                 `json:"name"`
	RoleListID *string                 `json:"role_list_id"`
	RoleList   *[]WebRoleGroupRoleLine `json:"role_list"`
}

// WebRoleGroupParameter ...
type WebRoleGroupParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WebRoleGroupOrderBy ...
	WebRoleGroupOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// WebRoleGroupOrderByrByString ...
	WebRoleGroupOrderByrByString = []string{
		"def._name",
	}

	// WebRoleGroupSelectStatement ...
	WebRoleGroupSelectStatement = `SELECT def.id, def._name
	FROM role_group def
	`

	// WebRoleGroupWhereStatement ...
	WebRoleGroupWhereStatement = `WHERE def._name IS not NULL and def.deleted_at is null and def.is_mysm = 1 `
)
