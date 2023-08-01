package viewmodel

type PromoItemLineVM struct {
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
