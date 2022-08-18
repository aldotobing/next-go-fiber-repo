package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IUserRepository ...
type IUserRepository interface {
	SelectAll(c context.Context, parameter models.UserParameter) ([]models.User, error)
	FindAll(ctx context.Context, parameter models.UserParameter) ([]models.User, int, error)
	FindByID(c context.Context, parameter models.UserParameter) (models.User, error)
	FindByCode(c context.Context, parameter models.UserParameter) (models.User, error)
	FindByEmail(c context.Context, parameter models.UserParameter) (models.User, error)
	FindByPhone(c context.Context, parameter models.UserParameter) (models.User, error)
}

// UserRepository ...
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository ...
func NewUserRepository(DB *sql.DB) IUserRepository {
	return &UserRepository{DB: DB}
}

// Scan rows
func (repository UserRepository) scanRows(rows *sql.Rows) (res models.User, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository UserRepository) scanRow(row *sql.Row) (res models.User, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository UserRepository) SelectAll(c context.Context, parameter models.UserParameter) (data []models.User, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}
	if str.IsValidUUID(parameter.RoleID) {
		conditionString += ` AND def.role_id = '` + parameter.RoleID + `'`
	}
	if str.IsValidUUID(parameter.ProfilePhotoID) {
		conditionString += ` AND def.profile_photo_id = '` + parameter.ProfilePhotoID + `'`
	}

	statement := models.UserSelectStatement + ` ` + models.UserWhereStatement +
		` AND (LOWER(def."email") LIKE $1 OR LOWER(def."name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, parameter.Search)
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
func (repository UserRepository) FindAll(ctx context.Context, parameter models.UserParameter) (data []models.User, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}
	if str.IsValidUUID(parameter.RoleID) {
		conditionString += ` AND def.role_id = '` + parameter.RoleID + `'`
	}
	if str.IsValidUUID(parameter.ProfilePhotoID) {
		conditionString += ` AND def.profile_photo_id = '` + parameter.ProfilePhotoID + `'`
	}

	query := models.UserSelectStatement + ` ` + models.UserWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."email") LIKE $1 OR LOWER(def."name") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "users" def ` + models.UserWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository UserRepository) FindByID(c context.Context, parameter models.UserParameter) (data models.User, err error) {
	statement := models.UserSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCode ...
func (repository UserRepository) FindByCode(c context.Context, parameter models.UserParameter) (data models.User, err error) {
	statement := models.UserSelectStatement + ` WHERE def.deleted_at IS NULL AND def.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByEmail ...
func (repository UserRepository) FindByEmail(c context.Context, parameter models.UserParameter) (data models.User, err error) {
	statement := models.UserSelectStatement + ` WHERE def.deleted_at IS NULL AND def.email = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Email)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByPhone ...
func (repository UserRepository) FindByPhone(c context.Context, parameter models.UserParameter) (data models.User, err error) {
	statement := models.UserSelectStatement + ` WHERE def.deleted_at IS NULL AND def.phone = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Phone)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
