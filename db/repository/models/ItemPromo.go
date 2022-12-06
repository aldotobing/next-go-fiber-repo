package models

// ItemPromoPromo ...
type ItemPromo struct {
	PromoID            *string `json:"promo_id"`
	PromoLineID        *string `json:"promo_line_id"`
	ItemID             *string `json:"item_id"`
	ItemCode           *string `json:"item_code"`
	ItemName           *string `json:"item_name"`
	ItemDescription    *string `json:"item_description"`
	ItemCategoryID     *string `json:"item_category_id"`
	ItemCategoryName   *string `json:"item_category_name"`
	ItemPicture        *string `json:"item_picture"`
	UomID              *string `json:"uom_id"`
	UomName            *string `json:"uom_name"`
	UomLineConversion  *string `json:"uom_line_conversion"`
	ItemPrice          *string `json:"item_price"`
	PriceListVersionID *string `json:"price_list_version_id"`
	DiscPercent        *string `json:"disc_percent"`
	DiscAmount         *string `json:"disc_amount"`
	MinValue           *string `json:"min_value"`
	MinQty             *string `json:"min_qty"`
	CustMaxQty         *string `json:"cust_max_qty"`
	GlobalMaxQty       *string `json:"global_max_qty"`
	Description        *string `json:"description"`
	StartDate          *string `json:"start_date"`
	EndDate            *string `json:"end_date"`
}

// ItemPromoParameter ...
type ItemPromoParameter struct {
	PromoID   string `json:"promo_id"`
	ItemID    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
	ExceptId  string `json:"except_id"`
}

var (
	// ItemPromoOrderBy ...
	ItemPromoOrderBy = []string{"pil.item_id", "i._name", "i.created_date"}
	// ItemPromoOrderByrByString ...
	ItemPromoOrderByrByString = []string{
		"i._name",
	}

	// ItemPromoSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ItemPromo CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemPromoSelectStatement = `
	SELECT 
		PR.ID AS PROMO_ID,
		PRL.ID AS PROMO_LINE_ID,
		PIL.ITEM_ID AS ITEM_ID,
		I._NAME AS ITEM_NAME,
		I.CODE AS ITEM_CODE,
		I.DESCRIPTION AS I_DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		I.ITEM_PICTURE AS ITEM_PICTURE,
		U.ID AS UOM_ID,
		U._NAME AS UOM_NAME,
		IUL.CONVERSION AS IUL_CONVERSION,
		IP.PRICE AS ITEM_PRICE,
		IP.PRICE_LIST_VERSION_ID AS PRICE_LIST_VERSION_ID,
		PRL.DISC_PCT AS DISC_PERCENT,
		PRL.DISC_AMT AS DISC_AMOUNT,
		PRL.MINIMUM_VALUE AS MINIMUM_VALUE,
		PRL.MINIMUM_QTY AS MINIMUM_QTY,
		PRL.CUSTOMER_MAX_QTY AS CUS_MAX_QTY,
		PRL.GLOBAL_MAX_QTY AS GLOBAL_MAX_QTY,
		PRL.DESCRIPTION AS DESCRIPTION,
		PR.START_DATE AS START_DATE,
		PR.END_DATE AS END_DATE
	FROM PROMO_ITEM_LINE PIL
	LEFT JOIN PROMO_LINE PRL ON PRL.ID = PIL.PROMO_LINE_ID	
	LEFT JOIN PROMO PR ON PR.ID = PRL.PROMO_ID
	LEFT JOIN ITEM I ON I.ID = PIL.ITEM_ID
	LEFT JOIN UOM U ON U.ID = PIL.UOM_ID
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
	LEFT JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = I.ID
	JOIN ITEM_PRICE IP ON IP.UOM_ID = U.ID
	AND IP.ITEM_ID = IUL.ITEM_ID
	`

	// ItemPromoWhereStatement ...
	ItemPromoWhereStatement = ` WHERE i.created_date IS not NULL `
)
