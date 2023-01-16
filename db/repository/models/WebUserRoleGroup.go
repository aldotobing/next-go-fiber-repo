package models

// WebUserRoleGroup ...
type WebUserRoleGroup struct {
	ID            *string `json:"id"`
	UserID        *string `json:"user_id"`
	RoleGroupID   *string `json:"role_group_id"`
	RoleGroupName *string `json:"role_group_name"`
}

// WebUserRoleGroupParameter ...
type WebUserRoleGroupParameter struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WebUserRoleGroupOrderBy ...
	WebUserRoleGroupOrderBy = []string{"def.id", "rg._name", "def.created_date"}
	// WebUserRoleGroupOrderByrByString ...
	WebUserRoleGroupOrderByrByString = []string{
		"rg._name",
	}

	// WebUserRoleGroupSelectStatement ...
	WebUserRoleGroupSelectStatement = `SELECT def.id, def.user_id, def.role_group_id,rg._name
	FROM user_role_group def
	join role_group rg on rg.id = def.role_group_id
	`

	// WebUserRoleGroupWhereStatement ...
	WebUserRoleGroupWhereStatement = `WHERE rg._name IS not NULL and rg.is_mysm = 1 `
)
