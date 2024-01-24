package models

import "database/sql"

// CouponRedeem ...
type CouponRedeem struct {
	ID                string         `json:"id"`
	CouponID          string         `json:"coupon_id"`
	CouponName        string         `json:"coupon_name"`
	CouponDescription string         `json:"coupon_description"`
	CustomerID        string         `json:"customer_id"`
	CustomerName      string         `json:"customer_name"`
	Redeem            string         `json:"redeem"`
	RedeemAt          sql.NullString `json:"redeem_at"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         sql.NullString `json:"updated_at"`
	DeletedAt         sql.NullString `json:"deleted_at"`
}

// CouponRedeemParameter ...
type CouponRedeemParameter struct {
	ID         string `json:"id"`
	Now        string `json:"now"`
	Search     string `json:"search"`
	ShowAll    string `json:"show_all"`
	CustomerID string `json:"customer_id"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
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
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			CP._NAME,
			CP.DESCRIPTION,
			C.CUSTOMER_NAME
		FROM COUPON_REDEEM DEF
		LEFT JOIN COUPONS CP ON CP.ID = DEF.COUPON_ID
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
	`
	CouponRedeemWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
