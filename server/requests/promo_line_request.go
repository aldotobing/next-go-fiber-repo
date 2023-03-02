package requests

// PromoLineRequest ...
type PromoLineRequest struct {
	ID              string `json:"id"`
	PromoID         string `json:"promo_id"`
	GlobalMaxQty    string `json:"global_max_qty"`
	CustomerMaxQty  string `json:"customer_max_qty"`
	DiscPercent     string `json:"disc_percent"`
	DiscAmount      string `json:"disc_amount"`
	MinimumValue    string `json:"minimum_value"`
	MaximumValue    string `json:"maximum_value"`
	Multiply        string `json:"multiply"`
	Description     string `json:"description"`
	MinimumQty      string `json:"minimum_qty"`
	MaximumQty      string `json:"maximum_qty"`
	MinimumQtyUomID string `json:"minimum_qty_uom_id"`
	PromoType       string `json:"promo_type"`
	Strata          string `json:"strata"`
}
