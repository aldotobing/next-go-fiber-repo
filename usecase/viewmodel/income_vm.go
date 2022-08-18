package viewmodel

// IncomeVM ....
type IncomeVM struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MappingName string  `json:"mapping_name"`
	MinValue    float64 `json:"min_value"`
	MaxValue    float64 `json:"max_value"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   string  `json:"deleted_at"`
}
