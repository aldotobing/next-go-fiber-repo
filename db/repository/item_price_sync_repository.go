package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IItemPriceSyncRepository ...
type IItemPriceSyncRepository interface {
	FindByID(c context.Context, parameter models.ItemPriceSyncParameter) (models.ItemPriceSync, error)
	FindItemPrice(c context.Context, parameter models.ItemPriceSyncParameter) (models.ItemPriceSync, error)
	Add(c context.Context, model *models.ItemPriceSync) (*string, error)
	Edit(c context.Context, model *models.ItemPriceSync) (*string, error)
}

// ItemPriceSyncRepository ...
type ItemPriceSyncRepository struct {
	DB *sql.DB
}

// NewItemPriceSyncRepository ...
func NewItemPriceSyncRepository(DB *sql.DB) IItemPriceSyncRepository {
	return &ItemPriceSyncRepository{DB: DB}
}

// Scan rows
func (repository ItemPriceSyncRepository) scanRows(rows *sql.Rows) (res models.ItemPriceSync, err error) {
	err = rows.Scan(
		&res.PriceListId, &res.PriceListCode, &res.ID, &res.PriceListVersionID, &res.PriceListVersionCode, &res.ItemId, &res.UomId, &res.Price, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemPriceSyncRepository) scanRow(row *sql.Row) (res models.ItemPriceSync, err error) {
	err = row.Scan(
		&res.PriceListId, &res.PriceListCode, &res.ID, &res.PriceListVersionID, &res.PriceListVersionCode, &res.ItemId, &res.UomId, &res.Price, &res.CreatedDate, &res.ModifiedDate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ItemPriceSyncRepository) FindItemPrice(c context.Context, parameter models.ItemPriceSyncParameter) (data models.ItemPriceSync, err error) {
	statement := models.ItemPriceSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND lower(plv.description) = $1 and lower(pl.code) = $2 and def.item_id = $3 and def.uom_id = $4`

	row := repository.DB.QueryRowContext(c, statement, strings.ToLower(parameter.PriceListVersionCode), strings.ToLower(parameter.PriceListCode), parameter.ItemId, parameter.UomId)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemPriceSyncRepository) FindByID(c context.Context, parameter models.ItemPriceSyncParameter) (data models.ItemPriceSync, err error) {
	statement := models.ItemPriceSyncSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ItemPriceSyncRepository) Add(c context.Context, model *models.ItemPriceSync) (res *string, err error) {
	statement := `INSERT INTO item_price (
		price_list_version_id, item_id ,uom_id,price,
		created_date,modified_date
		)
	VALUES (
		(
			select plvs.id from price_list_version plvs where lower(plvs.description) = $1 and plvs.price_list_id = (select pls.id from price_list pls where lower(pls.code) = $2)
		),
		$3, $4, $5, $6, $7
		) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		strings.ToLower(*str.NullString(model.PriceListVersionCode)), strings.ToLower(*str.NullString(model.PriceListCode)),
		str.NullString(model.ItemId), str.NullString(model.UomId), str.NullString(model.Price),
		str.NullString(model.CreatedDate),
		str.NullString(model.ModifiedDate),
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ItemPriceSyncRepository) Edit(c context.Context, model *models.ItemPriceSync) (res *string, err error) {
	statement := `UPDATE item_price SET 
	price_list_version_id = 
	(
		select plvs.id from price_list_version plvs where lower(plvs.description) = $1 and plvs.price_list_id = (select pls.id from price_list pls where lower(pls.code) = $2)
	), 
	item_id = $3, uom_id=$4, price = $5,
	 modified_date = $6
	WHERE id = $7 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		strings.ToLower(*str.NullString(model.PriceListVersionCode)), strings.ToLower(*str.NullString(model.PriceListCode)),
		str.NullString(model.ItemId), str.NullString(model.UomId), str.NullString(model.Price),
		str.NullString(model.ModifiedDate),
		model.ID,
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
