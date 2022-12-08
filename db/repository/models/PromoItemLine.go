package models

// PromoItemLinePromo ...
type PromoItemLine struct {
	ID                 *string `json:"id"`
	ItemID             *string `json:"item_id"`
	UomLineConversion  *string `json:"uom_line_conversion"`
	PromoID            *string `json:"promo_id"`
	PromoLineID        *string `json:"promo_line_id"`
	PromoName          *string `json:"promo_name"`
	ItemCode           *string `json:"item_code"`
	ItemName           *string `json:"item_name"`
	ItemDescription    *string `json:"item_description"`
	ItemCategoryID     *string `json:"item_category_id"`
	ItemCategoryName   *string `json:"item_category_name"`
	ItemPicture        *string `json:"item_picture"`
	Qty                *string `json:"item_qty"`
	UomID              *string `json:"uom_id"`
	UomName            *string `json:"uom_name"`
	ItemPrice          *string `json:"item_price"`
	PriceListVersionID *string `json:"price_list_version_id"`
	GlobalMaxQty       *string `json:"global_max_qty"`
	CustomerMaxQty     *string `json:"customer_max_qty"`
	DiscPercent        *string `json:"disc_percent"`
	DiscAmount         *string `json:"disc_amount"`
	MinValue           *string `json:"min_value"`
	MinQty             *string `json:"min_qty"`
	Description        *string `json:"description"`
	Multiply           *string `json:"multiply"`
	MinQtyUomID        *string `json:"min_qty_uom_id"`
	PromoType          *string `json:"promo_type"`
	Strata             *string `json:"strata"`
	StartDate          *string `json:"start_date"`
	EndDate            *string `json:"end_date"`
}

// PromoItemLine ...
type PromoItemLineBreakDown struct {
	ID          *string `json:"id"`
	PromoLineID *string `json:"promo_line_id"`
	ItemID      *string `json:"item_id"`
	UomID       *string `json:"uom_id"`
	Qty         *string `json:"item_qty"`
}

// PromoItemLineParameter ...
type PromoItemLineParameter struct {
	ID          string `json:"id"`
	PromoID     string `json:"promo_id"`
	PromoLineID string `json:"promo_line_id"`
	ItemID      string `json:"item_id"`
	CustomerID  string `json:"customer_id"`
	UomID       string `json:"uom_id"`
	ItemName    string `json:"item_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Search      string `json:"search"`
	Page        int    `json:"page"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	By          string `json:"by"`
	Sort        string `json:"sort"`
	ExceptId    string `json:"except_id"`
}

var (
	// PromoItemLineOrderBy ...
	PromoItemLineOrderBy = []string{"pil.id", "i._name", "i.created_date"}
	// PromoItemLineOrderByrByString ...
	PromoItemLineOrderByrByString = []string{
		"i._name",
	}

	// PromoItemLineSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--PromoItemLine CONVERSION > 1 (HIGHEST UOM)
	*/
	PromoItemLineSelectStatement = `
	SELECT 
		PIL.ID AS ID,
		I.ID AS ITEM_ID,
		IUL.CONVERSION AS IUL_CONVERSION,
		PR.ID AS PROMO_ID,
		PL.ID AS PROMO_LINE_ID,
		PR._NAME AS PROMO_NAME,
		I.CODE AS ITEM_CODE,
		I._NAME AS ITEM_NAME,
		I.DESCRIPTION AS ITEM_DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		I.ITEM_PICTURE AS ITEM_PICTURE,
		PIL.QTY AS ITEM_LINE_QTY,
		PIL.UOM_ID AS UOM_ID,
		UOM._NAME AS UOM_NAME,
		IP.PRICE AS ITEM_PRICE,
		IP.PRICE_LIST_VERSION_ID AS PRICE_LIST_VERSION_ID,
		PL.GLOBAL_MAX_QTY AS GLOBAL_MAX_QTY,
		PL.CUSTOMER_MAX_QTY AS CUSTOMER_MAX_QTY,
		PL.DISC_PCT AS DISC_PERCENT,
		PL.DISC_AMT AS DISC_AMOUNT,
		PL.MINIMUM_VALUE AS MINIMUM_VALUE,
		PL.MINIMUM_QTY AS PL_MIN_QTY,
		PL.DESCRIPTION AS PL_DESC,
		PL.MULTIPLY AS MULTIPLY,
		PL.MINIMUM_QTY_UOM_ID AS PL_MIN_QTY_UOM_ID,
		PL.PROMO_TYPE AS PL_PROMO_TYPE,
		PL.STRATA AS PL_STRATA,
		PR.START_DATE AS START_DATE,
		PR.END_DATE AS END_DATE
	FROM ITEM I 
	LEFT JOIN PROMO_ITEM_LINE PIL ON PIL.ITEM_ID = I.ID 
	LEFT JOIN UOM UOM ON UOM.ID = PIL.UOM_ID
	LEFT JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
	JOIN ITEM_PRICE IP ON IP.ITEM_ID = PIL.ITEM_ID
	LEFT JOIN PROMO_LINE PL ON PL.ID = PIL.PROMO_LINE_ID
	LEFT JOIN PROMO PR ON PR.ID = PL.PROMO_ID 
	LEFT JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = PIL.ITEM_ID AND IUL.UOM_ID = PIL.UOM_ID
	`

	// PromoItemLineWhereStatement ...
	PromoItemLineWhereStatement = ` WHERE i.created_date IS not NULL `
)
