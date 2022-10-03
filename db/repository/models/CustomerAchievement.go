package models

// CustomerAchievement ...
type CustomerAchievement struct {
	ID           *string `json:"customer_id"`
	Code         *string `json:"customer_code"`
	CustomerName *string `json:"customer_name"`
	Achievement  *string `json:"achievement"`
}

// CustomerAchievementParameter ...
type CustomerAchievementParameter struct {
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
	// CustomerAchievementOrderBy ...
	CustomerAchievementOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date"}
	// CustomerAchievementOrderByrByString ...
	CustomerAchievementOrderByrByString = []string{
		"cus.customer_name",
	}

	// CustomerAchievementSelectStatement ...

	CustomerAchievementSelectStatement = `
	SELECT 
		SIH.CUST_BILL_TO_ID AS CUSTOMER_ID,
		CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
		CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
		COALESCE (SUM(NET_AMOUNT), 0) AS ACHIEVEMENT
	FROM SALES_INVOICE_HEADER SIH
	LEFT JOIN CUSTOMER CUS ON CUS.ID = SIH.CUST_BILL_TO_ID
	`
	// CustomerAchievementWhereStatement ...
	CustomerAchievementWhereStatement = ` 
	WHERE cus.created_date IS not NULL 
	AND date_trunc('month', transaction_date) = date_trunc('month', current_timestamp)
	`
)
