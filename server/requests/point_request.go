package requests

// PointRequest ...
type PointRequest struct {
	PointType         string              `json:"point_type" validate:"required"`
	InvoiceDocumentNo string              `json:"InvoiceDocumentNo"`
	UserID            string              `json:"user_id"`
	Point             string              `json:"point"`
	CustomerID        string              `json:"customer_id"`
	Note              string              `json:"note"`
	CustomerCodes     []PointCustomerCode `json:"customer_codes" validate:"required"`
}

type PointCustomerCode struct {
	Point        string `json:"point"`
	CustomerCode string `json:"customer_code"`
}
