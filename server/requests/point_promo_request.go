package requests

// PointPromoRequest ...
type PointPromoRequest struct {
	ID                 string                    `json:"id"`
	StartDate          string                    `json:"start_date"`
	EndDate            string                    `json:"end_date"`
	CreatedAt          string                    `json:"created_at"`
	UpdatedAt          string                    `json:"updated_at"`
	DeletedAt          string                    `json:"deleted_at"`
	Multiplicator      bool                      `json:"multiplicator"`
	PointConversion    string                    `json:"poin_conversion"`
	QuantityConversion string                    `json:"quantity_conversion"`
	PromoType          string                    `json:"promo_type"`
	Strata             []PointPromoStrataRequest `json:"strata"`
	Items              []string                  `json:"items"`
}

type PointPromoStrataRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Point string `json:"point"`
}
