package models

// ItemDetailsDetails ...
type ItemDetails struct {
	ID                      *string `json:"item_id"`
	Code                    *string `json:"item_code"`
	Name                    *string `json:"item_name"`
	Description             *string `json:"item_description"`
	ItemDetailsCategoryId   *string `json:"item_category_id"`
	ItemDetailsCategoryName *string `json:"item_category_name"`
	ItemDetailsPicture      *string `json:"item_picture"`
	UomID                   *string `json:"uom_id"`
	UomName                 *string `json:"uom_name"`
	UomLineConversion       *string `json:"uom_line_conversion"`
	ItemDetailsPrice        *string `json:"item_price"`
	PriceListVersionId      *string `json:"price_list_verison_id"`
}

// ItemDetailsParameter ...
type ItemDetailsParameter struct {
	ID                    string `json:"item_id"`
	Code                  string `json:"item_code"`
	Name                  string `json:"item_name"`
	ItemDetailsCategoryId string `json:"item_category_id"`
	PriceListVersionId    string `json:"price_list_version_id"`
	Search                string `json:"search"`
	Page                  int    `json:"page"`
	Offset                int    `json:"offset"`
	Limit                 int    `json:"limit"`
	By                    string `json:"by"`
	Sort                  string `json:"sort"`
	ExceptId              string `json:"except_id"`
}

var (
	// ItemDetailsOrderBy ...
	ItemDetailsOrderBy = []string{"def.id", "def._name", "def.created_date", "iul.conversion"}
	// ItemDetailsOrderByrByString ...
	ItemDetailsOrderByrByString = []string{
		"def._name",
	}

	// ItemDetailsSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ItemDetails CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemDetailsSelectStatement = `
	SELECT DEF.ID AS DEF_ID,
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
		IP.PRICE_LIST_VERSION_ID AS PRICE_LIST_VERSION_ID
	FROM ITEM_UOM_LINE IUL
	LEFT JOIN ITEM DEF ON IUL.ITEM_ID = DEF.ID
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID
	LEFT JOIN UOM UOM ON UOM.ID = IUL.UOM_ID
	JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID AND IP.ITEM_ID = IUL.ITEM_ID
	`

	// ItemDetailsWhereStatement ...
	ItemDetailsWhereStatement = ` WHERE def.created_date IS not NULL AND IUL.CONVERSION > 1 `
)
