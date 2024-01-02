package models

import "database/sql"

const (
	PointTypeWithdraw = "Withdraw"
	PointTypeLoyalty  = "Loyalty"
	PointTypePromo    = "Promo"
	PointTypeCashback = "Cashback"
)

// Point ...
type Point struct {
	ID            string         `json:"id"`
	PointType     string         `json:"point_type"`
	PointTypeName string         `json:"point_type_name"`
	InvoiceID     sql.NullString `json:"invoice_id"`
	Point         string         `json:"point"`
	CustomerID    string         `json:"customer_id"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
	DeletedAt     sql.NullString `json:"deleted_at"`
	ExpiredAt     sql.NullString `json:"expired_at"`

	Customer          WebCustomer    `json:"customer"`
	InvoiceDocumentNo sql.NullString `json:"invoice_document_no"`
}

// PointGetBalance ...
type PointGetBalance struct {
	Withdraw string `json:"withdraw"`
	Loyalty  string `json:"loyalty"`
	Promo    string `json:"promo"`
	Cashback string `json:"cashback"`
}

// PointParameter ...
type PointParameter struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	PointType  string `json:"point_type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Renewal    string `json:"renewal"`
	Search     string `json:"search"`
	ShowAll    string `json:"show_all"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	PointOrderBy          = []string{"def.id", "def._name"}
	PointOrderByrByString = []string{
		"def._name",
	}

	PointSelectStatement = `SELECT 
			DEF.ID, 
			DEF.POINT_TYPE, 
			PT._NAME,
			DEF.INVOICE_ID,
			DEF.POINT,
			DEF.CUSTOMER_ID,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.EXPIRED_AT,
			C.CUSTOMER_NAME,
			C.CUSTOMER_CODE,
			B.BRANCH_CODE,
			B._NAME,
			R._NAME,
			SIH.DOCUMENT_NO
		FROM POINTS DEF
		LEFT JOIN POINT_TYPE PT ON PT.ID = DEF.POINT_TYPE
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
		LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
		LEFT JOIN REGION R ON R.ID = B.REGION_ID
		LEFT JOIN SALES_INVOICE_HEADER SIH ON SIH.ID = DEF.INVOICE_ID
	`

	PointWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
