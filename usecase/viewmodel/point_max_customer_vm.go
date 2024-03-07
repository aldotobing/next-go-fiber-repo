package viewmodel

// PointMaxCustomerVM ....
type PointMaxCustomerVM struct {
	ID              string `json:"id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	CustomerCode    string `json:"customer_code"`
	CustomerName    string `json:"customer_name"`
	MonthlyMaxPoint string `json:"monthly_max_point"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
