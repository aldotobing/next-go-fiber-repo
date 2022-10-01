package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerTargetYearRepository ...
type ICustomerTargetYearRepository interface {
	SelectAll(c context.Context, parameter models.CustomerTargetYearParameter) ([]models.CustomerTargetYear, error)
	FindAll(ctx context.Context, parameter models.CustomerTargetYearParameter) ([]models.CustomerTargetYear, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerTargetYearParameter) (models.CustomerTargetYear, error)
	// 	Edit(c context.Context, model *models.CustomerTargetYear) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerTargetYear) (*string, error)
}

// CustomerTargetYearRepository ...
type CustomerTargetYearRepository struct {
	DB *sql.DB
}

// NewCustomerTargetYearRepository ...
func NewCustomerTargetYearRepository(DB *sql.DB) ICustomerTargetYearRepository {
	return &CustomerTargetYearRepository{DB: DB}
}

// Scan rows
func (repository CustomerTargetYearRepository) scanRows(rows *sql.Rows) (res models.CustomerTargetYear, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTarget,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository CustomerTargetYearRepository) scanRow(row *sql.Row) (res models.CustomerTargetYear, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTarget,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerTargetYearRepository) SelectAll(c context.Context, parameter models.CustomerTargetYearParameter) (data []models.CustomerTargetYear, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerTargetYearSelectStatement + ` ` +
		models.CustomerTargetYearWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1) ` +
		conditionString +
		` GROUP BY cus.id` +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort +
		` ), 0) ` +
		` AS TARGET ` +
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
func (repository CustomerTargetYearRepository) FindAll(ctx context.Context, parameter models.CustomerTargetYearParameter) (data []models.CustomerTargetYear, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	query := models.CustomerTargetYearSelectStatement + ` ` + models.CustomerTargetYearWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerTargetYearWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerTargetYearRepository) FindByID(c context.Context, parameter models.CustomerTargetYearParameter) (data models.CustomerTargetYear, err error) {
	statement := models.CustomerTargetYearSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	//fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerTargetYearRepository) Edit(c context.Context, model *models.CustomerTargetYear) (res *string, err error) {
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
// 		model.CustomerTargetYearName,
// 		model.CustomerTargetYearAddress,
// 		model.CustomerTargetYearPhone,
// 		model.CustomerTargetYearEmail,
// 		model.CustomerTargetYearCpName,
// 		model.CustomerTargetYearProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerTargetYearRepository) EditAddress(c context.Context, model *models.CustomerTargetYear) (res *string, err error) {
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
// 		model.CustomerTargetYearName,
// 		model.CustomerTargetYearAddress,
// 		model.CustomerTargetYearProvinceID,
// 		model.CustomerTargetYearCityID,
// 		model.CustomerTargetYearDistrictID,
// 		model.CustomerTargetYearSubdistrictID,
// 		model.CustomerTargetYearPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
