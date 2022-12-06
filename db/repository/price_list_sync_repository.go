package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IPriceListSyncRepository ...
type IPriceListSyncRepository interface {
	FindByID(c context.Context, parameter models.PriceListSyncParameter) (models.PriceListSync, error)
	FindByCode(c context.Context, parameter models.PriceListSyncParameter) (models.PriceListSync, error)
	Add(c context.Context, model *models.PriceListSync) (*string, error)
	Edit(c context.Context, model *models.PriceListSync) (*string, error)
}

// PriceListSyncRepository ...
type PriceListSyncRepository struct {
	DB *sql.DB
}

// NewPriceListSyncRepository ...
func NewPriceListSyncRepository(DB *sql.DB) IPriceListSyncRepository {
	return &PriceListSyncRepository{DB: DB}
}

// Scan rows
func (repository PriceListSyncRepository) scanRows(rows *sql.Rows) (res models.PriceListSync, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name, &res.PriceListPrint, &res.PriceListBranchID, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository PriceListSyncRepository) scanRow(row *sql.Row) (res models.PriceListSync, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name, &res.PriceListPrint, &res.PriceListBranchID, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository PriceListSyncRepository) FindByCode(c context.Context, parameter models.PriceListSyncParameter) (data models.PriceListSync, err error) {
	statement := models.PriceListSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(def.code) = $1`

	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Code))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository PriceListSyncRepository) FindByID(c context.Context, parameter models.PriceListSyncParameter) (data models.PriceListSync, err error) {
	statement := models.PriceListSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository PriceListSyncRepository) Add(c context.Context, model *models.PriceListSync) (res *string, err error) {
	statement := `INSERT INTO price_list (
		_name, code ,created_date,modified_date,
		print_price_list, branch_id
	)
	VALUES (
		$1, $2, $3, $4, $5, $6
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code), str.NullString(model.CreatedDate), str.NullString(model.ModifiedDate),
		str.NullString(model.PriceListPrint), str.NullString(model.PriceListBranchID),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository PriceListSyncRepository) Edit(c context.Context, model *models.PriceListSync) (res *string, err error) {
	statement := `UPDATE price_list SET 
	_name = $1, code = $2,  modified_date = $3 ,
	print_price_list = $4, branch_id = $5
	
	WHERE id = $6 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code), str.NullString(model.ModifiedDate),
		str.NullString(model.PriceListPrint), str.NullString(model.PriceListBranchID),
		model.ID,
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
