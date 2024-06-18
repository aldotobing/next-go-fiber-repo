package requests

// CouponRedeemRequest ...
type CouponRedeemRequest struct {
	CustomerID string `json:"customer_id" validate:"required"`
	CouponID   string `json:"coupon_id" validate:"required"`
	Otp        string `json:"otp"`
}

type CouponRedeemOTPRequest struct {
	CustomerID   string            `json:"customer_id" validate:"required"`
	CouponRedeem []CouponRedeemOTP `json:"coupon_redeem" validate:"required"`
	Otp          string            `json:"otp"`
}

type CouponRedeemOTP struct {
	CouponID string `json:"coupon_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
