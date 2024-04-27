package models

import "database/sql"

const (
	PromoTypeStrata = "Strata"
	PromoTypePoint  = "Point"
)

// PointPromo ...
type PointPromo struct {
	ID                 string         `json:"id"`
	StartDate          string         `json:"start_date"`
	EndDate            string         `json:"end_date"`
	CreatedAt          string         `json:"created_at"`
	UpdatedAt          sql.NullString `json:"updated_at"`
	DeletedAt          sql.NullString `json:"deleted_at"`
	Multiplicator      bool           `json:"multiplicator"`
	PointConversion    sql.NullString `json:"poin_conversion"`
	QuantityConversion sql.NullString `json:"quantity_conversion"`
	PromoType          sql.NullString `json:"promo_type"`
	Strata             sql.NullString `json:"strata"`
	Items              string
	Image              sql.NullString
	Title              sql.NullString
	Description        sql.NullString
}

// PointPromoParameter ...
type PointPromoParameter struct {
	ID            string `json:"id"`
	Month         string `json:"month"`
	Year          string `json:"year"`
	CustomerID    string `json:"customer_id"`
	PointType     string `json:"point_type"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Renewal       string `json:"renewal"`
	Now           bool
	Search        string `json:"search"`
	ShowAll       string `json:"show_all"`
	RegionID      string `json:"region_id"`
	RegionGroupID string `json:"region_group_id"`
	BranchID      string `json:"branch_id"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	By            string `json:"by"`
	Sort          string `json:"sort"`
}

var (
	PointPromoOrderBy          = []string{"def.id", "def._name"}
	PointPromoOrderByrByString = []string{
		"def._name",
	}

	PointPromoSelectStatement = `SELECT 
			DEF.ID, 
			DEF.START_DATE,
			DEF.END_DATE,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.MULTIPLICATOR,
			DEF.POINT_CONVERSION,
			DEF.QUANTITY_CONVERSION,
			DEF.PROMO_TYPE,
			DEF.STRATA,
			array_to_string((array_agg(I.ID || '#sep#' || I._NAME || '#sep#' || I.ITEM_PICTURE || '#sep#' || PPI.UOM_ID || '#sep#' || PPI.UOM_NAME || '#sep#' || PPI.CONVERTION || '#sep#' || PPI.QTY order by PPI.ID asc)),'|') AS Items,
			DEF.IMAGE_URL,
			DEF.TITLE,
			DEF.DESCRIPTION
		FROM POINT_PROMO DEF
		LEFT JOIN POINT_PROMO_ITEM PPI ON PPI.PROMO_ID = DEF.ID and PPI.DELETED_AT IS NULL
		LEFT JOIN ITEM I ON I.ID = PPI.ITEM_ID
	`

	PointPromoWhereStatement = `WHERE DEF.DELETED_AT IS NULL `
)
