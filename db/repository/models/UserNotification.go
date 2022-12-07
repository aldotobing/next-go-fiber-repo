package models

type UserNotification struct {
	ID        *string `json:"id"`
	UserID    *string `json:"user_id"`
	RowID     *string `json:"row_id"`
	Type      *string `json:"type_notification"`
	Text      *string `json:"notification_text"`
	CreatedAt *string `json:"created_date"`
	UpdatedAt *string `json:"modified_date"`
	DeletedAt *string `json:"deleted_date"`
	CreatedBy *int    `json:"created_by"`
	UpdatedBy *int    `json:"modified_by"`
	DeletedBy *int    `json:"deleted_by"`
	Status    *string `json:"notification_status"`
	Title     *string `json:"notification_title"`
}

type UserNotificationParameter struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	UserNotificationOrderBy = []string{"def.id", "def.created_date", "def.modified_date"}

	UserNotificationByrByString = []string{
		"def.id",
	}

	UserNotificationSelectStatement = ` select 
	def.id, def.user_id, def.row_id, def.type_notification, def.notification_text,
	def.created_date, def.modified_date,def.deleted_date,
	def.created_by, def.modified_by, def.deleted_by,
	coalesce(def.notification_status,'un read'), def.notification_title
	from user_notification def `

	UserNotificationWhereStatement = ` WHERE def.deleted_date IS NULL `
)
