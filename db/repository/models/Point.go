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
			DEF.DELETED_AT
		FROM POINTS DEF
		LEFT JOIN POINT_TYPE PT ON PT.ID = DEF.POINT_TYPE
	`
	PointWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
