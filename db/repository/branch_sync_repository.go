package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IBranchSyncRepository ...
type IBranchSyncRepository interface {
	FindByID(c context.Context, parameter models.BranchSyncParameter) (models.BranchSync, error)
	FindByCode(c context.Context, parameter models.BranchSyncParameter) (models.BranchSync, error)
	Add(c context.Context, model *models.BranchSync) (*string, error)
	Edit(c context.Context, model *models.BranchSync) (*string, error)
}

// BranchSyncRepository ...
type BranchSyncRepository struct {
	DB *sql.DB
}

// NewBranchSyncRepository ...
func NewBranchSyncRepository(DB *sql.DB) IBranchSyncRepository {
	return &BranchSyncRepository{DB: DB}
}

// Scan rows
func (repository BranchSyncRepository) scanRows(rows *sql.Rows) (res models.BranchSync, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository BranchSyncRepository) scanRow(row *sql.Row) (res models.BranchSync, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository BranchSyncRepository) FindByCode(c context.Context, parameter models.BranchSyncParameter) (data models.BranchSync, err error) {
	statement := models.BranchSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(def.branch_code) = $1`
	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Code))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository BranchSyncRepository) FindByID(c context.Context, parameter models.BranchSyncParameter) (data models.BranchSync, err error) {
	statement := models.BranchSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository BranchSyncRepository) Add(c context.Context, model *models.BranchSync) (res *string, err error) {
	statement := `INSERT INTO branch (
		_name, code, branch_picture, branch_category_id,
		active, parent_id,have_variant,alias,
		description,keterangan,created_date,modified_date,
		url_video
	)
	VALUES (
		$1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11, $12,$13
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository BranchSyncRepository) Edit(c context.Context, model *models.BranchSync) (res *string, err error) {
	statement := `UPDATE branch SET 
	_name = $1, code = $2, branch_picture = $3, branch_category_id = $4, 
	active = $5, parent_id = $6 ,have_variant = $7, alias = $8 ,
	description = $9, keterangan = $10 , modified_date = $11 ,
	url_video = $12
	
	WHERE id = $13 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code),
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
