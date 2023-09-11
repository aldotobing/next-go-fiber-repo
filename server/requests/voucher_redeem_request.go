package requests

// VoucherRedeemRequest ...
type VoucherRedeemRequest struct {
	CustomerID         string `json:"customer_id"`
	Redeem             string `json:"redeem"`
	RedeemToDocumentNo string `json:"redeem_to_doc_no"`
	VoucherID          string `json:"voucher_id"`
}

// VoucherRedeemBulkRequest ...
type VoucherRedeemBulkRequest struct {
	VouchersRedeem []VoucherRedeemRequest `json:"vouchers_redeem"`
}
