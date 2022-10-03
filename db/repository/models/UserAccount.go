package models

type UserAccount struct {
	ID                 string  `json:"id_user"`
	CustomerID         *string `json:"customer_id"`
	Name               *string `json:"name"`
	Code               *string `json:"code"`
	Phone              *string `json:"phone"`
	PriceListID        *string `json:"price_list_id"`
	PriceListVersionID *string `json:"price_list_version_id"`
	CustomerTypeID     *string `json:"customer_type_id"`
	CustomerLevelName  *string `json:"customer_level_name"`
	CustomerAddress    *string `json:"customer_address"`
	SalesmanID         *string `json:"salesman_id"`
	SalesmanName       *string `json:"salesman_name"`
	SalesmanCode       *string `json:"salesman_code"`
}

type UserAccountParameter struct {
	ID            string `json:"id_user"`
	CustomerID    string `json:"customer_id"`
	BranchID      string `json:"id_branch"`
	FbId          string `json:"id_facebook"`
	GoogleId      string `json:"id_google"`
	AppleId       string `json:"id_apple"`
	Name          string `json:"name"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Gender        string `json:"gender"`
	PhoneNo       string `json:"phone_no"`
	Code          string `json:"code_cus"`
	Email         string `json:"email_user"`
	NoTelp        string `json:"no_telp"`
	Address       string `json:"address_user"`
	Level         int    `json:"level"`
	BirtDate      string `json:"birthdate"`
	BirtDatePlace string `json:"birthplace"`
	ReferalCode   string `json:"referral_code"`

	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// UserAccountOrderBy ...
	UserAccountOrderBy = []string{"def.id_user", "def.name", "def.username", "def.created_at_user", "def.updated_at_user"}
	// UserAccountOrderByrByString ...
	UserAccountOrderByrByString = []string{
		"def.username",
	}

	// UserAccountSelectStatement ...
	UserAccountSelectStatement = ` select def.id as user_id,cus.id as cus_id,
	cus.customer_name,
	cus.customer_code,cus.customer_phone,
	cus.price_list_id,
	(select plv.id from price_list_version plv where plv.price_list_id = pl.id and now()::date between plv.start_date and plv.end_date) as version_id,
	cus.customer_type_id,cl._name cus_level_name,cus.customer_address,
	s.id as salesman_id,s.salesman_code,s.salesman_name
	from _user def 
	join customer cus on cus.id = def.partner_id 
	left join price_list pl on pl.id = cus.price_list_id
	left join customer_level cl on cl.id = cus.customer_level_id
	left join salesman s on s.id = cus.salesman_id
	`

	// UserAccountSelectStatement ...
	AdminUserAccountSelectStatement = ` select def.id as user_id,null as cus_id,
	def.login,
	null,null,
	null,
	null,
	null,null,null,
	null,null,null
	from _user def
	`

	// UserAccountWhereStatement ...
	UserAccountWhereStatement = ` def.created_date is not null `
)
