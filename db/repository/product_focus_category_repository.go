package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IProductFocusCategoryRepository ...
type IProductFocusCategoryRepository interface {
	SelectAll(c context.Context, parameter models.ProductFocusCategoryParameter) ([]models.ProductFocusCategory, error)
	FindAll(ctx context.Context, parameter models.ProductFocusCategoryParameter) ([]models.ProductFocusCategory, int, error)
	FindByID(c context.Context, parameter models.ProductFocusCategoryParameter) (models.ProductFocusCategory, error)
	FindByBranchID(c context.Context, parameter models.ProductFocusCategoryParameter) (models.ProductFocusCategory, error)
	// Add(c context.Context, model *models.ProductFocusCategory) (*string, error)
	// Edit(c context.Context, model *models.ProductFocusCategory) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ProductFocusCategoryRepository ...
type ProductFocusCategoryRepository struct {
	DB *sql.DB
}

// NewProductFocusCategoryRepository ...
func NewProductFocusCategoryRepository(DB *sql.DB) IProductFocusCategoryRepository {
	return &ProductFocusCategoryRepository{DB: DB}
}

// Scan rows
func (repository ProductFocusCategoryRepository) scanRows(rows *sql.Rows) (res models.ProductFocusCategory, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name, &res.Foto,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ProductFocusCategoryRepository) scanRow(row *sql.Row) (res models.ProductFocusCategory, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name, &res.Foto,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ProductFocusCategoryRepository) SelectAll(c context.Context, parameter models.ProductFocusCategoryParameter) (data []models.ProductFocusCategory, err error) {
	conditionString := ``

	statement := models.ProductFocusCategorySelectStatement + ` ` + models.ProductFocusCategoryWhereStatement +
		` AND (LOWER(ic."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository ProductFocusCategoryRepository) FindAll(ctx context.Context, parameter models.ProductFocusCategoryParameter) (data []models.ProductFocusCategory, count int, err error) {
	conditionString := ``

	query := models.ProductFocusCategorySelectStatement + ` ` + models.ProductFocusCategoryWhereStatement + ` ` + conditionString + `
		AND (LOWER(ic."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "product_focus" def ` + models.ProductFocusCategoryWhereStatement + ` ` +
		conditionString + ` AND (LOWER(ic."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ProductFocusCategoryRepository) FindByID(c context.Context, parameter models.ProductFocusCategoryParameter) (data models.ProductFocusCategory, err error) {
	statement := models.ProductFocusCategorySelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByBranchID ...
func (repository ProductFocusCategoryRepository) FindByBranchID(c context.Context, parameter models.ProductFocusCategoryParameter) (data models.ProductFocusCategory, err error) {
	statement := models.ProductFocusCategorySelectStatement + ` WHERE def.created_date IS NOT NULL AND def.branch_id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
