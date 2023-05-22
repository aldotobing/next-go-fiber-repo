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
		&res.CustomerGender,
		&res.CustomerActiveStatus,
		&res.CustomerLatitude,
		&res.CustomerLongitude,
		&res.CustomerBranchCode,
		&res.CustomerBranchName,
		&res.CustomerBranchArea,
		&res.CustomerBranchAddress,
		&res.CustomerBranchLat,
		&res.CustomerBranchLng,
		&res.CustomerBranchPicPhoneNo,
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
		&res.CustomerLevelID,
		&res.CustomerUserID,
		&res.CustomerUserName,
		&res.ModifiedBy,
		&res.ModifiedDate,
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
		&res.CustomerGender,
		&res.CustomerActiveStatus,
		&res.CustomerLatitude,
		&res.CustomerLongitude,
		&res.CustomerBranchCode,
		&res.CustomerBranchName,
		&res.CustomerBranchArea,
		&res.CustomerBranchAddress,
		&res.CustomerBranchLat,
		&res.CustomerBranchLng,
		&res.CustomerBranchPicPhoneNo,
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
		&res.CustomerLevelID,
		&res.CustomerUserID,
		&res.CustomerUserName,
		&res.ModifiedBy,
		&res.ModifiedDate,
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

	if parameter.BranchId != "" {
		conditionString += ` AND C.BRANCH_ID= ` + parameter.BranchId
	}

	if parameter.PhoneNumber != "" {
		conditionString += ` AND c.customer_phone LIKE '%` + parameter.PhoneNumber + `%'`
	}

	statement := models.WebCustomerSelectStatement + ` ` + models.WebCustomerWhereStatement +
		` AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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

	if parameter.BranchId != "" {
		conditionString += ` AND C.BRANCH_ID= ` + parameter.BranchId
	}

	if parameter.PhoneNumber != "" {
		conditionString += ` AND c.customer_phone LIKE '%` + parameter.PhoneNumber + `%'`
	}

	query := models.WebCustomerSelectStatement + ` ` + models.WebCustomerWhereStatement + ` ` + conditionString + `
		AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`

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
		conditionString + ` AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1 )`
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
		user_id = $3,
		customer_phone = $4,
		customer_religion = $5,
		customer_nik = $6,
		customer_level_id = $7,
		customer_gender = $8,
		customer_code = $9,
		customer_email = $10,
		customer_birthdate = $11,
		customer_profile_picture = $12,
		customer_photo_ktp = $13,
		customer_cp_name = $14,
		modified_date = now(),
		modified_by = $15
	WHERE id = $16
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerUserID,
		model.CustomerPhone,
		model.CustomerReligion,
		model.CustomerNik,
		model.CustomerLevelID,
		model.CustomerGender,
		model.Code,
		model.CustomerEmail,
		model.CustomerBirthDate,
		model.CustomerProfilePicture,
		model.CustomerPhotoKtp,
		model.CustomerCpName,
		model.ModifiedBy,
		model.ID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository WebCustomerRepository) Add(c context.Context, model *models.WebCustomer) (res *string, err error) {
	statement := `INSERT INTO customer (
			customer_name, customer_address, customer_phone, customer_email,
			customer_cp_name, customer_profile_picture, created_date, modified_date, 
			tax_calc_method, branch_id, customer_code, device_id, 
			salesman_id, user_id, customer_religion, customer_nik,
			customer_level_id, customer_gender, customer_birthdate
		)
	VALUES (
			$1, $2, $3, $4,
			$5, $6, now(), now(),
			$7, $8, $9, 99, 
			$10, $11, $12, $13,
			$14, $15, $16
		) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName, model.CustomerAddress, model.CustomerPhone, model.CustomerEmail,
		model.CustomerCpName, model.CustomerProfilePicture,
		model.CustomerTaxCalcMethod, model.CustomerBranchID, model.Code,
		model.CustomerSalesmanID, model.CustomerUserID, model.CustomerReligion, model.CustomerNik,
		model.CustomerLevelID, model.CustomerGender, model.CustomerBirthDate,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
