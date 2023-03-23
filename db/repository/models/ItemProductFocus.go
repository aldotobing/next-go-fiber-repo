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

// ItemProductFocusV2 ...
type ItemProductFocusV2 struct {
	ID             *string `json:"item_id"`
	Code           *string `json:"item_code"`
	Name           *string `json:"item_name"`
	Description    *string `json:"item_description"`
	ItemCategory   *string `json:"item_category"`
	ItemPicture    *string `json:"item_picture"`
	AdditionalData *string `json:"additional_data"`
	MultiplyData   *string `json:"multiply_data"`
}

// ItemProductFocusParameter ...
type ItemProductFocusParameter struct {
	ID                 string `json:"item_id"`
	Code               string `json:"item_code"`
	Name               string `json:"item_name"`
	ItemCategoryId     string `json:"item_category_id"`
	PriceListVersionId string `json:"price_list_version_id"`
	CustomerTypeId     string `json:"customer_type_id"`
	CustomerID         string `json:"customer_id"`
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

	ItemProductFocusV2SelectStatement = `with temp_data as(
		select DEF.ID,
		array_to_string((array_agg(U.ID || '#sep#' || u."_name" || '#sep#' || IUL."conversion" || '#sep#' || IUL."visibility" order by iul."conversion" asc)),'|') AS MULTIPLY_DATA
		from item def
		    left JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID 
			left JOIN UOM U ON U.ID = IUL.UOM_ID
	    WHERE def.created_date IS NOT NULL
		AND DEF.ACTIVE = 1
		AND DEF.HIDE = 0
		AND (LOWER(def."_name") LIKE LOWER($1))
		group by def.id 
		order by DEF.ID asc
	)
	SELECT I.ID,
		I.CODE,
		I._NAME AS ITEM_NAME,
		I.DESCRIPTION,
		I.ITEM_PICTURE AS ITEM_PICTURE,
		array_to_string((array_agg(distinct ic.id || '#sep#' ||ic."_name")),'|') AS category,
		array_to_string((array_agg(U.ID || '#sep#' || u."_name" || '#sep#' || IUL.conversion::text || '#sep#' || ip.price::text || '#sep#' || ip.price_list_version_id || '#sep#' || IUL.visibility order by iul."conversion" asc)),'|') AS additional_data,
        td.MULTIPLY_DATA
	FROM PRODUCT_FOCUS DEF
	JOIN ITEM I ON I.ID = DEF.ITEM_ID
	JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
	JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = i.ID
	JOIN UOM U ON U.ID = IUL.UOM_ID
	JOIN ITEM_PRICE IP ON ip.item_id = iul.item_id and ip.uom_id = iul.uom_id  
	left join TEMP_DATA TD on TD.ID = I.ID`

	// ItemProductFocusWhereStatement ...
	ItemProductFocusV2WhereStatement = ` WHERE def.created_date IS not NULL AND I.HIDE = 0 `
)
