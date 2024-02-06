package viewmodel

// CouponRedeemVM ....
type CouponRedeemVM struct {
	ID                    string `json:"id"`
	CouponID              string `json:"coupon_id"`
	CustomerID            string `json:"customer_id"`
	Redeem                string `json:"redeem"`
	RedeemAt              string `json:"redeem_at"`
	RedeemedToDocumentNo  string `json:"redeem_to_doc_no"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	DeletedAt             string `json:"deleted_at"`
	ExpiredAt             string `json:"expired_at"`
	CouponName            string `json:"coupon_name"`
	CouponDescription     string `json:"coupon_description"`
	CouponPointConversion string `json:"coupon_point_conversion"`
	CustomerName          string `json:"customer_name"`
}

type CouponRedeemReportVM struct {
	ID                    string `json:"id"`
	CouponID              string `json:"coupon_id"`
	CustomerID            string `json:"customer_id"`
	Redeem                string `json:"redeem"`
	RedeemAt              string `json:"redeem_at"`
	RedeemedToDocumentNo  string `json:"redeem_to_doc_no"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	DeletedAt             string `json:"deleted_at"`
	ExpiredAt             string `json:"expired_at"`
	CouponName            string `json:"coupon_name"`
	CouponDescription     string `json:"coupon_description"`
	CouponPointConversion string `json:"coupon_point_conversion"`
	CustomerName          string `json:"customer_name"`
	BranchName            string `json:"branch_name"`
	BranchCode            string `json:"branch_code"`
	RegionName            string `json:"region_name"`
	RegionGroupName       string `json:"region_group_name"`
}
