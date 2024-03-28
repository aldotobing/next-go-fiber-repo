package viewmodel

// LottaryVM ....
type LottaryVM struct {
	ID              string `json:"id"`
	SerialNo        string `json:"serial_no"`
	Status          string `json:"status"`
	CustomerCode    string `json:"customer_code"`
	CustomerID      string `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	Year            string `json:"_year"`
	Quartal         string `json:"_quartal"`
	Sequence        string `json:"_sequence"`
	BranchName      string `json:"customer_branch_name"`
	RegionCode      string `json:"customer_region_code"`
	RegionName      string `json:"customer_region_name"`
	RegionGroup     string `json:"customer_region_group"`
	CustomerType    string `json:"customer_type_name"`
	CustomerLevel   string `json:"customer_level_name"`
	CustomerAddress string `json:"customer_address"`
	CustomerCpName  string `json:"customer_cp_name"`
}
