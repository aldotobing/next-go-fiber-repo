package viewmodel

// PointPromoVM ....
type PointPromoVM struct {
	ID                 string               `json:"id"`
	StartDate          string               `json:"start_date"`
	EndDate            string               `json:"end_date"`
	CreatedAt          string               `json:"created_at"`
	UpdatedAt          string               `json:"updated_at"`
	DeletedAt          string               `json:"deleted_at"`
	Multiplicator      bool                 `json:"multiplicator"`
	PointConversion    string               `json:"poin_conversion"`
	QuantityConversion string               `json:"quantity_conversion"`
	PromoType          string               `json:"promo_type"`
	Strata             []PointPromoStrataVM `json:"strata"`
	Items              []PointPromoItemVM   `json:"items"`
	Image              string               `json:"image"`
	Title              string               `json:"title"`
	Description        string               `json:"description"`
}

// PointPromoStrataVM ...
type PointPromoStrataVM struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Point string `json:"point"`
}
