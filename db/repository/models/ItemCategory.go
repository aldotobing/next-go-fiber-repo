package models

// ItemCategory ...
type ItemCategory struct {
	ID   *string `json:"id"`
	Code *string `json:"code"`
	Name *string `json:"name"`
}

// ItemCategoryParameter ...
type ItemCategoryParameter struct {
	ID     string `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// ItemCategoryOrderBy ...
	ItemCategoryOrderBy = []string{"def.id", "def._name", "def.created_date", "def.sequenc"}
	// ItemCategoryOrderByrByString ...
	ItemCategoryOrderByrByString = []string{
		"def._name",
	}

	// ItemCategorySelectStatement ...
	ItemCategorySelectStatement = `SELECT def.id, def.code, def._name
	FROM item_category def
	`

	// ItemCategoryWhereStatement ...
	ItemCategoryWhereStatement = ` WHERE def.created_date IS not NULL `
)
