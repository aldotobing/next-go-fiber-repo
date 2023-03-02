package models

// Doctor ...
type WebSalesman struct {
	ID                    *string `json:"salesman_id"`
	Code                  *string `json:"salesman_code"`
	PartnerName           *string `json:"salesman_name"`
	PartnerPhone          *string `json:"salesman_phone"`
	PartnerUserID         *string `json:"salesman_user_id"`
	PartnerUserName       *string `json:"salesman_user_name"`
	PartnerProfilePicture *string `json:"salesman_profile_picture"`
}

// DoctorParameter ...
type WebSalesmanParameter struct {
	ID             string `json:"salesman_id"`
	Code           string `json:"salesman_code"`
	Name           string `json:"salesman_name"`
	BranchID       string `json:"branch_id"`
	CustomerTypeId string `json:"custome_type_id"`
	UserId         string `json:"admin_user_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	WebSalesmanOrderBy = []string{"def.id", "def.salesman_name", "def.created_date"}
	// CustomerOrderByrByString ...
	WebSalesmanOrderByrByString = []string{
		"def.salesman_name",
	}

	// CustomerSelectStatement ...

	WebSalesmanSelectStatement = `
	select def.id,def.salesman_code, def.salesman_name, def.salesman_phone_no,
	us.id,us.login
		from salesman def
		left join _user us on us.id = def.user_id
	`

	// CustomerWhereStatement ...
	WebSalesmanWhereStatement = `where def.created_date is not null  `
)
