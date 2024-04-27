package viewmodel

// PointRuleVM ....
type PointRuleVM struct {
	ID              string `json:"id"`
	StartDate       string `json:"start_dates"`
	EndDate         string `json:"end_date"`
	MinOrder        string `json:"min_order"`
	PointConversion string `json:"point_conversion"`
	MonthlyMaxPoint string `json:"monthly_max_point"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`

	Customers []PointRuleCustomerVM `json:"customers"`
}

type PointRuleCustomerVM struct {
	CustomerCode string `json:"customer_code"`
	Value        string `json:"value"`
}
