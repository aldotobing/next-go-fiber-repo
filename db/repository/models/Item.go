package models

import "encoding/json"

// Item ...
type Item struct {
	ID               *string          `json:"item_id"`
	Code             *string          `json:"code"`
	Name             *string          `json:"item_name"`
	Description      *string          `json:"item_description"`
	ItemCategoryId   *string          `json:"item_category_id"`
	ItemCategoryName *string          `json:"item_category_name"`
	ItemPicture      *string          `json:"item_picture"`
	Uom              *json.RawMessage `json:"item_uom"`
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
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ITEM CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemSelectStatement = `
	SELECT DEF.ID,
		DEF.CODE,
		DEF._NAME,
		DEF.DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		DEF.ITEM_PICTURE AS ITEM_PICTURE,
			(SELECT JSON_AGG(T) AS ITEM_UOM
				FROM
					(SELECT UOM.ID::varchar(255) AS UOM_ID,
							UOM._NAME::varchar(255) AS UOM_NAME,
							IUL.CONVERSION::varchar(255) AS IUL_CONVERSION,
							IP.PRICE::varchar(255) AS ITEM_PRICE,
							IP.PRICE_LIST_VERSION_ID::varchar(255) AS PRICE_LIST_VERSION_ID
						FROM ITEM_UOM_LINE IUL
						JOIN ITEM I ON I.ID = IUL.ITEM_ID
						JOIN UOM UOM ON UOM.ID = IUL.UOM_ID
						JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID
						AND IP.ITEM_ID = IUL.ITEM_ID
						WHERE IUL.ITEM_ID = DEF.ID
							AND IUL.CONVERSION > 1 ) T)
	FROM ITEM DEF
	JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	`

	// ItemWhereStatement ...
	ItemWhereStatement = ` WHERE def.created_date IS not NULL `
)
