package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IMaritalStatusRepository ...
type IMaritalStatusRepository interface {
	SelectAll(c context.Context, parameter models.MaritalStatusParameter) ([]models.MaritalStatus, error)
	FindAll(ctx context.Context, parameter models.MaritalStatusParameter) ([]models.MaritalStatus, int, error)
	FindByID(c context.Context, parameter models.MaritalStatusParameter) (models.MaritalStatus, error)
	FindByMappingName(c context.Context, parameter models.MaritalStatusParameter) (models.MaritalStatus, error)
	Add(c context.Context, model *models.MaritalStatus) (string, error)
	Edit(c context.Context, model *models.MaritalStatus) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// MaritalStatusRepository ...
type MaritalStatusRepository struct {
	DB *sql.DB
}

// NewMaritalStatusRepository ...
func NewMaritalStatusRepository(DB *sql.DB) IMaritalStatusRepository {
	return &MaritalStatusRepository{DB: DB}
}

// Scan rows
func (repository MaritalStatusRepository) scanRows(rows *sql.Rows) (res models.MaritalStatus, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.FillSpouseName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository MaritalStatusRepository) scanRow(row *sql.Row) (res models.MaritalStatus, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.FillSpouseName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository MaritalStatusRepository) SelectAll(c context.Context, parameter models.MaritalStatusParameter) (data []models.MaritalStatus, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.MaritalStatusSelectStatement + ` ` + models.MaritalStatusWhereStatement +
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
func (repository MaritalStatusRepository) FindAll(ctx context.Context, parameter models.MaritalStatusParameter) (data []models.MaritalStatus, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.MaritalStatusSelectStatement + ` ` + models.MaritalStatusWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "marital_statuses" def ` + models.MaritalStatusWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository MaritalStatusRepository) FindByID(c context.Context, parameter models.MaritalStatusParameter) (data models.MaritalStatus, err error) {
	statement := models.MaritalStatusSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository MaritalStatusRepository) FindByMappingName(c context.Context, parameter models.MaritalStatusParameter) (data models.MaritalStatus, err error) {
	statement := models.MaritalStatusSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository MaritalStatusRepository) Add(c context.Context, model *models.MaritalStatus) (res string, err error) {
	statement := `INSERT INTO marital_statuses (name, mapping_name, fill_spouse_name, status)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.FillSpouseName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository MaritalStatusRepository) Edit(c context.Context, model *models.MaritalStatus) (res string, err error) {
	statement := `UPDATE marital_statuses SET name = $1, mapping_name = $2, fill_spouse_name = $3, status = $4 WHERE id = $5 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.FillSpouseName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository MaritalStatusRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE marital_statuses SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
