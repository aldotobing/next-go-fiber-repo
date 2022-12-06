package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemCategoryRepository ...
type IItemCategoryRepository interface {
	SelectAll(c context.Context, parameter models.ItemCategoryParameter) ([]models.ItemCategory, error)
	FindAll(ctx context.Context, parameter models.ItemCategoryParameter) ([]models.ItemCategory, int, error)
	FindByID(c context.Context, parameter models.ItemCategoryParameter) (models.ItemCategory, error)
	// Add(c context.Context, model *models.ItemCategory) (*string, error)
	// Edit(c context.Context, model *models.ItemCategory) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemCategoryRepository ...
type ItemCategoryRepository struct {
	DB *sql.DB
}

// NewItemCategoryRepository ...
func NewItemCategoryRepository(DB *sql.DB) IItemCategoryRepository {
	return &ItemCategoryRepository{DB: DB}
}

// Scan rows
func (repository ItemCategoryRepository) scanRows(rows *sql.Rows) (res models.ItemCategory, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Image,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemCategoryRepository) scanRow(row *sql.Row) (res models.ItemCategory, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Image,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemCategoryRepository) SelectAll(c context.Context, parameter models.ItemCategoryParameter) (data []models.ItemCategory, err error) {
	conditionString := ``

	statement := models.ItemCategorySelectStatement + ` ` + models.ItemCategoryWhereStatement +
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
func (repository ItemCategoryRepository) FindAll(ctx context.Context, parameter models.ItemCategoryParameter) (data []models.ItemCategory, count int, err error) {
	conditionString := ``

	query := models.ItemCategorySelectStatement + ` ` + models.ItemCategoryWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "item_category" def ` + models.ItemCategoryWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemCategoryRepository) FindByID(c context.Context, parameter models.ItemCategoryParameter) (data models.ItemCategory, err error) {
	statement := models.ItemCategorySelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
