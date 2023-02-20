package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerRepository ...
type IUserCheckinActivityRepository interface {
	SelectAll(c context.Context, parameter models.UserCheckinActivityParameter) ([]models.UserCheckinActivity, error)
	FindAll(ctx context.Context, parameter models.UserCheckinActivityParameter) ([]models.UserCheckinActivity, int, error)
	FindByID(c context.Context, parameter models.UserCheckinActivityParameter) (models.UserCheckinActivity, error)
	Add(c context.Context, model *models.UserCheckinActivity) (*string, error)
	FindActiveCheckin(c context.Context, parameter models.UserCheckinActivityParameter) (models.UserCheckinActivity, error)
}

// CustomerRepository ...
type UserCheckinActivityRepository struct {
	DB *sql.DB
}

// NewCustomerRepository ...
func NewUserCheckinActivityRepository(DB *sql.DB) IUserCheckinActivityRepository {
	return &UserCheckinActivityRepository{DB: DB}
}

// Scan rows
func (repository UserCheckinActivityRepository) scanRows(rows *sql.Rows) (res models.UserCheckinActivity, err error) {
	err = rows.Scan(
		&res.ID,
		&res.UserID,
		&res.CheckinTime,
		&res.Login,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository UserCheckinActivityRepository) scanRow(row *sql.Row) (res models.UserCheckinActivity, err error) {
	err = row.Scan(
		&res.ID,
		&res.UserID,
		&res.CheckinTime,
		&res.Login,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository UserCheckinActivityRepository) SelectAll(c context.Context, parameter models.UserCheckinActivityParameter) (data []models.UserCheckinActivity, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	statement := models.UserCheckinActivitySelectStatement + ` ` + models.UserCheckinActivityWhereStatement +
		` AND (LOWER(usr.login) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print

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
func (repository UserCheckinActivityRepository) FindAll(ctx context.Context, parameter models.UserCheckinActivityParameter) (data []models.UserCheckinActivity, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND def.id = '` + parameter.ID + `'`
	}

	query := models.UserCheckinActivitySelectStatement + ` ` + models.UserCheckinActivityWhereStatement + ` ` + conditionString + `
		AND (LOWER(usr.login) LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`

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
		from user_checkin_activity def
		left join _user usr on usr.id = def.user_id   ` + models.UserCheckinActivityWhereStatement + ` ` +
		conditionString + ` AND (LOWER(usr.login) LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository UserCheckinActivityRepository) FindByID(c context.Context, parameter models.UserCheckinActivityParameter) (data models.UserCheckinActivity, err error) {
	statement := models.UserCheckinActivitySelectStatement + ` WHERE def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByID ...
func (repository UserCheckinActivityRepository) FindActiveCheckin(c context.Context, parameter models.UserCheckinActivityParameter) (data models.UserCheckinActivity, err error) {
	statement := models.UserCheckinActivitySelectStatement + ` WHERE def.user_id = $1 and now()::date = def.checkin_time::date  order by def.id desc limit 1 `
	row := repository.DB.QueryRowContext(c, statement, parameter.UserId)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository UserCheckinActivityRepository) Add(c context.Context, model *models.UserCheckinActivity) (res *string, err error) {
	statement := `INSERT INTO user_checkin_activity (user_id, checkin_time,
		created_date,modified_date)
	VALUES ($1,now(),now(),now()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.UserID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
