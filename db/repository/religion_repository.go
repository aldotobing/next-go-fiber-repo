package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IReligionRepository ...
type IReligionRepository interface {
	SelectAll(c context.Context, parameter models.ReligionParameter) ([]models.Religion, error)
	FindAll(ctx context.Context, parameter models.ReligionParameter) ([]models.Religion, int, error)
	FindByID(c context.Context, parameter models.ReligionParameter) (models.Religion, error)
	FindByMappingName(c context.Context, parameter models.ReligionParameter) (models.Religion, error)
	Add(c context.Context, model *models.Religion) (string, error)
	Edit(c context.Context, model *models.Religion) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// ReligionRepository ...
type ReligionRepository struct {
	DB *sql.DB
}

// NewReligionRepository ...
func NewReligionRepository(DB *sql.DB) IReligionRepository {
	return &ReligionRepository{DB: DB}
}

// Scan rows
func (repository ReligionRepository) scanRows(rows *sql.Rows) (res models.Religion, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository ReligionRepository) scanRow(row *sql.Row) (res models.Religion, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ReligionRepository) SelectAll(c context.Context, parameter models.ReligionParameter) (data []models.Religion, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.ReligionSelectStatement + ` ` + models.ReligionWhereStatement +
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
func (repository ReligionRepository) FindAll(ctx context.Context, parameter models.ReligionParameter) (data []models.Religion, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.ReligionSelectStatement + ` ` + models.ReligionWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "religions" def ` + models.ReligionWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository ReligionRepository) FindByID(c context.Context, parameter models.ReligionParameter) (data models.Religion, err error) {
	statement := models.ReligionSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository ReligionRepository) FindByMappingName(c context.Context, parameter models.ReligionParameter) (data models.Religion, err error) {
	statement := models.ReligionSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ReligionRepository) Add(c context.Context, model *models.Religion) (res string, err error) {
	statement := `INSERT INTO religions (name, mapping_name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ReligionRepository) Edit(c context.Context, model *models.Religion) (res string, err error) {
	statement := `UPDATE religions SET name = $1, mapping_name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository ReligionRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE religions SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
