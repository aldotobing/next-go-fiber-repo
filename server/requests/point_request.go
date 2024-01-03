package requests

// PointRequest ...
type PointRequest struct {
	PointType     string              `json:"point_type" validate:"required"`
	InvoiceID     string              `json:"invoice_id" validate:"required"`
	Point         string              `json:"point" validate:"required"`
	CustomerID    string              `json:"customer_id"`
	CustomerCodes []PointCustomerCode `json:"customer_codes" validate:"required"`
}

type PointCustomerCode struct {
	CustomerCode string `json:"customer_code"`
}
