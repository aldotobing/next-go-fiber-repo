package requests

// BroadcastRequest ...
type BroadcastRequest struct {
	Title          string `json:"title" validate:"required"`
	Body           string `json:"body" validate:"required"`
	BroadcastDate  string `json:"broadcast_date"`
	BroadcastTime  string `json:"broadcast_time"`
	BranchID       string `json:"branch_id"`
	RegionID       string `json:"region_id"`
	RegionGroupID  string `json:"region_group_id"`
	CustomerTypeID string `json:"customer_type_id"`
}
