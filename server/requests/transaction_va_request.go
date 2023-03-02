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

type InquiryBodyRequest struct {
	Language        string `json:"language"`
	TransactionDate string `json:"trxDateTime"`
	TransmisionDate string `json:"transmissionDateTime"`
	CompanyCode     string `json:"companyCode"`
	VaReChanelId    string `json:"channelID"`
	Billkey1        string `json:"billKey1"` //code va yang diambil di next
	Billkey2        string `json:"billKey2"`
	Billkey3        string `json:"billKey3"`
	Reference1      string `json:"reference1"`
	Reference2      string `json:"reference2"`
	Reference3      string `json:"reference3"`
	PaymentAmount   string `json:"paymentAmount"`
	Currency        string `json:"currency"`
	TransactionID   string `json:"transactionID"`
}

type InquiryVaRequest struct {
	InquiryBody InquiryBodyRequest `json:"InquiryRequest"`
}
