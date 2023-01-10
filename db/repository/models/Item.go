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

// ItemParameter ...
type ItemParameter struct {
	ID                 string `json:"item_id"`
	Code               string `json:"item_code"`
	Name               string `json:"item_name"`
	ItemCategoryId     string `json:"item_category_id"`
	PriceListVersionId string `json:"price_list_version_id"`
	UomID              string `json:"uom_id"`
	CustomerTypeId     string `json:"customer_type_id"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
	ExceptId           string `json:"except_id"`
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

	ItemSelectStatement = ` select def.id,def.code as item_code, def._name,def.description as i_descript,
	def.item_category_id as cat_id, ic._name as ic_name,
	u.id as uom_id, u._name as uom_name, 
	iul.conversion as konversi, (x.price * iul.conversion) as harga,
	x.plv_id as price_list_version_id, def.item_picture
	from item def
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = def.ITEM_CATEGORY_ID
	join item_uom_line iul on def.id = iul.item_id
	join uom u on u.id = iul.uom_id
	join 
	( select ip.item_id as i_id, iuls.uom_id as u_uom,  ip.price ,ip.price_list_version_id as plv_id
	 from item_price ip 
	 join item_uom_line iuls on iuls.item_id = ip.item_id and iuls.uom_id = ip.uom_id
	 where iuls.conversion = 1 and ip.price_list_version_id = $1
	 )x on x.i_id = def.id  `

	// ItemWhereStatement ...
	ItemWhereStatement = ` WHERE def.created_date IS NOT NULL AND IUL.VISIBILITY = 1 AND DEF.ACTIVE = 1 AND DEF.HIDE = 0`
)
