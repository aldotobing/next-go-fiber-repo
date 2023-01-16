package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebUserRoleGroupRepository ...
type IWebUserRoleGroupRepository interface {
	SelectAll(c context.Context, parameter models.WebUserRoleGroupParameter) ([]models.WebUserRoleGroup, error)
	FindAll(ctx context.Context, parameter models.WebUserRoleGroupParameter) ([]models.WebUserRoleGroup, int, error)
	FindByID(c context.Context, parameter models.WebUserRoleGroupParameter) (models.WebUserRoleGroup, error)
}

// WebUserRoleGroupRepository ...
type WebUserRoleGroupRepository struct {
	DB *sql.DB
}

// NewWebUserRoleGroupRepository ...
func NewWebUserRoleGroupRepository(DB *sql.DB) IWebUserRoleGroupRepository {
	return &WebUserRoleGroupRepository{DB: DB}
}

// Scan rows
func (repository WebUserRoleGroupRepository) scanRows(rows *sql.Rows) (res models.WebUserRoleGroup, err error) {
	err = rows.Scan(
		&res.ID, &res.UserID, &res.RoleGroupID, &res.RoleGroupName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebUserRoleGroupRepository) scanRow(row *sql.Row) (res models.WebUserRoleGroup, err error) {
	err = row.Scan(
		&res.ID, &res.UserID, &res.RoleGroupID, &res.RoleGroupName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebUserRoleGroupRepository) SelectAll(c context.Context, parameter models.WebUserRoleGroupParameter) (data []models.WebUserRoleGroup, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += ` and def.user_id = ` + parameter.UserID
	}
	statement := models.WebUserRoleGroupSelectStatement + ` ` + models.WebUserRoleGroupWhereStatement +
		` AND (LOWER(rg."_name") LIKE $1  ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebUserRoleGroupRepository) FindAll(ctx context.Context, parameter models.WebUserRoleGroupParameter) (data []models.WebUserRoleGroup, count int, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += ` and def.user_id = ` + parameter.UserID
	}

	query := models.WebUserRoleGroupSelectStatement + ` ` + models.WebUserRoleGroupWhereStatement + ` ` + conditionString + `
		AND (LOWER(rg."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*)  FROM user_role_group def
	join role_group rg on rg.id = def.role_group_id ` + models.WebUserRoleGroupWhereStatement + ` ` +
		conditionString + ` AND (LOWER(rg."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebUserRoleGroupRepository) FindByID(c context.Context, parameter models.WebUserRoleGroupParameter) (data models.WebUserRoleGroup, err error) {
	statement := models.WebUserRoleGroupSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
