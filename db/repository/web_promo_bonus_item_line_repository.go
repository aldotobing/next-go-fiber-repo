package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IWebPromoBonusItemLineRepository ...
type IWebPromoBonusItemLineRepository interface {
	SelectAll(c context.Context, parameter models.WebPromoBonusItemLineParameter) ([]models.WebPromoBonusItemLine, error)
	FindAll(ctx context.Context, parameter models.WebPromoBonusItemLineParameter) ([]models.WebPromoBonusItemLine, int, error)
	FindByID(c context.Context, parameter models.WebPromoBonusItemLineParameter) (models.WebPromoBonusItemLine, error)
	Add(c context.Context, model *models.WebPromoBonusItemLineBreakDown) (*string, error)
	// Edit(c context.Context, model *models.WebPromoBonusItemLine) (*string, error)
	Delete(c context.Context, id string) (string, error)
}

// WebPromoBonusItemLineRepository ...
type WebPromoBonusItemLineRepository struct {
	DB *sql.DB
}

// NewWebPromoBonusItemLineRepository ...
func NewWebPromoBonusItemLineRepository(DB *sql.DB) IWebPromoBonusItemLineRepository {
	return &WebPromoBonusItemLineRepository{DB: DB}
}

// Scan rows
func (repository WebPromoBonusItemLineRepository) scanRows(rows *sql.Rows) (res models.WebPromoBonusItemLine, err error) {
	err = rows.Scan(
		&res.ID,
		&res.ItemID,
		&res.PromoID,
		&res.PromoLineID,
		&res.PromoName,
		&res.ItemCode,
		&res.ItemName,
		&res.Qty,
		&res.UomID,
		&res.UomName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebPromoBonusItemLineRepository) scanRow(row *sql.Row) (res models.WebPromoBonusItemLine, err error) {
	err = row.Scan(
		&res.ID,
		&res.ItemID,
		&res.PromoID,
		&res.PromoLineID,
		&res.PromoName,
		&res.ItemCode,
		&res.ItemName,
		&res.Qty,
		&res.UomID,
		&res.UomName,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebPromoBonusItemLineRepository) SelectAll(c context.Context, parameter models.WebPromoBonusItemLineParameter) (data []models.WebPromoBonusItemLine, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = '` + parameter.PromoID + `'`
	}
	if parameter.PromoLineID != "" {
		conditionString += ` AND PIL.PROMO_LINE_ID = '` + parameter.PromoLineID + `'`
	}
	if parameter.ID != "" {
		conditionString += ` AND PIL.ID = '` + parameter.ID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND PIL.UOM_ID = '` + parameter.UomID + `'`
	}

	statement := models.WebPromoBonusItemLineSelectStatement + ` ` + models.WebPromoBonusItemLineWhereStatement +
		` AND (LOWER(i."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebPromoBonusItemLineRepository) FindAll(ctx context.Context, parameter models.WebPromoBonusItemLineParameter) (data []models.WebPromoBonusItemLine, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = '` + parameter.PromoID + `'`
	}

	if parameter.PromoLineID != "" {
		conditionString += ` AND PIL.PROMO_LINE_ID = '` + parameter.PromoLineID + `'`
	}

	if parameter.ID != "" {
		conditionString += ` AND PIL.ID = '` + parameter.ID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND PIL.UOM_ID = '` + parameter.UomID + `'`
	}

	query := models.WebPromoBonusItemLineSelectStatement + ` ` + models.WebPromoBonusItemLineWhereStatement + ` ` +
		`AND (LOWER(i."_name") LIKE $1)` + conditionString + ` ` + `ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT count(*)
		FROM
		promo_bonus_line pil 
				LEFT JOIN promo_line prl ON prl.id = pil.promo_line_id 
				LEFT JOIN promo pr ON pr.id = prl.promo_id 
				LEFT JOIN item i ON i.id = pil.item_id 
				LEFT JOIN uom u ON u.id = pil.uom_id
				` +
		models.WebPromoBonusItemLineWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebPromoBonusItemLineRepository) FindByID(c context.Context, parameter models.WebPromoBonusItemLineParameter) (data models.WebPromoBonusItemLine, err error) {
	statement := models.WebPromoBonusItemLineSelectStatement + ` WHERE i.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ItemID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository WebPromoBonusItemLineRepository) Add(c context.Context, model *models.WebPromoBonusItemLineBreakDown) (res *string, err error) {
	statement := `INSERT INTO promo_bonus_line (
		promo_line_id, 
		item_id, 
		uom_id, 
		qty, created_date) VALUES ($1, $2, $3, $4, now()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		model.PromoLineID,
		str.NullString(model.ItemID),
		str.NullString(model.UomID),
		str.NullString(model.Qty)).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository WebPromoBonusItemLineRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `
		DELETE FROM promo_bonus_line WHERE id = $1 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		id).Scan(&res)

	return
}
