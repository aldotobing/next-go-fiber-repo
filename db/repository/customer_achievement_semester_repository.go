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

// ICustomerAchievementSemester ...
type ICustomerAchievementSemester interface {
	SelectAll(c context.Context, parameter models.CustomerAchievementSemesterParameter) ([]models.CustomerAchievementSemester, error)
	FindAll(ctx context.Context, parameter models.CustomerAchievementSemesterParameter) ([]models.CustomerAchievementSemester, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerAchievementSemesterParameter) (models.CustomerAchievementSemester, error)
	// 	Edit(c context.Context, model *models.CustomerAchievementSemester) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerAchievementSemester) (*string, error)
}

// CustomerAchievementSemester ...
type CustomerAchievementSemester struct {
	DB *sql.DB
}

// NewCustomerAchievementSemester ...
func NewCustomerAchievementSemesterRepository(DB *sql.DB) ICustomerAchievementSemester {
	return &CustomerAchievementSemester{DB: DB}
}

// Scan rows
func (repository CustomerAchievementSemester) scanRows(rows *sql.Rows) (res models.CustomerAchievementSemester, err error) {
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
func (repository CustomerAchievementSemester) scanRow(row *sql.Row) (res models.CustomerAchievementSemester, err error) {
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
func (repository CustomerAchievementSemester) SelectAll(c context.Context, parameter models.CustomerAchievementSemesterParameter) (data []models.CustomerAchievementSemester, err error) {
	conditionString := ``

	conditionStringSemester := ` and extract(year from sih.transaction_date) = (select extract (year from now())) `

	/*
		SET QUARTER
	*/
	month := time.Now().Month()
	semester := int(math.Ceil(float64(month) / 6))

	if semester == 1 {
		conditionStringSemester += ` AND extract (month from SIH.TRANSACTION_DATE) in (1,2,3,4,5,6)  `
	}
	if semester == 2 {
		conditionStringSemester += ` AND extract (month from SIH.TRANSACTION_DATE) in (7,8,9,10,11,12) `
	}

	/*
		END QUARTER
	*/
	groupByString := ` GROUP BY CUS.ID, CUSTOMER_CODE, CUSTOMER_NAME `

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerAchievementSemesterSelectStatement + ` ` + models.CustomerAchievementSemesterWhereStatement + conditionStringSemester +
		` AND (LOWER(cus."customer_name") LIKE $1) ` + conditionString + groupByString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort +
		` ), 0) AS ACHIEVEMENT ` +
		` FROM CUSTOMER CUS ` +
		` WHERE CUS.ID = '` + parameter.ID + `'`

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
func (repository CustomerAchievementSemester) FindAll(ctx context.Context, parameter models.CustomerAchievementSemesterParameter) (data []models.CustomerAchievementSemester, count int, err error) {
	conditionString := ``

	conditionStringSemester := ` and extract(year from sih.transaction_date) = (select extract (year from now())) `

	/*
		SET QUARTER
	*/
	currentTime := time.Now()
	month := currentTime.Month()
	quarter := int(math.Ceil(float64(month) / 6))

	if quarter == 1 {
		conditionStringSemester += ` AND extract (month from SIH.TRANSACTION_DATE) in (1,2,3,4,5,6)  `
	}
	if quarter == 2 {
		conditionStringSemester += ` AND extract (month from SIH.TRANSACTION_DATE) in (7,8,9,10,11,12) `
	}
	/*
		END QUARTER
	*/

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	query := models.CustomerAchievementSemesterSelectStatement + ` ` + models.CustomerAchievementSemesterWhereStatement + conditionStringSemester + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerAchievementSemesterWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerAchievementSemester) FindByID(c context.Context, parameter models.CustomerAchievementSemesterParameter) (data models.CustomerAchievementSemester, err error) {
	statement := models.CustomerAchievementSemesterSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerAchievementSemester) Edit(c context.Context, model *models.CustomerAchievementSemester) (res *string, err error) {
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
// 		model.CustomerAchievementSemesterName,
// 		model.CustomerAchievementSemesterAddress,
// 		model.CustomerAchievementSemesterPhone,
// 		model.CustomerAchievementSemesterEmail,
// 		model.CustomerAchievementSemesterCpName,
// 		model.CustomerAchievementSemesterProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerAchievementSemester) EditAddress(c context.Context, model *models.CustomerAchievementSemester) (res *string, err error) {
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
// 		model.CustomerAchievementSemesterName,
// 		model.CustomerAchievementSemesterAddress,
// 		model.CustomerAchievementSemesterProvinceID,
// 		model.CustomerAchievementSemesterCityID,
// 		model.CustomerAchievementSemesterDistrictID,
// 		model.CustomerAchievementSemesterSubdistrictID,
// 		model.CustomerAchievementSemesterPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
