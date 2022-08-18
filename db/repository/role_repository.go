package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IRoleRepository ...
type IRoleRepository interface {
	SelectAll(c context.Context, parameter models.RoleParameter) ([]models.Role, error)
	FindAll(ctx context.Context, parameter models.RoleParameter) ([]models.Role, int, error)
	FindByID(c context.Context, parameter models.RoleParameter) (models.Role, error)
	Add(c context.Context, model *models.Role) (string, error)
	Edit(c context.Context, model *models.Role) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// RoleRepository ...
type RoleRepository struct {
	DB *sql.DB
}

// NewRoleRepository ...
func NewRoleRepository(DB *sql.DB) IRoleRepository {
	return &RoleRepository{DB: DB}
}

// Scan rows
func (repository RoleRepository) scanRows(rows *sql.Rows) (res models.Role, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository RoleRepository) scanRow(row *sql.Row) (res models.Role, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository RoleRepository) SelectAll(c context.Context, parameter models.RoleParameter) (data []models.Role, err error) {
	conditionString := ``

	statement := models.RoleSelectStatement + ` ` + models.RoleWhereStatement +
		` AND (LOWER(def."code") LIKE $1 OR LOWER(def."name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository RoleRepository) FindAll(ctx context.Context, parameter models.RoleParameter) (data []models.Role, count int, err error) {
	conditionString := ``

	query := models.RoleSelectStatement + ` ` + models.RoleWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."code") LIKE $1 OR LOWER(def."name") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "roles" def ` + models.RoleWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository RoleRepository) FindByID(c context.Context, parameter models.RoleParameter) (data models.Role, err error) {
	statement := models.RoleSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository RoleRepository) Add(c context.Context, model *models.Role) (res string, err error) {
	statement := `INSERT INTO roles (code, name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository RoleRepository) Edit(c context.Context, model *models.Role) (res string, err error) {
	statement := `UPDATE roles SET code = $1, name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository RoleRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE roles SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
