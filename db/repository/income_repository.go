package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IIncomeRepository ...
type IIncomeRepository interface {
	SelectAll(c context.Context, parameter models.IncomeParameter) ([]models.Income, error)
	FindAll(ctx context.Context, parameter models.IncomeParameter) ([]models.Income, int, error)
	FindByID(c context.Context, parameter models.IncomeParameter) (models.Income, error)
	FindByMappingName(c context.Context, parameter models.IncomeParameter) (models.Income, error)
	Add(c context.Context, model *models.Income) (string, error)
	Edit(c context.Context, model *models.Income) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// IncomeRepository ...
type IncomeRepository struct {
	DB *sql.DB
}

// NewIncomeRepository ...
func NewIncomeRepository(DB *sql.DB) IIncomeRepository {
	return &IncomeRepository{DB: DB}
}

// Scan rows
func (repository IncomeRepository) scanRows(rows *sql.Rows) (res models.Income, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.MinValue, &res.MaxValue, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository IncomeRepository) scanRow(row *sql.Row) (res models.Income, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.MinValue, &res.MaxValue, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository IncomeRepository) SelectAll(c context.Context, parameter models.IncomeParameter) (data []models.Income, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.IncomeSelectStatement + ` ` + models.IncomeWhereStatement +
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
func (repository IncomeRepository) FindAll(ctx context.Context, parameter models.IncomeParameter) (data []models.Income, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.IncomeSelectStatement + ` ` + models.IncomeWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "incomes" def ` + models.IncomeWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository IncomeRepository) FindByID(c context.Context, parameter models.IncomeParameter) (data models.Income, err error) {
	statement := models.IncomeSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository IncomeRepository) FindByMappingName(c context.Context, parameter models.IncomeParameter) (data models.Income, err error) {
	statement := models.IncomeSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository IncomeRepository) Add(c context.Context, model *models.Income) (res string, err error) {
	statement := `INSERT INTO incomes (name, mapping_name, min_value, max_value, status)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.MinValue, model.MaxValue, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository IncomeRepository) Edit(c context.Context, model *models.Income) (res string, err error) {
	statement := `UPDATE incomes SET name = $1, mapping_name = $2, min_value = $3, max_value = $4, status = $5 WHERE id = $6 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.MinValue, model.MaxValue, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository IncomeRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE incomes SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
