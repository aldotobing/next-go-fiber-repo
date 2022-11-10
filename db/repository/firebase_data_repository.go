package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IFireStoreUserRepository ...
type IFireStoreUserRepository interface {
	SelectAll(c context.Context, parameter models.FireStoreUserParameter) ([]models.FireStoreUser, error)
}

// FireStoreUserRepository ...
type FireStoreUserRepository struct {
	DB *sql.DB
}

// NewFireStoreUserRepository ...
func NewFireStoreUserRepository(DB *sql.DB) IFireStoreUserRepository {
	return &FireStoreUserRepository{DB: DB}
}

// Scan rows
func (repository FireStoreUserRepository) scanRows(rows *sql.Rows) (res models.FireStoreUser, err error) {
	err = rows.Scan(
		&res.UID,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository FireStoreUserRepository) scanRow(row *sql.Row) (res models.FireStoreUser, err error) {
	err = row.Scan(
		&res.UID,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository FireStoreUserRepository) SelectAll(c context.Context, parameter models.FireStoreUserParameter) (data []models.FireStoreUser, err error) {

	statement := models.FireStoreUserSelectStatement + ` ` + models.FireStoreUserWhereStatement +
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
