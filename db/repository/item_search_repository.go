package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemSearchRepository ...
type IItemSearchRepository interface {
	SelectAll(c context.Context, parameter models.ItemSearchParameter) ([]models.ItemSearch, error)
	FindAll(ctx context.Context, parameter models.ItemSearchParameter) ([]models.ItemSearch, int, error)
	FindByID(c context.Context, parameter models.ItemSearchParameter) (models.ItemSearch, error)
	// Add(c context.Context, model *models.ItemSearch) (*string, error)
	// Edit(c context.Context, model *models.ItemSearch) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemSearchRepository ...
type ItemSearchRepository struct {
	DB *sql.DB
}

// NewItemSearchRepository ...
func NewItemSearchRepository(DB *sql.DB) IItemSearchRepository {
	return &ItemSearchRepository{DB: DB}
}

// Scan rows
func (repository ItemSearchRepository) scanRows(rows *sql.Rows) (res models.ItemSearch, err error) {
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
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemSearchRepository) scanRow(row *sql.Row) (res models.ItemSearch, err error) {
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
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemSearchRepository) SelectAll(c context.Context, parameter models.ItemSearchParameter) (data []models.ItemSearch, err error) {
	conditionString := ``
	conditionStringPriceListVersion := ``

	if parameter.Name != "" {
		conditionString += ` or (LOWER (ic."_name") like ` + `'%` + strings.ToLower(parameter.Name) + `%')`
	}

	// if parameter.PriceListVersionId != "" {
	// 	conditionStringPriceListVersion += ` AND ip.price_list_version_id = '` + parameter.PriceListVersionId + `'`
	// }

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id NOT IN (83, 307, 393) `
	}

	statement := models.ItemSearchSelectStatement + ` ` + models.ItemSearchWhereStatement +
		` AND ((LOWER(def."_name") LIKE $2 )) ` + conditionString + conditionStringPriceListVersion + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, parameter.PriceListVersionId, "%"+strings.ToLower(parameter.Name)+"%")

	fmt.Println("select ALL : " + statement)
	// fmt.Println("select ALL PARAM : " + parameter.Name)

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
func (repository ItemSearchRepository) FindAll(ctx context.Context, parameter models.ItemSearchParameter) (data []models.ItemSearch, count int, err error) {
	conditionString := ``

	if parameter.Name != "" {
		conditionString += ` or LOWER (ic."_name") like` + `'%` + parameter.Name + `%'` + `'`
	}

	/*
		customerType 7 = Apotek Lokal
		customerType 15 = MT LOKAL INDEPENDEN
		defId 83 = TOLAK ANGIN CAIR /D5
		Tampilkan TAC D5 hanya pada kedua customerType di atas
	*/
	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id NOT IN (83, 307, 393) `
	}

	query := models.ItemSearchSelectStatement + ` ` + models.ItemSearchWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	//fmt.Println(query)

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

	query = `SELECT COUNT(*) FROM "item" def ` + models.ItemSearchWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemSearchRepository) FindByID(c context.Context, parameter models.ItemSearchParameter) (data models.ItemSearch, err error) {
	conditionString := ``

	statement := models.ItemSearchSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1` + conditionString + ``

	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemSearchRepository) FindByCategoryID(c context.Context, parameter models.ItemSearchParameter) (data models.ItemSearch, err error) {
	statement := models.ItemSearchSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
