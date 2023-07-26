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
	Customername  *string `json:"customer_name"`
}

// DoctorParameter ...
type TransactionVAParameter struct {
	ID            string `json:"partner_id"`
	InvoiceCode   string `json:"invoice_code"`
	VACode        string `json:"va_code"`
	UserId        string `json:"admin_user_id"`
	Search        string `json:"search"`
	CurrentVaUser int    `json:"current_va_user"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	By            string `json:"by"`
	Sort          string `json:"sort"`
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
	select def.id,def.invoice_code, def.va_code, ROUND(def.amount) as _amount,def.va_pair_id,
	def.va_ref1,def.va_ref2, def.start_date, def.end_date, def.va_partner_code, def.paid_status,
	c.customer_name as cus_name
		from virtual_account_transaction def
		join sales_invoice_header sih on sih.document_no = def.invoice_code
		left join customer c on sih.cust_bill_to_id = c.id
	`

	// CustomerWhereStatement ...
	TransactionVAWhereStatement = `where def.created_date is not null  `
)
