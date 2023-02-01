package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerRepository ...
type IWebDoctorRepository interface {
	SelectAll(c context.Context, parameter models.WebDoctorParameter) ([]models.WebDoctor, error)
	FindAll(ctx context.Context, parameter models.WebDoctorParameter) ([]models.WebDoctor, int, error)
	FindByID(c context.Context, parameter models.WebDoctorParameter) (models.WebDoctor, error)
	Edit(c context.Context, model *models.WebDoctor) (*string, error)
	Add(c context.Context, model *models.WebDoctor) (*string, error)
}

// CustomerRepository ...
type WebDoctorRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewWebDoctorRepository(DB *sql.DB) IWebDoctorRepository {
	return &WebDoctorRepository{DB: DB}
}

// Scan rows
func (repository WebDoctorRepository) scanRows(rows *sql.Rows) (res models.WebDoctor, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.DoctorName,
		&res.DoctorAddress,
		&res.DoctorPhone,
		&res.DoctorUserID,
		&res.DoctorUserName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebDoctorRepository) scanRow(row *sql.Row) (res models.WebDoctor, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.DoctorName,
		&res.DoctorAddress,
		&res.DoctorPhone,
		&res.DoctorUserID,
		&res.DoctorUserName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebDoctorRepository) SelectAll(c context.Context, parameter models.WebDoctorParameter) (data []models.WebDoctor, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	statement := models.WebDoctorSelectStatement + ` ` + models.WebDoctorWhereStatement +
		` AND (LOWER(def."_name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository WebDoctorRepository) FindAll(ctx context.Context, parameter models.WebDoctorParameter) (data []models.WebDoctor, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	query := models.WebDoctorSelectStatement + ` ` + models.WebDoctorWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `select count(*)
	from partner def
	left join _user us on us.id = def.user_id  ` + models.WebDoctorWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebDoctorRepository) FindByID(c context.Context, parameter models.WebDoctorParameter) (data models.WebDoctor, err error) {
	statement := models.WebDoctorSelectStatement + ` WHERE def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	// fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebDoctorRepository) Edit(c context.Context, model *models.WebDoctor) (res *string, err error) {
	statement := `UPDATE partner SET 
	_name = $1, 
	address = $2, 
	user_id = $3,
	phone_no = $4 
	WHERE id = $5 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.DoctorName,
		model.DoctorAddress,
		model.DoctorUserID,
		model.DoctorPhone,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository WebDoctorRepository) Add(c context.Context, model *models.WebDoctor) (res *string, err error) {
	statement := `INSERT INTO partner (code, _name,address, phone_no, 
		company_id,device_id,created_date,modified_date)
	VALUES ($1, $2, $3, $4, 2,99,now(),now()) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.DoctorName,
		model.DoctorAddress, model.DoctorPhone,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
