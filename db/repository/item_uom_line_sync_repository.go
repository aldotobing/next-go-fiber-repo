package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IItemUomLineSyncRepository ...
type IItemUomLineSyncRepository interface {
	FindByID(c context.Context, parameter models.ItemUomLineSyncParameter) (models.ItemUomLineSync, error)
	FindByItemmAndUomCode(c context.Context, parameter models.ItemUomLineSyncParameter) (models.ItemUomLineSync, error)
	Add(c context.Context, model *models.ItemUomLineSync) (*string, error)
	Edit(c context.Context, model *models.ItemUomLineSync) (*string, error)
}

// ItemUomLineSyncRepository ...
type ItemUomLineSyncRepository struct {
	DB *sql.DB
}

// NewItemUomLineSyncRepository ...
func NewItemUomLineSyncRepository(DB *sql.DB) IItemUomLineSyncRepository {
	return &ItemUomLineSyncRepository{DB: DB}
}

// Scan rows
func (repository ItemUomLineSyncRepository) scanRows(rows *sql.Rows) (res models.ItemUomLineSync, err error) {
	err = rows.Scan(
		&res.ID, &res.ItemCode, &res.ItemName,
		&res.UomCode, &res.UomName, &res.UomConversion,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemUomLineSyncRepository) scanRow(row *sql.Row) (res models.ItemUomLineSync, err error) {
	err = row.Scan(
		&res.ID, &res.ItemCode, &res.ItemName,
		&res.UomCode, &res.UomName, &res.UomConversion,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository ItemUomLineSyncRepository) FindByItemmAndUomCode(c context.Context, parameter models.ItemUomLineSyncParameter) (data models.ItemUomLineSync, err error) {
	statement := models.ItemUomLineSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(i.code) = $1  AND lower(uo.code) = $2 `
	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.ItemCode), strings.ToLower(parameter.UomCode))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemUomLineSyncRepository) FindByID(c context.Context, parameter models.ItemUomLineSyncParameter) (data models.ItemUomLineSync, err error) {
	statement := models.ItemUomLineSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ItemUomLineSyncRepository) Add(c context.Context, model *models.ItemUomLineSync) (res *string, err error) {
	fmt.Println("insert item data")
	statement := `INSERT INTO item_uom_line (
		item_id, uom_id, conversion
	)
	VALUES (
		(select id from item where code = $1), (select id from uom where code = $2), $3
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.ItemCode), str.NullString(model.UomCode), str.NullString(model.UomConversion),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ItemUomLineSyncRepository) Edit(c context.Context, model *models.ItemUomLineSync) (res *string, err error) {
	statement := `UPDATE item_uom_line SET 
	conversion = $1, modified_date = now()	
	WHERE id = $2`
	// item_id = (select id from item where code = $2) and uom_id = (select id from uom where code = $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.UomConversion), str.NullString(model.ID),
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
