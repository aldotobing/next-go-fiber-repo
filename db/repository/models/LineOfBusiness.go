package models

// LineOfBusiness ...
type LineOfBusiness struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

// LineOfBusinessParameter ...
type LineOfBusinessParameter struct {
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
	// LineOfBusinessOrderBy ...
	LineOfBusinessOrderBy = []string{"def.id", "def.name", "def.mapping_name", "def.created_at", "def.updated_at"}
	// LineOfBusinessOrderByrByString ...
	LineOfBusinessOrderByrByString = []string{
		"def.name", "def.mapping_name",
	}

	// LineOfBusinessSelectStatement ...
	LineOfBusinessSelectStatement = `SELECT def.id, def.name, def.mapping_name, def.status, def.created_at, def.updated_at, def.deleted_at
	FROM line_of_businesses def`

	// LineOfBusinessWhereStatement ...
	LineOfBusinessWhereStatement = `WHERE def.deleted_at IS NULL`
)
