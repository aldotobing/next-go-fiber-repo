package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IBranchRepository ...
type IBranchRepository interface {
	SelectAll(c context.Context, parameter models.BranchParameter) ([]models.Branch, error)
	FindAll(ctx context.Context, parameter models.BranchParameter) ([]models.Branch, int, error)
	FindByID(c context.Context, parameter models.BranchParameter) (models.Branch, error)
	Update(c context.Context, in models.Branch) (res *string, err error)
}

// BranchRepository ...
type BranchRepository struct {
	DB *sql.DB
}

// NewBranchRepository ...
func NewBranchRepository(DB *sql.DB) IBranchRepository {
	return &BranchRepository{DB: DB}
}

// Scan rows
func (repository BranchRepository) scanRows(rows *sql.Rows) (res models.Branch, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Name,
		&res.Code,
		&res.Area,
		&res.RegionID,
		&res.RegionName,
		&res.RegionGroupID,
		&res.RegionGroupName,
		&res.PICPhoneNo,
		&res.PICName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository BranchRepository) scanRow(row *sql.Row) (res models.Branch, err error) {
	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Code,
		&res.Area,
		&res.RegionID,
		&res.RegionName,
		&res.RegionGroupID,
		&res.RegionGroupName,
		&res.PICPhoneNo,
		&res.PICName,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository BranchRepository) SelectAll(c context.Context, parameter models.BranchParameter) (data []models.Branch, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += " and def.id in (select branch_id from user_branch where user_id = " + parameter.UserID + ")"
	}

	if parameter.RegionID != "" && parameter.RegionID != "0" {
		conditionString += " and def.region_id = '" + parameter.RegionID + "'"
	}

	statement := models.BranchSelectStatement + ` ` + models.BranchWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository BranchRepository) FindAll(ctx context.Context, parameter models.BranchParameter) (data []models.Branch, count int, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += " and def.id in (select branch_id from user_branch where user_id = " + parameter.UserID + ")"
	}
	query := models.BranchSelectStatement + ` ` + models.BranchWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "branch" def ` + models.BranchWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository BranchRepository) FindByID(c context.Context, parameter models.BranchParameter) (data models.Branch, err error) {
	statement := models.BranchSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository BranchRepository) Update(c context.Context, in models.Branch) (res *string, err error) {
	statement := `UPDATE branch SET 
			pic_phone_no = $1
			pic_name = $2
		WHERE id = $3
		RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		in.PICPhoneNo,
		in.PICName,
		in.ID).Scan(&res)
	if err != nil {
		return
	}

	return
}
