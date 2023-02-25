package models

// ShoppingCart ...
type ShoppingCart struct {
	ID               *string `json:"shopping_cart_id"`
	ItemID           *string `json:"item_id"`
	ItemName         *string `json:"item_name"`
	ItemCategoryID   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	CustomerID       *string `json:"customer_id"`
	CustomerName     *string `json:"customer_name"`
	UomID            *string `json:"uom_id"`
	UomName          *string `json:"uom_name"`
	Price            *string `json:"item_price"`
	CreatedBy        *string `json:"created_by"`
	Qty              *string `json:"qty"`
	StockQty         *string `json:"stock_qty"`
	CreatedAt        *string `json:"created_at"`
	ModifiedBy       *string `json:"modified_by"`
	ModifiedAt       *string `json:"modified_at"`
	ItemPicture      *string `json:"img_source"`
	TotalPrice       *string `json:"total_price"`
}

// ShoppingCartTotalPrice ...
type ShoppingCheckouAble struct {
	IsAble   *string `json:"is_able"`
	MinOmzet *string `json:"min_omzet"`
}

type GroupedShoppingCart struct {
	CategoryID        string         `json:"category_id"`
	CategoryName      string         `json:"CategoryName"`
	ListShoppingChart []ShoppingCart `json:"list_sopping_cart"`
}

// ShoppingCartParameter ...
type ShoppingCartParameter struct {
	ID             string `json:"shopping_cart_id"`
	ListLine       string `json:"list_line"`
	CustomerID     string `json:"customer_id"`
	ItemCategoryID string `json:"item_category_id"`
	ItemID         string `json:"item_id"`
	ItemName       string `json:"item_name"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// ShoppingCartOrderBy ...
	ShoppingCartOrderBy = []string{"def.id", "def.item_id", "def.created_date"}
	// ShoppingCartOrderByrByString ...
	ShoppingCartOrderByrByString = []string{
		"def.item_id",
	}

	// ShoppingCartSelectStatement ...
	ShoppingCartSelectStatement = `select def.id, cus.id as c_id,cus.customer_name,
	it.id as i_id,it._name as i_name, uom.id as uo_id, uom._name as uo_name,
	def.qty::integer, def.stock_qty::integer,
	def.price, it.item_category_id, ic._name as cat_name, it.item_picture,def.total_price::bigint
	from cart def
	join customer cus on cus.id = def.customer_id
	join item it on it.id = def.item_id
	join uom uom on uom.id = def.uom_id
	join item_category ic on ic.id=  it.item_category_id
	`

	GroupedShoppingCartSelectStatement = `select it.item_category_id, ic._name as cat_name
	from cart def
	join customer cus on cus.id = def.customer_id
	join item it on it.id = def.item_id
	join uom uom on uom.id = def.uom_id
	join item_category ic on ic.id=  it.item_category_id
	`

	// ShoppingCartWhereStatement ...
	ShoppingCartWhereStatement = ` WHERE def.created_date IS not NULL `
)
