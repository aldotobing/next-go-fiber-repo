package viewmodel

// DistrictVM ....
type DistrictVM struct {
	ID        string `json:"id"`
	CityID    string `json:"city_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
