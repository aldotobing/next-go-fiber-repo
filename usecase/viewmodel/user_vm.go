package viewmodel

// UserVM ....
type UserVM struct {
	ID        string `json:"user_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type AccountOpeningVM struct {
	ID               string          `json:"id"`
	UserID           string          `json:"user_id"`
	User             UserVM          `json:"user"`
	BirthPlace       string          `json:"birth_place"`
	BirthPlaceCityID string          `json:"birth_place_city_id"`
	BirthPlaceCity   CityVM          `json:"birth_place_city"`
	BirthDate        string          `json:"birth_date"`
	Email            string          `json:"email"`
	EmailValidAt     string          `json:"email_valid_at"`
	GenderID         string          `json:"gender_id"`
	Gender           GenderVM        `json:"gender"`
	MaritalStatusID  string          `json:"marital_status_id"`
	MaritalStatus    MaritalStatusVM `json:"marital_status"`
	MotherName       string          `json:"mother_name"`
	Name             string          `json:"name"`
	Phone            string          `json:"phone"`
	PhoneValidAt     string          `json:"phone_valid_at"`
	Status           bool            `json:"status"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
	DeletedAt        string          `json:"deleted_at"`
}
