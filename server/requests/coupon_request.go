package requests

// CouponRequest ...
type CouponRequest struct {
	StartDate       string `json:"start_date" validate:"required"`
	EndDate         string `json:"end_date" validate:"required"`
	PointConversion string `json:"point_conversion" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
}
