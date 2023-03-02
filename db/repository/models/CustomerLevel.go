package models

// CustomerLevel ...
type CustomerLevel struct {
	ID           *string `json:"id"`
	Code         *string `json:"code"`
	Name         *string `json:"name"`
	CreatedDate  *string `json:"created_date"`
	ModifiedDate *string `json:"modified_date"`
}

// CustomerLevelParameter ...
type CustomerLevelParameter struct {
	Search string `json:"search"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// CustomerLevelOrderBy ...
	CustomerLevelOrderBy = []string{"def.id", "def.code", "def._name", "def.created_at", "def.modified_date"}
	// CustomerLevelOrderByrByStrings ...
	CustomerLevelOrderByrByString = []string{
		"def._name",
	}

	// CustomerLevelSelectStatement ...
	CustomerLevelSelectStatement = `
	SELECT 
		def.id,
		def.code,
		def._name, 
		def.created_date,
		def.modified_date
	FROM customer_level def
	`

	// CustomerLevelWhereStatement ...
	CustomerLevelWhereStatement = ` WHERE def.created_date IS NOT NULL `
)
