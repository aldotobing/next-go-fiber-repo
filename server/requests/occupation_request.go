package requests

// OccupationRequest ...
type OccupationRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
