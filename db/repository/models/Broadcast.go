package models

import "database/sql"

// Broadcast ...
type Broadcast struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	Body           string         `json:"body"`
	BroadcastDate  string         `json:"broadcast_date"`
	BroadcastTime  string         `json:"broadcast_time"`
	RepeatEveryDay bool           `json:"repeat_every_day"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      sql.NullString `json:"updated_at"`
	DeletedAt      sql.NullString `json:"deleted_at"`
	Parameter      sql.NullString `json:"parameter"`
}

// BroadcastParameter ...
type BroadcastParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	BroadcastOrderBy = []string{"b.id", "b.title"}
	// CustomerOrderByrByString ...
	BroadcastOrderByrByString = []string{
		"b.title",
	}

	// CustomerSelectStatement ...

	BroadcastSelectStatement = `SELECT 
			B.ID, 
			B.TITLE, B.BODY, 
			B.BROADCAST_DATE, 
			B.BROADCAST_TIME,
			B.REPEAT_EVERY_DAY,
			B.CREATED_AT,
			B.UPDATED_AT,
			B.DELETED_AT,
			B.PARAMETER
		FROM BROADCAST B
	`

	// CustomerWhereStatement ...
	BroadcastWhereStatement = `WHERE B.DELETED_AT IS NULL `
)
