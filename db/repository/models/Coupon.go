package models

import "database/sql"

// Coupon ...
type Coupon struct {
	ID              string         `json:"id"`
	StartDate       string         `json:"start_dates"`
	EndDate         string         `json:"end_date"`
	PointConversion string         `json:"point_conversion"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       sql.NullString `json:"updated_at"`
	DeletedAt       sql.NullString `json:"deleted_at"`
	Interval        sql.NullInt64  `json:"interval"`
	PhotoURL        sql.NullString `json:"photo_url"`
}

// CouponParameter ...
type CouponParameter struct {
	ID      string   `json:"id"`
	IDs     []string `json:"ids"`
	Now     string   `json:"now"`
	Search  string   `json:"search"`
	ShowAll string   `json:"show_all"`
	Page    int      `json:"page"`
	Offset  int      `json:"offset"`
	Limit   int      `json:"limit"`
	By      string   `json:"by"`
	Sort    string   `json:"sort"`
}

var (
	CouponOrderBy          = []string{"def.id"}
	CouponOrderByrByString = []string{}

	CouponSelectStatement = `SELECT 
			DEF.ID, 
			DEF.START_DATE,
			DEF.END_DATE,
			DEF.POINT_CONVERSION,
			DEF._NAME,
			DEF.DESCRIPTION,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.INTERVAL_DAY,
			DEF.PHOTO_URL
		FROM COUPONS DEF
	`
	CouponWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
