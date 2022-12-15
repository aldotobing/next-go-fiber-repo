package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IDashboardWebRepository ...
type IDashboardWebRepository interface {
	GetData(c context.Context, parameter models.DashboardWebParameter) (models.DashboardWeb, error)
}

// DashboardWebRepository ...
type DashboardWebRepository struct {
	DB *sql.DB
}

// NewDashboardWebRepository ...
func NewDashboardWebRepository(DB *sql.DB) IDashboardWebRepository {
	return &DashboardWebRepository{DB: DB}
}

// Scan row
func (repository DashboardWebRepository) scanRow(row *sql.Row) (res models.DashboardWeb, err error) {
	err = row.Scan(
		&res.TotalActiveUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) (data models.DashboardWeb, err error) {
	statement := models.DashboardWebSelectStatement
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
