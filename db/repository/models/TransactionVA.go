package models

// Doctor ...
type TransactionVA struct {
	ID            *string `json:"va_id"`
	InvoiceCode   *string `json:"invoice_code"`
	VACode        *string `json:"va_code"`
	VAPartnerCode *string `json:"va_partner_code"`
	Amount        *string `json:"amount"`
	VaPairID      *string `json:"va_pair_id"`
	VaRef1        *string `json:"va_ref1"`
	VaRef2        *string `json:"va_ref2"`
	StartDate     *string `json:"start_date"`
	EndDate       *string `json:"end_date"`
	PaidStatus    *string `json:"paid_status"`
}

// DoctorParameter ...
type TransactionVAParameter struct {
	ID          string `json:"partner_id"`
	InvoiceCode string `json:"invoice_code"`
	VACode      string `json:"va_code"`
	UserId      string `json:"admin_user_id"`
	Search      string `json:"search"`
	Page        int    `json:"page"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	By          string `json:"by"`
	Sort        string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	TransactionVAOrderBy = []string{"def.id", "def.invoice_code", "def.created_date"}
	// CustomerOrderByrByString ...
	TransactionVAOrderByrByString = []string{
		"def.invoice_code",
	}

	// CustomerSelectStatement ...

	TransactionVASelectStatement = `
	select def.id,def.invoice_code, def.va_code, def.amount,def.va_pair_id,
	def.va_ref1,def.va_ref2, def.start_date, def.end_date, def.va_partner_code, def.paid_status
		from virtual_account_transaction def
	`

	// CustomerWhereStatement ...
	TransactionVAWhereStatement = `where def.created_date is not null  `
)