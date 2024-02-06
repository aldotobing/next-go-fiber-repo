package models

import "database/sql"

// CouponRedeem ...
type CouponRedeem struct {
	ID                    string         `json:"id"`
	CouponID              string         `json:"coupon_id"`
	CouponName            string         `json:"coupon_name"`
	CouponDescription     string         `json:"coupon_description"`
	CouponPointConversion string         `json:"coupon_point_conversion"`
	CouponPhotoURL        sql.NullString `json:"coupon_photo_url"`
	CustomerID            string         `json:"customer_id"`
	CustomerName          string         `json:"customer_name"`
	Redeem                string         `json:"redeem"`
	RedeemAt              sql.NullString `json:"redeem_at"`
	RedeemedToDocumentNo  sql.NullString `json:"redeem_to_doc_no"`
	CreatedAt             string         `json:"created_at"`
	UpdatedAt             sql.NullString `json:"updated_at"`
	DeletedAt             sql.NullString `json:"deleted_at"`
	ExpiredAt             sql.NullString `json:"expired_at"`
}

// CouponRedeemReport ...
type CouponRedeemReport struct {
	ID                    string         `json:"id"`
	CouponID              string         `json:"coupon_id"`
	CouponName            string         `json:"coupon_name"`
	CouponDescription     string         `json:"coupon_description"`
	CouponPointConversion string         `json:"coupon_point_conversion"`
	CustomerID            string         `json:"customer_id"`
	CustomerName          string         `json:"customer_name"`
	CustomerCode          string         `json:"customer_code"`
	Redeem                string         `json:"redeem"`
	RedeemAt              sql.NullString `json:"redeem_at"`
	RedeemedToDocumentNo  sql.NullString `json:"redeem_to_doc_no"`
	CreatedAt             string         `json:"created_at"`
	UpdatedAt             sql.NullString `json:"updated_at"`
	DeletedAt             sql.NullString `json:"deleted_at"`
	ExpiredAt             sql.NullString `json:"expired_at"`
	BranchName            string         `json:"branch_name"`
	BranchCode            string         `json:"branch_code"`
	RegionName            string         `json:"region_name"`
	RegionGroupName       string         `json:"region_group_name"`
	CustomerLevelName     sql.NullString `json:"customer_level_name"`
}

// CouponRedeemParameter ...
type CouponRedeemParameter struct {
	ID                   string `json:"id"`
	Now                  string `json:"now"`
	Search               string `json:"search"`
	ShowAll              string `json:"show_all"`
	CustomerID           string `json:"customer_id"`
	RedeemedToDocumentNo string `json:"redeem_to_doc_no"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	BranchID             string `json:"branch_id"`
	RegionID             string `json:"region_id"`
	RegionGroupID        string `json:"region_group_id"`
	Page                 int    `json:"page"`
	Offset               int    `json:"offset"`
	Limit                int    `json:"limit"`
	By                   string `json:"by"`
	Sort                 string `json:"sort"`
}

var (
	CouponRedeemOrderBy          = []string{"def.id"}
	CouponRedeemOrderByrByString = []string{}

	CouponRedeemSelectStatement = `SELECT 
			DEF.ID, 
			DEF.COUPON_ID,
			DEF.CUSTOMER_ID,
			DEF.REDEEMED,
			DEF.REDEEMED_AT,
			DEF.REDEEM_TO_DOC_NO,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.EXPIRED_AT,
			CP._NAME,
			CP.DESCRIPTION,
			CP.POINT_CONVERSION,
			CP.PHOTO_URL,
			C.CUSTOMER_NAME
		FROM COUPON_REDEEM DEF
		LEFT JOIN COUPONS CP ON CP.ID = DEF.COUPON_ID
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
	`
	CouponRedeemWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
