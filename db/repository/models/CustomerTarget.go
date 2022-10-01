package models

// CustomerTarget ...
type CustomerTarget struct {
	ID             *string `json:"customer_id"`
	Code           *string `json:"customer_code"`
	CustomerName   *string `json:"customer_name"`
	CustomerTarget *string `json:"customer_target"`
}

// CustomerTargetParameter ...
type CustomerTargetParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	CustomerName   string `json:"customer_name"`
	CustomerTarget string `json:"customer_target"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// CustomerTargetOrderBy ...
	CustomerTargetOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date", "bmt._month"}
	// CustomerTargetOrderByrByString ...
	CustomerTargetOrderByrByString = []string{
		"cus.customer_name",
	}

	// CustomerTargetSelectStatement ...

	CustomerTargetSelectStatement = `
	SELECT CUS.ID AS CUSTOMER_ID,
	CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
	CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
	COALESCE ((SELECT (SCT.SALES_TURNOVER)
				FROM CUSTOMER CUS
				JOIN SALESMAN S ON S.ID = CUS.SALESMAN_ID
				LEFT JOIN SALESMAN_CUSTOMER_TARGET SCT ON SCT.CUSTOMER_ID = CUS.ID
				LEFT JOIN SALESMAN_TARGET_LINE STL ON SCT.SALESMAN_TARGET_LINE_ID = STL.ID
				LEFT JOIN BRANCH_MONTHLY_TARGET BMT ON BMT.ID = STL.BRANCH_MONTHLY_TARGET_ID
				LEFT JOIN BRANCH_YEARLY_TARGET BYT ON BYT.ID = BMT.BRANCH_YEARLY_TARGET_ID
	`
	// CustomerTargetWhereStatement ...
	CustomerTargetWhereStatement = ` WHERE cus.created_date IS not NULL 
									and byt._year = (SELECT DATE_PART('year', now()::date)) 
									and bmt._month = (SELECT DATE_PART('month', now()::date)) 
									`
)
