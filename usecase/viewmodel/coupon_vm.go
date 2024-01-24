package viewmodel

// CouponVM ....
type CouponVM struct {
	ID              string `json:"id"`
	StartDate       string `json:"start_dates"`
	EndDate         string `json:"end_date"`
	PointConversion string `json:"point_conversion"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
