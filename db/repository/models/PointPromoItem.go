package models

import "database/sql"

// PointPromoItem ...
type PointPromoItem struct {
	ID        string
	ItemID    string
	ItemName  string
	CreatedAt string
	UpdatedAt sql.NullString
	DeletedAt sql.NullString
}
