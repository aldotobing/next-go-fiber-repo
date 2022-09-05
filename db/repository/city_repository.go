package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICityRepository ...
type ICityRepository interface {
	SelectAll(c context.Context, parameter models.CityParameter) ([]models.City, error)
	FindAll(ctx context.Context, parameter models.CityParameter) ([]models.City, int, error)
	FindByID(c context.Context, parameter models.CityParameter) (models.City, error)
	Add(c context.Context, model *models.City) (*string, error)
	Edit(c context.Context, model *models.City) (*string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
	AddDataBreakDown(c context.Context, model *models.MpCityDataBreakDown) (*string, error)
}

// CityRepository ...
type CityRepository struct {
	DB *sql.DB
}

// NewCityRepository ...
func NewCityRepository(DB *sql.DB) ICityRepository {
	return &CityRepository{DB: DB}
}

// Scan rows
func (repository CityRepository) scanRows(rows *sql.Rows) (res models.City, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CityRepository) scanRow(row *sql.Row) (res models.City, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CityRepository) SelectAll(c context.Context, parameter models.CityParameter) (data []models.City, err error) {
	conditionString := ``

	if parameter.ProvinceID != "" {
		conditionString += ` AND def.id_province = '` + parameter.ProvinceID + `'`
	}

	statement := models.CitySelectStatement + ` ` + models.CityWhereStatement +
		` AND (LOWER(def."name_city") LIKE $1 OR LOWER(p."name_province") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository CityRepository) FindAll(ctx context.Context, parameter models.CityParameter) (data []models.City, count int, err error) {
	conditionString := ``

	// if parameter.ProvinceID != "" {
	// 	conditionString += ` AND def.id_province = '` + parameter.ProvinceID + `'`
	// }

	query := models.CitySelectStatement + ` ` + models.CityWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	fmt.Println(query)

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

	query = `SELECT COUNT(*) FROM "city" def ` + models.CityWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CityRepository) FindByID(c context.Context, parameter models.CityParameter) (data models.City, err error) {
	statement := models.CitySelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository CityRepository) Add(c context.Context, model *models.City) (res *string, err error) {
	statement := `INSERT INTO mp_city (name_city,id_province, long_city, lat_city,
		created_at_city, created_by_city)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_city`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.Name, model.Name, model.Name,
		model.Name, model.Name).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository CityRepository) Edit(c context.Context, model *models.City) (res *string, err error) {
	statement := `UPDATE mp_city SET 
	name_city = $1, id_province = $2, long_city = $3, lat_city = $4, 
	updated_at_city = $5, updated_by_city = $6 WHERE id_city = $7 RETURNING id_city`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.Name,
		model.Name, model.Name, model.Name, model.Name, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository CityRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE mp_city SET updated_at_city = $1, deleted_at_city = $2 WHERE id_city = $3 RETURNING id_city`

	err = repository.DB.QueryRowContext(c, statement, now, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository CityRepository) AddDataBreakDown(c context.Context, model *models.MpCityDataBreakDown) (res *string, err error) {
	statement := `INSERT INTO mp_city (name_city, id_province, old_id
		,created_at_city,updated_at_city,
		created_by_city,updated_by_city,lat_city,long_city)
	VALUES ($1, 
		( select id_province from mp_province mpprov where mpprov.old_id = $2),
		$3, now(), now(), 1, 1,(select cast($4 as double precision))
		,(select cast($5 as double precision))
	
	) RETURNING id_city`

	println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.ProvinceID,
		model.OldID, model.LatCity, model.LongCity).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
