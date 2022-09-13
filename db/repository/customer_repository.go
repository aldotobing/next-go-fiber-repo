package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerRepository ...
type ICustomerRepository interface {
	SelectAll(c context.Context, parameter models.CustomerParameter) ([]models.Customer, error)
	FindAll(ctx context.Context, parameter models.CustomerParameter) ([]models.Customer, int, error)
	FindByID(c context.Context, parameter models.CustomerParameter) (models.Customer, error)
	Edit(c context.Context, model *models.Customer) (*string, error)
	EditAddress(c context.Context, model *models.Customer) (*string, error)
}

// CustomerRepository ...
type CustomerRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewCustomerRepository(DB *sql.DB) ICustomerRepository {
	return &CustomerRepository{DB: DB}
}

// Scan rows
func (repository CustomerRepository) scanRows(rows *sql.Rows) (res models.Customer, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerCpName,
		&res.CustomerAddress,
		&res.CustomerProvinceName,
		&res.CustomerCityName,
		&res.CustomerDistrictName,
		&res.CustomerSubdistrictName,
		&res.CustomerSalesmanCode,
		&res.CustomerSalesmanName,
		&res.CustomerSalesmanPhone,
		&res.CustomerSalesCycle,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
		&res.CustomerPhone,
		&res.Point,
		&res.GiftName,
		&res.Loyalty,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CustomerRepository) scanRow(row *sql.Row) (res models.Customer, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerCpName,
		&res.CustomerAddress,
		&res.CustomerProvinceName,
		&res.CustomerCityName,
		&res.CustomerDistrictName,
		&res.CustomerSubdistrictName,
		&res.CustomerSalesmanCode,
		&res.CustomerSalesmanName,
		&res.CustomerSalesmanPhone,
		&res.CustomerSalesCycle,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
		&res.CustomerPhone,
		&res.Point,
		&res.GiftName,
		&res.Loyalty,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerRepository) SelectAll(c context.Context, parameter models.CustomerParameter) (data []models.Customer, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	statement := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement +
		` AND (LOWER(c."customer_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print
	//fmt.Println(statement)

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
func (repository CustomerRepository) FindAll(ctx context.Context, parameter models.CustomerParameter) (data []models.Customer, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	query := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerRepository) FindByID(c context.Context, parameter models.CustomerParameter) (data models.Customer, err error) {
	statement := models.CustomerSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository CustomerRepository) Edit(c context.Context, model *models.Customer) (res *string, err error) {
	statement := `UPDATE customer SET 
	customer_name = $1, 
	customer_address = $2, 
	customer_phone = $3 
	WHERE id = $4 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerPhone,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (repository CustomerRepository) EditAddress(c context.Context, model *models.Customer) (res *string, err error) {
	statement := `UPDATE customer SET 
	customer_name = $1, 
	customer_address = $2, 
	customer_city_id = $3, 
	customer_district_id = $4, 
	customer_subdistrict_id = $5, 
	customer_postal_code = $6
	WHERE id = $7 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerCityID,
		model.CustomerDistrictID,
		model.CustomerSubdistrictID,
		model.CustomerPostalCode,
		model.ID).Scan(&res)

	fmt.Println(statement)
	if err != nil {
		return res, err
	}
	return res, err
}
