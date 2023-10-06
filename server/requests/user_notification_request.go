package requests

type UserNotificationRequest struct {
	UserID    string `json:"user_id"`
	RowID     string `json:"row_id"`
	Type      string `json:"type_notification"`
	Text      string `json:"notification_text"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_date"`
	UpdatedAt string `json:"modified_date"`
	DeletedAt string `json:"deleted_date"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"modified_by"`
	DeletedBy int    `json:"deleted_by"`
}
