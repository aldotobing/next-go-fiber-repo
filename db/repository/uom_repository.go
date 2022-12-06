package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IUomRepository ...
type IUomRepository interface {
	SelectAll(c context.Context, parameter models.UomParameter) ([]models.Uom, error)
	FindAll(ctx context.Context, parameter models.UomParameter) ([]models.Uom, int, error)
	FindByID(c context.Context, parameter models.UomParameter) (models.Uom, error)
}

// UomRepository ...
type UomRepository struct {
	DB *sql.DB
}

// NewUomRepository ...
func NewUomRepository(DB *sql.DB) IUomRepository {
	return &UomRepository{DB: DB}
}

// Scan rows
func (repository UomRepository) scanRows(rows *sql.Rows) (res models.Uom, err error) {
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
func (repository UomRepository) scanRow(row *sql.Row) (res models.Uom, err error) {
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
func (repository UomRepository) SelectAll(c context.Context, parameter models.UomParameter) (data []models.Uom, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += `AND u.id = ` + parameter.ID + ` `
	}

	statement := models.UomSelectStatement + ` ` + models.UomWhereStatement +
		` AND (LOWER(u."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository UomRepository) FindAll(ctx context.Context, parameter models.UomParameter) (data []models.Uom, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += `AND u.id = ` + parameter.ID + ` `
	}

	query := models.UomSelectStatement + ` ` + models.UomWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "uom" u ` + models.UomWhereStatement + ` ` +
		conditionString + ` AND (LOWER(u."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository UomRepository) FindByID(c context.Context, parameter models.UomParameter) (data models.Uom, err error) {
	statement := models.UomSelectStatement + ` WHERE u.created_date IS NOT NULL AND u.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
