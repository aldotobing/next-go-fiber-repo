package requests

// InvestmentPurposeRequest ...
type InvestmentPurposeRequest struct {
	Name        string `json:"name" validate:"required"`
	MappingName string `json:"mapping_name"`
	Status      bool   `json:"status"`
}
