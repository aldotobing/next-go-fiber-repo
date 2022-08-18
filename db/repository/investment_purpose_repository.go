package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IInvestmentPurposeRepository ...
type IInvestmentPurposeRepository interface {
	SelectAll(c context.Context, parameter models.InvestmentPurposeParameter) ([]models.InvestmentPurpose, error)
	FindAll(ctx context.Context, parameter models.InvestmentPurposeParameter) ([]models.InvestmentPurpose, int, error)
	FindByID(c context.Context, parameter models.InvestmentPurposeParameter) (models.InvestmentPurpose, error)
	FindByMappingName(c context.Context, parameter models.InvestmentPurposeParameter) (models.InvestmentPurpose, error)
	Add(c context.Context, model *models.InvestmentPurpose) (string, error)
	Edit(c context.Context, model *models.InvestmentPurpose) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// InvestmentPurposeRepository ...
type InvestmentPurposeRepository struct {
	DB *sql.DB
}

// NewInvestmentPurposeRepository ...
func NewInvestmentPurposeRepository(DB *sql.DB) IInvestmentPurposeRepository {
	return &InvestmentPurposeRepository{DB: DB}
}

// Scan rows
func (repository InvestmentPurposeRepository) scanRows(rows *sql.Rows) (res models.InvestmentPurpose, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository InvestmentPurposeRepository) scanRow(row *sql.Row) (res models.InvestmentPurpose, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository InvestmentPurposeRepository) SelectAll(c context.Context, parameter models.InvestmentPurposeParameter) (data []models.InvestmentPurpose, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.InvestmentPurposeSelectStatement + ` ` + models.InvestmentPurposeWhereStatement +
		` AND (LOWER(def."name") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository InvestmentPurposeRepository) FindAll(ctx context.Context, parameter models.InvestmentPurposeParameter) (data []models.InvestmentPurpose, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.InvestmentPurposeSelectStatement + ` ` + models.InvestmentPurposeWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."name") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "investment_purposes" def ` + models.InvestmentPurposeWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository InvestmentPurposeRepository) FindByID(c context.Context, parameter models.InvestmentPurposeParameter) (data models.InvestmentPurpose, err error) {
	statement := models.InvestmentPurposeSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository InvestmentPurposeRepository) FindByMappingName(c context.Context, parameter models.InvestmentPurposeParameter) (data models.InvestmentPurpose, err error) {
	statement := models.InvestmentPurposeSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository InvestmentPurposeRepository) Add(c context.Context, model *models.InvestmentPurpose) (res string, err error) {
	statement := `INSERT INTO investment_purposes (name, mapping_name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository InvestmentPurposeRepository) Edit(c context.Context, model *models.InvestmentPurpose) (res string, err error) {
	statement := `UPDATE investment_purposes SET name = $1, mapping_name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository InvestmentPurposeRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE investment_purposes SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
