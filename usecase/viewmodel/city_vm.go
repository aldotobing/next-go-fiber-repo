package viewmodel

// CityVM ....
type CityVM struct {
	ID          string `json:"id"`
	ProvinceID  string `json:"province_id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	MappingName string `json:"mapping_name"`
	Type        int    `json:"type"`
	Status      bool   `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
