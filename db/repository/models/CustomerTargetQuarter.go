package models

// CustomerTargetQuarterQuarter ...
type CustomerTargetQuarter struct {
	ID                    *string `json:"customer_id"`
	Code                  *string `json:"customer_code"`
	CustomerName          *string `json:"customer_name"`
	CustomerTargetQuarter *string `json:"customer_target"`
}

// CustomerTargetQuarterParameter ...
type CustomerTargetQuarterParameter struct {
	ID           string `json:"customer_id"`
	Code         string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
	Search       string `json:"search"`
	Page         int    `json:"page"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	By           string `json:"by"`
	Sort         string `json:"sort"`
}

var (
	// CustomerTargetQuarterOrderBy ...
	CustomerTargetQuarterOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date", "bmt._month"}
	// CustomerTargetQuarterOrderByrByString ...
	CustomerTargetQuarterOrderByrByString = []string{
		"cus.customer_name",
	}

	CustomerTargetQuarterGroupByrByString = []string{
		"cus.customer_name",
	}

	// CustomerTargetQuarterSelectStatement ...

	CustomerTargetQuarterSelectStatement = `
	SELECT
	CUS.ID AS CUSTOMER_ID,
	CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
	CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
		COALESCE ((SELECT COALESCE (SUM(SCT.SALES_TURNOVER), 0)
		FROM CUSTOMER CUS
		LEFT JOIN SALESMAN S ON S.ID = CUS.SALESMAN_ID
		LEFT JOIN SALESMAN_CUSTOMER_TARGET SCT ON SCT.CUSTOMER_ID = CUS.ID
		LEFT JOIN SALESMAN_TARGET_LINE STL ON SCT.SALESMAN_TARGET_LINE_ID = STL.ID
		LEFT JOIN BRANCH_MONTHLY_TARGET BMT ON BMT.ID = STL.BRANCH_MONTHLY_TARGET_ID
		LEFT JOIN BRANCH_YEARLY_TARGET BYT ON BYT.ID = BMT.BRANCH_YEARLY_TARGET_ID
	`
	// CustomerTargetQuarterWhereStatement ...
	CustomerTargetQuarterWhereStatement = ` WHERE byt._year = 2021
											and coalesce(cus.active, 1) = 1  `
)

//and cus.id = 5317 and bmt._month in (1, 2, 3)
//(SELECT DATE_PART('year', now()::date))
