package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerTargetRepository ...
type ICustomerTargetRepository interface {
	SelectAll(c context.Context, parameter models.CustomerTargetParameter) ([]models.CustomerTarget, error)
	FindAll(ctx context.Context, parameter models.CustomerTargetParameter) ([]models.CustomerTarget, int, error)
	// 	FindByID(c context.Context, parameter models.CustomerTargetParameter) (models.CustomerTarget, error)
	// 	Edit(c context.Context, model *models.CustomerTarget) (*string, error)
	// 	EditAddress(c context.Context, model *models.CustomerTarget) (*string, error)
}

// CustomerTargetRepository ...
type CustomerTargetRepository struct {
	DB *sql.DB
}

// NewCustomerTargetRepository ...
func NewCustomerTargetRepository(DB *sql.DB) ICustomerTargetRepository {
	return &CustomerTargetRepository{DB: DB}
}

// Scan rows
func (repository CustomerTargetRepository) scanRows(rows *sql.Rows) (res models.CustomerTarget, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTarget,
		&res.Month,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository CustomerTargetRepository) scanRow(row *sql.Row) (res models.CustomerTarget, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerTarget,
		&res.Month,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerTargetRepository) SelectAll(c context.Context, parameter models.CustomerTargetParameter) (data []models.CustomerTarget, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerTargetSelectStatement + ` ` + models.CustomerTargetWhereStatement +
		` AND (LOWER(c."customer_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print
	// fmt.Println(statement)

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
func (repository CustomerTargetRepository) FindAll(ctx context.Context, parameter models.CustomerTargetParameter) (data []models.CustomerTarget, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	query := models.CustomerTargetSelectStatement + ` ` + models.CustomerTargetWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerTargetWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerTargetRepository) FindByID(c context.Context, parameter models.CustomerTargetParameter) (data models.CustomerTarget, err error) {
	statement := models.CustomerTargetSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository CustomerTargetRepository) Edit(c context.Context, model *models.CustomerTarget) (res *string, err error) {
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
// 		model.CustomerTargetName,
// 		model.CustomerTargetAddress,
// 		model.CustomerTargetPhone,
// 		model.CustomerTargetEmail,
// 		model.CustomerTargetCpName,
// 		model.CustomerTargetProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// func (repository CustomerTargetRepository) EditAddress(c context.Context, model *models.CustomerTarget) (res *string, err error) {
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
// 		model.CustomerTargetName,
// 		model.CustomerTargetAddress,
// 		model.CustomerTargetProvinceID,
// 		model.CustomerTargetCityID,
// 		model.CustomerTargetDistrictID,
// 		model.CustomerTargetSubdistrictID,
// 		model.CustomerTargetPostalCode,
// 		model.ID).Scan(&res)

// 	fmt.Println(statement)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
