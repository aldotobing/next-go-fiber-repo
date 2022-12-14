package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebPromo ...
type IWebPromo interface {
	SelectAll(c context.Context, parameter models.WebPromoParameter) ([]models.WebPromo, error)
	FindAll(ctx context.Context, parameter models.WebPromoParameter) ([]models.WebPromo, int, error)
	Add(c context.Context, parameter *models.WebPromo) (*string, error)
	Delete(c context.Context, id string) (string, error)
	// 	Edit(c context.Context, model *models.WebPromo) (*string, error)
	// 	EditAddress(c context.Context, model *models.WebPromo) (*string, error)
}

// WebPromo ...
type WebPromo struct {
	DB *sql.DB
}

// NewWebPromo ...
func NewWebPromoRepository(DB *sql.DB) IWebPromo {
	return &WebPromo{DB: DB}
}

// Scan rows
func (repository WebPromo) scanRows(rows *sql.Rows) (res models.WebPromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
		&res.ShowInApp,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository WebPromo) scanRow(row *sql.Row) (res models.WebPromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
		&res.ShowInApp,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebPromo) SelectAll(c context.Context, parameter models.WebPromoParameter) (data []models.WebPromo, err error) {
	conditionString := ``

	if (parameter.StartDate != "") && (parameter.EndDate != "") {
		conditionString += ` AND start_date >= ` + `'` +
			parameter.StartDate + `'::date` + ` AND end_date <= ` + `'` + parameter.EndDate + `'::date` +
			` + INTERVAL ` + `'1 MONTH' `
	}

	statement := models.WebPromoSelectStatement + ` ` + models.WebPromoWhereStatement +
		` AND (LOWER(pc._name) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

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
func (repository WebPromo) FindAll(ctx context.Context, parameter models.WebPromoParameter) (data []models.WebPromo, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	query := models.WebPromoSelectStatement + ` ` + models.WebPromoWhereStatement + ` ` + conditionString + `
		AND (LOWER(cus."customer_name") LIKE $1  )` + `GROUP BY ` + `ORDER BY` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.WebPromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebPromo) FindByID(c context.Context, parameter models.WebPromoParameter) (data models.WebPromo, err error) {
	statement := models.WebPromoSelectStatement + ` WHERE pc.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository WebPromo) Edit(c context.Context, model *models.WebPromo) (res *string, err error) {
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
// 		model.WebPromoName,
// 		model.WebPromoAddress,
// 		model.WebPromoPhone,
// 		model.WebPromoEmail,
// 		model.WebPromoCpName,
// 		model.WebPromoProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

func (repository WebPromo) Add(c context.Context, model *models.WebPromo) (res *string, err error) {
	statement := `INSERT INTO promo (code, _name, description, url_banner,
		start_date, end_date, active,show_in_app)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.PromoName, model.PromoDescription, model.PromoUrlBanner,
		model.StartDate, model.EndDate, 1, model.ShowInApp).Scan(&res)

	fmt.Println("PROMO INSERT : " + statement)

	if err != nil {
		fmt.Println("INSERT PROMO BERHASIL! :)")
		return res, err
	}
	return res, err
}

// Delete ...
func (repository WebPromo) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE promo set active = 0 where id= $1 RETURNING id `

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
