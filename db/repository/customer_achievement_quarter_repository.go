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

// ICustomerAchievementQuarter ...
type ICustomerAchievementQuarter interface {
	SelectAll(c context.Context, parameter models.CustomerAchievementQuarterParameter) ([]models.CustomerAchievementQuarter, error)
	FindAll(ctx context.Context, parameter models.CustomerAchievementQuarterParameter) ([]models.CustomerAchievementQuarter, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerAchievementQuarterParameter) (models.CustomerAchievementQuarter, error)
	// 	Edit(c context.Context, model *models.CustomerAchievementQuarter) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerAchievementQuarter) (*string, error)
}

// CustomerAchievementQuarter ...
type CustomerAchievementQuarter struct {
	DB *sql.DB
}

// NewCustomerAchievementQuarter ...
func NewCustomerAchievementQuarterRepository(DB *sql.DB) ICustomerAchievementQuarter {
	return &CustomerAchievementQuarter{DB: DB}
}

// Scan rows
func (repository CustomerAchievementQuarter) scanRows(rows *sql.Rows) (res models.CustomerAchievementQuarter, err error) {
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
func (repository CustomerAchievementQuarter) scanRow(row *sql.Row) (res models.CustomerAchievementQuarter, err error) {
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
func (repository CustomerAchievementQuarter) SelectAll(c context.Context, parameter models.CustomerAchievementQuarterParameter) (data []models.CustomerAchievementQuarter, err error) {
	conditionString := ``

	conditionStringQuarter := ``

	quarterGetCurrentYear := ` and extract(year from sih.transaction_date) = (select extract (year from now())) `

	/*
		SET QUARTER
	*/
	currentTime := time.Now()
	month := currentTime.Month()
	quarter := int(math.Ceil(float64(month) / 3))

	if quarter == 1 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (1,2,3)  ` + quarterGetCurrentYear + ` `
	}
	if quarter == 2 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (4,5,6) ` + quarterGetCurrentYear + ` `
	}
	if quarter == 3 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (7,8,9) ` + quarterGetCurrentYear + ` `
	}
	if quarter == 4 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (10,11,12) ` + quarterGetCurrentYear + ` `
	}
	/*
		END QUARTER
	*/
	groupByString := ` GROUP BY CUS.ID, CUSTOMER_CODE, CUSTOMER_NAME `

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerAchievementQuarterSelectStatement + ` ` + models.CustomerAchievementQuarterWhereStatement + conditionStringQuarter +
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
func (repository CustomerAchievementQuarter) FindAll(ctx context.Context, parameter models.CustomerAchievementQuarterParameter) (data []models.CustomerAchievementQuarter, count int, err error) {
	conditionString := ``

	conditionStringQuarter := ``

	quarterGetCurrentYear := ` and extract(year from sih.transaction_date) = (select extract (year from now())) `

	/*
		SET QUARTER
	*/
	currentTime := time.Now()
	month := currentTime.Month()
	quarter := int(math.Ceil(float64(month) / 3))

	if quarter == 1 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (1,2,3)  ` + quarterGetCurrentYear + ` `
	}
	if quarter == 2 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (4,5,6) ` + quarterGetCurrentYear + ` `
	}
	if quarter == 3 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (7,8,9) ` + quarterGetCurrentYear + ` `
	}
	if quarter == 4 {

		conditionStringQuarter += ` AND extract (month from SIH.TRANSACTION_DATE) in (10,11,12) ` + quarterGetCurrentYear + ` `
	}
	/*
		END QUARTER
	*/

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	query := models.CustomerAchievementQuarterSelectStatement + ` ` + models.CustomerAchievementQuarterWhereStatement + conditionStringQuarter + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerAchievementQuarterWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerAchievementQuarter) FindByID(c context.Context, parameter models.CustomerAchievementQuarterParameter) (data models.CustomerAchievementQuarter, err error) {
	statement := models.CustomerAchievementQuarterSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerAchievementQuarter) Edit(c context.Context, model *models.CustomerAchievementQuarter) (res *string, err error) {
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
// 		model.CustomerAchievementQuarterName,
// 		model.CustomerAchievementQuarterAddress,
// 		model.CustomerAchievementQuarterPhone,
// 		model.CustomerAchievementQuarterEmail,
// 		model.CustomerAchievementQuarterCpName,
// 		model.CustomerAchievementQuarterProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerAchievementQuarter) EditAddress(c context.Context, model *models.CustomerAchievementQuarter) (res *string, err error) {
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
// 		model.CustomerAchievementQuarterName,
// 		model.CustomerAchievementQuarterAddress,
// 		model.CustomerAchievementQuarterProvinceID,
// 		model.CustomerAchievementQuarterCityID,
// 		model.CustomerAchievementQuarterDistrictID,
// 		model.CustomerAchievementQuarterSubdistrictID,
// 		model.CustomerAchievementQuarterPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
