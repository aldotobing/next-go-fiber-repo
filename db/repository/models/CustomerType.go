package models

// CustomerType ...
type CustomerType struct {
	ID   *string `json:"customer_type_id"`
	Code *string `json:"customer_type_code"`
	Name *string `json:"customer_type_name"`
}

// CustomerTypeParameter ...
type CustomerTypeParameter struct {
	ID       string `json:"customertype_id"`
	IDs      string `json:"customertype_ids"`
	UserID   string `json:"user_id"`
	RegionID string `json:"region_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
	ExceptId string `json:"except_id"`
}

var (
	// CustomerTypeOrderBy ...
	CustomerTypeOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// CustomerTypeOrderByrByString ...
	CustomerTypeOrderByrByString = []string{
		"def._name",
	}

	// CustomerTypeSelectStatement ...

	CustomerTypeSelectStatement = `
	select def.id,def.code,def._name 
	from customer_type def
		`

	// CustomerTypeWhereStatement ...
	CustomerTypeWhereStatement = ` WHERE def.created_date IS not NULL `
)
