package viewmodel

type OmzetValueVM struct {
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
	TotalQuantity   *string `json:"total_quantity"`
	TotalOmzet      *string `json:"total_omzet"`
}

type OmzetValueByRegionVM struct {
	TotalOmzet    *string            `json:"total_omzet"`
	TotalQuantity *string            `json:"total_quantity"`
	Area          []OmzetValueAreaVM `json:"area"`
}

type OmzetValueAreaVM struct {
	RegionID        *string `json:"region_id"`
	RegionName      *string `json:"region_name"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
}

type OmzetValueBranchVM struct {
	RegionID        *string `json:"region_id"`
	RegionName      *string `json:"region_name"`
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
	BranchID        *string `json:"branch_id"`
	BranchName      *string `json:"branch_name"`
	BranchCode      *string `json:"branch_code"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
}
