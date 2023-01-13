package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebRoleGroupRoleLineRepository ...
type IWebRoleGroupRoleLineRepository interface {
	SelectAll(c context.Context, parameter models.WebRoleGroupRoleLineParameter) ([]models.WebRoleGroupRoleLine, error)
	FindAll(ctx context.Context, parameter models.WebRoleGroupRoleLineParameter) ([]models.WebRoleGroupRoleLine, int, error)
	FindByID(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (models.WebRoleGroupRoleLine, error)
	Add(c context.Context, model *models.WebRoleGroupRoleLine) (*string, error)
	Delete(c context.Context, id string, now time.Time) (*string, error)
}

// WebRoleGroupRoleLineRepository ...
type WebRoleGroupRoleLineRepository struct {
	DB *sql.DB
}

// NewWebRoleGroupRoleLineRepository ...
func NewWebRoleGroupRoleLineRepository(DB *sql.DB) IWebRoleGroupRoleLineRepository {
	return &WebRoleGroupRoleLineRepository{DB: DB}
}

// Scan rows
func (repository WebRoleGroupRoleLineRepository) scanRows(rows *sql.Rows) (res models.WebRoleGroupRoleLine, err error) {
	err = rows.Scan(
		&res.ID, &res.RoleID, &res.RoleGroupID,
		&res.RoleName, &res.RoleGroupName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebRoleGroupRoleLineRepository) scanRow(row *sql.Row) (res models.WebRoleGroupRoleLine, err error) {
	err = row.Scan(
		&res.ID, &res.RoleID, &res.RoleGroupID,
		&res.RoleName, &res.RoleGroupName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebRoleGroupRoleLineRepository) SelectAll(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (data []models.WebRoleGroupRoleLine, err error) {
	conditionString := ``

	if parameter.RoleGroupID != "" {
		conditionString += ` and  def.role_group_id = ` + parameter.RoleGroupID
	}

	statement := models.WebRoleGroupRoleLineSelectStatement + ` ` + models.WebRoleGroupRoleLineWhereStatement +
		` AND (LOWER(rl."_name") LIKE $1  ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebRoleGroupRoleLineRepository) FindAll(ctx context.Context, parameter models.WebRoleGroupRoleLineParameter) (data []models.WebRoleGroupRoleLine, count int, err error) {
	conditionString := ``

	if parameter.RoleGroupID != "" {
		conditionString += ` and  def.role_group_id = ` + parameter.RoleGroupID
	}

	query := models.WebRoleGroupRoleLineSelectStatement + ` ` + models.WebRoleGroupRoleLineWhereStatement + ` ` + conditionString + `
		AND (LOWER(rl."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "role_group_role_line" def 
	join role rl on rl.id = def.role_id
	join role_group rg on rg.id = def.role_group_id
	` + models.WebRoleGroupRoleLineWhereStatement + ` ` +
		conditionString + ` AND (LOWER(rl."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebRoleGroupRoleLineRepository) FindByID(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (data models.WebRoleGroupRoleLine, err error) {
	statement := models.WebRoleGroupRoleLineSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository WebRoleGroupRoleLineRepository) Add(c context.Context, model *models.WebRoleGroupRoleLine) (res *string, err error) {
	statement := `INSERT INTO role_group_role_line (role_group_id,role_id,created_date, modified_date)
	VALUES ($1,$2, now(),now()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.RoleGroupID, model.RoleID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository WebRoleGroupRoleLineRepository) Delete(c context.Context, id string, now time.Time) (res *string, err error) {
	deletelinestatement := `delete from role_group_role_line WHERE id = $1`

	deletedRow, _ := repository.DB.QueryContext(c, deletelinestatement, id)
	deletedRow.Close()

	res = &id
	return res, err
}
