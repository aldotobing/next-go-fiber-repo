package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesInvoiceRepository ...
type ISalesInvoiceRepository interface {
	SelectAll(c context.Context, parameter models.SalesInvoiceParameter) ([]models.SalesInvoice, error)
	FindAll(ctx context.Context, parameter models.SalesInvoiceParameter) ([]models.SalesInvoice, int, error)
	FindByID(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	FindByDocumentNo(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	FindByCustomerId(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// Add(c context.Context, model *models.SalesInvoice) (*string, error)
	Edit(c context.Context, model *models.SalesInvoice) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// SalesInvoiceRepository ...
type SalesInvoiceRepository struct {
	DB *sql.DB
}

// NewSalesInvoiceRepository ...
func NewSalesInvoiceRepository(DB *sql.DB) ISalesInvoiceRepository {
	return &SalesInvoiceRepository{DB: DB}
}

// Scan rows
func (repository SalesInvoiceRepository) scanRows(rows *sql.Rows) (res models.SalesInvoice, err error) {
	err = rows.Scan(
		&res.ID, &res.CustomerName, &res.NoInvoice, &res.NoOrder, &res.TrasactionDate, &res.ModifiedDate, &res.JatuhTempo, &res.Status, &res.NetAmount, &res.OutStandingAmount, &res.InvoiceLine,
		&res.TotalPaid, &res.PaymentMethod,
	)

	return
}

// Scan row
func (repository SalesInvoiceRepository) scanRow(row *sql.Row) (res models.SalesInvoice, err error) {
	err = row.Scan(
		&res.ID, &res.CustomerName, &res.NoInvoice, &res.NoOrder, &res.TrasactionDate, &res.ModifiedDate, &res.JatuhTempo, &res.Status, &res.NetAmount, &res.OutStandingAmount, &res.InvoiceLine,
		&res.TotalPaid, &res.PaymentMethod,
	)

	return
}

// SelectAll ...
func (repository SalesInvoiceRepository) SelectAll(c context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, err error) {
	var conditionString string
	var args []interface{}
	var index int = 1

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_bill_to_id = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerID)
		index++
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND def.transaction_date BETWEEN $` + strconv.Itoa(index) + ` AND $` + strconv.Itoa(index+1)
		args = append(args, parameter.StartDate, parameter.EndDate)
		index += 2
	}

	if parameter.UserId != "" {
		conditionString += ` AND def.branch_id IN (SELECT ub.branch_id FROM user_branch ub WHERE ub.user_id = $` + strconv.Itoa(index) + `)`
		args = append(args, parameter.UserId)
		index++
	}

	statement := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement +
		` AND (LOWER(def."document_no") LIKE $` + strconv.Itoa(index) + `) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	args = append(args, "%"+strings.ToLower(parameter.NoInvoice)+"%")

	rows, err := repository.DB.QueryContext(c, statement, args...)

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

// func (repository SalesInvoiceRepository) SelectAll(c context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, err error) {
// 	conditionString := ``

// 	if parameter.CustomerID != "" {
// 		conditionString += ` AND def.cust_bill_to_id = ` + parameter.CustomerID + ``
// 	}
// 	if parameter.StartDate != "" && parameter.EndDate != "" {
// 		conditionString += ` AND def.transaction_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
// 	}
// 	if parameter.UserId != "" {
// 		conditionString += ` AND def.branch_id in (select ub.branch_id from user_branch ub where ub.user_id =  ` + parameter.UserId + `)`
// 	}

// 	statement := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement +
// 		` AND (LOWER(def."document_no") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

// 	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.NoInvoice)+"%")

// 	if err != nil {
// 		return data, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {

// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, temp)
// 	}

// 	return data, err
// }

// FindAll ...
func (repository SalesInvoiceRepository) FindAll(ctx context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, count int, err error) {
	var conditionString string
	var args, argsCount []interface{}
	var index int = 1

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_bill_to_id = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerID)
		argsCount = append(argsCount, parameter.CustomerID)
		index++
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND def.transaction_date BETWEEN $` + strconv.Itoa(index) + ` AND $` + strconv.Itoa(index+1)
		args = append(args, parameter.StartDate, parameter.EndDate)
		argsCount = append(argsCount, parameter.StartDate, parameter.EndDate)
		index += 2
	}

	if parameter.UserId != "" {
		conditionString += ` AND def.branch_id IN (SELECT ub.branch_id FROM user_branch ub WHERE ub.user_id = $` + strconv.Itoa(index) + `)`
		args = append(args, parameter.UserId)
		argsCount = append(argsCount, parameter.UserId)
		index++
	}

	query := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement + ` ` + conditionString +
		` AND (LOWER(def."document_no") LIKE $` + strconv.Itoa(index) + `) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $` + strconv.Itoa(index+1) + ` LIMIT $` + strconv.Itoa(index+2)
	args = append(args, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)

	rows, err := repository.DB.QueryContext(ctx, query, args...)
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

	argsCount = append(argsCount, "%"+strings.ToLower(parameter.Search)+"%")
	query = `SELECT COUNT(*) FROM "sales_invoice_header" def ` + models.SalesInvoiceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."document_no") LIKE $` + strconv.Itoa(index) + `)`

	err = repository.DB.QueryRowContext(ctx, query, argsCount...).Scan(&count)

	fmt.Println("Query:", query)
	fmt.Println("Args:", argsCount)

	return data, count, err
}

// func (repository SalesInvoiceRepository) FindAll(ctx context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, count int, err error) {
// 	conditionString := ``

// 	if parameter.CustomerID != "" {
// 		conditionString += ` AND def.cust_bill_to_id = '` + parameter.CustomerID + `'`
// 	}
// 	if parameter.StartDate != "" && parameter.EndDate != "" {
// 		conditionString += ` AND def.transaction_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
// 	}
// 	if parameter.UserId != "" {
// 		conditionString += ` AND def.branch_id in (select ub.branch_id from user_branch ub where ub.user_id =  '` + parameter.UserId + `')`
// 	}
// 	query := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement + ` ` + conditionString + `
// 		AND (LOWER(def."document_no") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
// 	fmt.Println(query)
// 	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
// 	if err != nil {
// 		return data, count, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, count, err
// 		}
// 		data = append(data, temp)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return data, count, err
// 	}

// 	query = `SELECT COUNT(*) FROM "sales_invoice_header" def ` + models.SalesInvoiceWhereStatement + ` ` +
// 		conditionString + ` AND (LOWER(def."document_no") LIKE $1)`
// 	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
// 	return data, count, err

// }

// FindByID ...
func (repository SalesInvoiceRepository) FindByID(c context.Context, parameter models.SalesInvoiceParameter) (data models.SalesInvoice, err error) {
	statement := models.SalesInvoiceSelectStatement + ` WHERE  def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByDocumentNo ...
func (repository SalesInvoiceRepository) FindByDocumentNo(c context.Context, parameter models.SalesInvoiceParameter) (data models.SalesInvoice, err error) {
	statement := models.SalesInvoiceSelectStatement + ` WHERE  def.document_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCustomerId ...
func (repository SalesInvoiceRepository) FindByCustomerId(c context.Context, parameter models.SalesInvoiceParameter) (data models.SalesInvoice, err error) {
	statement := models.SalesInvoiceSelectStatement + ` WHERE  def.cust_bill_to_id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository SalesInvoiceRepository) Edit(c context.Context, in *models.SalesInvoice) (out *string, err error) {
	statement := `UPDATE SALES_INVOICE_HEADER SET 
		modified_date = $1, 
		total_paid = $2,
		payment_method = $3
	WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.ModifiedDate,
		in.TotalPaid,
		in.PaymentMethod,
		in.ID,
	).Scan(&out)

	return
}
