package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebRegionAreaRepository ...
type IWebRegionAreaRepository interface {
	SelectAll(c context.Context, parameter models.WebRegionAreaParameter) ([]models.WebRegionArea, error)
	FindAll(ctx context.Context, parameter models.WebRegionAreaParameter) ([]models.WebRegionArea, int, error)
	FindByID(c context.Context, parameter models.WebRegionAreaParameter) (models.WebRegionArea, error)
}

// WebRegionAreaRepository ...
type WebRegionAreaRepository struct {
	DB *sql.DB
}

// NewWebRegionAreaRepository ...
func NewWebRegionAreaRepository(DB *sql.DB) IWebRegionAreaRepository {
	return &WebRegionAreaRepository{DB: DB}
}

// Scan rows
func (repository WebRegionAreaRepository) scanRows(rows *sql.Rows) (res models.WebRegionArea, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Code, &res.GroupName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebRegionAreaRepository) scanRow(row *sql.Row) (res models.WebRegionArea, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Code, &res.GroupName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebRegionAreaRepository) SelectAll(c context.Context, parameter models.WebRegionAreaParameter) (data []models.WebRegionArea, err error) {
	conditionString := ``

	statement := models.WebRegionAreaSelectStatement + ` ` + models.WebRegionAreaWhereStatement +
		` AND (LOWER(def."_name") LIKE $1  ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebRegionAreaRepository) FindAll(ctx context.Context, parameter models.WebRegionAreaParameter) (data []models.WebRegionArea, count int, err error) {
	conditionString := ``

	query := models.WebRegionAreaSelectStatement + ` ` + models.WebRegionAreaWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "role" def ` + models.WebRegionAreaWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebRegionAreaRepository) FindByID(c context.Context, parameter models.WebRegionAreaParameter) (data models.WebRegionArea, err error) {
	statement := models.WebRegionAreaSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
