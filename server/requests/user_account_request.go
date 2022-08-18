package requests

type UserAccountRegisterRequest struct {
	UserName    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Telepon     string `json:"telepon" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
}

type UserAccountLoginRequest struct {
	PhoneNo string `json:"phone_no" validate:"required"`
}

type UserAccountLoginBackEndRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserAccountRequest struct {
	Name      string `json:"name"`
	UserName  string `json:"username" validate:"required"`
	Level     int    `json:"level"`
	Email     string `json:"email" validate:"required"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Telepon   string `json:"phone_number" validate:"required"`
	Address   string `json:"address"`
	UpdatedBy int    `json:"updated_by"`
}
