package requests

// CustomerRequest ...
type TransactionVARequest struct {
	ID            string `json:"va_id"`
	InvoiceCode   string `json:"invoice_code"`
	VACode        string `json:"va_code"`
	Amount        string `json:"amount"`
	VaPairID      string `json:"va_pair_id"`
	VaRef1        string `json:"va_ref1"`
	VaRef2        string `json:"va_ref2"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	VAPartnerCode string `json:"va_partner_code"`
}
