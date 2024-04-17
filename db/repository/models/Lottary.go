package models

import "database/sql"

// Lottary ...
type Lottary struct {
	ID              string
	SerailNo        string
	Status          string
	CustomerCode    string
	CustomerName    sql.NullString
	CreatedAt       string
	UpdatedAt       sql.NullString
	DeletedAt       sql.NullString
	Year            sql.NullString
	Quartal         sql.NullString
	Sequence        sql.NullString
	BranchName      sql.NullString
	RegionCode      sql.NullString
	RegionName      sql.NullString
	RegionGroup     sql.NullString
	CustomerType    sql.NullString
	CustomerLevel   sql.NullString
	CustomerAddress sql.NullString
	CustomerCpName  sql.NullString
}

// LottaryParameter ...
type LottaryParameter struct {
	ID            string `json:"id"`
	Quartal       string `json:"quartal"`
	Year          string `json:"year"`
	CustomerID    string `json:"customer_id"`
	SerialNo      string `json:"serial_no"`
	StartDate     string `json:"start_date"`
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
	LottaryOrderBy          = []string{"def.id"}
	LottaryOrderByrByString = []string{}

	LottarySelectStatement = `SELECT 
			DEF.ID, 
			DEF.serial_no, 
			DEF.status,
			C.CUSTOMER_CODE,
			C.CUSTOMER_NAME,
			DEF.CREATED_DATE,
			DEF.MODIFIED_DATE,
			DEF.DELETED_DATE,
			DEF._year,
			DEF._quartal,
			DEF._sequence,
			b._name as b_name,
			r.code as r_code,
			r._name as r_name,
			r.group_name as r_g_name,
			c.customer_cp_name,
			cl._name as cus_level_name,
			ctp._name as ctp_name,
			c.customer_address
		FROM lottary DEF
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
		LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
		left join region r on r.id = b.region_id
		left join customer_level cl on cl.id = c.customer_level_id
		left join customer_type ctp on ctp.id= c.customer_type_id
		
	`

	LottaryWhereStatement = `WHERE DEF.DELETED_DATE IS NULL `
)
