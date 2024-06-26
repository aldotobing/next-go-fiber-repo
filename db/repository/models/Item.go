package models

import "encoding/json"

// Item ...
type Item struct {
	ID                 *string          `json:"item_id"`
	UOMLineID          *string          `json:"item_uom_line_id"`
	Code               *string          `json:"item_code"`
	Name               *string          `json:"item_name"`
	Description        *string          `json:"item_description"`
	ItemCategoryId     *string          `json:"item_category_id"`
	ItemCategoryName   *string          `json:"item_category_name"`
	ItemPicture        *string          `json:"item_picture"`
	UomID              *string          `json:"uom_id"`
	UomName            *string          `json:"uom_name"`
	UomLineConversion  *string          `json:"uom_line_conversion"`
	ItemPrice          *string          `json:"item_price"`
	PriceListVersionId *string          `json:"price_list_version_id"`
	Uom                *json.RawMessage `json:"item_uom"`
}

// ItemV2 ...
type ItemV2 struct {
	ID               *string `json:"id"`
	Name             *string `json:"_name"`
	Code             *string `json:"item_code"`
	Description      *string `json:"item_description"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	AdditionalData   *string `json:"additional_data"`
	MultiplyData     *string `json:"multiply_data"`
	ItemPicture      *string `json:"item_picture"`
}

// ItemParameter ...
type ItemParameter struct {
	ID                 string `json:"item_id"`
	IDs                string `json:"item_ids"`
	Code               string `json:"item_code"`
	Name               string `json:"item_name"`
	ItemCategoryId     string `json:"item_category_id"`
	ItemCategoryName   string `json:"item_category_name"`
	PriceListVersionId string `json:"price_list_version_id"`
	PriceListId        string `json:"price_list_id"`
	UomID              string `json:"uom_id"`
	CustomerTypeId     string `json:"customer_type_id"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
	ExceptId           string `json:"except_id"`
	CustomerID         string `json:"customer_id"`
}

var (
	// ItemOrderBy ...
	ItemOrderBy = []string{"def.id", "def._name", "def.created_date", "iul.conversion"}
	// ItemOrderByrByString ...
	ItemOrderByrByString = []string{
		"def._name",
	}

	// ItemSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ITEM CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemSelectStatement2 = `
	SELECT IP.id as uomline_id,DEF.ID AS DEF_ID,
		DEF.CODE AS DEF_CODE,
		DEF._NAME AS DEF_NAME,
		DEF.DESCRIPTION as DEF_DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		DEF.ITEM_PICTURE AS ITEM_PICTURE,
		UOM.ID AS UOM_ID,
		UOM._NAME AS UOM_NAME,
		IUL.CONVERSION AS IUL_CONVERSION,
		IP.PRICE AS ITEM_PRICE,
		IP.PRICE_LIST_VERSION_ID AS PRICE_LIST_VERSION_ID,
			(SELECT JSON_AGG(T) 
				FROM
					(SELECT UOM.ID::varchar(255) AS UOM_ID,
							UOM._NAME::varchar(255) AS UOM_NAME,
							IUL.CONVERSION::varchar(255) AS IUL_CONVERSION,
							IP.PRICE::varchar(255) AS ITEM_PRICE,
							IP.PRICE_LIST_VERSION_ID::varchar(255) AS PRICE_LIST_VERSION_ID
						FROM ITEM_UOM_LINE IUL
						JOIN ITEM I ON I.ID = IUL.ITEM_ID
						JOIN UOM UOM ON UOM.ID = IUL.UOM_ID
						JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID AND IP.ITEM_ID = IUL.ITEM_ID
						WHERE IUL.ITEM_ID = DEF.ID AND IUL.VISIBILITY = 1
						ORDER BY IUL.CONVERSION) T) AS ITEM_UOM
	FROM ITEM_UOM_LINE IUL
	LEFT JOIN ITEM DEF ON IUL.ITEM_ID = DEF.ID
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	LEFT JOIN UOM UOM ON UOM.ID = IUL.UOM_ID
	JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID AND IP.ITEM_ID = IUL.ITEM_ID
	`

	ItemSelectStatement = ` 
	SELECT
		DEF.ID,DEF.CODE AS ITEM_CODE, 
		DEF._NAME,DEF.DESCRIPTION AS I_DESCRIPT,
		DEF.ITEM_CATEGORY_ID AS CAT_ID, 
		IC._NAME AS IC_NAME,
		U.ID AS UOM_ID, 
		U._NAME AS UOM_NAME, 
		IUL.CONVERSION AS KONVERSI, 
		(X.PRICE * IUL.CONVERSION) AS HARGA,
		X.PLV_ID AS PRICE_LIST_VERSION_ID, DEF.ITEM_PICTURE
	FROM ITEM DEF
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	JOIN ITEM_UOM_LINE IUL ON DEF.ID = IUL.ITEM_ID
	JOIN UOM U ON U.ID = IUL.UOM_ID
	JOIN 
	(SELECT IP.ITEM_ID AS I_ID, IULS.UOM_ID AS U_UOM,  IP.PRICE ,IP.PRICE_LIST_VERSION_ID AS PLV_ID
	 FROM ITEM_PRICE IP 
	 JOIN ITEM_UOM_LINE IULS ON IULS.ITEM_ID = IP.ITEM_ID AND IULS.UOM_ID = IP.UOM_ID
	 WHERE IULS.CONVERSION = 1 AND IP.PRICE_LIST_VERSION_ID = $1) X ON X.I_ID = DEF.ID  `

	// ItemWhereStatement ...
	ItemWhereStatement = ` WHERE def.created_date IS NOT NULL AND IUL.VISIBILITY = 1 AND DEF.ACTIVE = 1 AND DEF.HIDE = 0 `

	ItemV2SelectStatement = `with temp_data as(
		select DEF.ID,
		array_to_string((array_agg(U.ID || '#sep#' || u."_name" || '#sep#' || IUL."conversion" || '#sep#' || IUL."visibility" order by iul."conversion" asc)),'|') AS MULTIPLY_DATA
		from item def
		    left JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID 
			left JOIN UOM U ON U.ID = IUL.UOM_ID
	    WHERE def.created_date IS NOT NULL
		{{ALL_PARAM}}
		AND (LOWER(def."_name") LIKE LOWER($1))
		group by def.id 
		order by DEF.ID asc
	)   
	SELECT
		DEF.ID,DEF.CODE AS ITEM_CODE,
		DEF._NAME,
		DEF.DESCRIPTION AS I_DESCRIPT,
		DEF.ITEM_CATEGORY_ID AS CAT_IHalobroD,
		array_to_string((array_agg(distinct ic."_name")),'|') AS category_name,
		array_to_string((array_agg(U.ID || '#sep#' || u."_name" || '#sep#' || IUL.conversion::text || '#sep#' || ip.modified_date || '#sep#' || ip.price::text || '#sep#' || ip.price_list_version_id || '#sep#' || IUL.visibility order by ip.modified_date desc)),'|') AS additional_data,
		td.MULTIPLY_DATA,
		DEF.ITEM_PICTURE
	FROM ITEM DEF
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	left JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID
	left join item_price ip on ip.item_id = iul.item_id and ip.uom_id = iul.uom_id
	left JOIN UOM U ON U.ID = IP.UOM_ID
	left join TEMP_DATA TD on TD.ID = DEF.ID
	WHERE def.created_date IS NOT NULL
		{{ALL_PARAM}}
		AND (LOWER(def."_name") LIKE LOWER($1)) `
)
