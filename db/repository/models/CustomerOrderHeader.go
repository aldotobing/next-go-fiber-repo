package models

// CustomerOrderHeader ...
type CustomerOrderHeader struct {
	ID                   *string             `json:"id_customer_order_header"`
	DocumentNo           *string             `json:"document_no"`
	TransactionDate      *string             `json:"transaction_date"`
	TransactionTime      *string             `json:"transaction_time"`
	CustomerID           *string             `json:"customer_id"`
	CustomerName         *string             `json:"customer_name"`
	TaxCalcMethod        *string             `json:"tax_calc_method"`
	SalesmanID           *string             `json:"salesman_id"`
	SalesmanName         *string             `json:"salesman_name"`
	PaymentTermsID       *string             `json:"payment_terms_id"`
	PaymentTermsName     *string             `json:"payment_terms_name"`
	ExpectedDeliveryDate *string             `json:"expected_delivery_date"`
	BranchID             *string             `json:"branch_id"`
	BranchName           *string             `json:"branch_name"`
	PriceLIstID          *string             `json:"price_list_id"`
	PriceLIstName        *string             `json:"price_list_name"`
	PriceLIstVersionID   *string             `json:"price_list_version_id"`
	PriceLIstVersionName *string             `json:"price_list_version_name"`
	Status               *string             `json:"status"`
	GrossAmount          *string             `json:"gross_amount"`
	TaxableAmount        *string             `json:"taxable_amount"`
	TaxAmount            *string             `json:"tax_amount"`
	RoundingAmount       *string             `json:"rounding_amount"`
	NetAmount            *string             `json:"net_amount"`
	DiscAmount           *string             `json:"disc_amount"`
	LineList             *string             `json:"line_list"`
	ListLine             []CustomerOrderLine `json:"list_line"`
	CustomerCode         *string             `json:"customer_code"`
	SalesmanCode         *string             `json:"salesman_code"`
	CustomerAddress      *string             `json:"customer_address"`
	ModifiedDate         *string             `json:"modified_date"`
	VoidReasonCode       *string             `json:"reason_code"`
	VoidReasonID         *string             `json:"reason_id"`
	VoidReasonText       *string             `json:"reason_text"`
	OrderSource          *string             `json:"order_source"`
	GlobalDiscAmount     *string             `json:"global_disc_amount"`
	OldPriceData         string              `json:"old_price_data"`
	PointPromo           string              `json:"point_promo"`
}

// CustomerOrderHeaderParameter ...
type CustomerOrderHeaderParameter struct {
	ID         string `json:"id_customer_order_header"`
	DocumentNo string `json:"document_no"`
	Status     string `json:"status"`
	UserID     string `json:"admin_user_id"`
	CustomerID string `json:"id_customer"`
	Search     string `json:"search"`
	DateParam  string `json:"date_param"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

var (
	// CustomerOrderHeaderOrderBy ...
	CustomerOrderHeaderOrderBy = []string{"def.id", "def.created_date"}
	// CustomerOrderHeaderOrderByrByString ...
	CustomerOrderHeaderOrderByrByString = []string{
		"def.id",
	}

	// CustomerOrderHeaderSelectStatement ...
	CustomerOrderHeaderSelectStatement = ` select 
	def.id as id_customer_order, def.document_no,to_char(def.transaction_date,'YYYY-MM-DD') as transaction_date ,to_char(def.transaction_time,'HH24:MI:SS') as transaction_time,
	def.cust_ship_to_id,cus.customer_name, def.tax_calc_method,
	cus.salesman_id, s.salesman_name, def.payment_terms_id,top._name as top_name,
	to_char(def.expected_delivery_date,'YYYY-MM-DD') as expected_d_date,b.id as b_id,b._name as b_name,
	pl.id as pl_id,pl._name as pl_name, plv.id as plv_id,plv.description,
	def.status,def.gross_amount,def.taxable_amount, def.tax_amount,
	def.rounding_amount, def.net_amount,def.disc_amount,
	cus.customer_code as c_code, s.salesman_code as s_code, cus.customer_address,to_char(def.modified_date,'YYYY-MM-DD') as modified_date,
	mtp._name as void_reason, 1 as order_source,
	coalesce(def.global_disc_amount,0)
	from customer_order_header def
	join customer cus on cus.id = def.cust_ship_to_id
	left join salesman s on s.id = cus.salesman_id
	left join term_of_payment top on top.id = def.payment_terms_id
	left join branch b on b.id = def.branch_id
	left join price_list pl on pl.id = def.price_list_id
	left join price_list_version plv on plv.id = def.price_list_version_id
	left join master_type mtp on mtp.id = def.void_reason_id and mtp._header ='Void Reason'
 	
	`

	CustomerOrderHeaderSFASelectStatement = ` select 
	def.id as id_customer_order, def.document_no,to_char(def.transaction_date,'YYYY-MM-DD') as transaction_date ,to_char(def.transaction_time,'HH:MI:SS') as transaction_time,
	def.cust_ship_to_id,cus.customer_name, def.tax_calc_method,
	cus.salesman_id, s.salesman_name, def.payment_terms_id,top._name as top_name,
	to_char(def.expected_delivery_date,'YYYY-MM-DD') as expected_d_date,b.id as b_id,b._name as b_name,
	pl.id as pl_id,pl._name as pl_name, plv.id as plv_id,plv.description,
	def.status,def.gross_amount,def.taxable_amount, def.tax_amount,
	def.rounding_amount, def.net_amount,def.disc_amount,
	cus.customer_code as c_code, s.salesman_code as s_code, cus.customer_address,to_char(def.modified_date,'YYYY-MM-DD') as modified_date,
	def.void_reason_notes as void_reason , 2 as order_source,
	coalesce(def.global_disc_amount,0)
	from sales_order_header def
	join customer cus on cus.id = def.cust_ship_to_id
	left join salesman s on s.id = cus.salesman_id
	left join term_of_payment top on top.id = def.payment_terms_id
	left join branch b on b.id = def.branch_id
	left join price_list pl on pl.id = def.price_list_id
	left join price_list_version plv on plv.id = def.price_list_version_id
 	
	`

	// CustomerOrderHeaderWhereStatement ...
	CustomerOrderHeaderWhereStatement = `WHERE def.created_date IS not NULL`
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
