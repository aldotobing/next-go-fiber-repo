package models

// Doctor ...
type UserCheckinActivity struct {
	ID           *string `json:"id"`
	UserID       *string `json:"user_id"`
	CheckinTime  *string `json:"checkin_time"`
	CreatedDate  *string `json:"created_date"`
	CreatedBy    *string `json:"created_by"`
	ModifiedDate *string `json:"modified_date"`
	ModifiedBy   *string `json:"modified_by"`
	Login        *string `json:"login"`
}

// DoctorParameter ...
type UserCheckinActivityParameter struct {
	ID            string `json:"id"`
	UserId        string `json:"user_id"`
	Search        string `json:"search"`
	CurrentVaUser int    `json:"current_va_user"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	By            string `json:"by"`
	Sort          string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	UserCheckinActivityOrderBy = []string{"def.id", "usr.login", "def.created_date"}
	// CustomerOrderByrByString ...
	UserCheckinActivityOrderByrByString = []string{
		"usr.login",
	}

	// CustomerSelectStatement ...

	UserCheckinActivitySelectStatement = `
	select def.id,def.user_id, def.checkin_time, usr.login
		from user_checkin_activity def
		left join _user usr on usr.id = def.user_id 		
	`

	// CustomerWhereStatement ...
	UserCheckinActivityWhereStatement = `where def.created_date is not null  `
)
