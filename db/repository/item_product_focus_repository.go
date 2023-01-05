package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemProductFocusRepository ...
type IItemProductFocusRepository interface {
	SelectAll(c context.Context, parameter models.ItemProductFocusParameter) ([]models.ItemProductFocus, error)
	FindAll(ctx context.Context, parameter models.ItemProductFocusParameter) ([]models.ItemProductFocus, int, error)
	FindByID(c context.Context, parameter models.ItemProductFocusParameter) (models.ItemProductFocus, error)
	// Add(c context.Context, model *models.ItemProductFocus) (*string, error)
	// Edit(c context.Context, model *models.ItemProductFocus) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemProductFocusRepository ...
type ItemProductFocusRepository struct {
	DB *sql.DB
}

// NewItemProductFocusRepository ...
func NewItemProductFocusRepository(DB *sql.DB) IItemProductFocusRepository {
	return &ItemProductFocusRepository{DB: DB}
}

// Scan rows
func (repository ItemProductFocusRepository) scanRows(rows *sql.Rows) (res models.ItemProductFocus, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name, &res.Description, &res.ItemCategoryId, &res.ItemCategoryName, &res.ItemPicture,
		&res.UomID, &res.UomName, &res.UomLineConversion, &res.ItemPrice, &res.PriceListVersionId,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemProductFocusRepository) scanRow(row *sql.Row) (res models.ItemProductFocus, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name, &res.Description, &res.ItemCategoryId, &res.ItemCategoryName, &res.ItemPicture,
		&res.UomID, &res.UomName, &res.UomLineConversion, &res.ItemPrice, &res.PriceListVersionId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemProductFocusRepository) SelectAll(c context.Context, parameter models.ItemProductFocusParameter) (data []models.ItemProductFocus, err error) {
	conditionString := ``

	if parameter.ItemCategoryId != "" {
		conditionString += ` AND i.Item_category_id = '` + parameter.ItemCategoryId + `'`
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionId + `'`
	}

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND i.id NOT IN (83, 307, 393) `
	}

	statement := models.ItemProductFocusSelectStatement + ` ` + models.ItemProductFocusWhereStatement +
		` AND (LOWER(i."_name") LIKE $1)  AND IUL.CONVERSION > 1 ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	fmt.Println(statement)
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
func (repository ItemProductFocusRepository) FindAll(ctx context.Context, parameter models.ItemProductFocusParameter) (data []models.ItemProductFocus, count int, err error) {
	conditionString := ``

	if parameter.ItemCategoryId != "" {
		conditionString += ` AND i.Item_category_id = '` + parameter.ItemCategoryId + `'`
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionId + `'`
	}

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND i.id NOT IN (83, 307, 393) `
	}

	query := models.ItemProductFocusSelectStatement + ` ` + models.ItemProductFocusWhereStatement + ` ` + conditionString + `
		AND (LOWER(i."_name") LIKE $1  )   AND IUL.CONVERSION > 1 ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "product_focus" def ` + models.ItemProductFocusWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)  AND IUL.CONVERSION > 1 `
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemProductFocusRepository) FindByID(c context.Context, parameter models.ItemProductFocusParameter) (data models.ItemProductFocus, err error) {
	statement := models.ItemProductFocusSelectStatement + ` WHERE def.created_date IS NOT NULL  AND IUL.CONVERSION > 1 AND def.item_id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
