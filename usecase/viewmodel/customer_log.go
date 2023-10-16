package viewmodel

type CustomerLogVM struct {
	ID                string `json:"id"`
	CustomerID        string `json:"customer_id"`
	CustomerCode      string `json:"customer_code"`
	CustomerName      string `json:"customer_name"`
	TypeChanges       string `json:"type_changes"`
	OldData           string `json:"old_data"`
	NewData           string `json:"new_data"`
	UserID            string `json:"user_id"`
	UserName          string `json:"user_name"`
	CreatedAt         string `json:"created_at"`
	BranchID          string `json:"branch_id"`
	BranchName        string `json:"branch_name"`
	RegionID          string `json:"region_id"`
	RegionName        string `json:"region_name"`
	RegionGroupID     string `json:"region_group_id"`
	RegionGroupName   string `json:"region_group_name"`
	CustomerLevelID   string `json:"customer_level_id"`
	CustomerLevelName string `json:"customer_level_name"`
}
