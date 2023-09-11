package models

import "database/sql"

// VoucherRedeem ...
type VoucherRedeem struct {
	ID               string         `json:"id"`
	CustomerID       string         `json:"customer_id"`
	Redeemed         string         `json:"redeemed"`
	RedeemedAt       sql.NullString `json:"redeemed_at"`
	RedeemedToDocNo  sql.NullString `json:"redeemed_to_doc_no"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        sql.NullString `json:"updated_at"`
	DeletedAt        sql.NullString `json:"deleted_at"`
	VoucherID        string         `json:"voucher_id"`
	VoucherName      string         `json:"voucher_name"`
	VoucherCashValue string         `json:"voucher_cash_value"`
}

// VoucherRedeemParameter ...
type VoucherRedeemParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	VoucherRedeemOrderBy          = []string{"def.id"}
	VoucherRedeemOrderByrByString = []string{}

	VoucherRedeemSelectStatement = `SELECT 
		DEF.ID, 
		DEF.CUSTOMER_ID, 
		DEF.REDEEMED, 
		DEF.REDEEMED_AT, 
		DEF.REDEEMED_TO_DOC_NO,
		DEF.CREATED_AT,
		DEF.UPDATED_AT,
		DEF.DELETED_AT,
		DEF.VOUCHER_ID,
		V._NAME,
		V.CASH_VALUE
	FROM VOUCHER_REDEEM DEF
	LEFT JOIN VOUCHER V ON V.ID = DEF.VOUCHER_ID
	`

	VoucherRedeemWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
