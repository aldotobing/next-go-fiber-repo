package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemPriceRepository ...
type IItemPriceRepository interface {
	SelectAll(c context.Context, parameter models.ItemPriceParameter) ([]models.ItemPrice, error)
	FindAll(ctx context.Context, parameter models.ItemPriceParameter) ([]models.ItemPrice, int, error)
	FindByID(c context.Context, parameter models.ItemPriceParameter) (models.ItemPrice, error)
	// Add(c context.Context, model *models.ItemPrice) (*string, error)
	// Edit(c context.Context, model *models.ItemPrice) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemPriceRepository ...
type ItemPriceRepository struct {
	DB *sql.DB
}

// NewItemPriceRepository ...
func NewItemPriceRepository(DB *sql.DB) IItemPriceRepository {
	return &ItemPriceRepository{DB: DB}
}

// Scan rows
func (repository ItemPriceRepository) scanRows(rows *sql.Rows) (res models.ItemPrice, err error) {
	err = rows.Scan(
		&res.ID,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.UomID,
		&res.UomName,
		&res.PriceListVersionID,
		&res.Price,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemPriceRepository) scanRow(row *sql.Row) (res models.ItemPrice, err error) {
	err = row.Scan(
		&res.ID,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.UomID,
		&res.UomName,
		&res.PriceListVersionID,
		&res.Price,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemPriceRepository) SelectAll(c context.Context, parameter models.ItemPriceParameter) (data []models.ItemPrice, err error) {
	conditionString := ``

	if parameter.ItemID != "" {
		conditionString += ` AND ip.item_id = '` + parameter.ItemID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND ip.uom_id = '` + parameter.UomID + `'`
	}

	if parameter.PriceListVersionID != "" {
		conditionString += ` AND ip.price_list_version_id =  '` + parameter.PriceListVersionID + `'`
	}

	statement := models.ItemPriceSelectStatement + ` ` + models.ItemPriceWhereStatement +
		` AND (LOWER(i."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	fmt.Println("Select All Query : " + statement)

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
func (repository ItemPriceRepository) FindAll(ctx context.Context, parameter models.ItemPriceParameter) (data []models.ItemPrice, count int, err error) {
	conditionString := ``

	if parameter.ItemID != "" {
		conditionString += ` AND ip.item_id = '` + parameter.ItemID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND ip.uom_id = '` + parameter.UomID + `'`
	}

	if parameter.PriceListVersionID != "" {
		conditionString += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionID + `'`
	}

	query := models.ItemPriceSelectStatement + ` ` + models.ItemPriceWhereStatement + ` ` + conditionString + `
		AND (LOWER(i."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	fmt.Println("Find All Query : " + query)
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

	query = `SELECT` +
		` COUNT(*) ` +
		` FROM "item_price" ip ` +
		` JOIN ITEM I ON I.ID = ip.ITEM_ID ` +
		` JOIN UOM U ON U.ID = ip.UOM_ID ` +
		models.ItemPriceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	fmt.Println("Count Query : " + query)

	return data, count, err
}

// FindByID ...
func (repository ItemPriceRepository) FindByID(c context.Context, parameter models.ItemPriceParameter) (data models.ItemPrice, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND ip.id = '` + parameter.ID + `'`
	}

	statement := models.ItemPriceSelectStatement + ` ip.id = $1` + conditionString + ``

	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Print("FindByID Query : " + statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemPriceRepository) FindByCategoryID(c context.Context, parameter models.ItemPriceParameter) (data models.ItemPrice, err error) {
	statement := models.ItemPriceSelectStatement + ` WHERE def.created_date IS NOT NULL AND ip.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
