package models

// Uom ...
type Uom struct {
	ID   *string `json:"id"`
	Code *string `json:"uom_code"`
	Name *string `json:"uom_name"`
}

// UomParameter ...
type UomParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// UomOrderBy ...
	UomOrderBy = []string{"u.id", "u._name", "u.created_date"}
	// UomOrderByrByString ...
	UomOrderByrByString = []string{
		"u._name",
	}

	// UomSelectStatement ...

	UomSelectStatement = `
	SELECT 
	u.id AS uom_id, 
	u.code AS uom_code, 
	u._name AS uom_name 
	FROM uom u `

	// UomWhereStatement ...
	UomWhereStatement = ` WHERE u.created_date IS not NULL `
)
