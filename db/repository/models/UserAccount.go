package models

type UserAccount struct {
	ID                 string  `json:"id_user"`
	CustomerID         *string `json:"customer_id"`
	Name               *string `json:"name"`
	Code               *string `json:"code"`
	Phone              *string `json:"phone"`
	UserAddress        *string `json:"user_address"`
	PriceListID        *string `json:"price_list_id"`
	PriceListVersionID *string `json:"price_list_version_id"`
	CustomerTypeID     *string `json:"customer_type_id"`
	CustomerLevelName  *string `json:"customer_level_name"`
	CustomerAddress    *string `json:"customer_address"`
	SalesmanID         *string `json:"salesman_id"`
	SalesmanName       *string `json:"salesman_name"`
	SalesmanCode       *string `json:"salesman_code"`
	FireStoreUID       *string `json:"firestore_uid"`
	FCMToken           *string `json:"fcm_token"`
	RoleIDList         *string `json:"role_id_list"`
	LoginCode          *string `json:"logincode"`
	ShowInApp          bool    `json:"show_in_app"`
	Active             string  `json:"active"`
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
	UserAccountSelectStatement = ` 
	select def.id,def.login,
	coalesce(
			(
			select STRING_AGG(id::character varying,',') from role where id in (
				select role_id from role_group_role_line 
				where role_group_id in( select role_group_id from user_role_group where user_id = def.id) 
				) and is_mysm=1
			),''
		),def.fcm_token,
		coalesce(c.show_in_apps, '0'), coalesce(c.active,'0')
	from _user def
	left join customer c on c.user_id = def.id

	`

	// UserAccountSelectStatement ...
	AdminUserAccountSelectStatement = ` select def.id,def.login,
		coalesce(
				(
				select STRING_AGG(id::character varying,',') from role where id in (
					select role_id from role_group_role_line 
					where role_group_id in( select role_group_id from user_role_group where user_id = def.id) 
					) and is_mysm=1
				),''
			),def.fcm_token
		from _user def
	`

	// UserAccountWhereStatement ...
	UserAccountWhereStatement = ` def.created_date is not null `
)
