package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesmanRepository ...
type ISalesmanRepository interface {
	SelectAll(c context.Context, parameter models.SalesmanParameter) ([]models.Salesman, error)
	FindAll(ctx context.Context, parameter models.SalesmanParameter) ([]models.Salesman, int, error)
	FindByID(c context.Context, parameter models.SalesmanParameter) (models.Salesman, error)
}

// SalesmanRepository ...
type SalesmanRepository struct {
	DB *sql.DB
}

// NewSalesmanRepository ...
func NewSalesmanRepository(DB *sql.DB) ISalesmanRepository {
	return &SalesmanRepository{DB: DB}
}

// Scan rows
func (repository SalesmanRepository) scanRows(rows *sql.Rows) (res models.Salesman, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesmanRepository) scanRow(row *sql.Row) (res models.Salesman, err error) {
	err = row.Scan(
		&res.ID,
		&res.Name,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SalesmanRepository) SelectAll(c context.Context, parameter models.SalesmanParameter) (data []models.Salesman, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += " and def.branch_id in (select branch_id from user_branch where user_id = " + parameter.UserID + ")"
	}

	statement := models.SalesmanSelectStatement + ` ` + models.SalesmanWhereStatement +
		` AND (LOWER(p."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository SalesmanRepository) FindAll(ctx context.Context, parameter models.SalesmanParameter) (data []models.Salesman, count int, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += " and def.branch_id in (select branch_id from user_branch where user_id = " + parameter.UserID + ")"
	}
	query := models.SalesmanSelectStatement + ` ` + models.SalesmanWhereStatement + ` ` + conditionString + `
		AND (LOWER(p."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "salesman" def 
	join partner p on p.id = def.partner_id
	` + models.SalesmanWhereStatement + ` ` +
		conditionString + ` AND (LOWER(p."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository SalesmanRepository) FindByID(c context.Context, parameter models.SalesmanParameter) (data models.Salesman, err error) {
	statement := models.SalesmanSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
