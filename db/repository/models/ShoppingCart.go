package models

// ShoppingCart ...
type ShoppingCart struct {
	ID           *string `json:"shopping_cart_id"`
	ItemID       *string `json:"item_id"`
	ItemName     *string `json:"item_name"`
	CustomerID   *string `json:"customer_id"`
	CustomerName *string `json:"customer_name"`
	UomID        *string `json:"uom_id"`
	UomName      *string `json:"uom_name"`
	Price        *string `json:"item_price"`
	CreatedBy    *string `json:"created_by"`
	Qty          *string `json:"qty"`
	StockQty     *string `json:"stock_qty"`
	CreatedAt    *string `json:"created_at"`
	ModifiedBy   *string `json:"modified_by"`
	ModifiedAt   *string `json:"modified_at"`
}

// ShoppingCartParameter ...
type ShoppingCartParameter struct {
	ID         string `json:"shopping_cart_id"`
	CustomerID string `json:"customer_id"`
	ItemID     string `json:"item_id"`
	ItemName   string `json:"item_name"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// ShoppingCartOrderBy ...
	ShoppingCartOrderBy = []string{"def.id", "def.item_id", "def.created_date"}
	// ShoppingCartOrderByrByString ...
	ShoppingCartOrderByrByString = []string{
		"def.item_id",
	}

	// ShoppingCartSelectStatement ...
	ShoppingCartSelectStatement = `select def.id, cus.id as c_id,p._name,
	it.id as i_id,it._name as i_name, uom.id as uo_id, uom._name as uo_name,
	def.qty, def.stock_qty,
	def.price
	from cart def
	join customer cus on cus.id = def.customer_id
	join partner p on p.id =cus.partner_id
	join item it on it.id = def.item_id
	join uom uom on uom.id = def.uom_id
	`

	// ShoppingCartWhereStatement ...
	ShoppingCartWhereStatement = ` WHERE def.created_date IS not NULL `
)
