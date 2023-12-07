package models

// CilentReturnInvoiceLine ...
type CilentReturnInvoiceLine struct {
	ID                 *string `json:"id_line"`
	HeaderID           *string `json:"header_id"`
	LineNo             *string `json:"line_no"`
	CategoryID         *string `json:"category_id"`
	ItemID             *string `json:"item_id"`
	ItemCode           *string `json:"item_code"`
	Qty                *string `json:"qty"`
	UomID              *string `json:"uom_id"`
	UomCode            *string `json:"uom_code"`
	StockQty           *string `json:"stock_qty"`
	UnitPrice          *string `json:"unit_price"`
	GrossAmount        *string `json:"gross_amount"`
	UseDiscAmount      *string `json:"use_disc_percent"`
	DiscPercent1       *string `json:"disc_percent1"`
	DiscPercent2       *string `json:"disc_percent2"`
	DiscPercent3       *string `json:"disc_percent3"`
	DiscPercent4       *string `json:"disc_percent4"`
	DiscPercent5       *string `json:"disc_percent5"`
	DiscountAmount     *string `json:"disc_amount"`
	TaxableAmount      *string `json:"taxable_amount"`
	TaxAmount          *string `json:"tax_amount"`
	RoundingAmount     *string `json:"rounding_amount"`
	NetAmount          *string `json:"net_amount"`
	StockQtyReplacment *string `json:"stock_qty_replacement"`
	UomCodeReplace     *string `json:"uom_code_replace"`
	UomBackOffice      *string `json:"uom_backoffice"`
}

// CilentReturnInvoiceLineParameter ...
type CilentReturnInvoiceLineParameter struct {
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
	// CilentReturnInvoiceLineOrderBy ...
	CilentReturnInvoiceLineOrderBy = []string{"def.id", "def.created_date"}
	// CilentReturnInvoiceLineOrderByrByString ...
	CilentReturnInvoiceLineOrderByrByString = []string{
		"def.id",
	}

	// CilentReturnInvoiceLineSelectStatement ...
	CilentReturnInvoiceLineSelectStatement = `  	
	`

	// CilentReturnInvoiceLineWhereStatement ...
	CilentReturnInvoiceLineWhereStatement = `WHERE def.created_date IS not NULL`
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
