package models

import "database/sql"

// CustomerLog ...
type CustomerLog struct {
	ID                string         `json:"id"`
	CustomerID        string         `json:"customer_id"`
	CustomerCode      string         `json:"customer_code"`
	CustomerName      string         `json:"customer_name"`
	OldData           string         `json:"old_data"`
	NewData           string         `json:"new_data"`
	UserID            sql.NullString `json:"user_id"`
	UserName          sql.NullString `json:"user_name"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         sql.NullString `json:"updated_at"`
	DeletedAt         sql.NullString `json:"deleted_at"`
	BranchID          sql.NullString `json:"branch_id"`
	BranchName        sql.NullString `json:"branch_name"`
	RegionID          sql.NullString `json:"region_id"`
	RegionName        sql.NullString `json:"region_name"`
	RegionGroupID     sql.NullString `json:"region_group_id"`
	RegionGroupName   sql.NullString `json:"region_group_name"`
	CustomerLevelID   sql.NullString `json:"customer_level_id"`
	CustomerLevelName sql.NullString `json:"customer_level_name"`
}

// CustomerLogParameter ...
type CustomerLogParameter struct {
	ID              string `json:"id"`
	UserId          string `json:"user_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	RegionGroupID   string `json:"region_group_id"`
	RegionID        string `json:"region_id"`
	CustomerLevelID string `json:"customer_level_id"`
	BranchID        string `json:"branch_id"`
	Search          string `json:"search"`
	Page            int    `json:"page"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	By              string `json:"by"`
	Sort            string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	CustomerLogOrderBy = []string{"def.id"}
	// CustomerOrderByrByString ...
	CustomerLogOrderByrByString = []string{}

	// CustomerLogSelectStatement ...
	CustomerLogSelectStatement = `
	SELECT 
		def.id, 
		def.customer_id, c.customer_code, c.customer_name,
		def.old_data, def.new_data, 
		def.user_id, u.login,
		def.created_at,
		b.id, b._name, r.id, r._name, r.group_id, r.group_name,
		cl.id, cl._name
	FROM
		customer_logs def
		left join customer c on c.id = def.customer_id
		left join branch b on b.id = c.branch_id
		left join region r on r.id = b.region_id
		left join _user u on u.id = def.user_id
		left join customer_level cl on cl.id = c.customer_level_id
	`

	// CustomerWhereStatement ...
	CustomerLogWhereStatement = ` WHERE c.created_date IS not NULL `
)
