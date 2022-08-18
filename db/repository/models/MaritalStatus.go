package models

// MaritalStatus ...
type MaritalStatus struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	MappingName    string  `json:"mapping_name"`
	FillSpouseName bool    `json:"fill_spouse_name"`
	Status         bool    `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      *string `json:"deleted_at"`
}

// MaritalStatusParameter ...
type MaritalStatusParameter struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	MappingName    string `json:"mapping_name"`
	FillSpouseName string `json:"fill_spouse_name"`
	Status         string `json:"status"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// MaritalStatusOrderBy ...
	MaritalStatusOrderBy = []string{"def.id", "def.name", "def.mapping_name", "def.created_at", "def.updated_at"}
	// MaritalStatusOrderByrByString ...
	MaritalStatusOrderByrByString = []string{
		"def.name", "def.mapping_name",
	}

	// MaritalStatusSelectStatement ...
	MaritalStatusSelectStatement = `SELECT def.id, def.name, def.mapping_name, def.fill_spouse_name, def.status, def.created_at, def.updated_at, def.deleted_at
	FROM marital_statuses def`

	// MaritalStatusWhereStatement ...
	MaritalStatusWhereStatement = `WHERE def.deleted_at IS NULL`
)
