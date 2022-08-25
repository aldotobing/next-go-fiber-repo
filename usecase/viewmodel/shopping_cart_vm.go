package viewmodel

type ShoppingCartVM struct {
	ID           string `json:"shopping_cart_id"`
	ItemID       string `json:"item_id"`
	ItemName     string `json:"item_name"`
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	UomID        string `json:"uom_id"`
	Price        string `json:"item_price"`
}
