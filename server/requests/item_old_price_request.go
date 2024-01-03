package requests

// ItemOldPriceRequest ...
type ItemOldPriceRequest struct {
	ItemCode     string `json:"item_code"`
	CustomerCode string `json:"customer_code"`
	StartDate    string `json:"start_date" validate:"required"`
	EndDate      string `json:"end_date" validate:"required"`
	Quantity     int    `json:"qty"`
}

// ItemOldPriceBulkRequest ...
type ItemOldPriceBulkRequest struct {
	OldPrice []ItemOldPriceRequest `json:"old_price"`
}
