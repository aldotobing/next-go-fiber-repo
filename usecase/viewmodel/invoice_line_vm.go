package viewmodel

type InvoiceLineVM struct {
	ItemName  string `json:"item_name"`
	UnitPrice string `json:"unit_price"`
	UomName   string `json:"uomname"`
	Quantity  string `json:"qty"`
}
