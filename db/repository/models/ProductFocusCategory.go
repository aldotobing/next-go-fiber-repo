package models

// ProductFocusCategory ...
type ProductFocusCategory struct {
	ID   *string `json:"id"`
	Code *string `json:"code"`
	Name *string `json:"name"`
	Foto *string `json:"foto"`
}

// ProductFocusCategoryParameter ...
type ProductFocusCategoryParameter struct {
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
	// ProductFocusCategoryOrderBy ...
	ProductFocusCategoryOrderBy = []string{"ic.id", "ic._name", "def.created_date"}
	// ProductFocusCategoryOrderByrByString ...
	ProductFocusCategoryOrderByrByString = []string{
		"ic.id",
	}

	// ProductFocusCategorySelectStatement ...
	ProductFocusCategorySelectStatement = `SELECT DISTINCT IC.ID AS IC_ID, IC.CODE AS IC_CODE,
	IC._NAME AS IC_NAME,
	null AS PICTURE
FROM PRODUCT_FOCUS DEF
JOIN ITEM I ON I.ID = DEF.ITEM_ID
JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID`

	// ProductFocusCategoryWhereStatement ...
	ProductFocusCategoryWhereStatement = ` WHERE def.created_date IS not NULL `
)
