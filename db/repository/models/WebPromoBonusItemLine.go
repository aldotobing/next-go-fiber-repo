package models

// WebPromoBonusItemLinePromo ...
type WebPromoBonusItemLine struct {
	ID          *string `json:"id"`
	ItemID      *string `json:"item_id"`
	PromoID     *string `json:"promo_id"`
	PromoLineID *string `json:"promo_line_id"`
	PromoName   *string `json:"promo_name"`
	ItemCode    *string `json:"item_code"`
	ItemName    *string `json:"item_name"`
	Qty         *string `json:"item_qty"`
	UomID       *string `json:"uom_id"`
	UomName     *string `json:"uom_name"`
}

// WebPromoBonusItemLine ...
type WebPromoBonusItemLineBreakDown struct {
	ID          *string `json:"id"`
	PromoLineID *string `json:"promo_line_id"`
	ItemID      *string `json:"item_id"`
	UomID       *string `json:"uom_id"`
	Qty         *string `json:"item_qty"`
}

// WebPromoBonusItemLineParameter ...
type WebPromoBonusItemLineParameter struct {
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
	// WebPromoBonusItemLineOrderBy ...
	WebPromoBonusItemLineOrderBy = []string{"pil.id", "i._name", "i.created_date"}
	// WebPromoBonusItemLineOrderByrByString ...
	WebPromoBonusItemLineOrderByrByString = []string{
		"i._name",
	}

	// WebPromoBonusItemLineSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--WebPromoBonusItemLine CONVERSION > 1 (HIGHEST UOM)
	*/
	WebPromoBonusItemLineSelectStatement = `
	SELECT 
		PIL.ID AS ID,
		I.ID AS ITEM_ID,
		PR.ID AS PROMO_ID,
		PL.ID AS PROMO_LINE_ID,
		PR._NAME AS PROMO_NAME,
		I.CODE AS ITEM_CODE,
		I._NAME AS ITEM_NAME,
		PIL.QTY AS ITEM_LINE_QTY,
		PIL.UOM_ID AS UOM_ID,
		UOM._NAME AS UOM_NAME
		
	FROM PROMO_BONUS_LINE PIL
	left join ITEM I on i.id = pil.item_id
	LEFT JOIN UOM UOM ON UOM.ID = PIL.UOM_ID
	LEFT JOIN PROMO_LINE PL ON PL.ID = PIL.PROMO_LINE_ID
	LEFT JOIN PROMO PR ON PR.ID = PL.PROMO_ID 
	`

	// WebPromoBonusItemLineWhereStatement ...
	WebPromoBonusItemLineWhereStatement = ` WHERE i.created_date IS not NULL `
)
