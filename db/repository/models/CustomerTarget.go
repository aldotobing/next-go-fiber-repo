package models

// CustomerTarget ...
type CustomerTarget struct {
	ID             *string `json:"customer_id"`
	Code           *string `json:"customer_code"`
	CustomerName   *string `json:"customer_name"`
	CustomerTarget *string `json:"customer_target"`
	Month          *string `json:"month"`
	Year           *string `json:"year"`
}

// CustomerTargetParameter ...
type CustomerTargetParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	CustomerName   string `json:"customer_name"`
	CustomerTarget string `json:"customer_target"`
	Month          string `json:"month"`
	Year           string `json:"year"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// CustomerTargetOrderBy ...
	CustomerTargetOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date"}
	// CustomerTargetOrderByrByString ...
	CustomerTargetOrderByrByString = []string{
		"cus.customer_name",
	}

	// CustomerTargetSelectStatement ...

	CustomerTargetSelectStatement = `
	SELECT 
		cus.id as customer_id, 
		cus.customer_code as customer_code,
		cus.customer_name as customer_name, 
		bmt._month as _month, 
		sct.sales_turnover as target
	FROM customer cus
		JOIN salesman s on s.id = cus.salesman_id
		LEFT JOIN salesman_customer_target sct on sct.customer_id = cus.id
		LEFT JOIN salesman_target_line stl on sct.salesman_target_line_id = stl.id
		LEFT JOIN branch_monthly_target bmt on bmt.id = stl.branch_monthly_target_id
		LEFT JOIN branch_yearly_target byt on byt.id = bmt.branch_yearly_target_id
   	order by customer_name , _month
	`

	// CustomerTargetWhereStatement ...
	CustomerTargetWhereStatement = ` WHERE cus.created_date IS not NULL and byt._year = 2021 `
)

//Customer Target where statement
// where byt._year = 2021 and s.id in (1) and coalesce(cus.active, 1) = 1 and cus.id = 5317
