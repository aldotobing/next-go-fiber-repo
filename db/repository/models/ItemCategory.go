package models

// ItemCategory ...
type ItemCategory struct {
	ID    *string `json:"id"`
	Code  *string `json:"code"`
	Name  *string `json:"name"`
	Image *string `json:"image"`
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
	ItemCategorySelectStatement = `SELECT def.id, def.code, def._name, ICI.IMG 
	FROM item_category def
	LEFT JOIN ITEM_CATEGORY_IMG ICI ON ICI.CATEGORY_ID = DEF.ID
	`

	// ItemCategoryWhereStatement ...
	ItemCategoryWhereStatement = ` WHERE def.created_date IS not NULL AND def.active = 1 `
)
