package models

// CustomerAchievementYear ...
type CustomerAchievementYear struct {
	ID           *string `json:"customer_id"`
	Code         *string `json:"customer_code"`
	CustomerName *string `json:"customer_name"`
	Achievement  *string `json:"achievement"`
}

// CustomerAchievementYearParameter ...
type CustomerAchievementYearParameter struct {
	ID           string `json:"customer_id"`
	Code         string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
	Achievement  string `json:"achievement"`
	Search       string `json:"search"`
	Page         int    `json:"page"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	By           string `json:"by"`
	Sort         string `json:"sort"`
}

var (
	// CustomerAchievementYearOrderBy ...
	CustomerAchievementYearOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date"}
	// CustomerAchievementYearOrderByrByString ...
	CustomerAchievementYearOrderByString = []string{
		"cus.customer_name",
	}

	// CustomerAchievementYearSelectStatement ...

	CustomerAchievementYearSelectStatement = `
	SELECT 
		SIH.CUST_BILL_TO_ID AS CUSTOMER_ID,
		CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
		CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
		COALESCE (SUM(NET_AMOUNT), 0) AS ACHIEVEMENT
	FROM SALES_INVOICE_HEADER SIH
	LEFT JOIN CUSTOMER CUS ON CUS.ID = SIH.CUST_BILL_TO_ID
	`
	// CustomerAchievementYearWhereStatement ...
	CustomerAchievementYearWhereStatement = ` 
	WHERE cus.created_date IS not NULL 
	AND DATE_TRUNC('year', TRANSACTION_DATE) = DATE_TRUNC('year', CURRENT_TIMESTAMP)
	`
)
