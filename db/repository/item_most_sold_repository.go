package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemMostSoldRepository ...
type IItemMostSoldRepository interface {
	SelectAll(c context.Context, parameter models.ItemMostSoldParameter) ([]models.ItemMostSold, error)
	FindAll(ctx context.Context, parameter models.ItemMostSoldParameter) ([]models.ItemMostSold, int, error)
	FindByID(c context.Context, parameter models.ItemMostSoldParameter) (models.ItemMostSold, error)
}

// ItemMostSoldRepository ...
type ItemMostSoldRepository struct {
	DB *sql.DB
}

// NewItemMostSoldRepository ...
func NewItemMostSoldRepository(DB *sql.DB) IItemMostSoldRepository {
	return &ItemMostSoldRepository{DB: DB}
}

// Scan rows
func (repository ItemMostSoldRepository) scanRows(rows *sql.Rows) (res models.ItemMostSold, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.ItemTotalSold,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemMostSoldRepository) scanRow(row *sql.Row) (res models.ItemMostSold, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.ItemTotalSold,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemMostSoldRepository) SelectAll(c context.Context, parameter models.ItemMostSoldParameter) (data []models.ItemMostSold, err error) {
	conditionString := ``

	if parameter.ItemCategoryId != "" {
		conditionString += ` AND i.item_category_id = '` + parameter.ItemCategoryId + `'`
	}

	statement := models.ItemMostSoldSelectStatement
	// + ` ` + models.ItemMostSoldWhereStatement +
	// 	` AND (LOWER(i._name) LIKE $1) ` + conditionString +` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement)
	//, "%"+strings.ToLower(parameter.Search)+"%")

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
func (repository ItemMostSoldRepository) FindAll(ctx context.Context, parameter models.ItemMostSoldParameter) (data []models.ItemMostSold, count int, err error) {
	conditionString := ``

	if parameter.ItemCategoryId != "" {
		conditionString += ` AND i.item_category_id = '` + parameter.ItemCategoryId + `'`
	}

	query := models.ItemMostSoldSelectStatement + ` ` + models.ItemMostSoldWhereStatement + ` ` + conditionString + `
		AND (LOWER(i."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT count (i.code) 
				FROM SALES_INVOICE_HEADER SIH
				JOIN SALES_INVOICE_LINE SIL ON
				SIL.HEADER_ID = SIH.ID
				JOIN ITEM I ON
				I.ID = SIL.ITEM_ID
				JOIN ITEM_CATEGORY IC ON
				IC.ID = I.ITEM_CATEGORY_ID
				WHERE SIH.TRANSACTION_DATE BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW() ` + models.ItemMostSoldWhereStatement + ` ` +
		conditionString + ` AND (LOWER i._name) LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemMostSoldRepository) FindByID(c context.Context, parameter models.ItemMostSoldParameter) (data models.ItemMostSold, err error) {
	statement := models.ItemMostSoldSelectStatement + ` WHERE i.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemMostSoldRepository) FindByCategoryID(c context.Context, parameter models.ItemMostSoldParameter) (data models.ItemMostSold, err error) {
	statement := models.ItemMostSoldSelectStatement + ` WHERE I.ITEM_CATEGORY_ID = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
