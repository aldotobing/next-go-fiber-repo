package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebUserRepository ...
type IWebUserRepository interface {
	SelectAll(c context.Context, parameter models.WebUserParameter) ([]models.WebUser, error)
	FindAll(ctx context.Context, parameter models.WebUserParameter) ([]models.WebUser, int, error)
	FindByID(c context.Context, parameter models.WebUserParameter) (models.WebUser, error)
	Add(c context.Context, model *models.WebUser) (*string, error)
	Edit(c context.Context, model *models.WebUser) (*string, error)
	Delete(c context.Context, id string, now time.Time) (*string, error)
}

// WebUserRepository ...
type WebUserRepository struct {
	DB *sql.DB
}

// NewWebUserRepository ...
func NewWebUserRepository(DB *sql.DB) IWebUserRepository {
	return &WebUserRepository{DB: DB}
}

// Scan rows
func (repository WebUserRepository) scanRows(rows *sql.Rows) (res models.WebUser, err error) {
	err = rows.Scan(
		&res.ID, &res.Login, &res.Password, &res.CompanyId,
		&res.Active, &res.FirestoreUID, &res.FcmToken,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository WebUserRepository) scanRow(row *sql.Row) (res models.WebUser, err error) {
	err = row.Scan(
		&res.ID, &res.Login, &res.Password, &res.CompanyId,
		&res.Active, &res.FirestoreUID, &res.FcmToken,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebUserRepository) SelectAll(c context.Context, parameter models.WebUserParameter) (data []models.WebUser, err error) {
	conditionString := ``

	statement := models.WebUserSelectStatement + ` ` + models.WebUserWhereStatement +
		` AND (LOWER(def."login") LIKE $1  ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

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
func (repository WebUserRepository) FindAll(ctx context.Context, parameter models.WebUserParameter) (data []models.WebUser, count int, err error) {
	conditionString := ``

	query := models.WebUserSelectStatement + ` ` + models.WebUserWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."login") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM _user def ` + models.WebUserWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."login") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebUserRepository) FindByID(c context.Context, parameter models.WebUserParameter) (data models.WebUser, err error) {
	statement := models.WebUserSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository WebUserRepository) Add(c context.Context, model *models.WebUser) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO _user (login,_password, created_date, modified_date ,company_id, active)
	VALUES ($1, $2, now(),now(),2, 1) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Login, model.Password).Scan(&res)

	if err != nil {
		return res, err
	}

	UserID := &res

	parts := strings.Split(*model.UserRoleGroupIDList, ",")
	if len(parts) >= 1 {
		for pi, _ := range parts {
			linestatement := `INSERT INTO user_role_group (user_id,role_group_id,created_date, modified_date)
					VALUES ($1,$2, now(),now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, UserID, parts[pi]).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Edit ...
func (repository WebUserRepository) Edit(c context.Context, model *models.WebUser) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `UPDATE _user SET 
	login = $1,_password = $2, modified_date = now()
	 WHERE id = $3 RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Login, model.Password,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	UserID := &res

	parts := strings.Split(*model.UserRoleGroupIDList, ",")

	if len(parts) >= 1 {
		deletelinestatement := `delete from user_role_group WHERE user_id = $1`

		deletedRow, _ := transaction.QueryContext(c, deletelinestatement, UserID)
		deletedRow.Close()

		for pi, _ := range parts {
			linestatement := `INSERT INTO user_role_group (user_id,role_group_id,created_date, modified_date)
						VALUES ($1,$2, now(),now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, UserID, parts[pi]).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (repository WebUserRepository) Delete(c context.Context, id string, now time.Time) (res *string, err error) {
	deletelinestatement := `delete from _user WHERE id = $1`

	deletedRow, _ := repository.DB.QueryContext(c, deletelinestatement, id)
	deletedRow.Close()

	res = &id
	return res, err
}
