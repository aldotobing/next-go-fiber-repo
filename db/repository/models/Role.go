package models

// Role ...
type Role struct {
	ID        string  `json:"id_role"`
	Name      *string `json:"name_role"`
	CreatedAt *string `json:"created_at_role"`
	UpdatedAt *string `json:"updated_at_role"`
	DeletedAt *string `json:"deleted_at_role"`
	CreatedBy *string `json:"created_by_role"`
	UpdatedBy *string `json:"updated_by_role"`
	DeletedBy *string `json:"deleted_by_role"`
}

// RoleParameter ...
type RoleParameter struct {
	ID     string `json:"id_role"`
	Name   string `json:"name_role"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// RoleOrderBy ...
	RoleOrderBy = []string{"def.id_role", "def.name_role", "def.created_at_role", "def.updated_at_role"}
	// RoleOrderByrByString ...
	RoleOrderByrByString = []string{
		"def.name_role",
	}

	// RoleSelectStatement ...
	RoleSelectStatement = `SELECT def.id_role, def.name_role, def.created_at_role, def.updated_at_role, def.deleted_at_role
	FROM roles def`

	// RoleWhereStatement ...
	RoleWhereStatement = `WHERE def.deleted_at_role IS NULL`
)
