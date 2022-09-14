package models

// ItemMostSold ...
type ItemMostSold struct {
	ID               *string `json:"item_id"`
	Code             *string `json:"item_code"`
	Name             *string `json:"item_name"`
	ItemPicture      *string `json:"item_picture"`
	ItemCategoryId   *string `json:"item_category_id"`
	ItemCategoryName *string `json:"item_category_name"`
	ItemTotalSold    *string `json:"item_total_sold"`
}

// ItemMostSoldParameter ...
type ItemMostSoldParameter struct {
	ID             string `json:"item_id"`
	Code           string `json:"item_code"`
	Name           string `json:"item_name"`
	ItemCategoryId string `json:"item_category_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// ItemMostSoldOrderBy ...
	ItemMostSoldOrderBy = []string{"i.id", "i._name", "i.created_date", "total_sold"}
	// ItemMostSoldOrderByrByString ...
	ItemMostSoldOrderByrByString = []string{
		"i._name",
	}

	// ItemMostSoldSelectStatement ...
	/*
		--TRANSAKSI 1 BULAN TERAKHIR
	*/
	ItemMostSoldSelectStatement = `
	SELECT 
		I.ID as ITEM_ID,
		I.CODE as ITEM_CODE,
		I._NAME AS ITEM_NAME, 
		I.ITEM_PICTURE AS ITEM_PICTURE, 
		I.ITEM_CATEGORY_ID AS CATEGORY_ID, 
		IC._NAME AS ITEM_CATEGORY_NAME, 
		SUM (SIL.STOCK_QTY)::INTEGER AS TOTAL_SOLD
	FROM SALES_INVOICE_HEADER SIH
		JOIN SALES_INVOICE_LINE SIL ON SIL.HEADER_ID = SIH.ID
		JOIN ITEM I ON I.ID = SIL.ITEM_ID
		JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID
	WHERE SIH.TRANSACTION_DATE BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
	GROUP BY I.ID, I._NAME, SIH.TRANSACTION_DATE, I.ITEM_PICTURE , I.ITEM_CATEGORY_ID, IC._NAME
	ORDER BY TOTAL_SOLD DESC
	`

	// ItemMostSoldWhereStatement ...
	ItemMostSoldWhereStatement = ` WHERE I.created_date IS NOT NULL `
)
