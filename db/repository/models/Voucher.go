package models

import "database/sql"

// Voucher ...
type Voucher struct {
	ID                string         `json:"id"`
	Code              string         `json:"code"`
	Name              string         `json:"name"`
	StartDate         string         `json:"start_date"`
	EndDate           string         `json:"end_date"`
	ImageURL          string         `json:"image_url"`
	VoucherCategoryID string         `json:"voucher_category_id"`
	CashValue         string         `json:"cash_value"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         sql.NullString `json:"updated_at"`
	DeletedAt         sql.NullString `json:"deleted_at"`
	Description       sql.NullString `json:"description"`
	TermAndCondition  sql.NullString `json:"term_and_condition"`
}

// VoucherParameter ...
type VoucherParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	VoucherOrderBy          = []string{"def.id", "def._name"}
	VoucherOrderByrByString = []string{
		"def._name",
	}

	VoucherSelectStatement = `SELECT 
			DEF.ID, 
			DEF.CODE, 
			DEF._NAME, 
			DEF.START_DATE, 
			DEF.END_DATE,
			DEF.IMAGE_URL,
			DEF.VOUCHER_CATEGORY_ID,
			DEF.CASH_VALUE,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.DESCRIPTION,
			DEF.TERM_AND_CONDITION
		FROM VOUCHER DEF
	`
	VoucherWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
