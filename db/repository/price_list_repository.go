package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IPriceListRepository ...
type IPriceListRepository interface {
	SelectAll(c context.Context, parameter models.PriceListParameter) ([]models.PriceList, error)
	FindAll(ctx context.Context, parameter models.PriceListParameter) ([]models.PriceList, int, error)
	FindByID(c context.Context, parameter models.PriceListParameter) (models.PriceList, error)
}

// PriceListRepository ...
type PriceListRepository struct {
	DB *sql.DB
}

// NewPriceListRepository ...
func NewPriceListRepository(DB *sql.DB) IPriceListRepository {
	return &PriceListRepository{DB: DB}
}

// Scan rows
func (repository PriceListRepository) scanRows(rows *sql.Rows) (res models.PriceList, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.BranchID,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository PriceListRepository) scanRow(row *sql.Row) (res models.PriceList, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.BranchID,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository PriceListRepository) SelectAll(c context.Context, parameter models.PriceListParameter) (data []models.PriceList, err error) {
	conditionString := ``

	if parameter.BranchID != "" {
		conditionString += ` AND pl.branch_id = ` + parameter.BranchID + ``
	}

	statement := models.PriceListSelectStatement + ` ` + models.PriceListWhereStatement +
		` AND (LOWER(pl."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository PriceListRepository) FindAll(ctx context.Context, parameter models.PriceListParameter) (data []models.PriceList, count int, err error) {
	conditionString := ``

	if parameter.BranchID != "" {
		conditionString += ` AND pl.branch_id = ` + parameter.BranchID + ``
	}

	query := models.PriceListSelectStatement + ` ` + models.PriceListWhereStatement + ` ` + conditionString + `
		AND (LOWER(pl."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "price_list" pl ` + models.PriceListWhereStatement + ` ` +
		conditionString + ` AND (LOWER(pl."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository PriceListRepository) FindByID(c context.Context, parameter models.PriceListParameter) (data models.PriceList, err error) {
	statement := models.PriceListSelectStatement + ` WHERE pl.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
