package models

// SalesmanType ...
type SalesmanType struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// SalesmanTypeParameter ...
type SalesmanTypeParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// SalesmanTypeOrderBy ...
	SalesmanTypeOrderBy = []string{"def.created_at", "def.id", "def._name", "def.code"}
	// SalesmanTypeOrderByrByString ...
	SalesmanTypeOrderByrByString = []string{
		"def._name",
	}

	// SalesmanTypeSelectStatement ...
	SalesmanTypeSelectStatement = `
		select def.id, def.code, def._name
		from salesman_type def
	`

	// SalesmanTypeWhereStatement ...
	SalesmanTypeWhereStatement = ` WHERE def.created_date IS not NULL `
)
