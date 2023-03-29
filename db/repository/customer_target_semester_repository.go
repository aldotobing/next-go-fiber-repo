package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerTargetSemesterRepository ...
type ICustomerTargetSemesterRepository interface {
	SelectAll(c context.Context, parameter models.CustomerTargetSemesterParameter) ([]models.CustomerTargetSemester, error)
	FindAll(ctx context.Context, parameter models.CustomerTargetSemesterParameter) ([]models.CustomerTargetSemester, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerTargetSemesterParameter) (models.CustomerTargetSemester, error)
	// 	Edit(c context.Context, model *models.CustomerTargetSemester) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerTargetSemester) (*string, error)
}

// CustomerTargetSemesterRepository ...
type CustomerTargetSemesterRepository struct {
	DB *sql.DB
}

// NewCustomerTargetSemesterRepository ...
func NewCustomerTargetSemesterRepository(DB *sql.DB) ICustomerTargetSemesterRepository {
	return &CustomerTargetSemesterRepository{DB: DB}
}

// Scan rows
func (repository CustomerTargetSemesterRepository) scanRows(rows *sql.Rows) (res models.CustomerTargetSemester, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTargetSemester,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository CustomerTargetSemesterRepository) scanRow(row *sql.Row) (res models.CustomerTargetSemester, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTargetSemester,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerTargetSemesterRepository) SelectAll(c context.Context, parameter models.CustomerTargetSemesterParameter) (data []models.CustomerTargetSemester, err error) {
	conditionString := ``
	conditionStringQuarter := ``

	/*
		SET QUARTER
	*/
	month := time.Now().Month()
	quarter := int(math.Ceil(float64(month) / 6))

	if quarter == 1 {
		conditionStringQuarter += ` and bmt._month in (1, 2, 3, 4, 5, 6) `
	}
	if quarter == 2 {
		conditionStringQuarter += ` and bmt._month in (7, 8, 9, 10, 11, 12) `
	}
	/*
		END QUARTER
	*/

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerTargetSemesterSelectStatement + ` ` +
		models.CustomerTargetSemesterWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1) ` +
		conditionString + conditionStringQuarter +
		` GROUP BY cus.id` + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort +
		` ), 0)` +
		` AS TARGET` +
		` FROM CUSTOMER CUS WHERE CUS.ID = '` + parameter.ID + `'`

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
func (repository CustomerTargetSemesterRepository) FindAll(ctx context.Context, parameter models.CustomerTargetSemesterParameter) (data []models.CustomerTargetSemester, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	query := models.CustomerTargetSemesterSelectStatement + ` ` + models.CustomerTargetSemesterWhereStatement + ` ` + conditionString + `
		AND (LOWER(c."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerTargetSemesterWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerTargetSemesterRepository) FindByID(c context.Context, parameter models.CustomerTargetSemesterParameter) (data models.CustomerTargetSemester, err error) {
	statement := models.CustomerTargetSemesterSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerTargetSemesterRepository) Edit(c context.Context, model *models.CustomerTargetSemester) (res *string, err error) {
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
// 		model.CustomerTargetSemesterName,
// 		model.CustomerTargetSemesterAddress,
// 		model.CustomerTargetSemesterPhone,
// 		model.CustomerTargetSemesterEmail,
// 		model.CustomerTargetSemesterCpName,
// 		model.CustomerTargetSemesterProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerTargetSemesterRepository) EditAddress(c context.Context, model *models.CustomerTargetSemester) (res *string, err error) {
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
// 		model.CustomerTargetSemesterName,
// 		model.CustomerTargetSemesterAddress,
// 		model.CustomerTargetSemesterProvinceID,
// 		model.CustomerTargetSemesterCityID,
// 		model.CustomerTargetSemesterDistrictID,
// 		model.CustomerTargetSemesterSubdistrictID,
// 		model.CustomerTargetSemesterPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
