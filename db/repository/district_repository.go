package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IDistrictRepository ...
type IDistrictRepository interface {
	SelectAll(c context.Context, parameter models.DistrictParameter) ([]models.District, error)
	FindAll(ctx context.Context, parameter models.DistrictParameter) ([]models.District, int, error)
	FindByID(c context.Context, parameter models.DistrictParameter) (models.District, error)
	FindByCode(c context.Context, parameter models.DistrictParameter) (models.District, error)
	Add(c context.Context, model *models.District) (string, error)
	Edit(c context.Context, model *models.District) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// DistrictRepository ...
type DistrictRepository struct {
	DB *sql.DB
}

// NewDistrictRepository ...
func NewDistrictRepository(DB *sql.DB) IDistrictRepository {
	return &DistrictRepository{DB: DB}
}

// Scan rows
func (repository DistrictRepository) scanRows(rows *sql.Rows) (res models.District, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository DistrictRepository) scanRow(row *sql.Row) (res models.District, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository DistrictRepository) SelectAll(c context.Context, parameter models.DistrictParameter) (data []models.District, err error) {
	conditionString := ``

	if parameter.CityID != "" {
		conditionString += ` AND def.city_id = '` + parameter.CityID + `'`
	}

	if parameter.IDs != "" {
		conditionString += ` AND def.id in (` + parameter.IDs + `)`
	}

	statement := models.DistrictSelectStatement + ` ` + models.DistrictWhereStatement +
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
func (repository DistrictRepository) FindAll(ctx context.Context, parameter models.DistrictParameter) (data []models.District, count int, err error) {
	conditionString := ``
	if parameter.CityID != "" {
		conditionString += ` AND def.city_id = '` + parameter.CityID + `'`
	}

	query := models.DistrictSelectStatement + ` ` + models.DistrictWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "district" def ` + models.DistrictWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository DistrictRepository) FindByID(c context.Context, parameter models.DistrictParameter) (data models.District, err error) {
	statement := models.DistrictSelectStatement + ` WHERE def._name IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCode ...
func (repository DistrictRepository) FindByCode(c context.Context, parameter models.DistrictParameter) (data models.District, err error) {
	statement := models.DistrictSelectStatement + ` WHERE def.deleted_at IS NULL AND def.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository DistrictRepository) Add(c context.Context, model *models.District) (res string, err error) {
	statement := `INSERT INTO districts (city_id, code, name, status)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository DistrictRepository) Edit(c context.Context, model *models.District) (res string, err error) {
	statement := `UPDATE districts SET city_id = $1, code = $2, name = $3, status = $4 WHERE id = $5 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository DistrictRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE districts SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
