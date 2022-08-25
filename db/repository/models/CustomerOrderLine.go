package models

// CustomerOrderLine ...
type CustomerOrderLine struct {
	ID   *string `json:"id_CustomerOrderLine"`
	Code *string `json:"code"`
	Name *string `json:"name_CustomerOrderLine"`
}

type MpCustomerOrderLineDataBreakDown struct {
	ID                    *string  `json:"id_CustomerOrderLine"`
	Name                  *string  `json:"name_CustomerOrderLine"`
	ProvinceID            *int     `json:"id_province"`
	OldID                 *int     `json:"old_id"`
	NationID              *int     `json:"id_nation"`
	LatCustomerOrderLine  *float64 `json:"lat_CustomerOrderLine"`
	LongCustomerOrderLine *float64 `json:"long_CustomerOrderLine"`
}

// CustomerOrderLineParameter ...
type CustomerOrderLineParameter struct {
	ID         string `json:"id_CustomerOrderLine"`
	ProvinceID string `json:"id_province"`
	Name       string `json:"name_CustomerOrderLine"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// CustomerOrderLineOrderBy ...
	CustomerOrderLineOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// CustomerOrderLineOrderByrByString ...
	CustomerOrderLineOrderByrByString = []string{
		"def._name",
	}

	// CustomerOrderLineSelectStatement ...
	CustomerOrderLineSelectStatement = `SELECT def.id,def.code,  def._name
	FROM CustomerOrderLine def
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
