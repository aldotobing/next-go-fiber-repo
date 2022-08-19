package models

// Item ...
type Item struct {
	ID   *string `json:"id_Item"`
	Code *string `json:"code"`
	Name *string `json:"name_Item"`
}

// ItemParameter ...
type ItemParameter struct {
	ID             string `json:"item_id"`
	Code           string `json:"item_code"`
	Name           string `json:"item_name"`
	ItemCategoryId string `json:"item_category_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// ItemOrderBy ...
	ItemOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// ItemOrderByrByString ...
	ItemOrderByrByString = []string{
		"def._name",
	}

	// ItemSelectStatement ...
	ItemSelectStatement = `SELECT def.id, def.code, def._name
	FROM Item def
	`

	// ItemWhereStatement ...
	ItemWhereStatement = ` WHERE def.created_date IS not NULL `
)
