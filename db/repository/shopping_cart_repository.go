package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IShoppingCartRepository ...
type IShoppingCartRepository interface {
	SelectAll(c context.Context, parameter models.ShoppingCartParameter) ([]models.ShoppingCart, error)
	FindAll(ctx context.Context, parameter models.ShoppingCartParameter) ([]models.ShoppingCart, int, error)
	FindByID(c context.Context, parameter models.ShoppingCartParameter) (models.ShoppingCart, error)
	Add(c context.Context, model *models.ShoppingCart) (*string, error)
	Edit(c context.Context, model *models.ShoppingCart) (*string, error)
	Delete(c context.Context, id string) (*string, error)
}

// ShoppingCartRepository ...
type ShoppingCartRepository struct {
	DB *sql.DB
}

// NewShoppingCartRepository ...
func NewShoppingCartRepository(DB *sql.DB) IShoppingCartRepository {
	return &ShoppingCartRepository{DB: DB}
}

// Scan rows
func (repository ShoppingCartRepository) scanRows(rows *sql.Rows) (res models.ShoppingCart, err error) {
	err = rows.Scan(
		&res.ID, &res.CustomerID, &res.CustomerName, &res.ItemID, &res.ItemName, &res.UomID, &res.UomName,
		&res.Qty, &res.StockQty, &res.Price,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ShoppingCartRepository) scanRow(row *sql.Row) (res models.ShoppingCart, err error) {
	err = row.Scan(
		&res.ID, &res.CustomerID, &res.CustomerName, &res.ItemID, &res.ItemName, &res.UomID, &res.UomName,
		&res.Qty, &res.StockQty, &res.Price,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ShoppingCartRepository) SelectAll(c context.Context, parameter models.ShoppingCartParameter) (data []models.ShoppingCart, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` and def.customer_id = ` + parameter.CustomerID
	}

	statement := models.ShoppingCartSelectStatement + ` ` + models.ShoppingCartWhereStatement +
		` AND (LOWER(it."_name") LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository ShoppingCartRepository) FindAll(ctx context.Context, parameter models.ShoppingCartParameter) (data []models.ShoppingCart, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` and def.customer_id = ` + parameter.CustomerID
	}

	query := models.ShoppingCartSelectStatement + ` ` + models.ShoppingCartWhereStatement + ` ` + conditionString + `
		AND (LOWER(it."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "ShoppingCart" def ` + models.ShoppingCartWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository ShoppingCartRepository) FindByID(c context.Context, parameter models.ShoppingCartParameter) (data models.ShoppingCart, err error) {
	statement := models.ShoppingCartSelectStatement + ` WHERE def.deleted_at_ShoppingCart IS NULL AND def.id_ShoppingCart = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository ShoppingCartRepository) Add(c context.Context, model *models.ShoppingCart) (res *string, err error) {
	statement := `INSERT INTO cart (customer_id,item_id, uom_id, price,
		created_date, created_by, qty , stock_qty )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.CustomerID, model.ItemID, model.UomID, model.Price,
		model.CreatedAt, model.CreatedBy, model.Qty, model.StockQty).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ShoppingCartRepository) Edit(c context.Context, model *models.ShoppingCart) (res *string, err error) {
	statement := `UPDATE cart SET 
	item_id = $1, uom_id = $2, price = $3, modified_date = $4, 
	modified_by = $5, qty = $6 ,stock_qty = $7  WHERE id = $8 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.ItemID, model.UomID,
		model.Price, model.ModifiedAt, model.ModifiedBy, model.Qty, model.StockQty, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository ShoppingCartRepository) Delete(c context.Context, id string) (res *string, err error) {
	statement := ` delete from  cart WHERE id = $1 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
