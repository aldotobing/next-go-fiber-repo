package models

// ItemProductFocus ...
type ItemProductFocus struct {
	ID                 *string `json:"item_id"`
	Code               *string `json:"item_code"`
	Name               *string `json:"item_name"`
	Description        *string `json:"item_description"`
	ItemCategoryId     *string `json:"item_category_id"`
	ItemCategoryName   *string `json:"item_category_name"`
	ItemPicture        *string `json:"item_picture"`
	UomID              *string `json:"uom_id"`
	UomName            *string `json:"uom_name"`
	UomLineConversion  *string `json:"uom_line_conversion"`
	ItemPrice          *string `json:"item_price"`
	PriceListVersionId *string `json:"price_list_version_id"`
}

// ItemProductFocusParameter ...
type ItemProductFocusParameter struct {
	ID                 string `json:"item_id"`
	Code               string `json:"item_code"`
	Name               string `json:"item_name"`
	ItemCategoryId     string `json:"item_category_id"`
	PriceListVersionId string `json:"price_list_version_id"`
	CustomerTypeId     string `json:"customer_type_id"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
}

var (
	// ItemProductFocusOrderBy ...
	ItemProductFocusOrderBy = []string{"def.item_id", "i._name", "def.created_date"}
	// ItemProductFocusOrderByrByString ...
	ItemProductFocusOrderByrByString = []string{
		"i._name",
	}

	// ItemProductFocusSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ItemProductFocus CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemProductFocusSelectStatement = `
	SELECT DISTINCT DEF.ITEM_ID,
	I.CODE,I._NAME AS ITEM_NAME,
	I.DESCRIPTION,
	IC.ID AS I_CATEGORY_ID,
	IC._NAME AS I_CATEGORY_NAME,
	I.ITEM_PICTURE AS ITEM_PICTURE,
	UOM.ID AS UOMID,
	UOM._NAME AS UOMNAME,
	IUL.CONVERSION AS UOMLINECONVERSION,
	IP.PRICE AS ITEMPRICE,
	IP.PRICE_LIST_VERSION_ID AS PRICELISTVERSIONID
FROM PRODUCT_FOCUS DEF
JOIN ITEM I ON I.ID = DEF.ITEM_ID
JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ITEM_ID
JOIN UOM UOM ON UOM.ID = IUL.UOM_ID
JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID
AND IP.ITEM_ID = IUL.ITEM_ID`

	// ItemProductFocusWhereStatement ...
	ItemProductFocusWhereStatement = ` WHERE def.created_date IS not NULL AND I.HIDE = 0 `
)
