package viewmodel

// VoucherVM ....
type VoucherRedeemVM struct {
	ID                   string `json:"id"`
	CustomerCode         string `json:"customer_code"`
	Redeemed             string `json:"redeemed"`
	RedeemedAt           string `json:"redeemed_at"`
	RedeemedToDocumentNo string `json:"redeemed_to_doc_no"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
	VoucherID            string `json:"voucher_id"`
	VoucherName          string `json:"voucher_name"`
	VoucherCashValue     string `json:"voucher_cash_value"`
	VoucherDescription   string `json:"voucher_description"`
	VoucherImageURL      string `json:"voucher_image_url"`
	VoucherStartDate     string `json:"voucher_start_date"`
	VoucherEndDate       string `json:"voucher_end_date"`
}
