package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerLevelRepository ...
type ICustomerLevelRepository interface {
	FindAll(ctx context.Context, parameter models.CustomerLevelParameter) ([]models.CustomerLevel, error)
}

// CustomerLevelRepository ...
type CustomerLevelRepository struct {
	DB *sql.DB
}

// NewCustomerLevelRepository ...
func NewCustomerLevelRepository(DB *sql.DB) ICustomerLevelRepository {
	return &CustomerLevelRepository{DB: DB}
}

// Scan rows
func (repository CustomerLevelRepository) scanRows(rows *sql.Rows) (res models.CustomerLevel, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.CreatedDate,
		&res.ModifiedDate,
	)

	return
}

// Scan row
func (repository CustomerLevelRepository) scanRow(row *sql.Row) (res models.CustomerLevel, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.CreatedDate,
		&res.ModifiedDate,
	)

	return
}

// FindAll ...
func (repository CustomerLevelRepository) FindAll(ctx context.Context, parameter models.CustomerLevelParameter) (data []models.CustomerLevel, err error) {
	conditionString := ``

	query := models.CustomerLevelSelectStatement + ` ` +
		models.CustomerLevelWhereStatement + ` ` +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.Query(query)
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
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, err
}
