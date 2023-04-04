package models

// CustomerAchievementSemester ...
type CustomerAchievementSemester struct {
	ID           *string `json:"customer_id"`
	Code         *string `json:"customer_code"`
	CustomerName *string `json:"customer_name"`
	Achievement  *string `json:"achievement"`
}

// CustomerAchievementSemesterParameter ...
type CustomerAchievementSemesterParameter struct {
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
	// CustomerAchievementSemesterOrderBy ...
	CustomerAchievementSemesterOrderBy = []string{"cus.id", "cus.customer_name", "cus.created_date"}
	// CustomerAchievementSemesterOrderByrByString ...
	CustomerAchievementSemesterOrderByrByString = []string{
		"cus.customer_name",
	}

	// CustomerAchievementSemesterSelectStatement ...

	CustomerAchievementSemesterSelectStatement = `
	SELECT 
		CUS.ID AS CUSTOMER_ID,
		CUS.CUSTOMER_CODE AS CUSTOMER_CODE,
		CUS.CUSTOMER_NAME AS CUSTOMER_NAME,
		COALESCE ((SELECT COALESCE(SUM(NET_AMOUNT), 0) 
	FROM SALES_INVOICE_HEADER SIH
	LEFT JOIN CUSTOMER CUS ON CUS.ID = SIH.CUST_BILL_TO_ID
	`
	// CustomerAchievementSemesterWhereStatement ...
	CustomerAchievementSemesterWhereStatement = ` 
	WHERE cus.created_date IS not NULL 
	`
)
