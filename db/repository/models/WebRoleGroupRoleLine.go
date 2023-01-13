package models

// WebRoleGroupRoleLine ...
type WebRoleGroupRoleLine struct {
	ID   *string `json:"id"`
	Name *string `json:"name_webrolegrouproleline"`
}

// WebRoleGroupRoleLineParameter ...
type WebRoleGroupRoleLineParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WebRoleGroupRoleLineOrderBy ...
	WebRoleGroupRoleLineOrderBy = []string{"def.id", "rl._name", "def.created_date"}
	// WebRoleGroupRoleLineOrderByrByString ...
	WebRoleGroupRoleLineOrderByrByString = []string{
		"rl._name",
	}

	// WebRoleGroupRoleLineSelectStatement ...
	WebRoleGroupRoleLineSelectStatement = `
	select def.id as rg_rl_id, def.role_id as role_id, def.role_group_id as role_group_id,
	rl._name as role_name, rg._name as role_group_name
	from role_group_role_line def
	join role rl on rl.id = def.role_id
	join role_group rg on rg.id = def.role_group_id
	`

	// WebRoleGroupRoleLineWhereStatement ...
	WebRoleGroupRoleLineWhereStatement = `WHERE def._name IS not NULL and def.deleted_at is null `
)
