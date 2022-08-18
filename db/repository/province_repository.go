package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IProvinceRepository ...
type IProvinceRepository interface {
	SelectAll(c context.Context, parameter models.ProvinceParameter) ([]models.Province, error)
	FindAll(ctx context.Context, parameter models.ProvinceParameter) ([]models.Province, int, error)
	FindByID(c context.Context, parameter models.ProvinceParameter) (models.Province, error)
	FindByCode(c context.Context, parameter models.ProvinceParameter) (models.Province, error)
	Add(c context.Context, model *models.Province) (string, error)
	Edit(c context.Context, model *models.Province) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
	AddDataBreakDown(c context.Context, model *models.MpProvinceDataBreakDown) (*string, error)
}

// ProvinceRepository ...
type ProvinceRepository struct {
	DB *sql.DB
}

// NewProvinceRepository ...
func NewProvinceRepository(DB *sql.DB) IProvinceRepository {
	return &ProvinceRepository{DB: DB}
}

// Scan rows
func (repository ProvinceRepository) scanRows(rows *sql.Rows) (res models.Province, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository ProvinceRepository) scanRow(row *sql.Row) (res models.Province, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ProvinceRepository) SelectAll(c context.Context, parameter models.ProvinceParameter) (data []models.Province, err error) {
	conditionString := ``

	statement := models.ProvinceSelectStatement + ` ` + models.ProvinceWhereStatement +
		` AND (LOWER(def."name_province") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository ProvinceRepository) FindAll(ctx context.Context, parameter models.ProvinceParameter) (data []models.Province, count int, err error) {
	conditionString := ``

	query := models.ProvinceSelectStatement + ` ` + models.ProvinceWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."name_province") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "mp_province" def ` + models.ProvinceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name_province") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository ProvinceRepository) FindByID(c context.Context, parameter models.ProvinceParameter) (data models.Province, err error) {
	statement := models.ProvinceSelectStatement + ` WHERE def.deleted_at_province IS NULL AND def.id_province = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCode ...
func (repository ProvinceRepository) FindByCode(c context.Context, parameter models.ProvinceParameter) (data models.Province, err error) {
	statement := models.ProvinceSelectStatement + ` WHERE def.deleted_at IS NULL AND def.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ProvinceRepository) Add(c context.Context, model *models.Province) (res string, err error) {
	statement := `INSERT INTO mp_province (code_province, name_province, id_nation,created_at_province,created_by_province)
	VALUES ($1, $2, $3, $4,$5) RETURNING id_province`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name, model.IdNation, model.CreatedAt, model.CreatedBy).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ProvinceRepository) Edit(c context.Context, model *models.Province) (res string, err error) {
	statement := `UPDATE mp_province SET code_province = $1, name_province = $2, updated_at_province =$3, updated_by_province= $4 
	WHERE id_province = $5 RETURNING id_province`

	err = repository.DB.QueryRowContext(c, statement, model.Code, model.Name, model.UpdatedAt, model.UpdatedBy, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository ProvinceRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE mp_province SET updated_at_province = $1, deleted_at_province = $2 WHERE id_province = $3 RETURNING id_province`

	err = repository.DB.QueryRowContext(c, statement, now, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository ProvinceRepository) AddDataBreakDown(c context.Context, model *models.MpProvinceDataBreakDown) (res *string, err error) {
	statement := `INSERT INTO mp_province (name_province, code_province, old_id
		,created_at_province,updated_at_province,
		created_by_province,updated_by_province,id_nation)
	VALUES ($1, $2, $3, now(), now(), 1, 1,
	( select id_nation from mp_nation mn where mn.old_id = $4)
	) RETURNING id_province`

	println(statement)

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.Code, model.OldID, model.NationID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
