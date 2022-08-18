package requests

// GenderRequest ...
type GenderRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
