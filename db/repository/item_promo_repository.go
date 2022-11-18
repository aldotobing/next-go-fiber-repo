package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemPromoRepository ...
type IItemPromoRepository interface {
	SelectAll(c context.Context, parameter models.ItemPromoParameter) ([]models.ItemPromo, error)
	FindAll(ctx context.Context, parameter models.ItemPromoParameter) ([]models.ItemPromo, int, error)
	FindByID(c context.Context, parameter models.ItemPromoParameter) (models.ItemPromo, error)
	// Add(c context.Context, model *models.ItemPromo) (*string, error)
	// Edit(c context.Context, model *models.ItemPromo) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemPromoRepository ...
type ItemPromoRepository struct {
	DB *sql.DB
}

// NewItemPromoRepository ...
func NewItemPromoRepository(DB *sql.DB) IItemPromoRepository {
	return &ItemPromoRepository{DB: DB}
}

// Scan rows
func (repository ItemPromoRepository) scanRows(rows *sql.Rows) (res models.ItemPromo, err error) {
	err = rows.Scan(
		&res.PromoID,
		&res.PromoLineID,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemDescription,
		&res.ItemCategoryID,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionID,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinValue,
		&res.MinQty,
		&res.CustMaxQty,
		&res.GlobalMaxQty,
		&res.Description,
		&res.StartDate,
		&res.EndDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemPromoRepository) scanRow(row *sql.Row) (res models.ItemPromo, err error) {
	err = row.Scan(
		&res.PromoID,
		&res.PromoLineID,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemDescription,
		&res.ItemCategoryID,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionID,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinValue,
		&res.MinQty,
		&res.CustMaxQty,
		&res.GlobalMaxQty,
		&res.Description,
		&res.StartDate,
		&res.EndDate,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemPromoRepository) SelectAll(c context.Context, parameter models.ItemPromoParameter) (data []models.ItemPromo, err error) {
	conditionString := ``

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND pr.start_date = '` + parameter.StartDate + `' AND pr.end_date = '` + parameter.EndDate + `'`
	}

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = '` + parameter.PromoID + `'`
	}

	statement := models.ItemPromoSelectStatement + ` ` + models.ItemPromoWhereStatement +
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
func (repository ItemPromoRepository) FindAll(ctx context.Context, parameter models.ItemPromoParameter) (data []models.ItemPromo, count int, err error) {
	conditionString := ``

	if parameter.StartDate != "" || parameter.EndDate != "" {
		conditionString += ` AND pr.start_date = '` + parameter.StartDate + `AND pr.end_date` + parameter.EndDate + `'`
	}

	query := models.ItemPromoSelectStatement + ` ` + models.ItemPromoWhereStatement + ` ` + conditionString + `
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

	query = `SELECT count(*)
		FROM
			promo_item_line pil 
			LEFT JOIN promo_line prl ON prl.id = pil.promo_line_id 
			LEFT JOIN promo pr ON pr.id = prl.promo_id 
			LEFT JOIN item i ON i.id = pil.item_id 
			LEFT JOIN uom u ON u.id = pil.uom_id ` +

		models.ItemPromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ItemPromoRepository) FindByID(c context.Context, parameter models.ItemPromoParameter) (data models.ItemPromo, err error) {
	statement := models.ItemPromoSelectStatement + ` WHERE i.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ItemID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
