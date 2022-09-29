package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IPriceListVersionSyncRepository ...
type IPriceListVersionSyncRepository interface {
	FindByID(c context.Context, parameter models.PriceListVersionSyncParameter) (models.PriceListVersionSync, error)
	FindByDescription(c context.Context, parameter models.PriceListVersionSyncParameter) (models.PriceListVersionSync, error)
	FindByDescriptionAndPricelistID(c context.Context, parameter models.PriceListVersionSyncParameter) (models.PriceListVersionSync, error)
	Add(c context.Context, model *models.PriceListVersionSync) (*string, error)
	Edit(c context.Context, model *models.PriceListVersionSync) (*string, error)
}

// PriceListVersionSyncRepository ...
type PriceListVersionSyncRepository struct {
	DB *sql.DB
}

// NewPriceListVersionSyncRepository ...
func NewPriceListVersionSyncRepository(DB *sql.DB) IPriceListVersionSyncRepository {
	return &PriceListVersionSyncRepository{DB: DB}
}

// Scan rows
func (repository PriceListVersionSyncRepository) scanRows(rows *sql.Rows) (res models.PriceListVersionSync, err error) {
	err = rows.Scan(
		&res.ID, &res.PriceListID, &res.StartDate, &res.EndDate, &res.Description, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository PriceListVersionSyncRepository) scanRow(row *sql.Row) (res models.PriceListVersionSync, err error) {
	err = row.Scan(
		&res.ID, &res.PriceListID, &res.StartDate, &res.EndDate, &res.Description, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByDescription ...
func (repository PriceListVersionSyncRepository) FindByDescription(c context.Context, parameter models.PriceListVersionSyncParameter) (data models.PriceListVersionSync, err error) {
	statement := models.PriceListVersionSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(def.description) = $1`

	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Description))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository PriceListVersionSyncRepository) FindByDescriptionAndPricelistID(c context.Context, parameter models.PriceListVersionSyncParameter) (data models.PriceListVersionSync, err error) {
	statement := models.PriceListVersionSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(def.description) = $1 and def.price_list_id = 
	(select pls.id from price_list pls where lower(pls.code) = $2 )
	`

	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Description), strings.ToLower(parameter.PriceListCode))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository PriceListVersionSyncRepository) FindByID(c context.Context, parameter models.PriceListVersionSyncParameter) (data models.PriceListVersionSync, err error) {
	statement := models.PriceListVersionSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository PriceListVersionSyncRepository) Add(c context.Context, model *models.PriceListVersionSync) (res *string, err error) {
	statement := `INSERT INTO price_list_version (
		price_list_id, start_date ,end_date,description,
		created_date,modified_date
		)
	VALUES (
		(
			select pls.id from price_list pls where lower(pls.code) = $1 
		), 
		$2, $3, $4, $5, $6
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		strings.ToLower(*str.NullString(model.PriceListCode)), str.NullString(model.StartDate), str.NullString(model.EndDate), str.NullString(model.Description),
		str.NullString(model.CreatedDate),
		str.NullString(model.ModifiedDate),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository PriceListVersionSyncRepository) Edit(c context.Context, model *models.PriceListVersionSync) (res *string, err error) {
	statement := `UPDATE price_list_version SET 
	price_list_id = 
	(
		select pls.id from price_list pls where lower(pls.code) = $1 
		), 
	start_date = $2, end_date=$3, description = $4,
	 modified_date = $5
	WHERE id = $6 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		strings.ToLower(*str.NullString(model.PriceListCode)), str.NullString(model.StartDate), str.NullString(model.EndDate), str.NullString(model.Description),
		str.NullString(model.ModifiedDate),
		model.ID,
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
