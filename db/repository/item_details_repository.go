package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemDetailsRepository ...
type IItemDetailsRepository interface {
	SelectAll(c context.Context, parameter models.ItemDetailsParameter) ([]models.ItemDetails, error)
	FindAll(ctx context.Context, parameter models.ItemDetailsParameter) ([]models.ItemDetails, int, error)
	FindByID(c context.Context, parameter models.ItemDetailsParameter) (models.ItemDetails, error)
	// Add(c context.Context, model *models.ItemDetails) (*string, error)
	// Edit(c context.Context, model *models.ItemDetails) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemDetailsRepository ...
type ItemDetailsRepository struct {
	DB *sql.DB
}

// NewItemDetailsRepository ...
func NewItemDetailsRepository(DB *sql.DB) IItemDetailsRepository {
	return &ItemDetailsRepository{DB: DB}
}

// Scan rows
func (repository ItemDetailsRepository) scanRows(rows *sql.Rows) (res models.ItemDetails, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemDetailsCategoryId,
		&res.ItemDetailsCategoryName,
		&res.ItemDetailsPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemDetailsPrice,
		&res.PriceListVersionId,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemDetailsRepository) scanRow(row *sql.Row) (res models.ItemDetails, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemDetailsCategoryId,
		&res.ItemDetailsCategoryName,
		&res.ItemDetailsPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemDetailsPrice,
		&res.PriceListVersionId,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemDetailsRepository) SelectAll(c context.Context, parameter models.ItemDetailsParameter) (data []models.ItemDetails, err error) {
	conditionString := ``
	conditionStringPriceListVersion := ``
	conditionStringException := ``

	if parameter.ItemDetailsCategoryId != "" {
		conditionString += ` AND def.item_category_id = '` + parameter.ItemDetailsCategoryId + `'`
	}

	if parameter.PriceListVersionId != "" {
		conditionStringPriceListVersion += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionId + `'`
	}

	if parameter.ExceptId != "" {
		conditionStringException += ` AND def.id <> '` + parameter.ExceptId + `'`
	}

	statement := models.ItemDetailsSelectStatement + ` ` + models.ItemDetailsWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + conditionStringPriceListVersion + conditionStringException + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	fmt.Println(statement)

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
func (repository ItemDetailsRepository) FindAll(ctx context.Context, parameter models.ItemDetailsParameter) (data []models.ItemDetails, count int, err error) {
	conditionString := ``
	conditionStringPriceListVersion := ``
	conditionStringException := ``

	if parameter.ItemDetailsCategoryId != "" {
		conditionString += ` AND def.item_category_id = '` + parameter.ItemDetailsCategoryId + `'`
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = '` + parameter.PriceListVersionId + `'`
	}

	query := models.ItemDetailsSelectStatement + ` ` + models.ItemDetailsWhereStatement + ` ` + conditionString + conditionStringPriceListVersion + conditionStringException + `
		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	fmt.Println(query)

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

	query = `SELECT COUNT(*) FROM "item" def ` + models.ItemDetailsWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemDetailsRepository) FindByID(c context.Context, parameter models.ItemDetailsParameter) (data models.ItemDetails, err error) {
	conditionString := ``

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = '` + parameter.PriceListVersionId + `'`
	}

	statement := models.ItemDetailsSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1` + conditionString + ``

	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemDetailsRepository) FindByCategoryID(c context.Context, parameter models.ItemDetailsParameter) (data models.ItemDetails, err error) {
	statement := models.ItemDetailsSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
