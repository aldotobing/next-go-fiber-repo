package models

// CustomerOrderLine ...
type CustomerOrderLine struct {
	ID             *string `json:"id_customer_order_line"`
	HeaderID       *string `json:"header_id"`
	CategoryName   *string `json:"item_category_name"`
	CategoryID     *string `json:"item_category_id"`
	ItemID         *string `json:"item_id"`
	ItemName       *string `json:"item_name"`
	UomID          *string `json:"uom_id"`
	UomName        *string `json:"uom_name"`
	QTY            *string `json:"qty"`
	StockQty       *string `json:"stock_qty"`
	UnitPrice      *string `json:"unit_price"`
	GrossAmount    *string `json:"gross_amount"`
	UseDiscPercent *string `json:"use_disc_percent"`
	DisPercent1    *string `json:"disc_percent1"`
	DisPercent2    *string `json:"disc_percent2"`
	DisPercent3    *string `json:"disc_percent3"`
	DisPercent4    *string `json:"disc_percent4"`
	DisPercent5    *string `json:"disc_percent5"`
	TaxableAmount  *string `json:"taxable_amount"`
	TaxAmount      *string `json:"tax_amount"`
	RoundingAmount *string `json:"rounding_amount"`
	NetAmount      *string `json:"net_amount"`
	SalesmanName   *string `json:"salesman_name"`
	SalesmanCode   *string `json:"salesman_code"`
	ItemPicture    *string `json:"item_picture"`
	FromPromo      *string `json:"from_promo"`
}

// CustomerOrderLineParameter ...
type CustomerOrderLineParameter struct {
	ID       string `json:"id_customer_order_line"`
	HeaderID string `json:"header_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
}

var (
	// CustomerOrderLineOrderBy ...
	CustomerOrderLineOrderBy = []string{"def.id", "def.created_date"}
	// CustomerOrderLineOrderByrByString ...
	CustomerOrderLineOrderByrByString = []string{
		"def.id",
	}

	// CustomerOrderLineSelectStatement ...
	CustomerOrderLineSelectStatement = `select 
	def.id as order_line_id, def.header_id, ic._name as cat_name, ic.id as ic_id,
	i.id as item_id, i._name as i_name,uo.id as uom_id, uo._name as uom_name,
	def.qty,def.stock_qty, def.unit_price,def.gross_amount,
	def.use_disc_percent,def.disc_percent1,def.disc_percent2,def.disc_percent3,
	def.disc_percent4,def.disc_percent5, def.taxable_amount, def.tax_amount,
	def.rounding_amount, def.net_amount, s.salesman_name, s.salesman_code, i.item_picture, coalesce(def.from_promo,0)
	from customer_order_line def
	join customer_order_header coh on coh.id = def.header_id
	join customer cus on cus.id = coh.cust_ship_to_id
	join item i on i.id = def.item_id
	join item_category ic on ic.id = i.item_category_id
	join uom uo on uo.id = def.uom_id
	join salesman s on s.id =cus.salesman_id	
	
	`

	SFACustomerOrderLineSelectStatement = `select 
	def.id as order_line_id, def.header_id, ic._name as cat_name, ic.id as ic_id,
	i.id as item_id, i._name as i_name,uo.id as uom_id, uo._name as uom_name,
	def.qty,def.stock_qty, def.unit_price,def.gross_amount,
	def.use_disc_percent,def.disc_percent1,def.disc_percent2,def.disc_percent3,
	def.disc_percent4,def.disc_percent5, def.taxable_amount, def.tax_amount,
	def.rounding_amount, def.net_amount, s.salesman_name, s.salesman_code, i.item_picture, 0
	from sales_order_line def
	join sales_order_header coh on coh.id = def.header_id
	join customer cus on cus.id = coh.cust_ship_to_id
	join item i on i.id = def.item_id
	join item_category ic on ic.id = i.item_category_id
	join uom uo on uo.id = def.uom_id
	join salesman s on s.id =cus.salesman_id	
	
	`

	// CustomerOrderLineWhereStatement ...
	CustomerOrderLineWhereStatement = `WHERE def.created_date IS not NULL`
)

// insert into customer_order_header(
// 	transaction_date, transaction_time,  cust_bill_to_id, cust_ship_to_id,payment_terms_id, expected_delivery_date,
// 	gross_amount,disc_amount,taxable_amount,tax_amount,rounding_amount,net_amount,tax_calc_method
// 	)values(
// 	now(),now(),1,1,1,now(),
// 		0,0,0,0,0,0,'E'
// 	)

// insert into customer_order_line (
// 	header_id, line_no,category_id,item_id,qty,
// uom_id, stock_qty, unit_price, gross_amount, use_disc_percent,
// disc_percent1, disc_percent2, disc_percent3, disc_percent4, disc_percent5,
// disc_amount,taxable_amount,tax_amount,rounding_amount,net_amount

// )values(
// 2,1,1,1,5,
// 1,3,25000,1,1,
// 	0,0,0,0,0,
// 	1,0,0,0,100000
// )
