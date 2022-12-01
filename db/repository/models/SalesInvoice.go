package models

import "encoding/json"

// SalesInvoice ...
type SalesInvoice struct {
	ID             *string          `json:"id"`
	CustomerName   *string          `json:"customer_name"`
	NoInvoice      *string          `json:"no_invoice"`
	NoOrder        *string          `json:"no_order"`
	TrasactionDate *string          `json:"transaction_date"`
	Status         *string          `json:"status"`
	NetAmount      *string          `json:"net_amount"`
	InvoiceLine    *json.RawMessage `json:"invoice_line"`
	ModifiedDate   *string          `json:"modified_date"`
}

// SalesInvoiceParameter ...
type SalesInvoiceParameter struct {
	ID         string `json:"id"`
	NoInvoice  string `json:"no_invoice"`
	CustomerID string `json:"customer_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	UserId     string `json:"user_id"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// SalesInvoiceOrderBy ...
	SalesInvoiceOrderBy = []string{"def.id", "def.document_no"}
	// SalesInvoiceOrderByrByString ...
	SalesInvoiceOrderByrByString = []string{
		"def.document_no",
	}

	// SalesInvoiceSelectStatement ...
	SalesInvoiceSelectStatement = `
	SELECT 
	DEF.ID as ID,
	P._NAME AS CUSTOMER_NAME,
	DEF.DOCUMENT_NO AS NO_INVOICE,
	SOH.DOCUMENT_NO AS NO_ORDER,
	DEF.TRANSACTION_DATE,
	DEF.MODIFIED_DATE,
	DEF.STATUS,
	DEF.NET_AMOUNT,
				(SELECT JSON_AGG(T) AS INVOICE_LINE
					FROM
									(SELECT I._NAME::VARCHAR(255) AS ITEM_NAME,
											SIL.UNIT_PRICE::VARCHAR(255) AS UNIT_PRICE,
											U._NAME::VARCHAR(255) AS UOMNAME,
											SIL.QTY::VARCHAR(255) AS QTY
										FROM SALES_INVOICE_HEADER subDEF
										JOIN SALES_INVOICE_LINE SIL ON SIL.HEADER_ID = subDEF.ID
										left JOIN SALES_ORDER_HEADER SOH ON SOH.ID = subDEF.SALES_ORDER_ID
										JOIN ITEM I ON I.ID = SIL.ITEM_ID
										JOIN UOM U ON U.ID = SIL.UOM_ID
										JOIN CUSTOMER C ON C.ID = subDEF.CUST_BILL_TO_ID
										WHERE SIL.HEADER_ID = subDEF.ID AND subDEF.CUST_BILL_TO_ID = DEF.CUST_BILL_TO_ID
										AND subdef.transaction_date = def.transaction_date) T)
FROM SALES_INVOICE_HEADER DEF
left JOIN SALES_ORDER_HEADER SOH ON SOH.ID = DEF.SALES_ORDER_ID
JOIN CUSTOMER C ON C.ID = DEF.CUST_BILL_TO_ID
JOIN PARTNER P ON P.ID = C.PARTNER_ID `

	// SalesInvoiceWhereStatement ...
	SalesInvoiceWhereStatement = ` where def.id is not null `
)
