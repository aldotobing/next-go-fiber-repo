package models

// Doctor ...
type WebDoctor struct {
	ID                   *string `json:"doctor_id"`
	Code                 *string `json:"doctor_code"`
	DoctorName           *string `json:"doctor_name"`
	DoctorPhone          *string `json:"doctor_phone"`
	DoctorAddress        *string `json:"doctor_address"`
	DoctorUserID         *string `json:"doctor_user_id"`
	DoctorUserName       *string `json:"doctor_user_name"`
	DoctorProfilePicture *string `json:"doctor_profile_picture"`
}

// DoctorParameter ...
type WebDoctorParameter struct {
	ID             string `json:"doctor_id"`
	Code           string `json:"doctor_code"`
	Name           string `json:"doctor_name"`
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
	WebDoctorOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// CustomerOrderByrByString ...
	WebDoctorOrderByrByString = []string{
		"def._name",
	}

	// CustomerSelectStatement ...

	WebDoctorSelectStatement = `
	select def.id,def.code, def._name, def.address,def.phone_no,
	us.id,us.login
		from partner def
	`

	// CustomerWhereStatement ...
	WebDoctorWhereStatement = `where def.created_date is not null and (select count(*) from role_group_role_line where role_group_id in(
		select role_group_id from user_role_group where user_id = us.id 
		) and role_id =111111002) >= 1  `
)
