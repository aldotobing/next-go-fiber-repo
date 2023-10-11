package models

import "database/sql"

// CustomerLog ...
type CustomerLog struct {
	ID           string         `json:"id"`
	CustomerID   string         `json:"customer_id"`
	CustomerCode string         `json:"customer_code"`
	CustomerName string         `json:"customer_name"`
	OldData      string         `json:"old_data"`
	NewData      string         `json:"new_data"`
	UserID       sql.NullString `json:"user_id"`
	UserName     sql.NullString `json:"user_name"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    sql.NullString `json:"updated_at"`
	DeletedAt    sql.NullString `json:"deleted_at"`
}

// CustomerLogParameter ...
type CustomerLogParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	Name           string `json:"customer_name"`
	CustomerTypeId string `json:"custome_type_id"`
	SalesmanTypeID string `json:"salesman_type_id"`
	UserId         string `json:"admin_user_id"`
	BranchId       string `json:"branch_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
	PhoneNumber    string `json:"phone_number"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	ShowInApp      string `json:"show_in_app"`
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
		def.created_at
	FROM
		customer_logs def
		left join customer c on c.id = def.customer_id
		left join branch b on b.id = c.branch_id
		left join region r on r.id = b.region_id
		left join _user u on u.id = def.user_id
	`

	// CustomerWhereStatement ...
	CustomerLogWhereStatement = ` WHERE c.created_date IS not NULL `
)
