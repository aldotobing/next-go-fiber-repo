package requests

// MaritalStatusRequest ...
type MaritalStatusRequest struct {
	Name           string `json:"name" validate:"required"`
	MappingName    string `json:"mapping_name"`
	FillSpouseName bool   `json:"fill_spouse_name"`
	Status         bool   `json:"status"`
}
