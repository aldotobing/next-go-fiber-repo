package requests

// PointRuleRequest ...
type PointRuleRequest struct {
	StartDate       string              `json:"start_date" validate:"required"`
	EndDate         string              `json:"end_date" validate:"required"`
	MinOrder        string              `json:"min_order" validate:"required"`
	PointConversion string              `json:"point_conversion" validate:"required"`
	MonthlyMaxPoint string              `json:"monthly_max_point" validate:"required"`
	Customers       []PointRuleCustomer `json:"customers"`
}

type PointRuleCustomer struct {
	CustomerCode string `json:"customer_code"`
	RegionID     string `json:"region_id"`
	Area         string `json:"area"`
	BranchID     string `json:"branch_id"`
}
