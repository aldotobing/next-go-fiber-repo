package viewmodel

// VoucherVM ....
type VoucherRedeemVM struct {
	ID                   string `json:"id"`
	CustomerID           string `json:"customer_id"`
	Redeemed             string `json:"redeemed"`
	RedeemedAt           string `json:"redeemed_at"`
	RedeemedToDocumentNo string `json:"redeemed_to_doc_no"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
	VoucherID            string `json:"voucher_id"`
	VoucherName          string `json:"voucher_name"`
	VoucherCashValue     string `json"voucher_cash_value"`
}
