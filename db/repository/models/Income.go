package models

// Income ...
type Income struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	MinValue    float64 `json:"min_value"`
	MaxValue    float64 `json:"max_value"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

// IncomeParameter ...
type IncomeParameter struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	MinValue    float64 `json:"min_value"`
	MaxValue    float64 `json:"max_value"`
	Status      string  `json:"status"`
	Search      string  `json:"search"`
	Page        int     `json:"page"`
	Offset      int     `json:"offset"`
	Limit       int     `json:"limit"`
	By          string  `json:"by"`
	Sort        string  `json:"sort"`
}

var (
	// IncomeOrderBy ...
	IncomeOrderBy = []string{"def.id", "def.name", "def.mapping_name", "def.created_at", "def.updated_at"}
	// IncomeOrderByrByString ...
	IncomeOrderByrByString = []string{
		"def.name", "def.mapping_name",
	}

	// IncomeSelectStatement ...
	IncomeSelectStatement = `SELECT def.id, def.name, def.mapping_name, def.min_value, def.max_value, def.status, def.created_at, def.updated_at, def.deleted_at
	FROM incomes def`

	// IncomeWhereStatement ...
	IncomeWhereStatement = `WHERE def.deleted_at IS NULL`
)
