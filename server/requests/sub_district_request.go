package requests

// SubDistrictRequest ...
type SubDistrictRequest struct {
	DistrictID string `json:"district_id" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Name       string `json:"name" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Status     bool   `json:"status"`
}
