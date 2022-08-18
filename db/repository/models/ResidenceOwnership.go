package models

// ResidenceOwnership ...
type ResidenceOwnership struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

// ResidenceOwnershipParameter ...
type ResidenceOwnershipParameter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MappingName string `json:"mapping_name"`
	Status      string `json:"status"`
	Search      string `json:"search"`
	Page        int    `json:"page"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	By          string `json:"by"`
	Sort        string `json:"sort"`
}

var (
	// ResidenceOwnershipOrderBy ...
	ResidenceOwnershipOrderBy = []string{"def.id", "def.name", "def.mapping_name", "def.created_at", "def.updated_at"}
	// ResidenceOwnershipOrderByrByString ...
	ResidenceOwnershipOrderByrByString = []string{
		"def.name", "def.mapping_name",
	}

	// ResidenceOwnershipSelectStatement ...
	ResidenceOwnershipSelectStatement = `SELECT def.id, def.name, def.mapping_name, def.status, def.created_at, def.updated_at, def.deleted_at
	FROM residence_ownerships def`

	// ResidenceOwnershipWhereStatement ...
	ResidenceOwnershipWhereStatement = `WHERE def.deleted_at IS NULL`
)
