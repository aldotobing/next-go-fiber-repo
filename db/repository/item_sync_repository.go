package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IItemSyncRepository ...
type IItemSyncRepository interface {
	FindByID(c context.Context, parameter models.ItemSyncParameter) (models.ItemSync, error)
	FindByCode(c context.Context, parameter models.ItemSyncParameter) (models.ItemSync, error)
	Add(c context.Context, model *models.ItemSync) (*string, error)
	Edit(c context.Context, model *models.ItemSync) (*string, error)
}

// ItemSyncRepository ...
type ItemSyncRepository struct {
	DB *sql.DB
}

// NewItemSyncRepository ...
func NewItemSyncRepository(DB *sql.DB) IItemSyncRepository {
	return &ItemSyncRepository{DB: DB}
}

// Scan rows
func (repository ItemSyncRepository) scanRows(rows *sql.Rows) (res models.ItemSync, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemSyncRepository) scanRow(row *sql.Row) (res models.ItemSync, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository ItemSyncRepository) FindByCode(c context.Context, parameter models.ItemSyncParameter) (data models.ItemSync, err error) {
	statement := models.ItemSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(def.code) = $1`
	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.Code))

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemSyncRepository) FindByID(c context.Context, parameter models.ItemSyncParameter) (data models.ItemSync, err error) {
	statement := models.ItemSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ItemSyncRepository) Add(c context.Context, model *models.ItemSync) (res *string, err error) {
	fmt.Println("insert item data")
	statement := `INSERT INTO item (
		_name, code, item_picture, item_category_id,
		active, parent_id,have_variant,alias,
		description,keterangan,created_date,modified_date,
		url_video
	)
	VALUES (
		$1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11, $12,$13
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code), str.NullString(model.ItemPicture), str.NullString(model.ItemCategoryId),
		str.NullString(model.ItemActive), str.NullString(model.ItemParentID), str.NullString(model.HaveVariant), str.NullString(model.ItemAlias),
		str.NullString(model.Description), str.NullString(model.Keterangan), str.NullString(model.CreatedDate), str.NullString(model.ModifiedDate),
		str.NullString(model.UrlVideo),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ItemSyncRepository) Edit(c context.Context, model *models.ItemSync) (res *string, err error) {
	statement := `UPDATE item SET 
	_name = $1, code = $2, item_picture = $3, item_category_id = $4, 
	active = $5, parent_id = $6 ,have_variant = $7, alias = $8 ,
	description = $9, keterangan = $10 , modified_date = $11 ,
	url_video = $12
	
	WHERE code = $13 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		str.NullString(model.Name), str.NullString(model.Code), str.NullString(model.ItemPicture), str.NullString(model.ItemCategoryId),
		str.NullString(model.ItemActive), str.NullString(model.ItemParentID), str.NullString(model.HaveVariant), str.NullString(model.ItemAlias),
		str.NullString(model.Description), str.NullString(model.Keterangan), str.NullString(model.ModifiedDate), str.NullString(model.UrlVideo),
		model.Code,
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
