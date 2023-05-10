package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebUserBranchRepository ...
type IWebUserBranchRepository interface {
	SelectAll(c context.Context, parameter models.WebUserBranchParameter) ([]models.WebUserBranch, error)
	FindAll(ctx context.Context, parameter models.WebUserBranchParameter) ([]models.WebUserBranch, int, error)
	FindByID(c context.Context, parameter models.WebUserBranchParameter) (models.WebUserBranch, error)
}

// WebUserBranchRepository ...
type WebUserBranchRepository struct {
	DB *sql.DB
}

// NewWebUserBranchRepository ...
func NewWebUserBranchRepository(DB *sql.DB) IWebUserBranchRepository {
	return &WebUserBranchRepository{DB: DB}
}

// Scan rows
func (repository WebUserBranchRepository) scanRows(rows *sql.Rows) (res models.WebUserBranch, err error) {
	err = rows.Scan(
		&res.ID, &res.UserID, &res.BranchID, &res.BranchName, &res.BranchCode,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebUserBranchRepository) scanRow(row *sql.Row) (res models.WebUserBranch, err error) {
	err = row.Scan(
		&res.ID, &res.UserID, &res.BranchID, &res.BranchName, &res.BranchCode,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebUserBranchRepository) SelectAll(c context.Context, parameter models.WebUserBranchParameter) (data []models.WebUserBranch, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += ` and def.user_id = ` + parameter.UserID
	}
	statement := models.WebUserBranchSelectStatement + ` ` + models.WebUserBranchWhereStatement +
		` AND (LOWER(br."_name") LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebUserBranchRepository) FindAll(ctx context.Context, parameter models.WebUserBranchParameter) (data []models.WebUserBranch, count int, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += ` and def.user_id = ` + parameter.UserID
	}

	query := models.WebUserBranchSelectStatement + ` ` + models.WebUserBranchWhereStatement + ` ` + conditionString + `
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
	join role_group rg on rg.id = def.role_group_id ` + models.WebUserBranchWhereStatement + ` ` +
		conditionString + ` AND (LOWER(rg."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebUserBranchRepository) FindByID(c context.Context, parameter models.WebUserBranchParameter) (data models.WebUserBranch, err error) {
	statement := models.WebUserBranchSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
