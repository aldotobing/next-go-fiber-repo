package viewmodel

type OmzetValueVM struct {
	RegionID      *string `json:"region_id"`
	RegionName    *string `json:"region_name"`
	TotalQuantity *string `json:"total_quantity"`
	TotalOmzet    *string `json:"total_omzet"`
}
