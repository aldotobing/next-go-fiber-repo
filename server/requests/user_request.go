package requests

// AccountOpeningRequest ...
type AccountOpeningRequest struct {
	UserID           string `json:"user_id" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Name             string `json:"name" validate:"required"`
	MaritalStatusID  string `json:"marital_status_id"`
	GenderID         string `json:"gender_id"`
	BirthPlace       string `json:"birth_place"`
	BirthPlaceCityID string `json:"birth_place_city_id"`
	BirthDate        string `json:"birth_date"`
	MotherName       string `json:"mother_name"`
	Phone            string `json:"phone" validate:"required"`
	Status           string `json:"status"`
}
