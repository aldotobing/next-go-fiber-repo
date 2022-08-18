package viewmodel

// SubDistrictVM ....
type SubDistrictVM struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	PostalCode string `json:"postal_code"`
	Status     bool   `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
