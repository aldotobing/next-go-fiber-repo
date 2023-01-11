package requests

// WebPromoBonusItemLineRequest ...
type WebPromoBonusItemLineRequest struct {
	PromoLineID string `json:"promo_line_id"`
	ItemID      string `json:"item_id"`
	UomID       string `json:"uom_id"`
	Qty         string `json:"qty"`
}
