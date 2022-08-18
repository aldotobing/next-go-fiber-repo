package requests

// ResidenceOwnershipRequest ...
type ResidenceOwnershipRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
