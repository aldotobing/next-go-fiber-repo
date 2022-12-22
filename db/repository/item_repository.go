package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemRepository ...
type IItemRepository interface {
	SelectAll(c context.Context, parameter models.ItemParameter) ([]models.Item, error)
	FindAll(ctx context.Context, parameter models.ItemParameter) ([]models.Item, int, error)
	FindByID(c context.Context, parameter models.ItemParameter) (models.Item, error)
	// Add(c context.Context, model *models.Item) (*string, error)
	// Edit(c context.Context, model *models.Item) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemRepository ...
type ItemRepository struct {
	DB *sql.DB
}

// NewItemRepository ...
func NewItemRepository(DB *sql.DB) IItemRepository {
	return &ItemRepository{DB: DB}
}

// Scan rows
func (repository ItemRepository) scanRows(rows *sql.Rows) (res models.Item, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionId,
		&res.Uom,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemRepository) scanRow(row *sql.Row) (res models.Item, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionId,
		&res.Uom,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemRepository) SelectAll(c context.Context, parameter models.ItemParameter) (data []models.Item, err error) {
	conditionString := ``

	if parameter.CustomerTypeId != "" && parameter.ItemCategoryId == "2" {
		//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
		conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
	} else {
		conditionString += ` AND def.item_category_id = ` + parameter.ItemCategoryId + ``
	}

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id <> 83 `
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionId + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND IUL.UOM_ID = '` + parameter.UomID + `'`
	}

	if parameter.ExceptId != "" {
		conditionString += ` AND def.id <> '` + parameter.ExceptId + `'`
	}

	statement := models.ItemSelectStatement + ` ` + models.ItemWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository ItemRepository) FindAll(ctx context.Context, parameter models.ItemParameter) (data []models.Item, count int, err error) {
	conditionString := ``

	if parameter.CustomerTypeId != "" && parameter.ItemCategoryId == "2" {
		//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
		conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
	} else if parameter.CustomerTypeId != "" {
		conditionString += ` AND def.item_category_id = '` + parameter.ItemCategoryId + `'`
	}

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id <> 83 `
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = '` + parameter.PriceListVersionId + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND IUL.UOM_ID = '` + parameter.UomID + `'`
	}

	if parameter.ExceptId != "" {
		conditionString += ` AND def.id <> '` + parameter.ExceptId + `'`
	}

	query := models.ItemSelectStatement + ` ` + models.ItemWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "item" DEF ` +
		`JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID ` +
		models.ItemWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemRepository) FindByID(c context.Context, parameter models.ItemParameter) (data models.Item, err error) {
	statement := models.ItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemRepository) FindByCategoryID(c context.Context, parameter models.ItemParameter) (data models.Item, err error) {
	statement := models.ItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
