package requests

// EducationLevelRequest ...
type EducationLevelRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
