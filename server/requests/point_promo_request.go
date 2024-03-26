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
	Items              []PointPromoItemRequest   `json:"items"`
	Image              string                    `json:"image"`
	Description        string                    `json:"description"`
}

type PointPromoStrataRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Point string `json:"point"`
}

type PointPromoItemRequest struct {
	ItemID     string `json:"item_id"`
	UomID      string `json:"uom_id"`
	UomName    string `json:"uom_name"`
	Convertion string `json:"convertion"`
}
