package models

// ItemSearchDetails ...
type ItemSearch struct {
	ID               *string `json:"item_id"`
	Code             *string `json:"item_code"`
	Name             *string `json:"item_name"`
	Description      *string `json:"item_description"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemPicture      *string `json:"item_picture"`
}

// ItemSearchParameter ...
type ItemSearchParameter struct {
	ID               string `json:"item_id"`
	Code             string `json:"item_code"`
	Name             string `json:"item_name"`
	ItemCategoryId   string `json:"item_category_id"`
	ItemCategoryName string `json:"item_category_name"`
	Search           string `json:"search"`
	Page             int    `json:"page"`
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	By               string `json:"by"`
	Sort             string `json:"sort"`
}

var (
	// ItemSearchOrderBy ...
	ItemSearchOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// ItemSearchOrderByrByString ...
	ItemSearchOrderByrByString = []string{
		"def._name",
	}

	// ItemSearchSelectStatement ...
	/*
		--SEARCH ITEM BERDASAR NAMA ATAU CATEGORY
	*/
	ItemSearchSelectStatement = `
	SELECT DEF.ID AS DEF_ID,
		DEF.CODE AS DEF_CODE,
		DEF._NAME AS DEF_NAME,
		DEF.DESCRIPTION as DEF_DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		DEF.ITEM_PICTURE AS ITEM_PICTURE
	FROM ITEM DEF 
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	`

	// ItemSearchWhereStatement ...
	ItemSearchWhereStatement = ` WHERE def.created_date IS not NULL `
)
