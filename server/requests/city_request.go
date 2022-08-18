package requests

// CityRequest ...
type CityRequest struct {
	ProvinceID string `json:"id_province" validate:"required"`
	Name       string `json:"name_city" validate:"required"`
	Long       string `json:"long_city" `
	Lat        string `json:"lat_city"`
	CreatedBy  int    `json:"created_by_city"`
	UpdatedBy  int    `json:"updated_by_city"`
	DeletedBy  int    `json:"deleted_by_city"`
}

type MpCityDataBreakDownRequest struct {
	Name       string  `json:"name"`
	ProvinceID int     `json:"provinsi_id"`
	OldID      int     `json:"id"`
	NationID   int     `json:"id_nation"`
	LatCity    float64 `json:"latitude"`
	LongCity   float64 `json:"longitude"`
}
