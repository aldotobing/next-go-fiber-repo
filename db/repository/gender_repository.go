package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IGenderRepository ...
type IGenderRepository interface {
	SelectAll(c context.Context, parameter models.GenderParameter) ([]models.Gender, error)
	FindAll(ctx context.Context, parameter models.GenderParameter) ([]models.Gender, int, error)
	FindByID(c context.Context, parameter models.GenderParameter) (models.Gender, error)
	FindByMappingName(c context.Context, parameter models.GenderParameter) (models.Gender, error)
	Add(c context.Context, model *models.Gender) (string, error)
	Edit(c context.Context, model *models.Gender) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// GenderRepository ...
type GenderRepository struct {
	DB *sql.DB
}

// NewGenderRepository ...
func NewGenderRepository(DB *sql.DB) IGenderRepository {
	return &GenderRepository{DB: DB}
}

// Scan rows
func (repository GenderRepository) scanRows(rows *sql.Rows) (res models.Gender, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository GenderRepository) scanRow(row *sql.Row) (res models.Gender, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository GenderRepository) SelectAll(c context.Context, parameter models.GenderParameter) (data []models.Gender, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.GenderSelectStatement + ` ` + models.GenderWhereStatement +
		` AND (LOWER(def."name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, parameter.Search)
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
func (repository GenderRepository) FindAll(ctx context.Context, parameter models.GenderParameter) (data []models.Gender, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.GenderSelectStatement + ` ` + models.GenderWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."name") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "genders" def ` + models.GenderWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository GenderRepository) FindByID(c context.Context, parameter models.GenderParameter) (data models.Gender, err error) {
	statement := models.GenderSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository GenderRepository) FindByMappingName(c context.Context, parameter models.GenderParameter) (data models.Gender, err error) {
	statement := models.GenderSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository GenderRepository) Add(c context.Context, model *models.Gender) (res string, err error) {
	statement := `INSERT INTO genders (name, mapping_name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository GenderRepository) Edit(c context.Context, model *models.Gender) (res string, err error) {
	statement := `UPDATE genders SET name = $1, mapping_name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository GenderRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE genders SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
