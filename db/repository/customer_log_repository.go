package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerLogRepository ...
type ICustomerLogRepository interface {
	SelectAll(c context.Context, parameter models.CustomerLogParameter) ([]models.CustomerLog, error)
	Add(c context.Context, oldIn, newIn interface{}, customerID string, userID int) error
	AddByUser(c context.Context, oldIn, newIn interface{}, customerID string) error
	AddBulk(c context.Context, model *models.WebCustomer) (*string, error)
}

// CustomerLogRepository ...
type CustomerLogRepository struct {
	DB *sql.DB
}

// NewCustomerLogRepository ...
func NewCustomerLogRepository(DB *sql.DB) ICustomerLogRepository {
	return &CustomerLogRepository{DB: DB}
}

// Scan rows
func (repository CustomerLogRepository) scanRows(rows *sql.Rows) (res models.CustomerLog, err error) {
	err = rows.Scan(
		&res.ID,
		&res.CustomerID,
		&res.CustomerCode,
		&res.CustomerName,
		&res.OldData,
		&res.NewData,
		&res.UserID,
		&res.UserName,
		&res.CreatedAt,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerLogRepository) SelectAll(c context.Context, in models.CustomerLogParameter) (data []models.CustomerLog, err error) {
	var conditionString string
	if in.StartDate != "" && in.EndDate != "" {
		conditionString += `AND def.created_at::date BETWEEN '` + in.StartDate + `'::date AND '` + in.EndDate + `'::date`
	} else {
		conditionString += `AND def.created_at::date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if in.UserId != "" {
		conditionString += ` AND c.branch_id in(
				select ub.branch_id  
				from user_branch ub
				where ub.user_id = ` + in.UserId + `
			) `
	}
	if in.RegionGroupID != "" {
		conditionString += ` AND R.GROUP_ID = ` + in.RegionGroupID
	}
	if in.RegionID != "" {
		conditionString += ` AND R.ID = ` + in.RegionID
	}
	if in.BranchID != "" {
		conditionString += ` AND B.ID = ` + in.BranchID
	}
	if in.CustomerLevelID != "" {
		conditionString += ` AND C.customer_level_id = ` + in.CustomerLevelID
	}
	statement := models.CustomerLogSelectStatement + ` ` + models.CustomerLogWhereStatement +
		conditionString + ` ORDER BY ` + in.By + ` ` + in.Sort
	rows, err := repository.DB.QueryContext(c, statement)
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

// Add ...
func (repository CustomerLogRepository) Add(c context.Context, oldIn, newIn interface{}, customerID string, userID int) (err error) {
	oldInJson, _ := json.Marshal(oldIn)
	newInJson, _ := json.Marshal(newIn)
	statement := `INSERT INTO customer_logs (
			customer_id, old_data, new_data, user_id,
			created_at, updated_at
		)
	VALUES (
			$1, $2, $3, $4,
			now(), now()
		)`
	err = repository.DB.QueryRowContext(c, statement,
		customerID, oldInJson, newInJson, userID,
	).Err()

	return
}

// AddByUser ...
func (repository CustomerLogRepository) AddByUser(c context.Context, oldIn, newIn interface{}, customerID string) (err error) {
	oldInJson, _ := json.Marshal(oldIn)
	newInJson, _ := json.Marshal(newIn)
	statement := `INSERT INTO customer_logs (
			customer_id, old_data, new_data,
			created_at, updated_at
		)
	VALUES (
			$1, $2, $3, $4,
			now(), now()
		)`
	err = repository.DB.QueryRowContext(c, statement,
		customerID, oldInJson, newInJson,
	).Err()

	return
}

// AddBulk ...
func (repository CustomerLogRepository) AddBulk(c context.Context, model *models.WebCustomer) (res *string, err error) {
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
