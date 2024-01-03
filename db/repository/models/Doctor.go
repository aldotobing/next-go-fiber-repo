package models

// Doctor ...
type Doctor struct {
	ID            *string `json:"doctor_id"`
	Code          *string `json:"doctor_code"`
	DoctorName    *string `json:"doctor_name"`
	DoctorAddress *string `json:"doctor_address"`
	DoctorPhone   *string `json:"doctor_phone"`
}

// DoctorParameter ...
type DoctorParameter struct {
	ID           string `json:"Doctor_id"`
	Code         string `json:"Doctor_code"`
	Phone        string `json:"Doctor_phone"`
	Name         string `json:"Doctor_name"`
	DoctorTypeId string `json:"custome_type_id"`
	UserId       string `json:"admin_user_id"`
	Search       string `json:"search"`
	Page         int    `json:"page"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	By           string `json:"by"`
	Sort         string `json:"sort"`
}

var (
	// DoctorOrderBy ...
	DoctorOrderBy = []string{"p.id", "p._name", "p.created_date"}
	// DoctorOrderByrByString ...
	DoctorOrderByrByString = []string{
		"p._name",
	}

	// DoctorSelectStatement ...

	DoctorSelectStatement = `
	
		select def.id,def.code, def._name, def.address,def.phone_no
		from partner def
	`

	// DoctorWhereStatement ...
	DoctorWhereStatement = ` where def.created_date is not null `
)

// '+62' || regexp_replace(SUBSTRING (Doctor_phone, 2, length(Doctor_phone)),'[^\w]+','','g' )
