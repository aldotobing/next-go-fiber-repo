package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IPromoItemLineRepository ...
type IPromoItemLineRepository interface {
	SelectAll(c context.Context, parameter models.PromoItemLineParameter) ([]models.PromoItemLine, error)
	FindAll(ctx context.Context, parameter models.PromoItemLineParameter) ([]models.PromoItemLine, int, error)
	FindByID(c context.Context, parameter models.PromoItemLineParameter) (models.PromoItemLine, error)
	Add(c context.Context, model *models.PromoItemLineBreakDown) (*string, error)
	// Edit(c context.Context, model *models.PromoItemLine) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// PromoItemLineRepository ...
type PromoItemLineRepository struct {
	DB *sql.DB
}

// NewPromoItemLineRepository ...
func NewPromoItemLineRepository(DB *sql.DB) IPromoItemLineRepository {
	return &PromoItemLineRepository{DB: DB}
}

// Scan rows
func (repository PromoItemLineRepository) scanRows(rows *sql.Rows) (res models.PromoItemLine, err error) {
	err = rows.Scan(
		&res.ID,
		&res.ItemID,
		&res.PromoID,
		&res.PromoLineID,
		&res.PromoName,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemDescription,
		&res.ItemCategoryID,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.Qty,
		&res.UomID,
		&res.UomName,
		&res.ItemPrice,
		&res.PriceListVersionID,
		&res.GlobalMaxQty,
		&res.CustomerMaxQty,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinValue,
		&res.MinQty,
		&res.Description,
		&res.Multiply,
		&res.MinQtyUomID,
		&res.PromoType,
		&res.Strata,
		&res.StartDate,
		&res.EndDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository PromoItemLineRepository) scanRow(row *sql.Row) (res models.PromoItemLine, err error) {
	err = row.Scan(
		&res.ID,
		&res.ItemID,
		&res.PromoID,
		&res.PromoLineID,
		&res.PromoName,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemDescription,
		&res.ItemCategoryID,
		&res.ItemCategoryName,
		&res.ItemPicture,
		&res.Qty,
		&res.UomID,
		&res.UomName,
		&res.ItemPrice,
		&res.PriceListVersionID,
		&res.GlobalMaxQty,
		&res.CustomerMaxQty,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinValue,
		&res.MinQty,
		&res.Description,
		&res.Multiply,
		&res.MinQtyUomID,
		&res.PromoType,
		&res.Strata,
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
func (repository PromoItemLineRepository) SelectAll(c context.Context, parameter models.PromoItemLineParameter) (data []models.PromoItemLine, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += `AND IP.PRICE_LIST_VERSION_ID =
	(SELECT id FROM price_list_version WHERE price_list_id = (SELECT price_list_id FROM customer WHERE id = ` + parameter.CustomerID + `` + `))` + ` `
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND pr.start_date = '` + parameter.StartDate + `' AND pr.end_date = '` + parameter.EndDate + `'`
	}

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = '` + parameter.PromoID + `'`
	}
	if parameter.PromoLineID != "" {
		conditionString += ` AND PRL.ID = '` + parameter.PromoLineID + `'`
	}
	if parameter.ID != "" {
		conditionString += ` AND PIL.ID = '` + parameter.ID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND PIL.UOM_ID = '` + parameter.UomID + `'`
	}

	statement := models.PromoItemLineSelectStatement + ` ` + models.PromoItemLineWhereStatement +
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
func (repository PromoItemLineRepository) FindAll(ctx context.Context, parameter models.PromoItemLineParameter) (data []models.PromoItemLine, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = (SELECT id FROM price_list_version WHERE price_list_id = (SELECT price_list_id FROM customer WHERE id = ` + parameter.CustomerID + `` + `))` + ` `
	}

	if parameter.StartDate != "" || parameter.EndDate != "" {
		conditionString += ` AND pr.start_date = '` + parameter.StartDate + `AND pr.end_date` + parameter.EndDate + `'`
	}

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = '` + parameter.PromoID + `'`
	}

	if parameter.PromoLineID != "" {
		conditionString += ` AND PRL.ID = '` + parameter.PromoLineID + `'`
	}

	if parameter.ID != "" {
		conditionString += ` AND PIL.ID = '` + parameter.ID + `'`
	}

	if parameter.UomID != "" {
		conditionString += ` AND PIL.UOM_ID = '` + parameter.UomID + `'`
	}

	query := models.PromoItemLineSelectStatement + ` ` + models.PromoItemLineWhereStatement + ` ` +
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
			promo_item_line pil 
				LEFT JOIN promo_line prl ON prl.id = pil.promo_line_id 
				LEFT JOIN promo pr ON pr.id = prl.promo_id 
				LEFT JOIN item i ON i.id = pil.item_id 
				LEFT JOIN uom u ON u.id = pil.uom_id 
				JOIN ITEM_PRICE IP ON IP.ITEM_ID = PIL.ITEM_ID
				` +
		models.PromoItemLineWhereStatement + ` ` +
		conditionString + ` AND (LOWER(i."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository PromoItemLineRepository) FindByID(c context.Context, parameter models.PromoItemLineParameter) (data models.PromoItemLine, err error) {
	statement := models.PromoItemLineSelectStatement + ` WHERE i.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ItemID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository PromoItemLineRepository) Add(c context.Context, model *models.PromoItemLineBreakDown) (res *string, err error) {
	statement := `INSERT INTO promo_item_line (
		promo_line_id, 
		item_id, 
		uom_id, 
		qty) VALUES ($1, $2, $3, $4) RETURNING id`

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
