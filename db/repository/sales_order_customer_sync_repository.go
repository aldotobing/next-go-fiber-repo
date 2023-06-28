package repository

import (
	"context"
	"database/sql"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ISalesOrderCustomerSyncRepository ...
type ISalesOrderCustomerSyncRepository interface {
	FindByID(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (models.SalesOrderCustomerSync, error)
	FindByDocumentNo(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (models.SalesOrderCustomerSync, error)
	InsertDataWithLine(c context.Context, model *models.SalesOrderCustomerSync) (res string, err error)
	RevisedSync(c context.Context, model *models.SalesOrderCustomerSync) (res string, err error)
}

// SalesOrderCustomerSyncRepository ...
type SalesOrderCustomerSyncRepository struct {
	DB *sql.DB
}

// NewSalesOrderCustomerSyncRepository ...
func NewSalesOrderCustomerSyncRepository(DB *sql.DB) ISalesOrderCustomerSyncRepository {
	return &SalesOrderCustomerSyncRepository{DB: DB}
}

// Scan rows
func (repository SalesOrderCustomerSyncRepository) scanRows(rows *sql.Rows) (res models.SalesOrderCustomerSync, err error) {
	err = rows.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesOrderCustomerSyncRepository) scanRow(row *sql.Row) (res models.SalesOrderCustomerSync, err error) {
	err = row.Scan(
		&res.ID, &res.DocumentNo, &res.CustomerCode,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository SalesOrderCustomerSyncRepository) FindByID(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (data models.SalesOrderCustomerSync, err error) {
	statement := models.SalesOrderCustomerSyncSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByDocumentNo ...
func (repository SalesOrderCustomerSyncRepository) FindByDocumentNo(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (data models.SalesOrderCustomerSync, err error) {
	statement := models.SalesOrderCustomerSyncSelectStatement + ` WHERE  def.document_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.DocumentNo)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository SalesOrderCustomerSyncRepository) InsertDataWithLine(c context.Context, model *models.SalesOrderCustomerSync) (res string, err error) {

	availableinvoice, _ := repository.FindByDocumentNo(c, models.SalesOrderCustomerSyncParameter{DocumentNo: *model.DocumentNo})

	statement := `
	insert into sales_order_header (
		document_no, document_type_id, transaction_date, transaction_time, cust_bill_to_id,
		tax_calc_method ,salesman_id ,payment_terms_id ,company_id ,
		branch_id ,price_list_id ,price_list_version_id ,status ,gross_amount ,
		disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount ,
		global_disc_amount,cust_ship_to_id,expected_delivery_date,void_reason_notes,
		created_date,modified_date,request_document_no
		)values(
		$1, $2, $3, $4, (select id from customer where customer_code = $5),
		$6, (select id from salesman where partner_id =(select id from partner where code = $7)), $8, $9, $10,
		$11 ,$12, $13, $14, $15,
		$16, $17, $18, $19, $20,
		(select id from customer where customer_code = $21),
		$22, $23, now(),now(),$24
		)
	RETURNING id`
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	if availableinvoice.ID != nil {

		deletelinestatement := `delete from sales_order_line WHERE header_id = $1`

		deletedRow, _ := transaction.QueryContext(c, deletelinestatement, availableinvoice.ID)
		deletedRow.Close()

		deleteheaderstatement := `delete from sales_order_header WHERE id = $1`

		deletedHeaderRow, _ := transaction.QueryContext(c, deleteheaderstatement, availableinvoice.ID)
		deletedHeaderRow.Close()

	}

	err = transaction.QueryRowContext(c, statement,
		model.DocumentNo, model.DocumentTypeID, model.TransactionDate, model.TransactionTime, model.CustomerCode,
		model.TaxCalcMethod, model.SalesmanCode, model.PaymentTermsID, model.CompanyID,
		model.BranchID, model.PriceLIstID, model.PriceLIstVersionID, str.EmptyString(*model.Status), str.EmptyString(*model.GrossAmount),
		model.DiscAmount, model.TaxableAmount, model.TaxAmount, model.RoundingAmount, model.NetAmount,
		model.GlobalDiscAmount, model.CustomerCode, model.ExpectedDeliveryDate, model.VoidReasonNotes, model.SalesRequestCode,
	).Scan(&res)

	if err != nil {
		fmt.Println("error insert header")
		return res, err
	}

	model.ID = &res

	if model.ListLine != nil && len(*model.ListLine) > 0 {
		for _, lineObject := range *model.ListLine {
			line_statement := `
			insert into sales_order_line(
				header_id ,line_no ,category_id ,item_id ,qty ,
				uom_id ,stock_qty ,unit_price ,gross_amount ,use_disc_percent ,
				disc_percent1 ,disc_percent2 ,disc_percent3 ,disc_percent4 ,disc_percent5 ,
				disc_amount ,taxable_amount ,tax_amount ,rounding_amount ,net_amount,
				location_id 
				
				)values(
				$1 ,$2 ,( select item_category_id from item where code = $3) ,( select id from item where code = $4 ) ,$5 ,
				( select id from uom where code =$6 ) ,$7 ,$8 ,$9 ,$10 ,
				$11 ,$12 ,$13 ,$14 ,$15 ,
				$16 ,$17 ,$18 ,$19 ,$20,
				$21 
				) returning id
			`
			var resLine string
			err = transaction.QueryRowContext(c, line_statement,
				model.ID, lineObject.LineNo, lineObject.ItemCode, lineObject.ItemCode, lineObject.Qty,
				lineObject.UomCode, lineObject.StockQty, lineObject.UnitPrice, str.EmptyStringToZero(*lineObject.GrossAmount), str.EmptyStringToZero(*lineObject.UseDiscPercent),
				str.EmptyStringToZero(*lineObject.DiscPercent1), str.EmptyStringToZero(*lineObject.DiscPercent2), str.EmptyStringToZero(*lineObject.DiscPercent3), str.EmptyStringToZero(*lineObject.DiscPercent4), str.EmptyStringToZero(*lineObject.DiscPercent5),
				str.EmptyStringToZero(*lineObject.DiscountAmount), str.EmptyStringToZero(*lineObject.TaxableAmount), str.EmptyStringToZero(*lineObject.TaxAmount), str.EmptyStringToZero(*lineObject.RoundingAmount), str.EmptyStringToZero(*lineObject.NetAmount),
				lineObject.LoacationID,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1, 1, 1,
				// 1, 1, 1,
			).Scan(&resLine)

			if err != nil {
				fmt.Println("error insert Line")
				return res, err
			}
		}
	}

	// if model.SalesRequestCode != nil {

	// 	updatechecoutstatus := ` update customer_order_header set status= 'submitted' where document_no = $1 `
	// 	updateCOStatusRow, _ := transaction.QueryContext(c, updatechecoutstatus, model.SalesRequestCode)
	// 	updateCOStatusRow.Close()

	// }

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

func (repository SalesOrderCustomerSyncRepository) RevisedSync(c context.Context, model *models.SalesOrderCustomerSync) (res string, err error) {

	statement := `
	update sales_order_header set status = $1, void_reason_notes = $2 where document_no = $3
	RETURNING id`
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	err = transaction.QueryRowContext(c, statement, model.Status, model.VoidReasonNotes,
		model.DocumentNo,
	).Scan(&res)

	if err != nil {
		return res, err
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
