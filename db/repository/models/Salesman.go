package models

// Salesman ...
type Salesman struct {
	ID   *string `json:"salesman_id"`
	Name *string `json:"salesman_name"`
}

// SalesmanParameter ...
type SalesmanParameter struct {
	ID       string `json:"salesman_id"`
	UserID   string `json:"user_id"`
	RegionID string `json:"region_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
}

var (
	// SalesmanOrderBy ...
	SalesmanOrderBy = []string{"def.id", "p._name", "def.created_date"}
	// SalesmanOrderByrByString ...
	SalesmanOrderByrByString = []string{
		"p._name",
	}

	// SalesmanSelectStatement ...

	SalesmanSelectStatement = `
	select def.id,p._name
	from salesman def
	join partner p on p.id = def.partner_id
		`

	// SalesmanWhereStatement ...
	SalesmanWhereStatement = ` WHERE def.created_date IS not NULL `
)
