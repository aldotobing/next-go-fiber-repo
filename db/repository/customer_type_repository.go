package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerTypeRepository ...
type ICustomerTypeRepository interface {
	SelectAll(c context.Context, parameter models.CustomerTypeParameter) ([]models.CustomerType, error)
	FindAll(ctx context.Context, parameter models.CustomerTypeParameter) ([]models.CustomerType, int, error)
	FindByID(c context.Context, parameter models.CustomerTypeParameter) (models.CustomerType, error)
}

// CustomerTypeRepository ...
type CustomerTypeRepository struct {
	DB *sql.DB
}

// NewCustomerTypeRepository ...
func NewCustomerTypeRepository(DB *sql.DB) ICustomerTypeRepository {
	return &CustomerTypeRepository{DB: DB}
}

// Scan rows
func (repository CustomerTypeRepository) scanRows(rows *sql.Rows) (res models.CustomerType, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CustomerTypeRepository) scanRow(row *sql.Row) (res models.CustomerType, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerTypeRepository) SelectAll(c context.Context, parameter models.CustomerTypeParameter) (data []models.CustomerType, err error) {
	conditionString := ``

	statement := models.CustomerTypeSelectStatement + ` ` + models.CustomerTypeWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository CustomerTypeRepository) FindAll(ctx context.Context, parameter models.CustomerTypeParameter) (data []models.CustomerType, count int, err error) {
	conditionString := ``

	query := models.CustomerTypeSelectStatement + ` ` + models.CustomerTypeWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "customer_type" def ` + models.CustomerTypeWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerTypeRepository) FindByID(c context.Context, parameter models.CustomerTypeParameter) (data models.CustomerType, err error) {
	statement := models.CustomerTypeSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
