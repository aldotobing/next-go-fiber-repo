package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

type IUserAccountRepository interface {
	FindByID(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByPhoneNo(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FindByEmailAndPass(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
	FIreStoreIDSync(c context.Context, model *models.UserAccount) (*string, error)
	FCMUpdate(c context.Context, model *models.UserAccount) (*string, error)
	FindByLoginName(c context.Context, parameter models.UserAccountParameter) (models.UserAccount, error)
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
		&res.ID, &res.LoginCode, &res.RoleIDList, &res.FCMToken,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository UserAccountRepository) scanRow(row *sql.Row) (res models.UserAccount, err error) {
	err = row.Scan(
		&res.ID, &res.LoginCode, &res.RoleIDList, &res.FCMToken, &res.ShowInApp, &res.Active,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserAccountRepository) FindByPhoneNo(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.created_date is not null AND cus.customer_phone = $1 AND lower(cus.customer_code) = $2`
	row := repository.DB.QueryRowContext(c, statement, parameter.PhoneNo, strings.ToLower(parameter.Code))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByLoginName(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE def.created_date is not null AND lower(def.login) = $1 `
	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Code))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
func (repository UserAccountRepository) FindByID(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.UserAccountSelectStatement + ` WHERE  cus.id = $1`

	row := repository.DB.QueryRowContext(c, statement, parameter.CustomerID)
	data, err = repository.scanRow(row)
	if err != nil {

		return data, err
	}

	return data, nil
}

func (repository UserAccountRepository) FindByEmailAndPass(c context.Context, parameter models.UserAccountParameter) (data models.UserAccount, err error) {
	statement := models.AdminUserAccountSelectStatement + ` WHERE def.created_date is not null AND def._password = $1 AND lower(def.login) = $2`
	row := repository.DB.QueryRowContext(c, statement, parameter.Password, strings.ToLower(parameter.Email))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository UserAccountRepository) FCMUpdate(c context.Context, model *models.UserAccount) (res *string, err error) {
	statement := `UPDATE _user SET
	fcm_token = $1
	WHERE id = $2
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		model.FCMToken,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (repository UserAccountRepository) FIreStoreIDSync(c context.Context, model *models.UserAccount) (res *string, err error) {
	statement := `UPDATE _user SET
	firestoreuid = $1
	WHERE id = $2
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.FireStoreUID,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
