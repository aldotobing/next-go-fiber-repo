package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWeebRoleRepository ...
type IWeebRoleRepository interface {
	SelectAll(c context.Context, parameter models.WeebRoleParameter) ([]models.WeebRole, error)
	FindAll(ctx context.Context, parameter models.WeebRoleParameter) ([]models.WeebRole, int, error)
	FindByID(c context.Context, parameter models.WeebRoleParameter) (models.WeebRole, error)
	Add(c context.Context, model *models.WeebRole) (*string, error)
	Edit(c context.Context, model *models.WeebRole) (*string, error)
	Delete(c context.Context, id string, now time.Time) (*string, error)
}

// WeebRoleRepository ...
type WeebRoleRepository struct {
	DB *sql.DB
}

// NewWeebRoleRepository ...
func NewWeebRoleRepository(DB *sql.DB) IWeebRoleRepository {
	return &WeebRoleRepository{DB: DB}
}

// Scan rows
func (repository WeebRoleRepository) scanRows(rows *sql.Rows) (res models.WeebRole, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Header,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WeebRoleRepository) scanRow(row *sql.Row) (res models.WeebRole, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Header,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WeebRoleRepository) SelectAll(c context.Context, parameter models.WeebRoleParameter) (data []models.WeebRole, err error) {
	conditionString := ``

	statement := models.WeebRoleSelectStatement + ` ` + models.WeebRoleWhereStatement +
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
func (repository WeebRoleRepository) FindAll(ctx context.Context, parameter models.WeebRoleParameter) (data []models.WeebRole, count int, err error) {
	conditionString := ``

	query := models.WeebRoleSelectStatement + ` ` + models.WeebRoleWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "role" def ` + models.WeebRoleWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WeebRoleRepository) FindByID(c context.Context, parameter models.WeebRoleParameter) (data models.WeebRole, err error) {
	statement := models.WeebRoleSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository WeebRoleRepository) Add(c context.Context, model *models.WeebRole) (res *string, err error) {
	statement := `INSERT INTO role (_name,_header, created_date, modified_date ,is_mysm, id)
	VALUES ($1, $2, now(),now(),1, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.Header, model.ID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository WeebRoleRepository) Edit(c context.Context, model *models.WeebRole) (res *string, err error) {
	statement := `UPDATE role SET 
	_name = $1,_header = $2, modified_date = now()
	 WHERE id = $3 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.Header,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository WeebRoleRepository) Delete(c context.Context, id string, now time.Time) (res *string, err error) {
	statement := `UPDATE role SET deleted_at = now() WHERE id = $1 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
