package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IBroadcastRepository ...
type IBroadcastRepository interface {
	SelectAll(c context.Context, parameter models.BroadcastParameter) ([]models.Broadcast, error)
	FindAll(ctx context.Context, parameter models.BroadcastParameter) ([]models.Broadcast, int, error)
	FindByID(c context.Context, parameter models.BroadcastParameter) (models.Broadcast, error)
	Add(c context.Context, model viewmodel.BroadcastVM) (string, error)
	Update(c context.Context, model viewmodel.BroadcastVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// BroadcastRepository ...
type BroadcastRepository struct {
	DB *sql.DB
}

// NewBroadcastRepository ...
func NewBroadcastRepository(DB *sql.DB) IBroadcastRepository {
	return &BroadcastRepository{DB: DB}
}

// Scan rows
func (repository BroadcastRepository) scanRows(rows *sql.Rows) (res models.Broadcast, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Title,
		&res.Body,
		&res.BroadcastDate,
		&res.BroadcastTime,
		&res.RepeatEveryDay,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Parameter,
	)

	return
}

// Scan row
func (repository BroadcastRepository) scanRow(row *sql.Row) (res models.Broadcast, err error) {
	err = row.Scan(
		&res.ID,
		&res.Title,
		&res.Body,
		&res.BroadcastDate,
		&res.BroadcastTime,
		&res.RepeatEveryDay,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Parameter,
	)

	return
}

// SelectAll ...
func (repository BroadcastRepository) SelectAll(c context.Context, parameter models.BroadcastParameter) (data []models.Broadcast, err error) {
	var conditionString string

	statement := models.BroadcastSelectStatement + models.BroadcastWhereStatement +
		` AND (LOWER(b."title") LIKE '%` + parameter.Search + `%') ` + conditionString +
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

// FindAll ...
func (repository BroadcastRepository) FindAll(ctx context.Context, parameter models.BroadcastParameter) (data []models.Broadcast, count int, err error) {
	var conditionString string

	statement := models.BroadcastSelectStatement + models.BroadcastWhereStatement +
		` AND (LOWER(b."title") LIKE '%` + parameter.Search + `%') ` + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.QueryContext(ctx, statement, parameter.Offset, parameter.Limit)
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

	countQuery := `SELECT COUNT(*) FROM BROADCAST B ` + models.BroadcastWhereStatement +
		` AND (LOWER(b."title") LIKE '%` + parameter.Search + `%') ` + conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository BroadcastRepository) FindByID(c context.Context, parameter models.BroadcastParameter) (data models.Broadcast, err error) {
	statement := models.BroadcastSelectStatement + ` WHERE B.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository BroadcastRepository) Add(c context.Context, in viewmodel.BroadcastVM) (res string, err error) {
	statement := `INSERT INTO BROADCAST (
			TITLE, 
			BODY,
			BROADCAST_DATE,
			BROADCAST_TIME,
			REPEAT_EVERY_DAY,
			CREATED_AT,
			UPDATED_AT,
			PARAMETER
		)
	VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), $6) RETURNING id`
	param, _ := json.Marshal(in.Parameter)
	err = repository.DB.QueryRowContext(c, statement,
		in.Title,
		in.Body,
		in.BroadcastDate,
		in.BroadcastTime,
		in.RepeatEveryDay,
		string(param),
	).Scan(&res)

	return
}

// Update ...
func (repository BroadcastRepository) Update(c context.Context, in viewmodel.BroadcastVM) (res string, err error) {
	statement := `UPDATE BROADCAST SET 
		TITLE = $1, 
		BODY = $2, 
		BROADCAST_DATE = $3, 
		BROADCAST_TIME = $4,
		REPEAT_EVERY_DAY = $5,
		UPDATED_AT = now(),
		PARAMETER = $6
	WHERE id = $7
	RETURNING id`
	param, _ := json.Marshal(in.Parameter)
	err = repository.DB.QueryRowContext(c, statement,
		in.Title,
		in.Body,
		in.BroadcastDate,
		in.BroadcastTime,
		in.RepeatEveryDay,
		string(param),
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository BroadcastRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE BROADCAST SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
