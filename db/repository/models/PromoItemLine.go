package models

// PromoItemLinePromo ...
type PromoItemLine struct {
	ID               *string `json:"id"`
	PromoID          *string `json:"promo_id"`
	PromoName        *string `json:"promo_name"`
	PromoLineID      *string `json:"promo_line_id"`
	ItemID           *string `json:"item_id"`
	ItemCode         *string `json:"item_code"`
	ItemName         *string `json:"item_name"`
	ItemDescription  *string `json:"item_description"`
	ItemCategoryID   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemPicture      *string `json:"item_picture"`
	Qty              *string `json:"item_qty"`
	UomID            *string `json:"uom_id"`
	UomName          *string `json:"uom_name"`
	ItemPrice        *string `json:"item_price"`
	DiscPercent      *string `json:"disc_percent"`
	DiscAmount       *string `json:"disc_amount"`
	MinValue         *string `json:"min_value"`
	MinQty           *string `json:"min_qty"`
	CustMaxQty       *string `json:"cust_max_qty"`
	GlobalMaxQty     *string `json:"global_max_qty"`
	Description      *string `json:"description"`
	StartDate        *string `json:"start_date"`
	EndDate          *string `json:"end_date"`
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
		PR.ID AS PROMO_ID,
		PR._NAME AS PROMO_NAME,
		PRL.ID AS PROMO_LINE_ID,
		PIL.ITEM_ID AS ITEM_ID,
		PIL.QTY AS ITEM_LINE_QTY,
		I._NAME AS ITEM_NAME,
		I.CODE AS ITEM_CODE,
		I.DESCRIPTION AS I_DESCRIPTION,
		IC.ID AS I_CATEGORY_ID,
		IC._NAME AS I_CATEGORY_NAME,
		I.ITEM_PICTURE AS ITEM_PICTURE,
		U.ID AS UOM_ID,
		U._NAME AS UOM_NAME,
		COALESCE (IP.PRICE, 0) AS ITEM_PRICE,
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
	LEFT JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = PIL.ITEM_ID AND IUL.UOM_ID = PIL.UOM_ID 
	`

	// PromoItemLineWhereStatement ...
	PromoItemLineWhereStatement = ` WHERE i.created_date IS not NULL `
)
