package requests

// DistrictRequest ...
type DistrictRequest struct {
	CityID string `json:"city_id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Status bool   `json:"status"`
}
