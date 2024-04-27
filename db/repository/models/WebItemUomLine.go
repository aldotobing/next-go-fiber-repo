package models

// WebItemUomLine ...
type WebItemUomLine struct {
	ID                *string `json:"item_uom_line_id"`
	ItemID            *string `json:"item_id"`
	ItemCode          *string `json:"item_code"`
	ItemName          *string `json:"item_name"`
	ItemCategoryId    *string `json:"item_category_id"`
	ItemCategoryName  *string `json:"item_category_name"`
	ItemUomID         *string `json:"item_uom_id"`
	ItemUomName       *string `json:"item_uom_name"`
	ItemUomConversion *string `json:"item_uom_conversion"`
	Visibility        string  `json:"visibility"`
}

// WebItemUomLineParameter ...
type WebItemUomLineParameter struct {
	ID       string `json:"item_uom_line_id"`
	Code     string `json:"item_code"`
	Name     string `json:"item_name"`
	ItemID   string `json:"item_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
	ExceptId string `json:"except_id"`
}

var (
	// WebItemUomLineOrderBy ...
	WebItemUomLineOrderBy = []string{"def.id", "def._name", "def.created_date", "def.conversion"}
	// WebItemUomLineOrderByrByString ...
	WebItemUomLineOrderByrByString = []string{
		"def._name",
	}

	// WebItemUomLineSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--WEBITEMUOMLINE CONVERSION > 1 (HIGHEST UOM)
	*/
	WebItemUomLineSelectStatement = `
	SELECT def.id as item__uom_line_id, i.id as item_id, u.id as uom_id, ic.id as category_id,
	i.code as item_code,i._name as item_name,ic._name as category_name,
	u._name as uom_name, def.conversion as uom_conv,
	def.visibility
	FROM ITEM_UOM_LINE def
	left join item i on i.id = def.item_id
	left join uom u on u.id = def.uom_id
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
	`

	// WebItemUomLineWhereStatement ...
	WebItemUomLineWhereStatement = ` WHERE def.created_date IS NOT NULL `
)
