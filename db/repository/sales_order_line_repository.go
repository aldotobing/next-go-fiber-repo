package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesOrderLineRepository ...
type ISalesOrderLineRepository interface {
	SelectAll(c context.Context, parameter models.SalesOrderLineParameter) ([]models.SalesOrderLine, error)
	FindAll(ctx context.Context, parameter models.SalesOrderLineParameter) ([]models.SalesOrderLine, int, error)
	FindByID(c context.Context, parameter models.SalesOrderLineParameter) (models.SalesOrderLine, error)
}

// SalesOrderLineRepository ...
type SalesOrderLineRepository struct {
	DB *sql.DB
}

// NewSalesOrderLineRepository ...
func NewSalesOrderLineRepository(DB *sql.DB) ISalesOrderLineRepository {
	return &SalesOrderLineRepository{DB: DB}
}

// Scan rows
func (repository SalesOrderLineRepository) scanRows(rows *sql.Rows) (res models.SalesOrderLine, err error) {
	err = rows.Scan(
		&res.ID, &res.HeaderID, &res.CategoryName, &res.CategoryID,
		&res.ItemID, &res.ItemName, &res.UomID, &res.UomName,
		&res.QTY, &res.StockQty, &res.UnitPrice, &res.GrossAmount,
		&res.UseDiscPercent, &res.DisPercent1, &res.DisPercent2, &res.DisPercent3,
		&res.DisPercent4, &res.DisPercent5, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.SalesmanName, &res.SalesmanCode, &res.ItemPicture,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesOrderLineRepository) scanRow(row *sql.Row) (res models.SalesOrderLine, err error) {
	err = row.Scan(
		&res.ID, &res.HeaderID, &res.CategoryName, &res.CategoryID,
		&res.ItemID, &res.ItemName, &res.UomID, &res.UomName,
		&res.QTY, &res.StockQty, &res.UnitPrice, &res.GrossAmount,
		&res.UseDiscPercent, &res.DisPercent1, &res.DisPercent2, &res.DisPercent3,
		&res.DisPercent4, &res.DisPercent5, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.SalesmanName, &res.SalesmanCode, &res.ItemPicture,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SalesOrderLineRepository) SelectAll(c context.Context, parameter models.SalesOrderLineParameter) (data []models.SalesOrderLine, err error) {
	conditionString := ``

	if parameter.HeaderID != "" {
		conditionString += ` AND def.header_id = '` + parameter.HeaderID + `'`
	}

	statement := models.SalesOrderLineSelectStatement + ` ` + models.SalesOrderLineWhereStatement +
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
func (repository SalesOrderLineRepository) FindAll(ctx context.Context, parameter models.SalesOrderLineParameter) (data []models.SalesOrderLine, count int, err error) {
	conditionString := ``

	if parameter.HeaderID != "" {
		conditionString += ` AND def.header_id = '` + parameter.HeaderID + `'`
	}

	query := models.SalesOrderLineSelectStatement + ` ` + models.SalesOrderLineWhereStatement + ` ` + conditionString + `
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
			left join price_list_version plv on plv.id = def.price_list_version_id ` + models.SalesOrderLineWhereStatement + ` ` +
		conditionString + ` AND (LOWER(cus."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository SalesOrderLineRepository) FindByID(c context.Context, parameter models.SalesOrderLineParameter) (data models.SalesOrderLine, err error) {
	statement := models.SalesOrderLineSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
