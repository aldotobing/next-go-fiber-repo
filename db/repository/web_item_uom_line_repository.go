package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebItemUomLineRepository ...
type IWebItemUomLineRepository interface {
	SelectAll(c context.Context, parameter models.WebItemUomLineParameter) ([]models.WebItemUomLine, error)
	FindAll(ctx context.Context, parameter models.WebItemUomLineParameter) ([]models.WebItemUomLine, int, error)
	FindByID(c context.Context, parameter models.WebItemUomLineParameter) (models.WebItemUomLine, error)
	// Add(c context.Context, model *models.Item) (*string, error)
	// Edit(c context.Context, model *models.Item) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// WebItemUomLineRepository ...
type WebItemUomLineRepository struct {
	DB *sql.DB
}

// NewWebItemUomLineRepository ...
func NewWebItemUomLineRepository(DB *sql.DB) IWebItemUomLineRepository {
	return &WebItemUomLineRepository{DB: DB}
}

// Scan rows
func (repository WebItemUomLineRepository) scanRows(rows *sql.Rows) (res models.WebItemUomLine, err error) {
	err = rows.Scan(
		&res.ID,
		&res.ItemID,
		&res.ItemUomID,
		&res.ItemCategoryId,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemCategoryName,
		&res.ItemUomName,
		&res.ItemUomConversion,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebItemUomLineRepository) scanRow(row *sql.Row) (res models.WebItemUomLine, err error) {
	err = row.Scan(
		&res.ID,
		&res.ItemID,
		&res.ItemUomID,
		&res.ItemCategoryId,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemCategoryName,
		&res.ItemUomName,
		&res.ItemUomConversion,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebItemUomLineRepository) SelectAll(c context.Context, parameter models.WebItemUomLineParameter) (data []models.WebItemUomLine, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
	}

	if parameter.ItemID != "" {
		conditionString += ` AND def.item_id = '` + parameter.ItemID + `'`
	}

	statement := models.WebItemUomLineSelectStatement + ` ` + models.WebItemUomLineWhereStatement +
		` AND (LOWER(i."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebItemUomLineRepository) FindAll(ctx context.Context, parameter models.WebItemUomLineParameter) (data []models.WebItemUomLine, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
	}

	if parameter.ItemID != "" {
		conditionString += ` AND def.item_id = '` + parameter.ItemID + `'`
	}

	query := models.WebItemUomLineSelectStatement + ` ` + models.WebItemUomLineWhereStatement + ` ` + conditionString + `
		AND (LOWER(i."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM ITEM_UOM_LINE DEF 
		left join item i on i.id = def.item_id
		left join uom u on u.id = def.uom_id
		LEFT JOIN ITEM_CATEGORY IC ON IC.ID = I.ITEM_CATEGORY_ID

		` +
		models.WebItemUomLineWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebItemUomLineRepository) FindByID(c context.Context, parameter models.WebItemUomLineParameter) (data models.WebItemUomLine, err error) {
	statement := models.WebItemUomLineSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository WebItemUomLineRepository) FindByCategoryID(c context.Context, parameter models.WebItemUomLineParameter) (data models.WebItemUomLine, err error) {
	statement := models.WebItemUomLineSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
