package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesmanTypeRepository ...
type ISalesmanTypeRepository interface {
	SelectAll(c context.Context, parameter models.SalesmanParameter) ([]models.SalesmanType, error)
}

// SalesmanTypeRepository ...
type SalesmanTypeRepository struct {
	DB *sql.DB
}

// NewSalesmanTypeRepository ...
func NewSalesmanTypeRepository(DB *sql.DB) ISalesmanTypeRepository {
	return &SalesmanTypeRepository{DB: DB}
}

// Scan rows
func (repository SalesmanTypeRepository) scanRows(rows *sql.Rows) (res models.SalesmanType, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
	)

	return
}

// Scan row
func (repository SalesmanTypeRepository) scanRow(row *sql.Row) (res models.Salesman, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
	)

	return
}

// SelectAll ...
func (repository SalesmanTypeRepository) SelectAll(c context.Context, in models.SalesmanParameter) (out []models.SalesmanType, err error) {
	conditionString := ``

	statement := models.SalesmanTypeSelectStatement + ` ` + models.SalesmanTypeWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + in.By + ` ` + in.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(in.Search)+"%")
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanRows(rows)
		if err != nil {
			return out, err
		}
		out = append(out, temp)
	}

	return
}
