package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ICilentInvoiceRepository ...
type ICilentInvoiceRepository interface {
	SelectAll(c context.Context, parameter models.CilentInvoiceParameter) ([]models.CilentInvoice, error)
	FindAll(ctx context.Context, parameter models.CilentInvoiceParameter) ([]models.CilentInvoice, int, error)
	FindByID(c context.Context, parameter models.CilentInvoiceParameter) (models.CilentInvoice, error)
	FindByDocumentNo(c context.Context, parameter models.CilentInvoiceParameter) (models.CilentInvoice, error)
	InsertDataWithLine(c context.Context, model *models.CilentInvoice) (res string, finishFlag bool, err error)
}

// CilentInvoiceRepository ...
type CilentInvoiceRepository struct {
	DB *sql.DB
}

// NewCilentInvoiceRepository ...
func NewCilentInvoiceRepository(DB *sql.DB) ICilentInvoiceRepository {
	return &CilentInvoiceRepository{DB: DB}
}

// Scan rows
func (repository CilentInvoiceRepository) scanRows(rows *sql.Rows) (res models.CilentInvoice, err error) {
	err = rows.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CilentInvoiceRepository) scanRow(row *sql.Row) (res models.CilentInvoice, err error) {
	err = row.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CilentInvoiceRepository) SelectAll(c context.Context, parameter models.CilentInvoiceParameter) (data []models.CilentInvoice, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.DateParam != "" {
		conditionString += ` AND def.modified_date > '` + parameter.DateParam + `'`
	}

	statement := models.CilentInvoiceSelectStatement + ` ` + models.CilentInvoiceWhereStatement +
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

// FindAll ...
func (repository CilentInvoiceRepository) FindAll(ctx context.Context, parameter models.CilentInvoiceParameter) (data []models.CilentInvoice, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	query := models.CilentInvoiceSelectStatement + ` ` + models.CilentInvoiceWhereStatement + ` ` + conditionString + `
		AND (LOWER(cus."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = `select 
			count(*)
			from customer_order_header def
			join customer cus on cus.id = def.cust_ship_to_id
			left join salesman s on s.id = def.salesman_id
			left join term_of_payment top on top.id = def.payment_terms_id
			left join branch b on b.id = def.branch_id
			left join price_list pl on pl.id = def.price_list_id
			left join price_list_version plv on plv.id = def.price_list_version_id ` + models.CilentInvoiceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(cus."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CilentInvoiceRepository) FindByID(c context.Context, parameter models.CilentInvoiceParameter) (data models.CilentInvoice, err error) {
	statement := models.CilentInvoiceSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByDocumentNo ...
func (repository CilentInvoiceRepository) FindByDocumentNo(c context.Context, parameter models.CilentInvoiceParameter) (data models.CilentInvoice, err error) {
	statement := models.CilentInvoiceSelectStatement + ` WHERE  def.document_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.DocumentNo)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository CilentInvoiceRepository) InsertDataWithLine(c context.Context, model *models.CilentInvoice) (res string, finishFlag bool, err error) {

	// Using a map to organize the print logic
	fields := map[string]*string{
		"ID":                   model.ID,
		"DocumentNo":           model.DocumentNo,
		"DocumentTypeID":       model.DocumentTypeID,
		"TransactionDate":      model.TransactionDate,
		"TransactionTime":      model.TransactionTime,
		"CustomerID":           model.CustomerID,
		"CustomerCode":         model.CustomerCode,
		"CustomerName":         model.CustomerName,
		"TaxCalcMethod":        model.TaxCalcMethod,
		"SalesmanID":           model.SalesmanID,
		"SalesmanCode":         model.SalesmanCode,
		"SalesRequestCode":     model.SalesRequestCode,
		"TransactionPoint":     model.TransactionPoint,
		"SalesmanName":         model.SalesmanName,
		"PaymentTermsID":       model.PaymentTermsID,
		"PaymentTermsName":     model.PaymentTermsName,
		"SalesOrderID":         model.SalesOrderID,
		"CompanyID":            model.CompanyID,
		"BranchID":             model.BranchID,
		"BranchName":           model.BranchName,
		"PriceListID":          model.PriceLIstID,
		"PriceListName":        model.PriceLIstName,
		"PriceListVersionID":   model.PriceLIstVersionID,
		"PriceListVersionName": model.PriceLIstVersionName,
		"Status":               model.Status,
		"GrossAmount":          model.GrossAmount,
		"TaxableAmount":        model.TaxableAmount,
		"TaxAmount":            model.TaxAmount,
		"RoundingAmount":       model.RoundingAmount,
		"OutstandingAmount":    model.OutstandingAmount,
		"NetAmount":            model.NetAmount,
		"DueDate":              model.DueDate,
		"DiscAmount":           model.DiscAmount,
		"PaidAmount":           model.PaidAmount,
		"NoPPN":                model.NoPPN,
		"GlobalDiscAmount":     model.GlobalDiscAmount,
		"InvoiceDate":          model.InvoiceDate,
		"PaidDate":             model.PaidDate,
	}

	for fieldName, fieldValue := range fields {
		if fieldValue != nil {
			fmt.Printf("%s: %s\n", fieldName, *fieldValue)
		} else {
			fmt.Printf("%s: nil\n", fieldName)
		}
	}

	availableinvoice, _ := repository.FindByDocumentNo(c, models.CilentInvoiceParameter{DocumentNo: *model.DocumentNo})

	statement := `
	insert into sales_invoice_header (
		document_no, document_type_id, transaction_date, transaction_time, cust_bill_to_id,
		tax_calc_method ,salesman_id ,payment_terms_id ,sales_order_id ,company_id ,
		branch_id ,price_list_id ,price_list_version_id ,status ,gross_amount ,
		disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
		outstanding_amount ,paid_amount ,due_date ,no_ppn ,global_disc_amount,
		transaction_point ,transaction_source_document_no, invoice_date,paid_date
		)values(
		$1, $2, $3, $4, (select id from customer where customer_code = $5),
		$6, (select id from salesman where partner_id =(select id from partner where code = $7)), $8, $9, $10,
		$11 ,(select id from price_list where code = $12), (select id from price_list_version where description = $13 and price_list_id = (select id from price_list where code = $12 ) ) , $14, $15,
		$16, $17, $18, $19, $20,
		$21, $22, $23, $24, $25,
		$26, $27, $28, $29 
		)
	RETURNING id`
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, finishFlag, err
	}
	defer transaction.Rollback()

	if availableinvoice.ID != nil {

		deletelinestatement := `delete from sales_invoice_line WHERE header_id = $1`

		deletedRow, _ := transaction.QueryContext(c, deletelinestatement, availableinvoice.ID)
		deletedRow.Close()

		deleteheaderstatement := `delete from sales_invoice_header WHERE id = $1`

		deletedHeaderRow, _ := transaction.QueryContext(c, deleteheaderstatement, availableinvoice.ID)
		deletedHeaderRow.Close()

	}
	//oustanding amount = net amount
	err = transaction.QueryRowContext(c, statement,
		model.DocumentNo, model.DocumentTypeID, model.TransactionDate, model.TransactionTime, model.CustomerCode,
		model.TaxCalcMethod, model.SalesmanCode, model.PaymentTermsID, model.SalesOrderID, model.CompanyID,
		model.BranchID, str.EmptyString(*model.PriceListCode), str.EmptyString(*model.PriceListVersionCode), str.EmptyString(*model.Status), str.EmptyString(*model.GrossAmount),
		model.DiscAmount, model.TaxableAmount, model.TaxAmount, model.RoundingAmount, model.NetAmount,
		str.EmptyString(*model.OutstandingAmount), str.EmptyString(*model.PaidAmount), model.DueDate, model.NoPPN, model.GlobalDiscAmount,
		str.EmptyString(*model.TransactionPoint), str.NullString(model.SalesRequestCode), str.NullString(model.InvoiceDate), str.NullString(model.PaidDate),
	).Scan(&res)

	if err != nil {
		return res, finishFlag, err
	}

	fmt.Printf("Successfully inserted invoice with DocumentNo: %s and generated ID: %s\n", *model.DocumentNo, res)

	model.ID = &res

	if model.ListLine != nil && len(*model.ListLine) > 0 {
		for _, lineObject := range *model.ListLine {
			line_statement := `
			insert into sales_invoice_line( 
				header_id ,line_no ,category_id ,item_id ,qty ,
				uom_id ,stock_qty ,unit_price ,gross_amount ,use_disc_percent ,
				disc_percent1 ,disc_percent2 ,disc_percent3 ,disc_percent4 ,disc_percent5 ,
				disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
				sales_order_line_id ,debt ,paid 
				)values(
				$1 ,$2 ,$3 ,( select id from item where code = $4 ) ,$5 ,
				( select id from uom where code =$6 ) ,$7 ,$8 ,$9 ,$10 ,
				$11 ,$12 ,$13 ,$14 ,$15 ,
				$16 ,$17 ,$18 ,$19 ,$20 ,
				$21 ,$22 ,$23
				) returning id
			`
			var resLine string
			err = transaction.QueryRowContext(c, line_statement,
				model.ID, lineObject.LineNo, lineObject.CategoryID, lineObject.ItemCode, lineObject.Qty,
				lineObject.UomCode, lineObject.StockQty, lineObject.UnitPrice, str.EmptyStringToZero(*lineObject.GrossAmount), str.EmptyStringToZero(*lineObject.UseDiscAmount),
				str.EmptyStringToZero(*lineObject.DiscPercent1), str.EmptyStringToZero(*lineObject.DiscPercent2), str.EmptyStringToZero(*lineObject.DiscPercent3), str.EmptyStringToZero(*lineObject.DiscPercent4), str.EmptyStringToZero(*lineObject.DiscPercent5),
				str.EmptyStringToZero(*lineObject.DiscountAmount), str.EmptyStringToZero(*lineObject.TaxableAmount), str.EmptyStringToZero(*lineObject.TaxAmount), str.EmptyStringToZero(*lineObject.RoundingAmount), str.EmptyStringToZero(*lineObject.NetAmount),
				str.NullString(lineObject.SalesOrderLineID), str.NullString(lineObject.Debt), str.NullString(lineObject.Paid),
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1,
			).Scan(&resLine)

			if err != nil {
				return res, finishFlag, err
			}
		}
	}

	if model.SalesRequestCode != nil {

		updatechecoutstatus := ` update customer_order_header set status= 'finish' where document_no = $1 `
		updateCOStatusRow, _ := transaction.QueryContext(c, updatechecoutstatus, model.SalesRequestCode)
		updateCOStatusRow.Close()

		finishFlag = true
	}

	if err = transaction.Commit(); err != nil {
		return res, finishFlag, err
	}

	return res, finishFlag, err
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
