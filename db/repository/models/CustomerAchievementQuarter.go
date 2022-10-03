package models

// CustomerAchievementQuarter ...
type CustomerAchievementQuarter struct {
	ID           *string `json:"customer_id"`
	Code         *string `json:"customer_code"`
	CustomerName *string `json:"customer_name"`
	Achievement  *string `json:"achievement"`
}

// CustomerAchievementQuarterParameter ...
type CustomerAchievementQuarterParameter struct {
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
	// CustomerAchievementQuarterOrderBy ...
	CustomerAchievementQuarterOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date"}
	// CustomerAchievementQuarterOrderByrByString ...
	CustomerAchievementQuarterOrderByrByString = []string{
		"cus.customer_name",
	}

	// CustomerAchievementQuarterSelectStatement ...

	CustomerAchievementQuarterSelectStatement = `
	SELECT 
		SIH.CUST_BILL_TO_ID AS CUSTOMER_ID,
		CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
		CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
		COALESCE (SUM(NET_AMOUNT), 0) AS ACHIEVEMENT
	FROM SALES_INVOICE_HEADER SIH
	LEFT JOIN CUSTOMER CUS ON CUS.ID = SIH.CUST_BILL_TO_ID
	`
	// CustomerAchievementQuarterWhereStatement ...
	CustomerAchievementQuarterWhereStatement = ` 
	WHERE cus.created_date IS not NULL 
	AND SIH.TRANSACTION_DATE > NOW() - INTERVAL '2 months'
	`
)
