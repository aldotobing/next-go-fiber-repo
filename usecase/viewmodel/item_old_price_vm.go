package viewmodel

// ItemOldPriceVM ....
type ItemOldPriceVM struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	CustomerCode string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
	ItemID       string `json:"item_id"`
	ItemCode     string `json:"item_code"`
	ItemName     string `json:"item_name"`
	UomID        string `json:"uom_id"`
	UomName      string `json:"uom_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Quantity     int    `json:"qty"`
	PriceListID  string `json:"price_list_id"`
	SellPrice    string `json:"sell_price"`
	PreservedQty string `json:"preserved_qty"`
	InvoiceQty   string `json:"invoice_qty"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
