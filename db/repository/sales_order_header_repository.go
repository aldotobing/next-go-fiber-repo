package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesOrderHeaderRepository ...
type ISalesOrderHeaderRepository interface {
	SelectAll(c context.Context, parameter models.SalesOrderHeaderParameter) ([]models.SalesOrderHeader, error)
	FindAll(ctx context.Context, parameter models.SalesOrderHeaderParameter) ([]models.SalesOrderHeader, int, error)
	FindByID(c context.Context, parameter models.SalesOrderHeaderParameter) (models.SalesOrderHeader, error)
}

// SalesOrderHeaderRepository ...
type SalesOrderHeaderRepository struct {
	DB *sql.DB
}

// NewSalesOrderHeaderRepository ...
func NewSalesOrderHeaderRepository(DB *sql.DB) ISalesOrderHeaderRepository {
	return &SalesOrderHeaderRepository{DB: DB}
}

// Scan rows
func (repository SalesOrderHeaderRepository) scanRows(rows *sql.Rows) (res models.SalesOrderHeader, err error) {
	err = rows.Scan(
		&res.ID, &res.DocumentNo, &res.TransactionDate, &res.TransactionTime,
		&res.CustomerID, &res.CustomerName, &res.TaxCalcMethod,
		&res.SalesmanID, &res.SalesmanName, &res.PaymentTermsID, &res.PaymentTermsName,
		&res.ExpectedDeliveryDate, &res.BranchID, &res.BranchName,
		&res.PriceLIstID, &res.PriceLIstName, &res.PriceLIstVersionID, &res.PriceLIstVersionName,
		&res.Status, &res.GrossAmount, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.DiscAmount,
		&res.CustomerCode, &res.SalesmanCode, &res.CustomerAddress, &res.ModifiedDate,
		&res.VoidReasonText, &res.OrderSource,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesOrderHeaderRepository) scanRow(row *sql.Row) (res models.SalesOrderHeader, err error) {
	err = row.Scan(
		&res.ID, &res.DocumentNo, &res.TransactionDate, &res.TransactionTime,
		&res.CustomerID, &res.CustomerName, &res.TaxCalcMethod,
		&res.SalesmanID, &res.SalesmanName, &res.PaymentTermsID, &res.PaymentTermsName,
		&res.ExpectedDeliveryDate, &res.BranchID, &res.BranchName,
		&res.PriceLIstID, &res.PriceLIstName, &res.PriceLIstVersionID, &res.PriceLIstVersionName,
		&res.Status, &res.GrossAmount, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.DiscAmount,
		&res.CustomerCode, &res.SalesmanCode, &res.CustomerAddress, &res.ModifiedDate,
		&res.VoidReasonText, &res.OrderSource,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SalesOrderHeaderRepository) SelectAll(c context.Context, parameter models.SalesOrderHeaderParameter) (data []models.SalesOrderHeader, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.DateParam != "" {
		conditionString += ` AND def.modified_date > '` + parameter.DateParam + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	statement := models.SalesOrderHeaderSelectStatement + ` ` + models.SalesOrderHeaderWhereStatement +
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
func (repository SalesOrderHeaderRepository) FindAll(ctx context.Context, parameter models.SalesOrderHeaderParameter) (data []models.SalesOrderHeader, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	query := models.SalesOrderHeaderSelectStatement + ` ` + models.SalesOrderHeaderWhereStatement + ` ` + conditionString + `
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
			left join price_list_version plv on plv.id = def.price_list_version_id ` + models.SalesOrderHeaderWhereStatement + ` ` +
		conditionString + ` AND (LOWER(cus."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository SalesOrderHeaderRepository) FindByID(c context.Context, parameter models.SalesOrderHeaderParameter) (data models.SalesOrderHeader, err error) {
	statement := models.SalesOrderHeaderSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
