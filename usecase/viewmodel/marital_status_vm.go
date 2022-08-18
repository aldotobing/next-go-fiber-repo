package viewmodel

// MaritalStatusVM ....
type MaritalStatusVM struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	MappingName    string `json:"mapping_name"`
	FillSpouseName bool   `json:"fill_spouse_name"`
	Status         bool   `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
