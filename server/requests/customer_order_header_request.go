package requests

type CustomerOrderHeaderRequest struct {
	ID                   string `json:"id_customer_order_header"`
	DocumentNo           string `json:"document_no"`
	TransactionDate      string `json:"transaction_date"`
	TransactionTime      string `json:"transaction_time"`
	CustomerID           string `json:"customer_id"`
	CustomerName         string `json:"customer_name"`
	TaxCalcMethod        string `json:"tax_calc_method"`
	SalesmanID           string `json:"salesman_id"`
	SalesmanName         string `json:"salesman_name"`
	PaymentTermsID       string `json:"payment_terms_id"`
	PaymentTermsName     string `json:"payment_terms_name"`
	ExpectedDeliveryDate string `json:"expected_delivery_date"`
	BranchID             string `json:"branch_id"`
	BranchName           string `json:"branch_name"`
	PriceLIstID          string `json:"price_list_id"`
	PriceLIstName        string `json:"price_list_name"`
	PriceLIstVersionID   string `json:"price_list_version_id"`
	PriceLIstVersionName string `json:"price_list_version_name"`
	Status               string `json:"status"`
	GrossAmount          string `json:"gross_amount"`
	TaxableAmount        string `json:"taxable_amount"`
	TaxAmount            string `json:"tax_amount"`
	RoundingAmount       string `json:"rounding_amount"`
	NetAmount            string `json:"net_amount"`
	DiscAmount           string `json:"disc_amount"`
	LineList             string `json:"line_list"`
	VoucherRedeemID      string `json:"voucher_redeem_id"`
	OldPriceID           string `json:"old_price_id"`
	OldPriceQuantity     string `json:"old_price_qty"`
	CouponRedeemID       string `json:"coupon_redeem_id"`
	PointPromo           string `json:"point_promo"`
}
