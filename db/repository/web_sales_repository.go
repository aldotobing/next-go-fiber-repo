package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ICustomerRepository ...
type IWebSalesmanRepository interface {
	SelectAll(c context.Context, parameter models.WebSalesmanParameter) ([]models.WebSalesman, error)
	FindAll(ctx context.Context, parameter models.WebSalesmanParameter) ([]models.WebSalesman, int, error)
	FindByID(c context.Context, parameter models.WebSalesmanParameter) (models.WebSalesman, error)
	Edit(c context.Context, model *models.WebSalesman) (*string, error)
	Add(c context.Context, model *models.WebSalesman) (*string, error)
}

// CustomerRepository ...
type WebSalesmanRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewWebSalesmanRepository(DB *sql.DB) IWebSalesmanRepository {
	return &WebSalesmanRepository{DB: DB}
}

// Scan rows
func (repository WebSalesmanRepository) scanRows(rows *sql.Rows) (res models.WebSalesman, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.PartnerName,
		&res.PartnerPhone,
		&res.PartnerUserID,
		&res.PartnerUserName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebSalesmanRepository) scanRow(row *sql.Row) (res models.WebSalesman, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.PartnerName,
		&res.PartnerPhone,
		&res.PartnerUserID,
		&res.PartnerUserName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebSalesmanRepository) SelectAll(c context.Context, parameter models.WebSalesmanParameter) (data []models.WebSalesman, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	statement := models.WebSalesmanSelectStatement + ` ` + models.WebSalesmanWhereStatement +
		` AND (LOWER(def."salesman_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebSalesmanRepository) FindAll(ctx context.Context, parameter models.WebSalesmanParameter) (data []models.WebSalesman, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	query := models.WebSalesmanSelectStatement + ` ` + models.WebSalesmanWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."salesman_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	fmt.Println(query)
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

	query = `select count(*)
	from salesman def
	left join _user us on us.id = def.user_id  ` + models.WebSalesmanWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."salesman_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebSalesmanRepository) FindByID(c context.Context, parameter models.WebSalesmanParameter) (data models.WebSalesman, err error) {
	statement := models.WebSalesmanSelectStatement + ` WHERE def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	// fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebSalesmanRepository) Edit(c context.Context, model *models.WebSalesman) (res *string, err error) {
	fmt.Println("user id nya", *model.PartnerUserID)
	statement := `UPDATE salesman SET 
	salesman_name = $1, 
	user_id = $2,
	salesman_phone_no = $3
	WHERE id = $4 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.PartnerName,
		str.NullOrEmtyString(model.PartnerUserID),
		model.PartnerPhone,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository WebSalesmanRepository) Add(c context.Context, model *models.WebSalesman) (res *string, err error) {
	statement := `INSERT INTO salesman (salesman_code, salesman_name, salesman_phone_no, 
		created_date,modified_date,user_id, location_id)
	VALUES ($1, $2, $3, now(), now(), $4, 2) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.PartnerName,
		model.PartnerPhone, str.NullOrEmtyString(model.PartnerUserID),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
