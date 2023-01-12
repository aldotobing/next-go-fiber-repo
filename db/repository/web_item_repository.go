package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebItemRepository ...
type IWebItemRepository interface {
	SelectAll(c context.Context, parameter models.WebItemParameter) ([]models.WebItem, error)
	FindAll(ctx context.Context, parameter models.WebItemParameter) ([]models.WebItem, int, error)
	FindByID(c context.Context, parameter models.WebItemParameter) (models.WebItem, error)
	// Add(c context.Context, model *models.Item) (*string, error)
	Edit(c context.Context, model *models.WebItem) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// WebItemRepository ...
type WebItemRepository struct {
	DB *sql.DB
}

// NewWebItemRepository ...
func NewWebItemRepository(DB *sql.DB) IWebItemRepository {
	return &WebItemRepository{DB: DB}
}

// Scan rows
func (repository WebItemRepository) scanRows(rows *sql.Rows) (res models.WebItem, err error) {
	err = rows.Scan(
		&res.ID,
		&res.ItemCategoryId,
		&res.Code,
		&res.Name,
		&res.ItemPicture,
		&res.ItemCategoryName,
		&res.ItemActive,
		&res.Description,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebItemRepository) scanRow(row *sql.Row) (res models.WebItem, err error) {
	err = row.Scan(
		&res.ID,
		&res.ItemCategoryId,
		&res.Code,
		&res.Name,
		&res.ItemPicture,
		&res.ItemCategoryName,
		&res.ItemActive,
		&res.Description,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebItemRepository) SelectAll(c context.Context, parameter models.WebItemParameter) (data []models.WebItem, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
	}

	if parameter.ItemCategoryId != "" {
		if parameter.ItemCategoryId == "2" {
			//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
		} else {
			conditionString += ` AND def.item_category_id = ` + parameter.ItemCategoryId + ``
		}
	}

	statement := models.WebItemSelectStatement + ` ` + models.WebItemWhereStatement +
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
func (repository WebItemRepository) FindAll(ctx context.Context, parameter models.WebItemParameter) (data []models.WebItem, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
	}

	if parameter.ItemCategoryId != "" {
		if parameter.ItemCategoryId == "2" {
			//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
		} else {
			conditionString += ` AND def.item_category_id = ` + parameter.ItemCategoryId + ``
		}
	}

	query := models.WebItemSelectStatement + ` ` + models.WebItemWhereStatement + ` ` + conditionString + `
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
		`LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID` +
		models.WebItemWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebItemRepository) FindByID(c context.Context, parameter models.WebItemParameter) (data models.WebItem, err error) {
	statement := models.WebItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository WebItemRepository) FindByCategoryID(c context.Context, parameter models.WebItemParameter) (data models.WebItem, err error) {
	statement := models.WebItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebItemRepository) Edit(c context.Context, model *models.WebItem) (res *string, err error) {
	statement := `UPDATE item SET 
	_name = $1, 
	item_picture = $2
	WHERE id = $3 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.Name,
		model.ItemPicture,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
