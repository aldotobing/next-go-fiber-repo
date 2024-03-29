package models

// CilentReturnInvoice ...
type CilentReturnInvoice struct {
	ID                   *string                    `json:"id_invoice_header"`
	DocumentNo           *string                    `json:"document_no"`
	DocumentTypeID       *string                    `json:"document_type_id"`
	TransactionDate      *string                    `json:"transaction_date"`
	TransactionTime      *string                    `json:"transaction_time"`
	CustomerID           *string                    `json:"customer_id"`
	CustomerCode         *string                    `json:"customer_code"`
	CustomerName         *string                    `json:"customer_name"`
	TaxCalcMethod        *string                    `json:"tax_calc_method"`
	SalesmanID           *string                    `json:"salesman_id"`
	SalesmanCode         *string                    `json:"salesman_code"`
	ReturnType           *string                    `json:"return_type"`
	Notes                *string                    `json:"notes"`
	SalesmanName         *string                    `json:"salesman_name"`
	PaymentTermsID       *string                    `json:"payment_terms_id"`
	PaymentTermsName     *string                    `json:"payment_terms_name"`
	CompanyID            *string                    `json:"company_id"`
	BranchID             *string                    `json:"branch_id"`
	BranchName           *string                    `json:"branch_name"`
	PriceLIstID          *string                    `json:"price_list_id"`
	PriceLIstName        *string                    `json:"price_list_name"`
	PriceLIstVersionID   *string                    `json:"price_list_version_id"`
	PriceLIstVersionName *string                    `json:"price_list_version_name"`
	Status               *string                    `json:"status"`
	GrossAmount          *string                    `json:"gross_amount"`
	TaxableAmount        *string                    `json:"taxable_amount"`
	TaxAmount            *string                    `json:"tax_amount"`
	RoundingAmount       *string                    `json:"rounding_amount"`
	NetAmount            *string                    `json:"net_amount"`
	DiscAmount           *string                    `json:"disc_amount"`
	GlobalDiscAmount     *string                    `json:"global_disc_amount"`
	ListLine             *[]CilentReturnInvoiceLine `json:"list_line"`
	RefInvoice           *string                    `json:"ref_invoice"`
}

// CilentReturnInvoiceParameter ...
type CilentReturnInvoiceParameter struct {
	ID         string `json:"id_customer_order_header"`
	DocumentNo string `json:"document_no"`
	CustomerID string `json:"id_customer"`
	Search     string `json:"search"`
	DateParam  string `json:"date_param"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// CilentReturnInvoiceOrderBy ...
	CilentReturnInvoiceOrderBy = []string{"def.id", "def.created_date"}
	// CilentReturnInvoiceOrderByrByString ...
	CilentReturnInvoiceOrderByrByString = []string{
		"def.id",
	}

	// CilentReturnInvoiceSelectStatement ...
	CilentReturnInvoiceSelectStatement = ` select 
	def.id as sih_id,document_no,def.cust_bill_to_id
	from sales_return_invoice_header def
 	
	`

	// CilentReturnInvoiceWhereStatement ...
	CilentReturnInvoiceWhereStatement = `WHERE def.created_date IS not NULL `
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
