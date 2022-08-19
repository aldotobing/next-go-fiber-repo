package models

type UserAccount struct {
	ID   string  `json:"id_user"`
	Name *string `json:"name"`
	Code *string `json:"code"`
}

type UserAccountParameter struct {
	ID            string `json:"id_user"`
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
	UserAccountSelectStatement = `select def.id as user_id,p._name,p.code
	from _user def 
	join partner p on p.id = def.partner_id 
	`

	// UserAccountWhereStatement ...
	UserAccountWhereStatement = ` def.created_date is not null `
)
