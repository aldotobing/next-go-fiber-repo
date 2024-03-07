package requests

// PointMaxCustomerRequest ...
type PointMaxCustomerRequest struct {
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	CustomerCode    string `json:"customer_code"`
	MonthlyMaxPoint string `json:"monthly_max_point"`
}

type PointMaxCustomerRequestHeader struct {
	Detail []PointMaxCustomerRequest `json:"detail"`
}
