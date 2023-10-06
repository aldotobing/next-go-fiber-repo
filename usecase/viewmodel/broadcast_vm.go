package viewmodel

// BroadcastVM ....
type BroadcastVM struct {
	ID             string               `json:"id"`
	Title          string               `json:"title"`
	Body           string               `json:"body"`
	BroadcastDate  string               `json:"broadcast_date"`
	BroadcastTime  string               `json:"broadcast_time"`
	RepeatEveryDay bool                 `json:"repeat_every_day"`
	CreatedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
	DeletedAt      string               `json:"deleted_at"`
	Parameter      BroadcastParameterVM `json:"parameter"`
}

// BroadcastParameterVM ....
type BroadcastParameterVM struct {
	BranchID          string `json:"branch_id"`
	BranchName        string `json:"branch_name"`
	RegionID          string `json:"region_id"`
	RegionName        string `json:"region_name"`
	RegionGroupID     string `json:"region_group_id"`
	RegionGroupName   string `json:"region_group_name"`
	CustomerTypeID    string `json:"customer_type_id"`
	CustomerTypeName  string `json:"customer_type_name"`
	CustomerLevelID   string `json:"customer_level_id"`
	CustomerLevelName string `json:"customer_level_name"`
}
