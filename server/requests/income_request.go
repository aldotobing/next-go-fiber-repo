package requests

// IncomeRequest ...
type IncomeRequest struct {
	Name        string  `json:"name" validate:"required"`
	MappingName string  `json:"mapping_name"`
	MinValue    float64 `json:"min_value"`
	MaxValue    float64 `json:"max_value"`
	Status      bool    `json:"status"`
}
