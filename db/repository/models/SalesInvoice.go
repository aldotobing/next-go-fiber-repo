package models

import "encoding/json"

// SalesInvoice ...
type SalesInvoice struct {
	ID             *string          `json:"id"`
	NoInvoice      *string          `json:"no_invoice"`
	NoOrder        *string          `json:"no_order"`
	TrasactionDate *string          `json:"transaction_date"`
	Status         *string          `json:"status"`
	InvoiceLine    *json.RawMessage `json:"invoice_line"`
	NetAmount      *int             `json:"net_amount"`
}

// SalesInvoiceParameter ...
type SalesInvoiceParameter struct {
	ID        string `json:"id"`
	NoInvoice string `json:"no_invoice"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
}

var (
	// SalesInvoiceOrderBy ...
	SalesInvoiceOrderBy = []string{"def.id", "def.document_no"}
	// SalesInvoiceOrderByrByString ...
	SalesInvoiceOrderByrByString = []string{
		"def.document_no",
	}

	// SalesInvoiceSelectStatement ...
	SalesInvoiceSelectStatement = `SELECT DEF.DOCUMENT_NO AS NO_INVOICE,
	SOH.DOCUMENT_NO AS NO_ORDER,
	DEF.TRANSACTION_DATE,
	DEF.STATUS,
	DEF.NET_AMOUNT,
				(SELECT JSON_AGG(T) AS INVOICE_LINE
					FROM
									(SELECT I._NAME::VARCHAR(255) AS ITEM_NAME,
											SIL.UNIT_PRICE::VARCHAR(255) AS UNIT_PRICE,
											U._NAME::VARCHAR(255) AS UOMNAME,
											SIL.QTY::VARCHAR(255) AS QTY
										FROM SALES_INVOICE_HEADER DEF
										JOIN SALES_INVOICE_LINE SIL ON SIL.HEADER_ID = DEF.ID
										JOIN SALES_ORDER_HEADER SOH ON SOH.ID = DEF.SALES_ORDER_ID
										JOIN ITEM I ON I.ID = SIL.ITEM_ID
										JOIN UOM U ON U.ID = SIL.UOM_ID
										JOIN CUSTOMER C ON C.ID = DEF.CUST_BILL_TO_ID
										WHERE SIL.HEADER_ID = DEF.ID ) T)
FROM SALES_INVOICE_HEADER DEF
JOIN SALES_INVOICE_LINE SIL ON SIL.HEADER_ID = DEF.ID
JOIN SALES_ORDER_HEADER SOH ON SOH.ID = DEF.SALES_ORDER_ID
JOIN ITEM I ON I.ID = SIL.ITEM_ID
JOIN UOM U ON U.ID = SIL.UOM_ID
JOIN CUSTOMER C ON C.ID = DEF.CUST_BILL_TO_ID `

	// SalesInvoiceWhereStatement ...
	SalesInvoiceWhereStatement = ` where def.id is not null `
)
