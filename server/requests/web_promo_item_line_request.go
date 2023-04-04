package requests

// WebPromoItemLineRequest ...
type WebPromoItemLineRequest struct {
	PromoLineID string                               `json:"promo_line_id" validate:"required"`
	Items       []WebPromoItemLineItemDetailsRequest `json:"items"`
}

// WebPromoItemLineItemDetailsRequest ...
type WebPromoItemLineItemDetailsRequest struct {
	ItemID string `json:"item_id" validate:"required"`
	UomID  string `json:"uom_id" validate:"required"`
	Qty    string `json:"qty" validate:"required"`
}

// WebPromoItemLineRequest ...
type WebPromoItemLineAddByCategoryRequest struct {
	PromoLineID string `json:"promo_line_id" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
	Qty         string `json:"qty" validate:"required"`
}
