package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointMaxCustomerRepository ...
type IPointMaxCustomerRepository interface {
	SelectAll(c context.Context, parameter models.PointMaxCustomerParameter) ([]models.PointMaxCustomer, error)
	FindAll(ctx context.Context, parameter models.PointMaxCustomerParameter) ([]models.PointMaxCustomer, int, error)
	FindByID(c context.Context, parameter models.PointMaxCustomerParameter) (models.PointMaxCustomer, error)
	FindByCustomerCode(c context.Context, customerCode string) (models.PointMaxCustomer, error)
	Add(c context.Context, model []viewmodel.PointMaxCustomerVM) error
	Update(c context.Context, model viewmodel.PointMaxCustomerVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// PointMaxCustomerRepository ...
type PointMaxCustomerRepository struct {
	DB *sql.DB
}

// NewPointMaxCustomerRepository ...
func NewPointMaxCustomerRepository(DB *sql.DB) IPointMaxCustomerRepository {
	return &PointMaxCustomerRepository{DB: DB}
}

// Scan rows
func (repository PointMaxCustomerRepository) scanRows(rows *sql.Rows) (res models.PointMaxCustomer, err error) {
	err = rows.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.CustomerCode,
		&res.CustomerName,
		&res.MonthlyMaxPoint,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// Scan row
func (repository PointMaxCustomerRepository) scanRow(row *sql.Row) (res models.PointMaxCustomer, err error) {
	err = row.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.CustomerCode,
		&res.CustomerName,
		&res.MonthlyMaxPoint,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// SelectAll ...
func (repository PointMaxCustomerRepository) SelectAll(c context.Context, parameter models.PointMaxCustomerParameter) (data []models.PointMaxCustomer, err error) {
	var conditionString string

	if parameter.ShowAll != "1" {
		conditionString += `AND NOW() BETWEEN DEF.START_DATE AND DEF.END_DATE`
	}

	statement := models.PointMaxCustomerSelectStatement + models.PointMaxCustomerWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort

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

// FindAll ...
func (repository PointMaxCustomerRepository) FindAll(ctx context.Context, parameter models.PointMaxCustomerParameter) (data []models.PointMaxCustomer, count int, err error) {
	var conditionString string

	if parameter.ShowAll != "1" {
		conditionString += `AND NOW() BETWEEN DEF.START_DATE AND DEF.END_DATE`
	}

	statement := models.PointMaxCustomerSelectStatement + models.PointMaxCustomerWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.QueryContext(ctx, statement, parameter.Offset, parameter.Limit)
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

	countQuery := `SELECT COUNT(*) FROM POINT_MAX_CUSTOMER def ` + models.PointMaxCustomerWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository PointMaxCustomerRepository) FindByID(c context.Context, parameter models.PointMaxCustomerParameter) (data models.PointMaxCustomer, err error) {
	statement := models.PointMaxCustomerSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// FindByCustomerCode ...
func (repository PointMaxCustomerRepository) FindByCustomerCode(c context.Context, customerCode string) (data models.PointMaxCustomer, err error) {
	statement := models.PointMaxCustomerSelectStatement + ` WHERE DEF.CUSTOMER_CODE = '` + customerCode + `'`
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository PointMaxCustomerRepository) Add(c context.Context, in []viewmodel.PointMaxCustomerVM) (err error) {
	var statementInsert string
	for _, datum := range in {
		if statementInsert == "" {
			statementInsert += `('` + datum.StartDate + `', '` + datum.EndDate + `', '` + datum.CustomerCode + `', '` + datum.MonthlyMaxPoint + `', NOW(), NOW())`
		} else {
			statementInsert += `, ('` + datum.StartDate + `', '` + datum.EndDate + `', '` + datum.CustomerCode + `', '` + datum.MonthlyMaxPoint + `', NOW(), NOW())`
		}
	}
	statement := `INSERT INTO POINT_MAX_CUSTOMER (
			START_DATE, 
			END_DATE,
			CUSTOMER_CODE,
			MONTHLY_MAX_POINT,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ` + statementInsert

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

// Update ...
func (repository PointMaxCustomerRepository) Update(c context.Context, in viewmodel.PointMaxCustomerVM) (res string, err error) {
	statement := `UPDATE POINT_MAX_CUSTOMER SET 
		START_DATE = $1, 
		END_DATE = $2, 
		CUSTOMER_CODE = $3, 
		MONTHLY_MAX_POINT = $4,
		UPDATED_AT = now()
	WHERE id = $5
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.CustomerCode,
		in.MonthlyMaxPoint,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository PointMaxCustomerRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINT_MAX_CUSTOMER SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
