package requests

// LoginRequest ....
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest ...
type RegisterRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Status      bool   `json:"status"`
	RoleID      string `json:"role_id"`
}

// VerifyMailRequest ...
type VerifyMailRequest struct {
	Email string `json:"email" validate:"required"`
}

// VerifyUserKeyRequest ...
type VerifyUserKeyRequest struct {
	Key string `json:"key" validate:"required"`
}
