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

type OmzetValueByBranchVM struct {
	TotalOmzet    *string              `json:"total_omzet"`
	TotalQuantity *string              `json:"total_quantity"`
	Branches      []OmzetValueBranchVM `json:"branches"`
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

type OmzetValueByCustomerVM struct {
	TotalOmzet    *string                `json:"total_omzet"`
	TotalQuantity *string                `json:"total_quantity"`
	Customers     []OmzetValueCustomerVM `json:"customers"`
}

type OmzetValueCustomerVM struct {
	RegionGroupName *string `json:"region_group_name"`
	RegionName      *string `json:"region_name"`
	BranchID        *string `json:"branch_id"`
	BranchClass     *string `json:"branch_class"`
	BranchName      *string `json:"branch_name"`
	BranchCode      *string `json:"branch_code"`
	CustomerID      *string `json:"customer_id"`
	CustomerCode    *string `json:"customer_code"`
	CustomerName    *string `json:"customer_name"`
	CustomerType    *string `json:"customer_type"`
	ProvinceName    *string `json:"customer_province_name"`
	CityName        *string `json:"customer_city_name"`
	CustomerLevel   *string `json:"customer_level"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
}
