package requests

// BroadcastRequest ...
type BroadcastRequest struct {
	Title             string   `json:"title" validate:"required"`
	Body              string   `json:"body" validate:"required"`
	BroadcastDate     string   `json:"broadcast_date"`
	BroadcastTime     string   `json:"broadcast_time"`
	BranchID          string   `json:"branch_id"`
	BranchName        string   `json:"branch_name"`
	RegionID          string   `json:"region_id"`
	RegionName        string   `json:"region_name"`
	RegionGroupID     string   `json:"region_group_id"`
	RegionGroupName   string   `json:"region_group_name"`
	CustomerTypeID    string   `json:"customer_type_id"`
	CustomerTypeName  string   `json:"customer_type_name"`
	CustomerLevelID   string   `json:"customer_level_id"`
	CustomerLevelName string   `json:"customer_level_name"`
	CustomerCode      []string `json:"customer_code"`
	RepeatEveryDay    bool     `json:"repeat_every_day"`
}
