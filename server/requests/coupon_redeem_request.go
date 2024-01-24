package requests

// CouponRedeemRequest ...
type CouponRedeemRequest struct {
	CustomerID string `json:"customer_id" validate:"required"`
	CouponID   string `json:"coupon_id" validate:"required"`
}
