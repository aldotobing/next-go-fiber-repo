package models

// CustomerTargetYearQuarter ...
type CustomerTargetYear struct {
	ID             *string `json:"customer_id"`
	Code           *string `json:"customer_code"`
	CustomerName   *string `json:"customer_name"`
	CustomerTarget *string `json:"customer_target"`
}

// CustomerTargetYearParameter ...
type CustomerTargetYearParameter struct {
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
	// CustomerTargetYearOrderBy ...
	CustomerTargetYearOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date", "bmt._month"}
	// CustomerTargetYearOrderByrByString ...
	CustomerTargetYearOrderByrByString = []string{
		"cus.customer_name",
	}

	CustomerTargetYearGroupByrByString = []string{
		"cus.customer_name",
	}

	// CustomerTargetYearSelectStatement ...

	CustomerTargetYearSelectStatement = `
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
	// CustomerTargetYearWhereStatement ...
	CustomerTargetYearWhereStatement = ` WHERE BYT._YEAR = (SELECT DATE_PART('year', now()::date)) AND COALESCE(CUS.ACTIVE, 1) = 1 `
)

//and cus.id = 5317 and bmt._month in (1, 2, 3)
//(SELECT DATE_PART('year', now()::date))
