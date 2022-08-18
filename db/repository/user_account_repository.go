package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

type IUserAccountRepository interface {
	FindByID(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByRefferalCode(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByEmail(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByEmailAndPassword(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByPhoneNo(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
}

type UserAccountRepository struct {
	DB *sql.DB
}

func NewUserAccountRepository(DB *sql.DB) IUserAccountRepository {
	return &UserAccountRepository{DB: DB}
}

// Scan rows
func (repository UserAccountRepository) scanRows(rows *sql.Rows) (res models.UserAccount, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Code,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository UserAccountRepository) scanRow(row *sql.Row) (res models.UserAccount, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Code,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserAccountRepository) FindByRefferalCode(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.referral_code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ReferalCode)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByEmail(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.email_user = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Email)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByEmailAndPassword(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.email_user = $1 and def.password = $2 `
	row := repository.DB.QueryRowContext(c, statement, parameter.Email, parameter.Password)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByPhoneNo(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.created_date is not null AND p.phone_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.PhoneNo)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByID(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.id_user = $1`

	row := repository.DB.QueryRowContext(c, statement, parameter.ID)
	data, err = repository.scanRow(row)
	if err != nil {

		return data, err
	}

	return data, nil
}
