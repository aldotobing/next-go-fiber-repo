package repository

import (
	"context"
	"database/sql"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IDoctorRepository ...
type IDoctorRepository interface {
	FindByCodeAndPhone(c context.Context, parameter models.DoctorParameter) (models.Doctor, error)
}

// DoctorRepository ...
type DoctorRepository struct {
	DB *sql.DB
}

// NewDoctorRepository ...
func NewDoctorRepository(DB *sql.DB) IDoctorRepository {
	return &DoctorRepository{DB: DB}
}

// Scan rows
func (repository DoctorRepository) scanRows(rows *sql.Rows) (res models.Doctor, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.DoctorName,
		&res.DoctorAddress,
		&res.DoctorPhone,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository DoctorRepository) scanRow(row *sql.Row) (res models.Doctor, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.DoctorName,
		&res.DoctorAddress,
		&res.DoctorPhone,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository DoctorRepository) FindByCodeAndPhone(c context.Context, parameter models.DoctorParameter) (data models.Doctor, err error) {
	statement := models.DoctorSelectStatement + ` WHERE def.code = $1 and def.phone_no = $2 `
	row := repository.DB.QueryRowContext(c, statement, parameter.Code, parameter.Phone)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
