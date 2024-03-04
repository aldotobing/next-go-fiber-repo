package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ICustomerRepository ...
type IWebCustomerRepository interface {
	SelectAll(c context.Context, parameter models.WebCustomerParameter) ([]models.WebCustomer, error)
	FindAll(ctx context.Context, parameter models.WebCustomerParameter) ([]models.WebCustomer, int, error)
	FindByID(c context.Context, parameter models.WebCustomerParameter) (models.WebCustomer, error)
	FindByIDNoCache(c context.Context, parameter models.WebCustomerParameter) (models.WebCustomer, error)
	FindByCodes(c context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, err error)
	Edit(c context.Context, model models.WebCustomer) (string, error)
	EditBulk(c context.Context, in requests.WebCustomerBulkRequest) error
	EditMaxPoint(c context.Context, in requests.WebCustomerMaxPointRequestHeader) error
	EditIndexPoint(c context.Context, in []viewmodel.PointRuleCustomerVM) error
	Add(c context.Context, model models.WebCustomer) (string, error)
	ReportSelect(c context.Context, parameter models.WebCustomerReportParameter) ([]models.WebCustomer, error)
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
		&res.CustomerBranchPicName,
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
		&res.CustomerUserToken,
		&res.CustomerUserFirstLoginTime,
		&res.ModifiedBy,
		&res.ModifiedDate,
		&res.CustomerPriceListID,
		&res.CustomerPriceListName,
		&res.ShowInApp,
		&res.IsDataComplete,
		&res.SalesmanTypeCode,
		&res.SalesmanTypeName,
		&res.CustomerAdminValidate,
		&res.IndexPoint,
		&res.CustomerPhotoKtpDashboard,
		&res.CustomerPhotoNpwp,
		&res.CustomerPhotoNpwpDashboard,
		&res.MonthlyMaxPoint,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// scanRowsReport
func (repository WebCustomerRepository) scanRowsReport(rows *sql.Rows) (res models.WebCustomer, err error) {
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
		&res.CustomerSalesCycle,
		&res.CustomerTypeId,
		&res.CustomerPhone,
		&res.CustomerTaxCalcMethod,
		&res.CustomerBranchID,
		&res.CustomerSalesmanID,
		&res.CustomerPhotoKtp,
		&res.CustomerNik,
		&res.CustomerLevelID,
		&res.CustomerUserID,
		&res.ModifiedDate,
		&res.CreatedDate,
		&res.RegionID,
		&res.RegionGroupID,
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
		&res.CustomerCityID,
		&res.CustomerDistrictID,
		&res.CustomerSubdistrictID,
		&res.CustomerSalesmanID,
		&res.ModifiedBy,
		&res.CustomerUserName,
		&res.IsDataComplete,
		&res.SalesmanTypeCode,
		&res.SalesmanTypeName,
		&res.CustomerUserToken,
		&res.CustomerUserFirstLoginTime,
		&res.CustomerAdminValidate,
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
		&res.CustomerBranchPicName,
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
		&res.CustomerUserToken,
		&res.CustomerUserFirstLoginTime,
		&res.ModifiedBy,
		&res.ModifiedDate,
		&res.CustomerPriceListID,
		&res.CustomerPriceListName,
		&res.ShowInApp,
		&res.IsDataComplete,
		&res.SalesmanTypeCode,
		&res.SalesmanTypeName,
		&res.CustomerAdminValidate,
		&res.IndexPoint,
		&res.CustomerPhotoKtpDashboard,
		&res.CustomerPhotoNpwp,
		&res.CustomerPhotoNpwpDashboard,
		&res.MonthlyMaxPoint,
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
		conditionString += ` AND c.customer_phone LIKE '` + parameter.PhoneNumber + `'`
	}

	if parameter.CustomerTypeId != "" {
		conditionString += ` AND c.customer_type_id = ` + parameter.CustomerTypeId
	}

	if parameter.SalesmanTypeID != "" {
		conditionString += ` AND st.id = ` + parameter.SalesmanTypeID
	}

	if parameter.Code != "" {
		conditionString += ` AND C.CUSTOMER_CODE IN (` + parameter.Code + `)`
	}

	var whereStatement string
	if parameter.ShowInApp == "" || parameter.ShowInApp == "1" {
		whereStatement = models.WebCustomerWhereStatement
	} else {
		whereStatement = models.WebCustomerWhereStatementAll
	}

	statement := models.WebCustomerSelectStatement + ` ` + whereStatement +
		` AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
// func (repository WebCustomerRepository) FindAll(ctx context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, count int, err error) {
// 	conditionString := ``

// 	if parameter.Search == "" {
// 		parameter.Search = ""
// 	}

// 	if parameter.ID != "" {
// 		conditionString += ` AND c.id = '` + parameter.ID + `'`
// 	}

// 	if parameter.UserId != "" {
// 		conditionString += ` AND C.BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.UserId + `) `
// 	}

// 	if parameter.BranchId != "" {
// 		conditionString += ` AND C.BRANCH_ID= ` + parameter.BranchId
// 	} else {
// 		parameter.BranchId = ""
// 	}

// 	if parameter.PhoneNumber != "" {
// 		conditionString += ` AND c.customer_phone LIKE '%` + parameter.PhoneNumber + `%'`
// 	} else {
// 		parameter.PhoneNumber = ""
// 	}

// 	query := models.WebCustomerSelectStatement + ` ` + models.WebCustomerWhereStatement + ` ` + conditionString + `
// 		AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`

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

// 	query = `SELECT COUNT(*) FROM "customer" c ` + models.WebCustomerWhereStatement + ` ` +
// 		conditionString + ` AND (LOWER(c.customer_name) LIKE $1 or LOWER(c.customer_code) LIKE $1 )`
// 	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
// 	return data, count, err
// }

func (repository WebCustomerRepository) FindAll(ctx context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, count int, err error) {
	conditionString := ``

	var args []interface{}
	var index int = 1

	if parameter.Search == "" {
		parameter.Search = ""
	}

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

	if parameter.BranchId != "" {
		conditionString += ` AND C.BRANCH_ID= $` + strconv.Itoa(index)
		args = append(args, parameter.BranchId)
		index++
	}

	if parameter.PhoneNumber != "" {
		conditionString += ` AND c.customer_phone LIKE $` + strconv.Itoa(index)
		args = append(args, "%"+parameter.PhoneNumber+"%")
		index++
	}

	args = append(args, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)

	var whereStatement string
	if parameter.ShowInApp == "" || parameter.ShowInApp == "1" {
		whereStatement = models.WebCustomerWhereStatement
	} else {
		whereStatement = models.WebCustomerWhereStatementAll
	}

	if parameter.Active != "" {
		conditionString += `AND C.ACTIVE = '` + parameter.Active + `'`
	}

	if parameter.IsDataComplete != "" {
		if parameter.IsDataComplete == "1" {
			conditionString += `AND C.is_data_completed = true`
		} else {
			conditionString += `AND C.is_data_completed = false`
		}
	}

	if parameter.AdminValidate == "1" {
		conditionString += `AND C.admin_validate = true`
	} else if parameter.AdminValidate == "0" {
		conditionString += `AND C.admin_validate = false`
	}

	if parameter.MonthlyMaxPoint != "" {
		conditionString += ` AND C.MONTHLY_MAX_POINT IS NOT NULL AND C.MONTHLY_MAX_POINT != 0`
	}

	query := models.WebCustomerSelectStatement + ` ` + whereStatement + ` ` + conditionString + `
		AND (LOWER(c.customer_name) LIKE $` + strconv.Itoa(index) + ` or LOWER(c.customer_code) LIKE $` + strconv.Itoa(index) + `) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $` + strconv.Itoa(index+1) + ` LIMIT $` + strconv.Itoa(index+2)

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

	queryCount := `SELECT COUNT(*) FROM "customer" c ` + whereStatement + ` ` + conditionString + ` AND (LOWER(c.customer_name) LIKE $` + strconv.Itoa(index) + ` or LOWER(c.customer_code) LIKE $` + strconv.Itoa(index) + `)`
	err = repository.DB.QueryRowContext(ctx, queryCount, args[:index]...).Scan(&count) // Reusing the args but slicing to the appropriate length for this query.
	return data, count, err
}

// FindByID ...
func (repository WebCustomerRepository) FindByID(c context.Context, parameter models.WebCustomerParameter) (data models.WebCustomer, err error) {
	statement := models.WebCustomerSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCodes ...
func (repository WebCustomerRepository) FindByCodes(c context.Context, parameter models.WebCustomerParameter) (data []models.WebCustomer, err error) {
	statement := models.WebCustomerSelectStatement + ` WHERE c.customer_code in (` + parameter.Code + `)`
	rows, err := repository.DB.QueryContext(c, statement)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, nil
}

// FindByID ...
func (repository WebCustomerRepository) FindByIDNoCache(c context.Context, parameter models.WebCustomerParameter) (data models.WebCustomer, err error) {
	statement := models.WebCustomerSelectStatement + ` WHERE c.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebCustomerRepository) Edit(c context.Context, model models.WebCustomer) (res string, err error) {
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
		modified_by = $15,
		show_in_apps = $16,
		admin_validate = $17,
		customer_photo_ktp_dashboard = $18,
		customer_photo_npwp = $19,
		customer_photo_npwp_dashboard = $20
	WHERE id = $21
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName.String,
		model.CustomerAddress.String,
		model.CustomerUserID.String,
		model.CustomerPhone.String,
		model.CustomerReligion.String,
		model.CustomerNik.String,
		model.CustomerLevelID.Int64,
		model.CustomerGender.String,
		model.Code.String,
		model.CustomerEmail.String,
		model.CustomerBirthDate.String,
		model.CustomerProfilePicture.String,
		model.CustomerPhotoKtp.String,
		model.CustomerCpName.String,
		model.UserID.Int64,
		model.ShowInApp.String,
		model.CustomerAdminValidate,
		model.CustomerPhotoKtpDashboard.String,
		model.CustomerPhotoNpwp.String,
		model.CustomerPhotoNpwpDashboard.String,
		model.ID.String).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// EditBulk ...
func (repo WebCustomerRepository) EditBulk(c context.Context, in requests.WebCustomerBulkRequest) (err error) {
	var customersCode string
	for _, datum := range in.Customers {
		if customersCode == "" {
			customersCode += `'` + datum.Code + `'`
		} else {
			customersCode += `, '` + datum.Code + `'`
		}
	}
	statement := `UPDATE customer SET 
		show_in_apps = $1,
		active = $2,
		modified_by = $3
	WHERE customer_code in (` + customersCode + `)`
	err = repo.DB.QueryRowContext(c, statement,
		in.ShowInApp,
		in.Active,
		in.UserID).Err()

	return
}

// EditMaxPoint ...
func (repo WebCustomerRepository) EditMaxPoint(c context.Context, in requests.WebCustomerMaxPointRequestHeader) (err error) {
	var customersCode string

	for _, datum := range in.Detail {
		if customersCode == "" {
			customersCode += `('` + datum.CustomerCode + `', ` + datum.MonthlyMaxPoint + `)`
		} else {
			customersCode += `, ('` + datum.CustomerCode + `', ` + datum.MonthlyMaxPoint + `)`
		}
	}

	statement := `update customer as c set
		monthly_max_point = val.column_a
	from (values
		` + customersCode + `
	) as val(column_b, column_a) 
	where c.customer_code = val.column_b;`

	err = repo.DB.QueryRowContext(c, statement).Err()

	return
}

func (repo WebCustomerRepository) EditIndexPoint(c context.Context, in []viewmodel.PointRuleCustomerVM) (err error) {
	var customersCode string

	for _, datum := range in {
		if customersCode == "" {
			customersCode += `('` + datum.CustomerCode + `', ` + datum.Value + `)`
		} else {
			customersCode += `, ('` + datum.CustomerCode + `', ` + datum.Value + `)`
		}
	}

	statement := `update customer as c set
		index_point = val.column_a
	from (values
		` + customersCode + `
	) as val(column_b, column_a) 
	where c.customer_code = val.column_b;`

	err = repo.DB.QueryRowContext(c, statement).Err()

	return
}

// Add ...
func (repository WebCustomerRepository) Add(c context.Context, model models.WebCustomer) (res string, err error) {
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

	customerUserID, _ := strconv.Atoi(model.CustomerUserID.String)

	err = repository.DB.QueryRowContext(c, statement,
		model.CustomerName.String, model.CustomerAddress.String, model.CustomerPhone.String, model.CustomerEmail.String,
		model.CustomerCpName.String, model.CustomerProfilePicture.String,
		model.CustomerTaxCalcMethod.String, model.CustomerBranchID.String, model.Code.String,
		model.CustomerSalesmanID.String, customerUserID, model.CustomerReligion.String, model.CustomerNik.String,
		int(model.CustomerLevelID.Int64), model.CustomerGender.String, model.CustomerBirthDate.String,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// ReportSelect ...
func (repository WebCustomerRepository) ReportSelect(c context.Context, parameter models.WebCustomerReportParameter) (data []models.WebCustomer, err error) {
	conditionString := ``

	var args []interface{}
	var index int = 1

	if parameter.RegionID != "" {
		conditionString += ` AND region_id = $` + strconv.Itoa(index)
		args = append(args, parameter.RegionID)
		index++
	}

	if parameter.RegionGroupID != "" {
		conditionString += ` AND region_group_id = $` + strconv.Itoa(index)
		args = append(args, parameter.RegionGroupID)
		index++
	}

	if parameter.BranchArea != "" {
		conditionString += ` AND LOWER(branch_area) LIKE LOWER($` + strconv.Itoa(index) + `)`
		args = append(args, "%"+parameter.BranchArea+"%")
		index++
	}

	if parameter.CustomerTypeID != "" {
		conditionString += ` AND cust_type_id = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerTypeID)
		index++
	}

	if parameter.BranchIDs != "" {
		ids := strings.Split(parameter.BranchIDs, ",")
		placeholders := make([]string, len(ids))
		for i, id := range ids {
			placeholders[i] = "$" + strconv.Itoa(index)
			args = append(args, id)
			index++
		}
		conditionString += ` AND c_branch_id IN (` + strings.Join(placeholders, ",") + `)`
	}

	if parameter.CustomerLevelID != "" {
		conditionString += ` AND CUSTOMER_LEVEL_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.CustomerLevelID)
		index++
	}

	if parameter.AdminUserID != "" {
		conditionString += ` AND C_BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = $` + strconv.Itoa(index) + `)`
		args = append(args, parameter.AdminUserID)
		index++
	}

	statement := `SELECT * FROM v_customer_report WHERE customer_created_date IS NOT NULL ` + conditionString
	rows, err := repository.DB.QueryContext(c, statement, args...)

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRowsReport(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

// func (repository WebCustomerRepository) ReportSelect(c context.Context, parameter models.WebCustomerReportParameter) (data []models.WebCustomer, err error) {
// 	conditionString := ``

// 	if parameter.RegionID != "" {
// 		conditionString += ` AND region_id = '` + parameter.RegionID + `'`
// 	}

// 	if parameter.RegionGroupID != "" {
// 		conditionString += ` AND region_group_id = '` + parameter.RegionGroupID + `'`
// 	}

// 	if parameter.BranchArea != "" {
// 		conditionString += ` AND LOWER(branch_area) LIKE LOWER('%` + parameter.BranchArea + `%')`
// 	}

// 	if parameter.CustomerTypeID != "" {
// 		conditionString += ` AND cust_type_id = '` + parameter.CustomerTypeID + `'`
// 	}

// 	if parameter.BranchIDs != "" {
// 		conditionString += ` AND c_branch_id IN (` + parameter.BranchIDs + `)`
// 	}

// 	if parameter.CustomerLevelID != "" {
// 		conditionString += ` AND CUSTOMER_LEVEL_ID = '` + parameter.CustomerLevelID + `'`
// 	}

// 	if parameter.AdminUserID != "" {
// 		conditionString += ` AND C_BRANCH_ID IN (SELECT BRANCH_ID FROM USER_BRANCH UB WHERE UB.USER_ID = ` + parameter.AdminUserID + `)`
// 	}

// 	statement := `select * from v_customer_report WHERE customer_created_date IS NOT NULL ` + conditionString
// 	rows, err := repository.DB.QueryContext(c, statement)

// 	if err != nil {
// 		return data, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {

// 		temp, err := repository.scanRowsReport(rows)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, temp)
// 	}

// 	return data, err
// }
