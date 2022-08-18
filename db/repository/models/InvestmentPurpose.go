package models

// InvestmentPurpose ...
type InvestmentPurpose struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

// InvestmentPurposeParameter ...
type InvestmentPurposeParameter struct {
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
	// InvestmentPurposeOrderBy ...
	InvestmentPurposeOrderBy = []string{"def.id", "def.name", "def.mapping_name", "def.created_at", "def.updated_at"}
	// InvestmentPurposeOrderByrByString ...
	InvestmentPurposeOrderByrByString = []string{
		"def.name", "def.mapping_name",
	}

	// InvestmentPurposeSelectStatement ...
	InvestmentPurposeSelectStatement = `SELECT def.id, def.name, def.mapping_name, def.status, def.created_at, def.updated_at, def.deleted_at
	FROM investment_purposes def`

	// InvestmentPurposeWhereStatement ...
	InvestmentPurposeWhereStatement = `WHERE def.deleted_at IS NULL`
)
