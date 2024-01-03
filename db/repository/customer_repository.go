package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
	BackendEdit(c context.Context, model *models.Customer) (*string, error)
	BackendAdd(c context.Context, model *models.Customer) (*string, error)
	FindByCodeAndPhone(c context.Context, parameter models.CustomerParameter) (models.Customer, error)
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
		&res.CustomerBranchPicName,
		&res.CustomerBranchPicPhoneNo,
		&res.CustomerRegionID,
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
		&res.CustomerPriceListID,
		&res.CustomerPriceListVersionID,
		&res.CustomerFCMToken,
		&res.CustomerPaymentTermsID,
		&res.CustomerPaymentTermsCode,
		&res.CustomerAdminValidate,
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
		&res.CustomerBranchPicName,
		&res.CustomerBranchPicPhoneNo,
		&res.CustomerRegionID,
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
		&res.CustomerPriceListID,
		&res.CustomerPriceListVersionID,
		&res.CustomerFCMToken,
		&res.CustomerPaymentTermsID,
		&res.CustomerPaymentTermsCode,
		&res.CustomerAdminValidate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerRepository) SelectAll(c context.Context, parameter models.CustomerParameter) (data []models.Customer, err error) {
	var conditionString string
	var args []interface{}
	var index int = 1

	if parameter.ID != "" {
		conditionString += ` AND c.id = $` + strconv.Itoa(index)
		args = append(args, parameter.ID)
		index++
	}

	if parameter.UserId != "" {
		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = $` + strconv.Itoa(index) + `) `
		args = append(args, parameter.UserId)
		index++
	}

	if parameter.FlagToken {
		conditionString += ` AND us.FCM_TOKEN IS NOT NULL `
	}

	if parameter.CustomerTypeId != "" {
		conditionString += ` AND C.CUSTOMER_TYPE_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerTypeId)
		index++
	}

	if parameter.BranchID != "" {
		conditionString += ` AND C.BRANCH_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.BranchID)
		index++
	}

	if parameter.RegionID != "" {
		conditionString += ` AND REG.ID = $` + strconv.Itoa(index)
		args = append(args, parameter.RegionID)
		index++
	}

	if parameter.RegionGroupID != "" {
		conditionString += ` AND C.GROUP_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.RegionGroupID)
		index++
	}

	if parameter.CustomerLevelId != "" {
		conditionString += ` AND C.CUSTOMER_LEVEL_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerLevelId)
		index++
	}

	if parameter.CustomerCodes != "" {
		conditionString += ` AND C.CUSTOMER_CODE in (` + parameter.CustomerCodes + `)`
	}

	if parameter.CustomerReligion != "" {
		conditionString += ` AND C.customer_religion = '` + parameter.CustomerReligion + `'`
	}

	statement := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement +
		` AND (LOWER(c."customer_name") LIKE $` + strconv.Itoa(index) + `) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	args = append(args, "%"+strings.ToLower(parameter.Search)+"%")

	rows, err := repository.DB.QueryContext(c, statement, args...)

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

// func (repository CustomerRepository) SelectAll(c context.Context, parameter models.CustomerParameter) (data []models.Customer, err error) {
// 	conditionString := ``

// 	if parameter.ID != "" {
// 		conditionString += ` AND c.id = '` + parameter.ID + `'`
// 	}

// 	if parameter.UserId != "" {
// 		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.UserId + `) `
// 	}

// 	if parameter.FlagToken {
// 		conditionString += ` AND us.FCM_TOKEN IS NOT NULL `
// 	}

// 	statement := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement +
// 		` AND (LOWER(c."customer_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
// 	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

// 	//print
// 	// fmt.Println(statement)

// 	if err != nil {
// 		return data, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {

// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, temp)
// 	}

// 	return data, err
// }

// FindAll ...
func (repository CustomerRepository) FindAll(ctx context.Context, parameter models.CustomerParameter) (data []models.Customer, count int, err error) {
	var conditionString string
	var args []interface{}
	var index int = 1

	if parameter.ID != "" {
		conditionString += ` AND c.id = $` + strconv.Itoa(index)
		args = append(args, parameter.ID)
		index++
	}

	if parameter.UserId != "" {
		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = $` + strconv.Itoa(index) + `) `
		args = append(args, parameter.UserId)
		index++
	}

	query := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement + ` ` + conditionString +
		` AND (LOWER(c."customer_name") LIKE $` + strconv.Itoa(index) + `) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $` + strconv.Itoa(index+1) + ` LIMIT $` + strconv.Itoa(index+2)

	args = append(args, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	rows, err := repository.DB.QueryContext(ctx, query, args...)

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

	countQuery := `SELECT COUNT(*) FROM "customer" c ` + models.CustomerWhereStatement + ` ` + conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(countQuery, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// func (repository CustomerRepository) FindAll(ctx context.Context, parameter models.CustomerParameter) (data []models.Customer, count int, err error) {
// 	conditionString := ``

// 	if parameter.ID != "" {
// 		conditionString += ` AND c.id = '` + parameter.ID + `'`
// 	}

// 	if parameter.UserId != "" {
// 		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.UserId + `) `
// 	}

// 	query := models.CustomerSelectStatement + ` ` + models.CustomerWhereStatement + ` ` + conditionString + `
// 		AND (LOWER(c."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
// 	// fmt.Println(query)
// 	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
// 	if err != nil {
// 		return data, count, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, count, err
// 		}
// 		data = append(data, temp)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return data, count, err
// 	}

// 	query = `SELECT COUNT(*) FROM "customer" c ` + models.CustomerWhereStatement + ` ` +
// 		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
// 	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
// 	return data, count, err
// }

// FindByID ...
func (repository CustomerRepository) FindByID(c context.Context, parameter models.CustomerParameter) (data models.Customer, err error) {
	statement := models.CustomerSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	// fmt.Println(statement)

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
	customer_phone = $3, 
	customer_email = $4,
	customer_cp_name = $5,
	customer_profile_picture = $6, 
	customer_photo_ktp = $7,
	customer_religion = $8,
	customer_birthdate = $9,
	customer_nik =$10
	WHERE id = $11 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerPhone,
		model.CustomerEmail,
		model.CustomerCpName,
		model.CustomerProfilePicture,
		model.CustomerPhotoKtp,
		model.CustomerReligion,
		model.CustomerBirthDate,
		model.CustomerNik,
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
	customer_province_id = $3,
	customer_city_id = $4, 
	customer_district_id = $5, 
	customer_subdistrict_id = $6, 
	customer_postal_code = $7
	WHERE id = $8 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName,
		model.CustomerAddress,
		model.CustomerProvinceID,
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

// Edit ...
func (repository CustomerRepository) BackendEdit(c context.Context, model *models.Customer) (res *string, err error) {
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
	salesman_id = $10 
	WHERE id = $11 
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
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository CustomerRepository) BackendAdd(c context.Context, model *models.Customer) (res *string, err error) {
	statement := `INSERT INTO customer (customer_name,customer_address, customer_phone, customer_email,
		customer_cp_name, customer_profile_picture, created_date, modified_date,tax_calc_method, branch_id,customer_code,device_id,
		salesman_id)
	VALUES ($1, $2, $3, $4, $5, $6, now(), now() ,$7, $8, $9,99, $10) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.CustomerName, model.CustomerAddress,
		model.CustomerPhone, model.CustomerEmail, model.CustomerCpName, model.CustomerProfilePicture,
		model.CustomerTaxCalcMethod, model.CustomerBranchID, model.Code, model.CustomerSalesmanID,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// FindByID ...
func (repository CustomerRepository) FindByCodeAndPhone(c context.Context, parameter models.CustomerParameter) (data models.Customer, err error) {
	statement := models.CustomerSelectStatement + ` WHERE c.customer_code = $1 and c.customer_phone = $2 `
	row := repository.DB.QueryRowContext(c, statement, parameter.Code, parameter.Phone)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
