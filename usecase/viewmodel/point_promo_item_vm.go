package viewmodel

// PointPromoItemVM ....
type PointPromoItemVM struct {
	ID         string `json:"id"`
	ItemName   string `json:"start_date"`
	UomID      string `json:"uom_id"`
	UomName    string `json:"uom_name"`
	Convertion string `json:"convertion"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
