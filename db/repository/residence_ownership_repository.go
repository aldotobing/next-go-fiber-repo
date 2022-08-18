package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IResidenceOwnershipRepository ...
type IResidenceOwnershipRepository interface {
	SelectAll(c context.Context, parameter models.ResidenceOwnershipParameter) ([]models.ResidenceOwnership, error)
	FindAll(ctx context.Context, parameter models.ResidenceOwnershipParameter) ([]models.ResidenceOwnership, int, error)
	FindByID(c context.Context, parameter models.ResidenceOwnershipParameter) (models.ResidenceOwnership, error)
	FindByMappingName(c context.Context, parameter models.ResidenceOwnershipParameter) (models.ResidenceOwnership, error)
	Add(c context.Context, model *models.ResidenceOwnership) (string, error)
	Edit(c context.Context, model *models.ResidenceOwnership) (string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
}

// ResidenceOwnershipRepository ...
type ResidenceOwnershipRepository struct {
	DB *sql.DB
}

// NewResidenceOwnershipRepository ...
func NewResidenceOwnershipRepository(DB *sql.DB) IResidenceOwnershipRepository {
	return &ResidenceOwnershipRepository{DB: DB}
}

// Scan rows
func (repository ResidenceOwnershipRepository) scanRows(rows *sql.Rows) (res models.ResidenceOwnership, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository ResidenceOwnershipRepository) scanRow(row *sql.Row) (res models.ResidenceOwnership, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.MappingName, &res.Status, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ResidenceOwnershipRepository) SelectAll(c context.Context, parameter models.ResidenceOwnershipParameter) (data []models.ResidenceOwnership, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	statement := models.ResidenceOwnershipSelectStatement + ` ` + models.ResidenceOwnershipWhereStatement +
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
func (repository ResidenceOwnershipRepository) FindAll(ctx context.Context, parameter models.ResidenceOwnershipParameter) (data []models.ResidenceOwnership, count int, err error) {
	conditionString := ``
	if parameter.Status != "" {
		conditionString += ` AND def.status = ` + str.StringToBoolString(parameter.Status)
	}

	query := models.ResidenceOwnershipSelectStatement + ` ` + models.ResidenceOwnershipWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "residence_ownerships" def ` + models.ResidenceOwnershipWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository ResidenceOwnershipRepository) FindByID(c context.Context, parameter models.ResidenceOwnershipParameter) (data models.ResidenceOwnership, err error) {
	statement := models.ResidenceOwnershipSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByMappingName ...
func (repository ResidenceOwnershipRepository) FindByMappingName(c context.Context, parameter models.ResidenceOwnershipParameter) (data models.ResidenceOwnership, err error) {
	statement := models.ResidenceOwnershipSelectStatement + ` WHERE def.deleted_at IS NULL AND def.mapping_name = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.MappingName)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ResidenceOwnershipRepository) Add(c context.Context, model *models.ResidenceOwnership) (res string, err error) {
	statement := `INSERT INTO residence_ownerships (name, mapping_name, status)
	VALUES ($1, $2, $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ResidenceOwnershipRepository) Edit(c context.Context, model *models.ResidenceOwnership) (res string, err error) {
	statement := `UPDATE residence_ownerships SET name = $1, mapping_name = $2, status = $3 WHERE id = $4 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Name, model.MappingName, model.Status, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository ResidenceOwnershipRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE residence_ownerships SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
