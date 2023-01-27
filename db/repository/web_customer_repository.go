package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerRepository ...
type IWebCustomerRepository interface {
	SelectAll(c context.Context, parameter models.WebCustomerParameter) ([]models.WebCustomer, error)
	FindAll(ctx context.Context, parameter models.WebCustomerParameter) ([]models.WebCustomer, int, error)
	FindByID(c context.Context, parameter models.WebCustomerParameter) (models.WebCustomer, error)
	Edit(c context.Context, model *models.WebCustomer) (*string, error)
	Add(c context.Context, model *models.WebCustomer) (*string, error)
}

// CustomerRepository ...
type WebCustomerRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewWebCustomerRepository(DB *sql.DB) IWebCustomerRepository {
	return &WebCustomerRepository{DB: DB}
}

// Scan rows
func (repository WebCustomerRepository) scanRows(rows *sql.Rows) (res models.WebCustomer, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerCpName,
		&res.CustomerAddress,
		&res.CustomerProfilePicture,
		&res.CustomerEmail,
		&res.CustomerBirthDate,
		&res.CustomerReligion,
		&res.CustomerActiveStatus,
		&res.CustomerLatitude,
		&res.CustomerLongitude,
		&res.CustomerBranchCode,
		&res.CustomerBranchName,
		&res.CustomerBranchArea,
		&res.CustomerBranchAddress,
		&res.CustomerBranchLat,
		&res.CustomerBranchLng,
		&res.CustomerRegionCode,
		&res.CustomerRegionName,
		&res.CustomerRegionGroup,
		&res.CustomerProvinceID,
		&res.CustomerProvinceName,
		&res.CustomerCityID,
		&res.CustomerCityName,
		&res.CustomerDistrictID,
		&res.CustomerDistrictName,
		&res.CustomerSubdistrictID,
		&res.CustomerSubdistrictName,
		&res.CustomerSalesmanCode,
		&res.CustomerSalesmanName,
		&res.CustomerSalesmanPhone,
		&res.CustomerSalesCycle,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
		&res.CustomerPhone,
		&res.CustomerPoint,
		&res.GiftName,
		&res.Loyalty,
		&res.CustomerTaxCalcMethod,
		&res.CustomerBranchID,
		&res.CustomerSalesmanID,
		&res.CustomerPhotoKtp,
		&res.CustomerNik,
		&res.CustomerLevel,
		&res.CustomerUserID,
		&res.CustomerUserName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebCustomerRepository) scanRow(row *sql.Row) (res models.WebCustomer, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.CustomerName,
		&res.CustomerCpName,
		&res.CustomerAddress,
		&res.CustomerProfilePicture,
		&res.CustomerEmail,
		&res.CustomerBirthDate,
		&res.CustomerReligion,
		&res.CustomerActiveStatus,
		&res.CustomerLatitude,
		&res.CustomerLongitude,
		&res.CustomerBranchCode,
		&res.CustomerBranchName,
		&res.CustomerBranchArea,
		&res.CustomerBranchAddress,
		&res.CustomerBranchLat,
		&res.CustomerBranchLng,
		&res.CustomerRegionCode,
		&res.CustomerRegionName,
		&res.CustomerRegionGroup,
		&res.CustomerProvinceID,
		&res.CustomerProvinceName,
		&res.CustomerCityID,
		&res.CustomerCityName,
		&res.CustomerDistrictID,
		&res.CustomerDistrictName,
		&res.CustomerSubdistrictID,
		&res.CustomerSubdistrictName,
		&res.CustomerSalesmanCode,
		&res.CustomerSalesmanName,
		&res.CustomerSalesmanPhone,
		&res.CustomerSalesCycle,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
		&res.CustomerPhone,
		&res.CustomerPoint,
		&res.GiftName,
		&res.Loyalty,
		&res.CustomerTaxCalcMethod,
		&res.CustomerBranchID,
		&res.CustomerSalesmanID,
		&res.CustomerPhotoKtp,
		&res.CustomerNik,
		&res.CustomerLevel,
		&res.CustomerUserID,
		&res.CustomerUserName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebCustomerRepository) SelectAll(c context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	if parameter.UserId != "" {
		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.UserId + `) `
	}

	statement := models.WebCustomerSelectStatement + ` ` + models.WebCustomerWhereStatement +
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
func (repository WebCustomerRepository) FindAll(ctx context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND c.id = '` + parameter.ID + `'`
	}

	if parameter.UserId != "" {
		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.UserId + `) `
	}

	query := models.WebCustomerSelectStatement + ` ` + models.WebCustomerWhereStatement + ` ` + conditionString + `
		AND (LOWER(c."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	// fmt.Println(query)
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.WebCustomerWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebCustomerRepository) FindByID(c context.Context, parameter models.WebCustomerParameter) (data models.WebCustomer, err error) {
	statement := models.WebCustomerSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	// fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebCustomerRepository) Edit(c context.Context, model *models.WebCustomer) (res *string, err error) {
	statement := `UPDATE customer SET 
	customer_name = $1, 
	customer_address = $2, 
	customer_phone = $3, 
	customer_email = $4,
	customer_cp_name = $5,
	customer_profile_picture = $6,
	tax_calc_method = $7,
	branch_id = $8,
	customer_code =$9,
	salesman_id = $10,
	user_id = $11 
	WHERE id = $12 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerPhone,
		model.CustomerEmail,
		model.CustomerCpName,
		model.CustomerProfilePicture,
		model.CustomerTaxCalcMethod,
		model.CustomerBranchID,
		model.Code,
		model.CustomerSalesmanID,
		model.CustomerUserID,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository WebCustomerRepository) Add(c context.Context, model *models.WebCustomer) (res *string, err error) {
	statement := `INSERT INTO customer (customer_name,customer_address, customer_phone, customer_email,
		customer_cp_name, customer_profile_picture, created_date, modified_date,tax_calc_method, branch_id,customer_code,device_id,
		salesman_id,user_id)
	VALUES ($1, $2, $3, $4, $5, $6, now(), now() ,$7, $8, $9,99, $10,$11) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.CustomerName, model.CustomerAddress,
		model.CustomerPhone, model.CustomerEmail, model.CustomerCpName, model.CustomerProfilePicture,
		model.CustomerTaxCalcMethod, model.CustomerBranchID, model.Code, model.CustomerSalesmanID, model.CustomerUserID,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
