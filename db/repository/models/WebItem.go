package models

// WebItem ...
type WebItem struct {
	ID               *string `json:"item_id"`
	Code             *string `json:"item_code"`
	Name             *string `json:"item_name"`
	ItemPicture      *string `json:"item_picture"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemHide         *string `json:"item_hide"`
	ItemActive       *string `json:"item_active"`
	ItemDescription  *string `json:"item_description"`
}

// WebItemSelectByCategory ...
type WebItemSelectByCategory struct {
	ID               *string `json:"item_id"`
	Code             *string `json:"item_code"`
	Name             *string `json:"item_name"`
	ItemPicture      *string `json:"item_picture"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemDescription  *string `json:"item_description"`
	UOMDetail        *string `json:"oum_detail"`
}

// WebItemParameter ...
type WebItemParameter struct {
	ID             string `json:"item_id"`
	Code           string `json:"item_code"`
	Name           string `json:"item_name"`
	ItemCategoryId string `json:"item_category_id"`
	ItemHide       string `json:"item_hide"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
	ExceptId       string `json:"except_id"`
}

var (
	// WebItemOrderBy ...
	WebItemOrderBy = []string{"def.id", "def._name", "def.created_date", "iul.conversion"}
	// WebItemOrderByrByString ...
	WebItemOrderByrByString = []string{
		"def._name",
	}

	// WebItemSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--WEBITEM CONVERSION > 1 (HIGHEST UOM)
	*/
	WebItemSelectStatement = `
	SELECT 
		DEF.ID AS item_id, 
		IC.ID AS category_id,
		DEF.CODE AS item_code,
		DEF._NAME AS item_name,
		(concat('` + ItemImagePath + `',def.item_picture)) AS item_picture,
		IC._NAME AS category_name,
		DEF.HIDE AS item_hide,
		DEf.ACTIVE AS item_active,
		DEF.DESCRIPTION AS description 
	FROM ITEM DEF
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	`

	// WebItemWhereStatement ...
	WebItemWhereStatement = ` WHERE def.created_date IS NOT NULL `
)
