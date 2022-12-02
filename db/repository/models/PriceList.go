package models

// PriceList ...
type PriceList struct {
	ID       *string `json:"id"`
	BranchID *string `json:"branch_id"`
	Code     *string `json:"price_list_code"`
	Name     *string `json:"price_list_name"`
}

// PriceListParameter ...
type PriceListParameter struct {
	ID       string `json:"id"`
	BranchID string `json:"branch_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
	ExceptId string `json:"except_id"`
}

var (
	// PriceListOrderBy ...
	PriceListOrderBy = []string{"pl.id", "pl._name", "pl.created_date"}
	// PriceListOrderByrByString ...
	PriceListOrderByrByString = []string{
		"pl._name",
	}

	// PriceListSelectStatement ...
	PriceListSelectStatement = `
	SELECT 
		PL.ID AS ID,
		PL.CODE AS PRICE_LIST_CODE,
		PL._NAME AS PRICE_LIST_NAME,
		PL.BRANCH_ID AS PRICE_LIST_BRANCH_ID
	FROM PRICE_LIST PL `

	// PriceListWhereStatement ...
	PriceListWhereStatement = ` WHERE pl.created_date IS not NULL `
)
