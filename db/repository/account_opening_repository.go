package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IAccountOpeningRepository ...
type IAccountOpeningRepository interface {
	SelectAll(c context.Context, parameter models.AccountOpeningParameter) ([]models.AccountOpening, error)
	FindAll(ctx context.Context, parameter models.AccountOpeningParameter) ([]models.AccountOpening, int, error)
	FindByID(c context.Context, parameter models.AccountOpeningParameter) (models.AccountOpening, error)
	FindByEmail(c context.Context, parameter models.AccountOpeningParameter) (models.AccountOpening, error)
	FindByPhone(c context.Context, parameter models.AccountOpeningParameter) (models.AccountOpening, error)
	Add(c context.Context, model *models.AccountOpening) (string, error)
	Edit(c context.Context, model *models.AccountOpening) (string, error)
	EditPhoneValidAt(c context.Context, model *models.AccountOpening) (string, error)
	EditEmailValidAt(c context.Context, model *models.AccountOpening) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// AccountOpeningRepository ...
type AccountOpeningRepository struct {
	DB *sql.DB
}

// NewAccountOpeningRepository ...
func NewAccountOpeningRepository(DB *sql.DB) IAccountOpeningRepository {
	return &AccountOpeningRepository{DB: DB}
}

// Scan rows
func (repository AccountOpeningRepository) scanRows(rows *sql.Rows) (res models.AccountOpening, err error) {
	err = rows.Scan(
		&res.ID, &res.UserID, &res.Email, &res.EmailValidAt, &res.Name, &res.MaritalStatusID, &res.GenderID,
		&res.BirthPlace, &res.BirthPlaceCityID, &res.BirthDate, &res.MotherName, &res.Phone, &res.PhoneValidAt,
		&res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository AccountOpeningRepository) scanRow(row *sql.Row) (res models.AccountOpening, err error) {
	err = row.Scan(
		&res.ID, &res.UserID, &res.Email, &res.EmailValidAt, &res.Name, &res.MaritalStatusID, &res.GenderID,
		&res.BirthPlace, &res.BirthPlaceCityID, &res.BirthDate, &res.MotherName, &res.Phone, &res.PhoneValidAt,
		&res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository AccountOpeningRepository) SelectAll(c context.Context, parameter models.AccountOpeningParameter) (data []models.AccountOpening, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}
	if str.IsValidUUID(parameter.UserID) {
		conditionString += ` AND def.user_id = '` + parameter.UserID + `'`
	}
	if str.IsValidUUID(parameter.GenderID) {
		conditionString += ` AND def.gender_id = '` + parameter.GenderID + `'`
	}
	if str.IsValidUUID(parameter.BirthPlaceCityID) {
		conditionString += ` AND def.birth_place_city_id = '` + parameter.BirthPlaceCityID + `'`
	}

	statement := models.AccountOpeningSelectStatement + ` ` + models.AccountOpeningWhereStatement +
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
func (repository AccountOpeningRepository) FindAll(ctx context.Context, parameter models.AccountOpeningParameter) (data []models.AccountOpening, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}
	if str.IsValidUUID(parameter.UserID) {
		conditionString += ` AND def.role_id = '` + parameter.UserID + `'`
	}
	if str.IsValidUUID(parameter.GenderID) {
		conditionString += ` AND def.gender_id = '` + parameter.GenderID + `'`
	}
	if str.IsValidUUID(parameter.BirthPlaceCityID) {
		conditionString += ` AND def.birth_place_city_id = '` + parameter.BirthPlaceCityID + `'`
	}

	query := models.AccountOpeningSelectStatement + ` ` + models.AccountOpeningWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "account_openings" def ` + models.AccountOpeningWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository AccountOpeningRepository) FindByID(c context.Context, parameter models.AccountOpeningParameter) (data models.AccountOpening, err error) {
	statement := models.AccountOpeningSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByEmail ...
func (repository AccountOpeningRepository) FindByEmail(c context.Context, parameter models.AccountOpeningParameter) (data models.AccountOpening, err error) {
	statement := models.AccountOpeningSelectStatement + ` WHERE def.deleted_at IS NULL AND def.email = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Email)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByPhone ...
func (repository AccountOpeningRepository) FindByPhone(c context.Context, parameter models.AccountOpeningParameter) (data models.AccountOpening, err error) {
	statement := models.AccountOpeningSelectStatement + ` WHERE def.deleted_at IS NULL AND def.phone = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Phone)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository AccountOpeningRepository) Add(c context.Context, model *models.AccountOpening) (res string, err error) {
	statement := `INSERT INTO account_openings (email, name, phone, status)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Email, model.Name, model.Phone, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository AccountOpeningRepository) Edit(c context.Context, model *models.AccountOpening) (res string, err error) {
	statement := `UPDATE account_openings SET name = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// EditPhoneValidAt ...
func (repository AccountOpeningRepository) EditPhoneValidAt(c context.Context, model *models.AccountOpening) (res string, err error) {
	statement := `UPDATE account_openings SET phone_valid_at = $1, status = $2 WHERE id = $3 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.PhoneValidAt, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// EditEmailValidAt ...
func (repository AccountOpeningRepository) EditEmailValidAt(c context.Context, model *models.AccountOpening) (res string, err error) {
	statement := `UPDATE AccountOpenings SET email_valid_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.EmailValidAt, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository AccountOpeningRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE AccountOpenings SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
