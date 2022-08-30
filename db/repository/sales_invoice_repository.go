package repository

import (
	"context"
	"database/sql"
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
	// Edit(c context.Context, model *models.SalesInvoice) (*string, error)
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
		&res.ID, &res.NoInvoice, &res.NoOrder, &res.TrasactionDate, &res.Status, &res.InvoiceLine, &res.NetAmount,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesInvoiceRepository) scanRow(row *sql.Row) (res models.SalesInvoice, err error) {
	err = row.Scan(
		&res.ID, &res.NoInvoice, &res.NoOrder, &res.TrasactionDate, &res.Status, &res.InvoiceLine, &res.NetAmount,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SalesInvoiceRepository) SelectAll(c context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, err error) {
	conditionString := ``

	statement := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement +
		` AND (LOWER(def."document_no") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository SalesInvoiceRepository) FindAll(ctx context.Context, parameter models.SalesInvoiceParameter) (data []models.SalesInvoice, count int, err error) {
	conditionString := ``

	query := models.SalesInvoiceSelectStatement + ` ` + models.SalesInvoiceWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."document_no") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "sales_invoice_header" def ` + models.SalesInvoiceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."document_no") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err

}

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
