package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerRepository ...
type IWebPartnerRepository interface {
	SelectAll(c context.Context, parameter models.WebPartnerParameter) ([]models.WebPartner, error)
	FindAll(ctx context.Context, parameter models.WebPartnerParameter) ([]models.WebPartner, int, error)
	FindByID(c context.Context, parameter models.WebPartnerParameter) (models.WebPartner, error)
	Edit(c context.Context, model *models.WebPartner) (*string, error)
	Add(c context.Context, model *models.WebPartner) (*string, error)
}

// CustomerRepository ...
type WebPartnerRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewWebPartnerRepository(DB *sql.DB) IWebPartnerRepository {
	return &WebPartnerRepository{DB: DB}
}

// Scan rows
func (repository WebPartnerRepository) scanRows(rows *sql.Rows) (res models.WebPartner, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.PartnerName,
		&res.PartnerAddress,
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
func (repository WebPartnerRepository) scanRow(row *sql.Row) (res models.WebPartner, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.PartnerName,
		&res.PartnerAddress,
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
func (repository WebPartnerRepository) SelectAll(c context.Context, parameter models.WebPartnerParameter) (data []models.WebPartner, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	statement := models.WebPartnerSelectStatement + ` ` + models.WebPartnerWhereStatement +
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
func (repository WebPartnerRepository) FindAll(ctx context.Context, parameter models.WebPartnerParameter) (data []models.WebPartner, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	query := models.WebPartnerSelectStatement + ` ` + models.WebPartnerWhereStatement + ` ` + conditionString + `
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
	left join _user us on us.id = def.user_id  ` + models.WebPartnerWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebPartnerRepository) FindByID(c context.Context, parameter models.WebPartnerParameter) (data models.WebPartner, err error) {
	statement := models.WebPartnerSelectStatement + ` WHERE def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	// fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebPartnerRepository) Edit(c context.Context, model *models.WebPartner) (res *string, err error) {
	statement := `UPDATE partner SET 
	_name = $1, 
	address = $2, 
	user_id = $3,
	phone_no = $4 
	WHERE id = $5 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.PartnerName,
		model.PartnerAddress,
		model.PartnerUserID,
		model.PartnerPhone,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Add ...
func (repository WebPartnerRepository) Add(c context.Context, model *models.WebPartner) (res *string, err error) {
	statement := `INSERT INTO partner (code, _name,address, phone_no, 
		company_id,device_id,created_date,modified_date)
	VALUES ($1, $2, $3, $4, 2,99,now(),now()) RETURNING id`

	fmt.Println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.PartnerName,
		model.PartnerAddress, model.PartnerPhone,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
