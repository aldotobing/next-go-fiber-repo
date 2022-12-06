package repository

import (
	"context"
	"database/sql"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IPromoLine ...
type IPromoLine interface {
	SelectAll(c context.Context, parameter models.PromoLineParameter) ([]models.PromoLine, error)
	FindAll(ctx context.Context, parameter models.PromoLineParameter) ([]models.PromoLine, int, error)
	Add(c context.Context, model *models.PromoLine) (*string, error)
	// 	Edit(c context.Context, model *models.PromoLine) (*string, error)
	// 	EditAddress(c context.Context, model *models.PromoLine) (*string, error)
}

// PromoLine ...
type PromoLine struct {
	DB *sql.DB
}

// NewPromoLine ...
func NewPromoLineRepository(DB *sql.DB) IPromoLine {
	return &PromoLine{DB: DB}
}

// Scan rows
func (repository PromoLine) scanRows(rows *sql.Rows) (res models.PromoLine, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PromoID,
		&res.GlobalMaxQty,
		&res.CustomerMaxQty,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinimumValue,
		&res.Multiply,
		&res.Description,
		&res.MinimumQty,
		&res.MinimumQtyUomID,
		&res.PromoType,
		&res.Strata,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository PromoLine) scanRow(row *sql.Row) (res models.PromoLine, err error) {
	err = row.Scan(
		&res.ID,
		&res.PromoID,
		&res.GlobalMaxQty,
		&res.CustomerMaxQty,
		&res.DiscPercent,
		&res.DiscAmount,
		&res.MinimumValue,
		&res.Multiply,
		&res.Description,
		&res.MinimumQty,
		&res.MinimumQtyUomID,
		&res.PromoType,
		&res.Strata,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository PromoLine) SelectAll(c context.Context, parameter models.PromoLineParameter) (data []models.PromoLine, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += `AND promo_id = ` + parameter.PromoID + ` `
	}

	if parameter.ID != "" {
		conditionString += `AND pl.id = ` + parameter.ID + ` `
	}

	statement := models.PromoLineSelectStatement + ` ` + models.PromoLineWhereStatement +
		` ` + conditionString + ` ORDER BY pl.id ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement)

	//print
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
func (repository PromoLine) FindAll(ctx context.Context, parameter models.PromoLineParameter) (data []models.PromoLine, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += `AND promo_id = ` + parameter.PromoID + ` `
	}

	if parameter.ID != "" {
		conditionString += `AND pl.id = ` + parameter.ID + ` `
	}

	query := models.PromoLineSelectStatement + ` ` + models.PromoLineWhereStatement + ` ` + conditionString + `
		ORDER BY pl.id ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.Query(query, parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	fmt.Print(query)

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

	query = `SELECT COUNT(*) FROM "promo_line" pl ` + models.PromoLineWhereStatement + ` ` +
		conditionString
	err = repository.DB.QueryRow(query).Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository PromoLine) FindByID(c context.Context, parameter models.PromoLineParameter) (data models.PromoLine, err error) {
	statement := models.PromoLineSelectStatement + ` WHERE pl.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository PromoLine) Edit(c context.Context, model *models.PromoLine) (res *string, err error) {
// 	statement := `UPDATE customer SET
// 	customer_name = $1,
// 	customer_address = $2,
// 	customer_phone = $3,
// 	customer_email = $4,
// 	customer_cp_name = $5,
// 	customer_profile_picture = $6
// 	WHERE id = $7
// 	RETURNING id`
// 	err = repository.DB.QueryRowContext(c, statement,
// 		model.PromoLineName,
// 		model.PromoLineAddress,
// 		model.PromoLinePhone,
// 		model.PromoLineEmail,
// 		model.PromoLineCpName,
// 		model.PromoLineProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

func (repository PromoLine) Add(c context.Context, model *models.PromoLine) (res *string, err error) {
	statement := `INSERT INTO promo_line (
		promo_id, 
		global_max_qty, 
		customer_max_qty, 
		disc_pct, 
		disc_amt, 
		minimum_value,
		multiply, 
		description,
		minimum_qty,
		minimum_qty_uom_id,
		promo_type, strata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.PromoID),
		str.NullString(model.GlobalMaxQty),
		str.NullString(model.CustomerMaxQty),
		str.NullString(model.DiscPercent),
		str.NullString(model.DiscAmount),
		str.NullString(model.MinimumValue),
		str.NullString(model.Multiply),
		str.NullString(model.Description),
		str.NullString(model.MinimumQty),
		str.NullOrEmtyString(model.MinimumQtyUomID),
		str.NullString(model.PromoType),
		str.NullString(model.Strata)).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
