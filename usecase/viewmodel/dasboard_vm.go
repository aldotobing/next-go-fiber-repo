package viewmodel

type OmzetValueVM struct {
	RegionID      *string `json:"region_id"`
	RegionName    *string `json:"region_name"`
	TotalQuantity *string `json:"total_quantity"`
	TotalOmzet    *string `json:"total_omzet"`
}

type OmzetValueByRegionVM struct {
	TotalOmzet    *string            `json:"total_omzet"`
	TotalQuantity *string            `json:"total_quantity"`
	Area          []OmzetValueAreaVM `json:"area"`
}

type OmzetValueAreaVM struct {
	ID        *string `json:"id"`
	Name      *string `json:"_name"`
	Quantity  *string `json:"quantity"`
	Omzet     *string `json:"omzet"`
	GroupID   *string `json:"group_id"`
	GroupName *string `json:"group_name"`
}
