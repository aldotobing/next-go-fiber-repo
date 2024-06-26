package models

import "encoding/json"

// SalesInvoice ...
type SalesInvoice struct {
	ID                    *string          `json:"invoice_id"`
	CustomerName          *string          `json:"customer_name"`
	NoInvoice             *string          `json:"no_invoice"`
	NoOrder               *string          `json:"no_order"`
	TrasactionDate        *string          `json:"transaction_date"`
	ModifiedDate          *string          `json:"modified_date"`
	ModifiedBy            *string          `json:"modified_by"`
	JatuhTempo            *string          `json:"jatuh_tempo"`
	Status                *string          `json:"status"`
	NetAmount             *string          `json:"net_amount"`
	OutStandingAmount     *string          `json:"outstanding_amount"`
	InvoiceLine           *json.RawMessage `json:"invoice_line"`
	TotalPaid             *string          `json:"total_paid"`
	PaymentMethod         *string          `json:"payment_method"`
	SourceDocumentNo      *string          `json:"source_document_no"`
	CustomerCode          *string          `json:"customer_code"`
	IDCustomerOrderHeader *string          `json:"id_customer_order_header"`
	CustomerID            *string          `json:"customer_id"`
	GlobalDiscAmount      *string          `json:"global_disc_amount"`
}

// SalesInvoiceParameter ...
type SalesInvoiceParameter struct {
	ID         string `json:"invoice_id"`
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
	SalesInvoiceOrderBy = []string{"def.id", "def.document_no", "def.created_date"}
	// SalesInvoiceOrderByrByString ...
	SalesInvoiceOrderByrByString = []string{
		"def.created_date",
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
	DEF.TRANSACTION_DATE + top.DAYS AS JATUH_TEMPO,
	DEF.STATUS,
	DEF.NET_AMOUNT,
	DEF.OUTSTANDING_AMOUNT,
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
										WHERE SUBDEF.id = DEF.ID) T),
	DEF.TOTAL_PAID,
	DEF.PAYMENT_METHOD,
	def.transaction_source_document_no,
	C.CUSTOMER_CODE,
	COH.ID AS ORDER_ID,
	C.ID,
	coalesce(DEF.GLOBAL_DISC_AMOUNT, 0) AS global_disc_amount
	FROM SALES_INVOICE_HEADER DEF
	left JOIN SALES_ORDER_HEADER SOH ON SOH.ID = DEF.SALES_ORDER_ID
	JOIN CUSTOMER C ON C.ID = DEF.CUST_BILL_TO_ID
	JOIN PARTNER P ON P.ID = C.PARTNER_ID 
	JOIN TERM_OF_PAYMENT TOP ON TOP.ID = DEF.PAYMENT_TERMS_ID
	LEFT JOIN CUSTOMER_ORDER_HEADER COH ON LEFT(COH.DOCUMENT_NO , 15) = LEFT(DEF.transaction_source_document_no, 15)`

	// SalesInvoiceWhereStatement ...
	SalesInvoiceWhereStatement = ` where def.id is not null `
)
