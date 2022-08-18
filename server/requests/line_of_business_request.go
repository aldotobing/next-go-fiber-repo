package requests

// LineOfBusinessRequest ...
type LineOfBusinessRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
