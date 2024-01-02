package viewmodel

// PointVM ....
type PointVM struct {
	ID                string `json:"id"`
	PointType         string `json:"point_type"`
	PointTypeName     string `json:"point_type_name"`
	InvoiceID         string `json:"invoice_id"`
	InvoiceDocumentNo string `json:"invoice_document_no"`
	Point             string `json:"point"`
	CustomerID        string `json:"customer_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	ExpiredAt         string `json:"expired_at"`

	DetailCustomer CustomerVM `json:"customer_detail,omitempty"`
	CustomerIDs    []string   `json:"customer_ids,omitempty"`
}

type PointBalanceVM struct {
	Balance string `json:"balance"`
}
