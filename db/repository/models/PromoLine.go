package models

// Promo Line ...
type PromoLine struct {
	ID              *string `json:"id"`
	PromoID         *string `json:"promo_id"` //dari table promo_content
	GlobalMaxQty    *string `json:"global_max_qty"`
	CustomerMaxQty  *string `json:"customer_max_qty"`
	DiscPercent     *string `json:"disc_percent"`
	DiscAmount      *string `json:"disc_amount"`
	MinimumValue    *string `json:"minimum_value"`
	MaximumValue    *string `json:"maximum_value"`
	Multiply        *string `json:"multiply"`
	Description     *string `json:"description"`
	MinimumQty      *string `json:"minimum_qty"`
	MaximumQty      *string `json:"maximum_qty"`
	MinimumQtyUomID *string `json:"minimum_qty_uom_id"`
	PromoType       *string `json:"promo_type"`
	Strata          *string `json:"strata"`
}

type PromoLineBreakdown struct {
	PromoID         *string `json:"promo_id"` //dari table promo_content
	GlobalMaxQty    *string `json:"global_max_qty"`
	CustomerMaxQty  *string `json:"customer_max_qty"`
	DiscPercent     *string `json:"disc_percent"`
	DiscAmount      *string `json:"disc_amount"`
	MinimumValue    *string `json:"minimum_value"`
	Multiply        *string `json:"multiply"`
	Description     *string `json:"description"`
	MinimumQty      *string `json:"minimum_qty"`
	MinimumQtyUomID *string `json:"minimum_qty_uom_id"`
	PromoType       *string `json:"promo_type"`
	Strata          *string `json:"strata"`
}

// PromoLineParameter ...
type PromoLineParameter struct {
	ID              string `json:"id"`
	PromoID         string `json:"promo_id"` //dari table promo_content
	GlobalMaxQty    string `json:"global_max_qty"`
	CustomerMaxQty  string `json:"customer_max_qty"`
	DiscPercent     string `json:"disc_percent"`
	DiscAmount      string `json:"disc_amount"`
	MinimumValue    string `json:"minimum_value"`
	Multiply        string `json:"multiply"`
	Description     string `json:"description"`
	MinimumQty      string `json:"minimum_qty"`
	MinimumQtyUomID string `json:"minimum_qty_uom_id"`
	PromoType       string `json:"promo_type"`
	Strata          string `json:"strata"`
	Search          string `json:"search"`
	Page            int    `json:"page"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	By              string `json:"by"`
	Sort            string `json:"sort"`
}

var (
	// PromoLineOrderBy ...
	PromoLineOrderBy = []string{"pl.id", "pl.created_date"}
	// PromoLineOrderByrByString ...
	PromoLineOrderByrByString = []string{
		"pl.id",
	}

	// PromoLineSelectStatement ...
	PromoLineSelectStatement = `
	SELECT 
		PL.ID AS ID, 
		PL.PROMO_ID AS PROMO_ID, 
		PL.GLOBAL_MAX_QTY AS GLOBAL_MAX_QTY,
		PL.CUSTOMER_MAX_QTY AS CUSTOMER_MAX_QTY,
		PL.DISC_PCT AS DISC_PERCENT,
		PL.DISC_AMT AS DISC_AMOUNT,
		PL.MINIMUM_VALUE AS MINIMUM_VALUE,
		PL.MAXIMUM_VALUE AS MAXIMUM_VALUE,
		PL.MULTIPLY AS MULTIPLY,
		PL.DESCRIPTION AS DESCRIPTION,
		PL.MINIMUM_QTY AS MINIMUM_QTY,
		PL.MAXIMUM_QTY AS MAXIMUM_QTY,
		PL.MINIMUM_QTY_UOM_ID AS MIN_QTY_UOM_ID,
		PL.PROMO_TYPE AS PROMO_TYPE,
		PL.STRATA AS STRATA
	FROM PROMO_LINE PL
	`
	// PromoLineWhereStatement ...
	PromoLineWhereStatement = ` 
	WHERE PL.ID IS NOT NULL
	`
)
