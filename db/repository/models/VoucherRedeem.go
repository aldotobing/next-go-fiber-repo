package models

import "database/sql"

// VoucherRedeem ...
type VoucherRedeem struct {
	ID                 string         `json:"id"`
	CustomerCode       string         `json:"customer_code"`
	Redeemed           string         `json:"redeemed"`
	RedeemedAt         sql.NullString `json:"redeemed_at"`
	RedeemedToDocNo    sql.NullString `json:"redeemed_to_doc_no"`
	CreatedAt          string         `json:"created_at"`
	UpdatedAt          sql.NullString `json:"updated_at"`
	DeletedAt          sql.NullString `json:"deleted_at"`
	VoucherID          string         `json:"voucher_id"`
	VoucherName        string         `json:"voucher_name"`
	VoucherCashValue   string         `json:"voucher_cash_value"`
	VoucherDescription sql.NullString `json:"voucher_description"`
	VoucherImageURL    string         `json:"voucher_image_url"`
	VoucherStartDate   string         `json:"voucher_start_date"`
	VoucherEndDate     string         `json:"voucher_end_date"`
}

// VoucherRedeemParameter ...
type VoucherRedeemParameter struct {
	ID           string `json:"id"`
	CustomerCode string `json:"customer_code"`
	DocumentNo   string `json:"document_no"`
	VoucherID    string `json:"voucher_id"`
	ShowAll      string `json:"show_all"`
	Search       string `json:"search"`
	Page         int    `json:"page"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	By           string `json:"by"`
	Sort         string `json:"sort"`
}

var (
	VoucherRedeemOrderBy          = []string{"def.id"}
	VoucherRedeemOrderByrByString = []string{}

	VoucherRedeemSelectStatement = `SELECT 
		DEF.ID, 
		DEF.CUSTOMER_CODE, 
		DEF.REDEEMED, 
		DEF.REDEEMED_AT, 
		DEF.REDEEMED_TO_DOC_NO,
		DEF.CREATED_AT,
		DEF.UPDATED_AT,
		DEF.DELETED_AT,
		DEF.VOUCHER_ID,
		V._NAME,
		V.CASH_VALUE,
		V.DESCRIPTION,
		V.IMAGE_URL,
		V.START_DATE,
		V.END_DATE
	FROM VOUCHER_REDEEM DEF
	LEFT JOIN VOUCHER V ON V.ID = DEF.VOUCHER_ID
	`

	VoucherRedeemWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
