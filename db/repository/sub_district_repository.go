package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISubDistrictRepository ...
type ISubDistrictRepository interface {
	SelectAll(c context.Context, parameter models.SubDistrictParameter) ([]models.SubDistrict, error)
	FindAll(ctx context.Context, parameter models.SubDistrictParameter) ([]models.SubDistrict, int, error)
	FindByID(c context.Context, parameter models.SubDistrictParameter) (models.SubDistrict, error)
	FindByCode(c context.Context, parameter models.SubDistrictParameter) (models.SubDistrict, error)
	Add(c context.Context, model *models.SubDistrict) (string, error)
	Edit(c context.Context, model *models.SubDistrict) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// SubDistrictRepository ...
type SubDistrictRepository struct {
	DB *sql.DB
}

// NewSubDistrictRepository ...
func NewSubDistrictRepository(DB *sql.DB) ISubDistrictRepository {
	return &SubDistrictRepository{DB: DB}
}

// Scan rows
func (repository SubDistrictRepository) scanRows(rows *sql.Rows) (res models.SubDistrict, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository SubDistrictRepository) scanRow(row *sql.Row) (res models.SubDistrict, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SubDistrictRepository) SelectAll(c context.Context, parameter models.SubDistrictParameter) (data []models.SubDistrict, err error) {
	conditionString := ``

	if parameter.DistrictID != "" {
		conditionString += ` AND def.district_id = '` + parameter.DistrictID + `'`
	}
	if parameter.IDs != "" {
		conditionString += ` AND def.id in (` + parameter.IDs + `)`
	}

	statement := models.SubDistrictSelectStatement + ` ` + models.SubDistrictWhereStatement +
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
func (repository SubDistrictRepository) FindAll(ctx context.Context, parameter models.SubDistrictParameter) (data []models.SubDistrict, count int, err error) {
	conditionString := ``
	if parameter.DistrictID != "" {
		conditionString += ` AND def.district_id = '` + parameter.DistrictID + `'`
	}

	query := models.SubDistrictSelectStatement + ` ` + models.SubDistrictWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "subdistrict" def ` + models.SubDistrictWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository SubDistrictRepository) FindByID(c context.Context, parameter models.SubDistrictParameter) (data models.SubDistrict, err error) {
	statement := models.SubDistrictSelectStatement + ` WHERE def._name IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCode ...
func (repository SubDistrictRepository) FindByCode(c context.Context, parameter models.SubDistrictParameter) (data models.SubDistrict, err error) {
	statement := models.SubDistrictSelectStatement + ` WHERE def.deleted_at IS NULL AND def.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository SubDistrictRepository) Add(c context.Context, model *models.SubDistrict) (res string, err error) {
	statement := `INSERT INTO subdistricts (city_id, code, name, status)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository SubDistrictRepository) Edit(c context.Context, model *models.SubDistrict) (res string, err error) {
	statement := `UPDATE subdistricts SET city_id = $1, code = $2, name = $3, status = $4 WHERE id = $5 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository SubDistrictRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE subdistricts SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
