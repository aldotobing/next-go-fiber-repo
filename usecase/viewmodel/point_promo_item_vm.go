package viewmodel

// PointPromoItemVM ....
type PointPromoItemVM struct {
	ID         string `json:"id"`
	ItemName   string `json:"item_name"`
	UomID      string `json:"uom_id"`
	UomName    string `json:"uom_name"`
	Image      string `json:"image"`
	Quantity   string `json:"quantity"`
	Convertion string `json:"convertion"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
