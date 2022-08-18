package models

// AccountOpening ...
type AccountOpening struct {
	ID               string        `json:"id"`
	UserID           string        `json:"user_id"`
	User             User          `json:"user"`
	Email            string        `json:"email"`
	EmailValidAt     string        `json:"email_valid_at"`
	Name             string        `json:"name"`
	MaritalStatusID  string        `json:"marital_status_id"`
	MaritalStatus    MaritalStatus `json:"marital_status"`
	GenderID         string        `json:"gender_id"`
	Gender           Gender        `json:"gender"`
	BirthPlace       string        `json:"birth_place"`
	BirthPlaceCityID string        `json:"birth_place_city_id"`
	City             City          `json:"city"`
	BirthDate        string        `json:"birth_date"`
	MotherName       string        `json:"mother_name"`
	Phone            string        `json:"phone"`
	PhoneValidAt     string        `json:"phone_valid_at"`
	Status           string        `json:"status"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	DeletedAt        *string       `json:"deleted_at"`
}

// AccountOpeningParameter ...
type AccountOpeningParameter struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	MaritalStatusID  string `json:"marital_status_id"`
	GenderID         string `json:"gender_id"`
	BirthPlaceCityID string `json:"birth_place_city_id"`
	MotherName       string `json:"mother_name"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Status           string `json:"status"`
	Search           string `json:"search"`
	Page             int    `json:"page"`
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	By               string `json:"by"`
	Sort             string `json:"sort"`
}

var (
	// AccountOpeningStatusPending ...
	AccountOpeningStatusPending = "Pending"
	// AccountOpeningStatusActive ...
	AccountOpeningStatusActive = "Active"
	// AccountOpeningCountryCodeDefault ...
	AccountOpeningCountryCodeDefault = "+62"
	// AccountOpeningOrderBy ...
	AccountOpeningOrderBy = []string{"def.id", "def.name", "def.email", "def.created_at", "def.updated_at"}
	// AccountOpeningOrderByrByString ...
	AccountOpeningOrderByrByString = []string{
		"def.name", "def.email",
	}

	// AccountOpeningSelectStatement ...
	AccountOpeningSelectStatement = `SELECT def.id, def.user_id, def.email, def.email_valid_at, def.name,
	def.marital_status_id, def.gender_id, def.birth_place, def.birth_place_city_id, def.birth_date, def.mother_name,
	def.phone, def.phone_valid_at, def.status, def.created_at, def.updated_at, def.deleted_at, u.code, u.name,
	ms.name, g.name, c.name
	FROM account_openings def
	LEFT JOIN users u ON u.id = def.user_id
	LEFT JOIN marital_statuses ms ON ms.id = def.marital_status_id
	LEFT JOIN genders g ON g.id = def.gender_id
	LEFT JOIN cities c ON c.id = def.birth_place_city_id`

	// AccountOpeningWhereStatement ...
	AccountOpeningWhereStatement = `WHERE def.deleted_at IS NULL`
)
