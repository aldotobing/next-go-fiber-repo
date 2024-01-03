package requests

// AccountOpeningRequest ...
type SalesInvoiceRequest struct {
	TotalPaid     string `json:"total_paid"`
	PaymentMethod string `json:"payment_method"`
}
