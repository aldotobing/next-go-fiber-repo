package models

type UserNotificationDetailParameter struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	RowID  string `json:"row_id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}
