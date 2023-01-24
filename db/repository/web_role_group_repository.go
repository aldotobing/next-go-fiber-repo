package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebRoleGroupRepository ...
type IWebRoleGroupRepository interface {
	SelectAll(c context.Context, parameter models.WebRoleGroupParameter) ([]models.WebRoleGroup, error)
	FindAll(ctx context.Context, parameter models.WebRoleGroupParameter) ([]models.WebRoleGroup, int, error)
	FindByID(c context.Context, parameter models.WebRoleGroupParameter) (models.WebRoleGroup, error)
	Add(c context.Context, model *models.WebRoleGroup) (*string, error)
	Edit(c context.Context, model *models.WebRoleGroup) (*string, error)
	Delete(c context.Context, id string, now time.Time) (*string, error)
}

// WebRoleGroupRepository ...
type WebRoleGroupRepository struct {
	DB *sql.DB
}

// NewWebRoleGroupRepository ...
func NewWebRoleGroupRepository(DB *sql.DB) IWebRoleGroupRepository {
	return &WebRoleGroupRepository{DB: DB}
}

// Scan rows
func (repository WebRoleGroupRepository) scanRows(rows *sql.Rows) (res models.WebRoleGroup, err error) {
	err = rows.Scan(
		&res.ID, &res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebRoleGroupRepository) scanRow(row *sql.Row) (res models.WebRoleGroup, err error) {
	err = row.Scan(
		&res.ID, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebRoleGroupRepository) SelectAll(c context.Context, parameter models.WebRoleGroupParameter) (data []models.WebRoleGroup, err error) {
	conditionString := ``

	statement := models.WebRoleGroupSelectStatement + ` ` + models.WebRoleGroupWhereStatement +
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
func (repository WebRoleGroupRepository) FindAll(ctx context.Context, parameter models.WebRoleGroupParameter) (data []models.WebRoleGroup, count int, err error) {
	conditionString := ``

	query := models.WebRoleGroupSelectStatement + ` ` + models.WebRoleGroupWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "role_group" def ` + models.WebRoleGroupWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebRoleGroupRepository) FindByID(c context.Context, parameter models.WebRoleGroupParameter) (data models.WebRoleGroup, err error) {
	statement := models.WebRoleGroupSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository WebRoleGroupRepository) Add(c context.Context, model *models.WebRoleGroup) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO role_group (_name,created_date, modified_date ,is_mysm)
	VALUES ($1, now(),now(),1) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Name).Scan(&res)

	if err != nil {
		return res, err
	}
	ReleGroupID := &res

	parts := strings.Split(*model.RoleListID, ",")
	if len(parts) >= 1 {
		for pi, _ := range parts {
			linestatement := `INSERT INTO role_group_role_line (role_group_id,role_id,created_date, modified_date)
					VALUES ($1,$2, now(),now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, ReleGroupID, parts[pi]).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository WebRoleGroupRepository) Edit(c context.Context, model *models.WebRoleGroup) (res *string, err error) {
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `UPDATE role_group SET 
	_name = $1, modified_date = now()
	 WHERE id = $2 RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Name,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	ReleGroupID := &res

	parts := strings.Split(*model.RoleListID, ",")

	if len(parts) >= 1 {
		deletelinestatement := `delete from role_group_role_line WHERE role_group_id = $1`

		deletedRow, _ := transaction.QueryContext(c, deletelinestatement, ReleGroupID)
		deletedRow.Close()

		for pi, _ := range parts {
			linestatement := `INSERT INTO role_group_role_line (role_group_id,role_id,created_date, modified_date)
						VALUES ($1,$2, now(),now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, ReleGroupID, parts[pi]).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (repository WebRoleGroupRepository) Delete(c context.Context, id string, now time.Time) (res *string, err error) {
	statement := `UPDATE role_group SET deleted_at = now() WHERE id = $1 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
