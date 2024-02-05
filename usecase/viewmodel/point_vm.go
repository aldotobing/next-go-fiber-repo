package viewmodel

// PointVM ....
type PointVM struct {
	ID                string `json:"id"`
	PointType         string `json:"point_type"`
	PointTypeName     string `json:"point_type_name"`
	InvoiceDocumentNo string `json:"invoice_document_no"`
	Point             string `json:"point"`
	CustomerID        string `json:"customer_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	ExpiredAt         string `json:"expired_at"`

	DetailCustomer CustomerVM `json:"customer_detail,omitempty"`
	CustomerIDs    []string   `json:"customer_ids,omitempty"`
	CustomerPoints []string   `json:"customer_points,omitempty"`
}

type PointBalanceVM struct {
	Balance string `json:"balance"`
}

// PointReportVM ....
type PointReportVM struct {
	BranchCode        string `json:"branch_code"`
	BranchName        string `json:"branch_name"`
	RegionName        string `json:"region_name"`
	RegionGroupName   string `json:"region_group_name"`
	PartnerCode       string `json:"partner_code"`
	PartnerName       string `json:"partner_name"`
	InvoiceDocumentNo string `json:"invoice_document_no"`
	NetAmount         string `json:"net_amount"`
	Point             string `json:"point"`
	TrasactionDate    string `json:"transaction_date"`
}
