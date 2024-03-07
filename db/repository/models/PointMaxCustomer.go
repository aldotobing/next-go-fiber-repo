package models

import "database/sql"

// PointMaxCustomer ...
type PointMaxCustomer struct {
	ID              string
	StartDate       string
	EndDate         string
	CustomerCode    string
	CustomerName    sql.NullString
	MonthlyMaxPoint string
	CreatedAt       string
	UpdatedAt       sql.NullString
	DeletedAt       sql.NullString
}

// PointMaxCustomerParameter ...
type PointMaxCustomerParameter struct {
	ID            string `json:"id"`
	Month         string `json:"month"`
	Year          string `json:"year"`
	CustomerID    string `json:"customer_id"`
	PointType     string `json:"point_type"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Renewal       string `json:"renewal"`
	Search        string `json:"search"`
	ShowAll       string `json:"show_all"`
	RegionID      string `json:"region_id"`
	RegionGroupID string `json:"region_group_id"`
	BranchID      string `json:"branch_id"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	By            string `json:"by"`
	Sort          string `json:"sort"`
}

var (
	PointMaxCustomerOrderBy          = []string{"def.id"}
	PointMaxCustomerOrderByrByString = []string{}

	PointMaxCustomerSelectStatement = `SELECT 
			DEF.ID, 
			DEF.START_DATE, 
			DEF.END_DATE,
			DEF.CUSTOMER_CODE,
			C.CUSTOMER_NAME,
			DEF.MONTHLY_MAX_POINT,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT
		FROM POINT_MAX_CUSTOMER DEF
		LEFT JOIN CUSTOMER C ON C.CUSTOMER_CODE = DEF.CUSTOMER_CODE
	`

	PointMaxCustomerWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
