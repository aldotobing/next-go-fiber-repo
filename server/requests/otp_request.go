package requests

// UserOtpRequest ...
type UserOtpRequest struct {
	Type  string `json:"type" validate:"required"`
	Phone string `json:"phone"`
	// CountryCode string `json:"country_code" validate:"required"`
	// Phone       string `json:"phone" validate:"required"`
}

// UserOtpSubmit ...
type UserOtpSubmit struct {
	Type string `json:"type"`
	Otp  string `json:"otp" validate:"required"`
}
