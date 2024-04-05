package viewmodel

// PointPromoItemVM ....
type PointPromoItemVM struct {
	ID         string `json:"item_id"`
	ItemName   string `json:"item_name"`
	UomID      string `json:"item_uom_id"`
	UomName    string `json:"item_uom_name"`
	Image      string `json:"image"`
	Quantity   string `json:"quantity"`
	Convertion string `json:"item_uom_conversion"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
