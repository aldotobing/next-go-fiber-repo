package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ICustomerRepository ...
type ITransactionVARepository interface {
	SelectAll(c context.Context, parameter models.TransactionVAParameter) ([]models.TransactionVA, error)
	FindAll(ctx context.Context, parameter models.TransactionVAParameter) ([]models.TransactionVA, int, error)
	FindByID(c context.Context, parameter models.TransactionVAParameter) (models.TransactionVA, error)
	Edit(c context.Context, model *models.TransactionVA) (*string, error)
	Add(c context.Context, model *models.TransactionVA) (*string, error)
	FindLastActiveVa(c context.Context, parameter models.TransactionVAParameter) (models.TransactionVA, error)
	FindByCode(c context.Context, parameter models.TransactionVAParameter) (models.TransactionVA, error)
	PaidTransaction(c context.Context, model *models.TransactionVA) (*string, error)
}

// CustomerRepository ...
type TransactionVARepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewTransactionVARepository(DB *sql.DB) ITransactionVARepository {
	return &TransactionVARepository{DB: DB}
}

// Scan rows
func (repository TransactionVARepository) scanRows(rows *sql.Rows) (res models.TransactionVA, err error) {
	err = rows.Scan(
		&res.ID,
		&res.InvoiceCode,
		&res.VACode,
		&res.Amount,
		&res.VaPairID,
		&res.VaRef1,
		&res.VaRef2,
		&res.StartDate,
		&res.EndDate,
		&res.VAPartnerCode,
		&res.PaidStatus,
		&res.Customername,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository TransactionVARepository) scanRow(row *sql.Row) (res models.TransactionVA, err error) {
	err = row.Scan(
		&res.ID,
		&res.InvoiceCode,
		&res.VACode,
		&res.Amount,
		&res.VaPairID,
		&res.VaRef1,
		&res.VaRef2,
		&res.StartDate,
		&res.EndDate,
		&res.VAPartnerCode,
		&res.PaidStatus,
		&res.Customername,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository TransactionVARepository) SelectAll(c context.Context, parameter models.TransactionVAParameter) (data []models.TransactionVA, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	statement := models.TransactionVASelectStatement + ` ` + models.TransactionVAWhereStatement +
		` AND (LOWER(def."invoice_code") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print

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
func (repository TransactionVARepository) FindAll(ctx context.Context, parameter models.TransactionVAParameter) (data []models.TransactionVA, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	query := models.TransactionVASelectStatement + ` ` + models.TransactionVAWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."invoice_code") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`

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

	query = `select count(*)
		from virtual_account_transaction def  ` + models.TransactionVAWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."invoice_code") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository TransactionVARepository) FindByID(c context.Context, parameter models.TransactionVAParameter) (data models.TransactionVA, err error) {
	statement := models.TransactionVASelectStatement + ` WHERE def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByID ...
func (repository TransactionVARepository) FindByCode(c context.Context, parameter models.TransactionVAParameter) (data models.TransactionVA, err error) {
	strFilter := `  `
	if parameter.CurrentVaUser == 1 {
		strFilter += ` and now()::date between def.start_date and def.end_date `
	}

	statement := models.TransactionVASelectStatement + ` WHERE def.va_code = $1  ` + strFilter
	row := repository.DB.QueryRowContext(c, statement, parameter.VACode)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByID ...
func (repository TransactionVARepository) FindLastActiveVa(c context.Context, parameter models.TransactionVAParameter) (data models.TransactionVA, err error) {
	statement := models.TransactionVASelectStatement + ` WHERE def.invoice_code = $1 and now() between def.start_date and def.end_date order by def.id desc limit 1 `
	row := repository.DB.QueryRowContext(c, statement, parameter.InvoiceCode)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository TransactionVARepository) Edit(c context.Context, model *models.TransactionVA) (res *string, err error) {

	statement := `UPDATE partner SET 
	_name = $1, 
	address = $2, 
	user_id = $3,
	phone_no = $4,
	email = $5 
	WHERE id = $6 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.InvoiceCode,
		model.InvoiceCode,
		str.NullOrEmtyString(model.InvoiceCode),
		model.InvoiceCode,
		model.InvoiceCode,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository TransactionVARepository) Add(c context.Context, model *models.TransactionVA) (res *string, err error) {
	statement := `INSERT INTO virtual_account_transaction (invoice_code, va_partner_code,
		created_date,modified_date)
	VALUES ($1, $2,now(),now()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.InvoiceCode, model.VAPartnerCode).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository TransactionVARepository) PaidTransaction(c context.Context, model *models.TransactionVA) (res *string, err error) {

	statement := `UPDATE virtual_account_transaction SET 
	va_pair_id = $1, 
	va_ref1 = $2, 
	va_ref2 = $3,
	paid_status = 'paid'
	WHERE id = $4 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		str.NullOrEmtyString(model.VaPairID),
		str.NullOrEmtyString(model.VaRef1),
		str.NullOrEmtyString(model.VaRef2),
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
