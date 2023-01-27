package models

// WebItem ...
type WebItem struct {
	ID               *string `json:"item_id"`
	Code             *string `json:"item_code"`
	Name             *string `json:"item_name"`
	ItemPicture      *string `json:"item_picture"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemActive       *string `json:"item_active"`
	Description      *string `json:"item_description"`
	ItemHide         *string `json:"item_hide"`
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
		def.id as item_id, 
		ic.id as category_id,
		def.code as item_code,
		def._name as item_name,
		(concat('` + ItemImagePath + `',def.item_picture)) AS item_picture,
		ic._name as category_name,
		def.hide as item_hide,
		def.active as item_active,
		def.description as description 
	FROM ITEM def
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	`

	// WebItemWhereStatement ...
	WebItemWhereStatement = ` WHERE def.created_date IS NOT NULL `
)
