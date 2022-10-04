package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerAchievementYear ...
type ICustomerAchievementYear interface {
	SelectAll(c context.Context, parameter models.CustomerAchievementYearParameter) ([]models.CustomerAchievementYear, error)
	FindAll(ctx context.Context, parameter models.CustomerAchievementYearParameter) ([]models.CustomerAchievementYear, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerAchievementYearParameter) (models.CustomerAchievementYear, error)
	// 	Edit(c context.Context, model *models.CustomerAchievementYear) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerAchievementYear) (*string, error)
}

// CustomerAchievementYear ...
type CustomerAchievementYear struct {
	DB *sql.DB
}

// NewCustomerAchievementYear ...
func NewCustomerAchievementYearRepository(DB *sql.DB) ICustomerAchievementYear {
	return &CustomerAchievementYear{DB: DB}
}

// Scan rows
func (repository CustomerAchievementYear) scanRows(rows *sql.Rows) (res models.CustomerAchievementYear, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.Achievement,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository CustomerAchievementYear) scanRow(row *sql.Row) (res models.CustomerAchievementYear, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.Achievement,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerAchievementYear) SelectAll(c context.Context, parameter models.CustomerAchievementYearParameter) (data []models.CustomerAchievementYear, err error) {
	conditionString := ``
	groupByString := ` GROUP BY CUS.ID, CUSTOMER_CODE, CUSTOMER_NAME `

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerAchievementYearSelectStatement + ` ` + models.CustomerAchievementYearWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1) ` + conditionString + groupByString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort +
		` ), 0) AS ACHIEVEMENT ` +
		` FROM CUSTOMER CUS ` +
		` WHERE CUS.ID = '` + parameter.ID + `'`

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
func (repository CustomerAchievementYear) FindAll(ctx context.Context, parameter models.CustomerAchievementYearParameter) (data []models.CustomerAchievementYear, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	query := models.CustomerAchievementYearSelectStatement + ` ` + models.CustomerAchievementYearWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerAchievementYearWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerAchievementYear) FindByID(c context.Context, parameter models.CustomerAchievementYearParameter) (data models.CustomerAchievementYear, err error) {
	statement := models.CustomerAchievementYearSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerAchievementYear) Edit(c context.Context, model *models.CustomerAchievementYear) (res *string, err error) {
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
// 		model.CustomerAchievementYearName,
// 		model.CustomerAchievementYearAddress,
// 		model.CustomerAchievementYearPhone,
// 		model.CustomerAchievementYearEmail,
// 		model.CustomerAchievementYearCpName,
// 		model.CustomerAchievementYearProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerAchievementYear) EditAddress(c context.Context, model *models.CustomerAchievementYear) (res *string, err error) {
// 	statement := `UPDATE customer SET
// 	customer_name = $1,
// 	customer_address = $2,
// 	customer_province_id = $3,
// 	customer_city_id = $4,
// 	customer_district_id = $5,
// 	customer_subdistrict_id = $6,
// 	customer_postal_code = $7
// 	WHERE id = $8
// 	RETURNING id`
// 	err = repository.DB.QueryRowContext(c, statement,
// 		model.CustomerAchievementYearName,
// 		model.CustomerAchievementYearAddress,
// 		model.CustomerAchievementYearProvinceID,
// 		model.CustomerAchievementYearCityID,
// 		model.CustomerAchievementYearDistrictID,
// 		model.CustomerAchievementYearSubdistrictID,
// 		model.CustomerAchievementYearPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
