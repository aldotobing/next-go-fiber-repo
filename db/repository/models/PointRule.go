package models

import "database/sql"

// PointRule ...
type PointRule struct {
	ID              string         `json:"id"`
	StartDate       string         `json:"start_dates"`
	EndDate         string         `json:"end_date"`
	MinOrder        string         `json:"min_order"`
	PointConversion string         `json:"point_conversion"`
	MonthlyMaxPoint string         `json:"monthly_max_point"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       sql.NullString `json:"updated_at"`
	DeletedAt       sql.NullString `json:"deleted_at"`

	Customer sql.NullString `json:"customer"`
}

// PointRuleParameter ...
type PointRuleParameter struct {
	ID      string `json:"id"`
	Search  string `json:"search"`
	ShowAll string `json:"show_all"`
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	By      string `json:"by"`
	Sort    string `json:"sort"`
}

var (
	PointRuleOrderBy          = []string{"def.id"}
	PointRuleOrderByrByString = []string{}

	PointRuleSelectStatement = `SELECT 
			DEF.ID, 
			DEF.START_DATE,
			DEF.END_DATE,
			DEF.MIN_ORDER,
			DEF.POINT_CONVERSION,
			DEF.MONTHLY_MAX_POINT,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.CUSTOMER
		FROM POINT_RULES DEF
	`
	PointRuleWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
