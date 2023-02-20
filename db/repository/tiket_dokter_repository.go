package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ITicketDokter ...
type ITicketDokter interface {
	SelectAll(c context.Context, parameter models.TicketDokterParameter) ([]models.TicketDokter, error)
	FindAll(ctx context.Context, parameter models.TicketDokterParameter) ([]models.TicketDokter, int, error)
	FindByID(c context.Context, parameter models.TicketDokterParameter) (models.TicketDokter, error)
	Add(c context.Context, parameter *models.TicketDokter) (*string, error)
	//Delete(c context.Context, id string) (string, error)
	Edit(c context.Context, model *models.TicketDokter) (*string, error)
	// 	EditAddress(c context.Context, model *models.TicketDokter) (*string, error)
}

// TicketDokter ...
type TicketDokter struct {
	DB *sql.DB
}

// NewTicketDokter ...
func NewTicketDokterRepository(DB *sql.DB) ITicketDokter {
	return &TicketDokter{DB: DB}
}

// Scan rows
func (repository TicketDokter) scanRows(rows *sql.Rows) (res models.TicketDokter, err error) {
	err = rows.Scan(
		&res.ID,
		&res.DoctorID,
		&res.DoctorName,
		&res.TicketCode,
		&res.CustomerID,
		&res.CustomerName,
		&res.CustomerHeight,
		&res.CustomerWeight,
		&res.CustomerAge,
		&res.CustomerPhone,
		&res.CustomerAltPhone,
		&res.CustomerProblem,
		&res.Solution,
		&res.Allergy,
		&res.Status,
		&res.StatusReason,
		&res.CreatedDate,
		&res.ModifiedDate,
		&res.CloseDate,
		&res.Description,
		&res.Hide,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository TicketDokter) scanRow(row *sql.Row) (res models.TicketDokter, err error) {
	err = row.Scan(
		&res.ID,
		&res.DoctorID,
		&res.DoctorName,
		&res.TicketCode,
		&res.CustomerID,
		&res.CustomerName,
		&res.CustomerHeight,
		&res.CustomerWeight,
		&res.CustomerAge,
		&res.CustomerPhone,
		&res.CustomerAltPhone,
		&res.CustomerProblem,
		&res.Solution,
		&res.Allergy,
		&res.Status,
		&res.StatusReason,
		&res.CreatedDate,
		&res.ModifiedDate,
		&res.CloseDate,
		&res.Description,
		&res.Hide,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository TicketDokter) SelectAll(c context.Context, parameter models.TicketDokterParameter) (data []models.TicketDokter, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND TD.ID = ` + parameter.ID
	}

	if parameter.CustomerID != "" {
		conditionString += ` AND TD.CUSTOMER_ID = ` + parameter.CustomerID
	}

	if parameter.DoctorID != "" {
		conditionString += ` AND TD.DOCTOR_ID = ` + parameter.DoctorID
	}

	if parameter.Status != "" {
		conditionString += ` AND TD.STATUS = ` + parameter.Status
	}

	statement := models.TicketDokterSelectStatement + ` ` + models.TicketDokterWhereStatement +
		` AND (LOWER(td.customer_name) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print
	fmt.Println(statement)

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
func (repository TicketDokter) FindAll(ctx context.Context, parameter models.TicketDokterParameter) (data []models.TicketDokter, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND TD.ID = ` + parameter.ID
	}

	if parameter.CustomerID != "" {
		conditionString += ` AND TD.CUSTOMER_ID = ` + parameter.CustomerID
	}

	if parameter.DoctorID != "" {
		conditionString += ` AND TD.DOCTOR_ID = ` + parameter.DoctorID
	}

	if parameter.Status != "" {
		conditionString += ` AND TD.STATUS = ` + parameter.Status
	}

	query := models.TicketDokterSelectStatement + ` ` + models.TicketDokterWhereStatement + ` ` + conditionString + `
		AND (LOWER(td."customer_name") LIKE $1  )` + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	fmt.Println("Find All Ticket" + query)

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

	query = `SELECT COUNT(*) FROM "ticket_dokter" td ` + models.TicketDokterWhereStatement + ` ` +
		conditionString + ` AND (LOWER(td."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository TicketDokter) FindByID(c context.Context, parameter models.TicketDokterParameter) (data models.TicketDokter, err error) {
	statement := models.TicketDokterSelectStatement + ` WHERE td.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository TicketDokter) Add(c context.Context, model *models.TicketDokter) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO ticket_dokter (
		customer_id, 
		customer_name, 
		height, 
		weight,
		age, 
		phone, 
		phone_alt, 
		problem, 
		solution, 
		allergy, 
		status, 
		created_date, 
		modified_date, 
		close_date, 
		doctor_id, 
		doctor_name, 
		description, 
		hide)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now(), now(), $12, $13, $14, $15, false) RETURNING id`

	err = transaction.QueryRowContext(c, statement,
		model.CustomerID,
		model.CustomerName,
		str.NullOrEmtyString(model.CustomerHeight),
		str.NullOrEmtyString(model.CustomerWeight),
		model.CustomerAge,
		str.NullOrEmtyString(model.CustomerPhone),
		str.NullOrEmtyString(model.CustomerAltPhone),
		str.NullOrEmtyString(model.CustomerProblem),
		str.NullOrEmtyString(model.Solution),
		model.Allergy,
		str.NullOrEmtyString(model.Status),
		str.NullOrEmtyString(model.CloseDate),
		model.DoctorID,
		model.DoctorName,
		model.Description).Scan(&res)

	//fmt.Println("TIKET DOKTER INSERT : " + statement)

	if err = transaction.Commit(); err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository TicketDokter) Edit(c context.Context, model *models.TicketDokter) (res *string, err error) {
	statement := `UPDATE ticket_dokter SET 
	status = $1, 
	solution = $2,
	close_date = now()
	WHERE id = $3 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.Status,
		model.Solution,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
