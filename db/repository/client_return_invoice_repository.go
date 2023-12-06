package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ICilentReturnInvoiceRepository ...
type ICilentReturnInvoiceRepository interface {
	SelectAll(c context.Context, parameter models.CilentReturnInvoiceParameter) ([]models.CilentReturnInvoice, error)
	FindByID(c context.Context, parameter models.CilentReturnInvoiceParameter) (models.CilentReturnInvoice, error)
	FindByDocumentNo(c context.Context, parameter models.CilentReturnInvoiceParameter) (models.CilentReturnInvoice, error)
	InsertDataWithLine(c context.Context, model *models.CilentReturnInvoice) (res string, err error)
}

// CilentReturnInvoiceRepository ...
type CilentReturnInvoiceRepository struct {
	DB *sql.DB
}

// NewCilentReturnInvoiceRepository ...
func NewCilentReturnInvoiceRepository(DB *sql.DB) ICilentReturnInvoiceRepository {
	return &CilentReturnInvoiceRepository{DB: DB}
}

// Scan rows
func (repository CilentReturnInvoiceRepository) scanRows(rows *sql.Rows) (res models.CilentReturnInvoice, err error) {
	err = rows.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CilentReturnInvoiceRepository) scanRow(row *sql.Row) (res models.CilentReturnInvoice, err error) {
	err = row.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CilentReturnInvoiceRepository) SelectAll(c context.Context, parameter models.CilentReturnInvoiceParameter) (data []models.CilentReturnInvoice, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.DateParam != "" {
		conditionString += ` AND def.modified_date > '` + parameter.DateParam + `'`
	}

	statement := models.CilentReturnInvoiceSelectStatement + ` ` + models.CilentReturnInvoiceWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

// FindByID ...
func (repository CilentReturnInvoiceRepository) FindByID(c context.Context, parameter models.CilentReturnInvoiceParameter) (data models.CilentReturnInvoice, err error) {
	statement := models.CilentReturnInvoiceSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByDocumentNo ...
func (repository CilentReturnInvoiceRepository) FindByDocumentNo(c context.Context, parameter models.CilentReturnInvoiceParameter) (data models.CilentReturnInvoice, err error) {
	statement := models.CilentReturnInvoiceSelectStatement + ` WHERE  def.document_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.DocumentNo)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository CilentReturnInvoiceRepository) InsertDataWithLine(c context.Context, model *models.CilentReturnInvoice) (res string, err error) {

	availableinvoice, _ := repository.FindByDocumentNo(c, models.CilentReturnInvoiceParameter{DocumentNo: *model.DocumentNo})

	statement := `
	insert into sales_return_invoice_header (
		document_no, document_type_id, transaction_date, transaction_time, cust_bill_to_id,
		tax_calc_method ,salesman_id ,payment_terms_id  ,company_id ,
		branch_id ,price_list_id ,price_list_version_id ,status ,gross_amount ,
		disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
		global_disc_amount,
		notes ,return_type, ref_invoice, created_date
		)values(
		$1, $2, $3, $4, (select id from customer where customer_code = $5),
		$6, (select id from salesman where partner_id =(select id from partner where code = $7)), $8, $9, $10,
		$11 ,$12, $13, $14, $15,
		$16, $17, $18, $19, $20,
		$21, $22, $23 , now()
		)
	RETURNING id`
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	if availableinvoice.ID != nil {

		deletelinestatement := `delete from sales_return_invoice_line WHERE header_id = $1`

		deletedRow, _ := transaction.QueryContext(c, deletelinestatement, availableinvoice.ID)
		deletedRow.Close()

		deleteheaderstatement := `delete from sales_return_invoice_header WHERE id = $1`

		deletedHeaderRow, _ := transaction.QueryContext(c, deleteheaderstatement, availableinvoice.ID)
		deletedHeaderRow.Close()

	}
	//oustanding amount = net amount
	err = transaction.QueryRowContext(c, statement,
		model.DocumentNo, model.DocumentTypeID, model.TransactionDate, model.TransactionTime, model.CustomerCode,
		model.TaxCalcMethod, model.SalesmanCode, model.PaymentTermsID, model.CompanyID,
		model.BranchID, model.PriceLIstID, model.PriceLIstVersionID, str.EmptyString(*model.Status), str.EmptyString(*model.GrossAmount),
		model.DiscAmount, model.TaxableAmount, model.TaxAmount, model.RoundingAmount,
		str.EmptyString(*model.NetAmount), model.GlobalDiscAmount,
		str.EmptyString(*model.Notes), str.NullString(model.ReturnType), str.NullString(model.RefInvoice),
	).Scan(&res)

	if err != nil {
		fmt.Println("error input data header ", err)
		return res, err
	}
	model.ID = &res

	if model.ListLine != nil && len(*model.ListLine) > 0 {
		for _, lineObject := range *model.ListLine {
			line_statement := `
			insert into sales_return_invoice_line( 
				header_id ,line_no ,category_id ,item_id ,qty ,
				uom_id ,stock_qty ,unit_price ,gross_amount ,use_disc_percent ,
				disc_percent1 ,disc_percent2 ,disc_percent3 ,disc_percent4 ,disc_percent5 ,
				disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
				stock_qty_replacement ,uom_id_replacement ,uom_backoffice 
				)values(
				$1 ,$2 ,$3 ,( select id from item where code = $4 ) ,$5 ,
				( select id from uom where code =$6 ) ,$7 ,$8 ,$9 ,$10 ,
				$11 ,$12 ,$13 ,$14 ,$15 ,
				$16 ,$17 ,$18 ,$19 ,$20 ,
				$21 ,( select id from uom where code =$22) ,$23
				) returning id
			`
			var resLine string
			err = transaction.QueryRowContext(c, line_statement,
				model.ID, lineObject.LineNo, lineObject.CategoryID, lineObject.ItemCode, lineObject.Qty,
				lineObject.UomCode, lineObject.StockQty, lineObject.UnitPrice, str.EmptyStringToZero(*lineObject.GrossAmount), str.EmptyStringToZero(*lineObject.UseDiscAmount),
				str.EmptyStringToZero(*lineObject.DiscPercent1), str.EmptyStringToZero(*lineObject.DiscPercent2), str.EmptyStringToZero(*lineObject.DiscPercent3), str.EmptyStringToZero(*lineObject.DiscPercent4), str.EmptyStringToZero(*lineObject.DiscPercent5),
				str.EmptyStringToZero(*lineObject.DiscountAmount), str.EmptyStringToZero(*lineObject.TaxableAmount), str.EmptyStringToZero(*lineObject.TaxAmount), str.EmptyStringToZero(*lineObject.RoundingAmount), str.EmptyStringToZero(*lineObject.NetAmount),
				str.NullString(lineObject.StockQtyReplacment), str.NullString(lineObject.UomCodeReplace), str.NullString(lineObject.UomBackOffice),
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1,
			).Scan(&resLine)

			if err != nil {
				fmt.Println("error input data line ", err)
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// insert into sales_invoice_header (
// 	document_no, document_type_id, transaction_date, transaction_time, cust_bill_to_id,
// 	tax_calc_method ,salesman_id ,payment_terms_id ,sales_order_id ,company_id ,
// 	branch_id ,price_list_id ,price_list_version_id ,status ,gross_amount ,
// 	disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
// 	outstanding_amount ,paid_amount ,due_date ,no_ppn ,global_disc_amount
// 	)values(
// 	$,$,$,$,$,
// 	$,$,$,$,$,
// 	$,$,$,$,$,
// 	$,$,$,$,$,
// 	$,$,$,$,?
// 	)
