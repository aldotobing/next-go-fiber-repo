package models

// CilentInvoiceLine ...
type CilentInvoiceLine struct {
	ID               *string `json:"id_line"`
	HeaderID         *string `json:"header_id"`
	LineNo           *string `json:"line_no"`
	CategoryID       *string `json:"category_id"`
	ItemID           *string `json:"item_id"`
	Qty              *string `json:"qty"`
	UomID            *string `json:"uom_id"`
	StockQty         *string `json:"stock_qty"`
	UnitPrice        *string `json:"unit_price"`
	GrossAmount      *string `json:"gross_amount"`
	UseDiscAmount    *string `json:"use_disc_percent"`
	DiscPercent1     *string `json:"disc_percent1"`
	DiscPercent2     *string `json:"disc_percent2"`
	DiscPercent3     *string `json:"disc_percent3"`
	DiscPercent4     *string `json:"disc_percent4"`
	DiscPercent5     *string `json:"disc_percent5"`
	DiscountAmount   *string `json:"disc_amount"`
	TaxableAmount    *string `json:"taxable_amount"`
	TaxAmount        *string `json:"tax_amount"`
	RoundingAmount   *string `json:"rounding_amount"`
	NetAmount        *string `json:"net_amount"`
	SalesOrderLineID *string `json:"sales_order_line_id"`
	Debt             *string `json:"debt"`
	Paid             *string `json:"paid"`
}

// CilentInvoiceLineParameter ...
type CilentInvoiceLineParameter struct {
	ID         string `json:"id_customer_order_header"`
	CustomerID string `json:"id_customer"`
	Search     string `json:"search"`
	DateParam  string `json:"date_param"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// CilentInvoiceLineOrderBy ...
	CilentInvoiceLineOrderBy = []string{"def.id", "def.created_date"}
	// CilentInvoiceLineOrderByrByString ...
	CilentInvoiceLineOrderByrByString = []string{
		"def.id",
	}

	// CilentInvoiceLineSelectStatement ...
	CilentInvoiceLineSelectStatement = ` select 
	def.id as id_customer_order, def.document_no,to_char(def.transaction_date,'YYYY-MM-DD') as transaction_date ,to_char(def.transaction_time,'HH:MI:SS') as transaction_time,
	 def.cust_ship_to_id,cus.customer_name, def.tax_calc_method,
	def.salesman_id, s.salesman_name, def.payment_terms_id,top._name as top_name,
	to_char(def.expected_delivery_date,'YYYY-MM-DD') as expected_d_date,b.id as b_id,b._name as b_name,
	pl.id as pl_id,pl._name as pl_name, plv.id as plv_id,plv.description,
	def.status,def.gross_amount,def.taxable_amount, def.tax_amount,
	def.rounding_amount, def.net_amount,def.disc_amount
	from customer_order_header def
	join customer cus on cus.id = def.cust_ship_to_id
	left join salesman s on s.id = def.salesman_id
	left join term_of_payment top on top.id = def.payment_terms_id
	left join branch b on b.id = def.branch_id
	left join price_list pl on pl.id = def.price_list_id
	left join price_list_version plv on plv.id = def.price_list_version_id
 	
	`

	// CilentInvoiceLineWhereStatement ...
	CilentInvoiceLineWhereStatement = `WHERE def.created_date IS not NULL`
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
