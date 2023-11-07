package models

import "database/sql"

// ItemOldPrice ...
type ItemOldPrice struct {
	ID           string         `json:"id"`
	CustomerID   string         `json:"customer_id"`
	CustomerCode string         `json:"customer_code"`
	CustomerName string         `json:"customer_name"`
	ItemID       string         `json:"item_id"`
	PriceListID  string         `json:"price_list_id"`
	SellPrice    string         `json:"sell_price"`
	Quantity     string         `json:"qty"`
	PreservedQty string         `json:"preserved_qty"`
	InvoiceQty   string         `json:"invoice_qty"`
	StartDate    string         `json:"start_date"`
	EndDate      string         `json:"end_date"`
	CreatedAt    sql.NullString `json:"created_at"`
	UpdatedAt    sql.NullString `json:"updated_at"`
	DeletedAt    sql.NullString `json:"deleted_at"`

	ItemCode    string         `json:"item_code"`
	ItemName    string         `json:"item_name"`
	ItemPicture sql.NullString `json:"item_picture"`
	UomID       string         `json:"uom_id"`
	UomName     string         `json:"uom_name"`
	Price       string         `json:"price"`
}

// ItemOldPriceParameter ...
type ItemOldPriceParameter struct {
	ID                 string `json:"id"`
	ItemID             string `json:"item_id"`
	ItemCode           string `json:"item_code"`
	ItemName           string `json:"item_name"`
	UomID              string `json:"uom_id"`
	UomName            string `json:"uom_name"`
	CustomerID         string `json:"customer_id"`
	PriceListVersionID string `json:"price_list_version_id"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
}

var (
	// ItemOldPriceOrderBy ...
	ItemOldPriceOrderBy = []string{"def.id", "def.created_at"}
	// ItemOldPriceOrderByrByString ...
	ItemOldPriceOrderByrByString = []string{
		"C.CUSTOMER_NAME",
	}

	ItemOldPriceSelectStatement = `
	SELECT DEF.ID, 
		DEF.CUSTOMER_ID, C.CUSTOMER_CODE, C.CUSTOMER_NAME,
		DEF.ITEM_ID, I.CODE, I._NAME, I.ITEM_PICTURE,
		DEF.PRICE_LIST_ID, DEF.SELL_PRICE, DEF.QTY, 
		DEF.UOM_ID, U._NAME,
		DEF.PRESERVED_QTY, DEF.INVOICED_QTY, DEF.START_DATE, DEF.END_DATE,
		DEF.CREATED_AT, DEF.UPDATED_AT, DEF.DELETED_AT
	FROM item_old_price def
	LEFT JOIN ITEM I ON I.ID = DEF.ITEM_ID
	LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
	LEFT JOIN UOM U ON U.ID = DEF.UOM_ID
	`

	// ItemOldPriceWhereStatement ...
	ItemOldPriceWhereStatement = ` WHERE DEF.DELETED_AT IS NULL `
)
