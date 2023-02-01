package models

// Doctor ...
type WebPartner struct {
	ID                    *string `json:"partner_id"`
	Code                  *string `json:"partner_code"`
	PartnerName           *string `json:"partner_name"`
	PartnerPhone          *string `json:"partner_phone"`
	PartnerAddress        *string `json:"partner_address"`
	PartnerUserID         *string `json:"partner_user_id"`
	PartnerUserName       *string `json:"partner_user_name"`
	PartnerProfilePicture *string `json:"partner_profile_picture"`
}

// DoctorParameter ...
type WebPartnerParameter struct {
	ID             string `json:"partner_id"`
	Code           string `json:"partner_code"`
	Name           string `json:"partner_name"`
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
	WebPartnerOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// CustomerOrderByrByString ...
	WebPartnerOrderByrByString = []string{
		"def._name",
	}

	// CustomerSelectStatement ...

	WebPartnerSelectStatement = `
	select def.id,def.code, def._name, def.address,def.phone_no,
	us.id,us.login
		from partner def
	`

	// CustomerWhereStatement ...
	WebPartnerWhereStatement = `where def.created_date is not null and def.is_mysm = 1  `
)
