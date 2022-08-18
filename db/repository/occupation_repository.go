package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IOccupationRepository ...
type IOccupationRepository interface {
	SelectAll(c context.Context, parameter models.OccupationParameter) ([]models.Occupation, error)
	FindAll(ctx context.Context, parameter models.OccupationParameter) ([]models.Occupation, int, error)
	FindByID(c context.Context, parameter models.OccupationParameter) (models.Occupation, error)
	FindByMappingName(c context.Context, parameter models.OccupationParameter) (models.Occupation, error)
	Add(c context.Context, model *models.Occupation) (string, error)
	Edit(c context.Context, model *models.Occupation) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// OccupationRepository ...
type OccupationRepository struct {
	DB *sql.DB
}

// NewOccupationRepository ...
func NewOccupationRepository(DB *sql.DB) IOccupationRepository {
	return &OccupationRepository{DB: DB}
}

// Scan rows
func (repository OccupationRepository) scanRows(rows *sql.Rows) (res models.Occupation, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository OccupationRepository) scanRow(row *sql.Row) (res models.Occupation, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository OccupationRepository) SelectAll(c context.Context, parameter models.OccupationParameter) (data []models.Occupation, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.OccupationSelectStatement + ` ` + models.OccupationWhereStatement +
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
func (repository OccupationRepository) FindAll(ctx context.Context, parameter models.OccupationParameter) (data []models.Occupation, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.OccupationSelectStatement + ` ` + models.OccupationWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "occupations" def ` + models.OccupationWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository OccupationRepository) FindByID(c context.Context, parameter models.OccupationParameter) (data models.Occupation, err error) {
	statement := models.OccupationSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository OccupationRepository) FindByMappingName(c context.Context, parameter models.OccupationParameter) (data models.Occupation, err error) {
	statement := models.OccupationSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository OccupationRepository) Add(c context.Context, model *models.Occupation) (res string, err error) {
	statement := `INSERT INTO occupations (name, mapping_name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository OccupationRepository) Edit(c context.Context, model *models.Occupation) (res string, err error) {
	statement := `UPDATE occupations SET name = $1, mapping_name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository OccupationRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE occupations SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
