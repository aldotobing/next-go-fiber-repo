package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	EditQuantity(c context.Context, model *models.ShoppingCart) (*string, error)
	Delete(c context.Context, id string) (*string, error)
	SelectAllForGroup(c context.Context, parameter models.ShoppingCartParameter) ([]models.GroupedShoppingCart, error)
	GetTotal(c context.Context, parameter models.ShoppingCartParameter) (models.ShoppingCheckouAble, error)
	SelectAllBonus(c context.Context, parameter models.ShoppingCartParameter) ([]models.ShoppingCartItemBonus, error)
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
		&res.Qty, &res.StockQty, &res.Price, &res.ItemCategoryID, &res.ItemCategoryName, &res.ItemPicture,
		&res.TotalPrice, &res.OldPriceID,
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
		&res.Qty, &res.StockQty, &res.Price, &res.ItemCategoryID, &res.ItemCategoryName, &res.ItemPicture,
		&res.TotalPrice, &res.OldPriceID,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository ShoppingCartRepository) scanIsAbleRow(row *sql.Row) (res models.ShoppingCheckouAble, err error) {
	err = row.Scan(
		&res.IsAble, &res.MinOmzet, &res.IsMinOrder, &res.MinOrder,
	)

	return
}

// Scan rows
func (repository ShoppingCartRepository) scanGroupedRows(rows *sql.Rows) (res models.GroupedShoppingCart, err error) {
	err = rows.Scan(
		&res.CategoryID, &res.CategoryName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository ShoppingCartRepository) scanGroupedRow(row *sql.Row) (res models.GroupedShoppingCart, err error) {
	err = row.Scan(
		&res.CategoryID, &res.CategoryName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ShoppingCartRepository) scanBonusRows(rows *sql.Rows) (res models.ShoppingCartItemBonus, err error) {
	err = rows.Scan(
		&res.ItemID, &res.ItemName, &res.ItemCode, &res.Qty, &res.UomName, &res.ItemPicture,
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

	if parameter.ItemCategoryID != "" {
		conditionString += ` and it.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ListID != "" {
		conditionString += ` AND def.id in (` + parameter.ListID + `)`
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
		created_date, created_by, qty , stock_qty,total_price, old_price )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	if *model.OldPriceID == "" {
		err = repository.DB.QueryRowContext(c, statement, model.CustomerID, model.ItemID, model.UomID, model.Price,
			model.CreatedAt, model.CreatedBy, model.Qty, model.StockQty, model.TotalPrice, nil).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement, model.CustomerID, model.ItemID, model.UomID, model.Price,
			model.CreatedAt, model.CreatedBy, model.Qty, model.StockQty, model.TotalPrice, model.OldPriceID).Scan(&res)
	}

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository ShoppingCartRepository) Edit(c context.Context, model *models.ShoppingCart) (res *string, err error) {
	statement := `UPDATE cart SET 
	item_id = $1, uom_id = $2, price = $3, modified_date = $4, 
	modified_by = $5, qty = $6 ,stock_qty = $7 ,total_price = $8 WHERE id = $9 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.ItemID, model.UomID,
		model.Price, model.ModifiedAt, model.ModifiedBy, model.Qty, model.StockQty, model.TotalPrice, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// EditQuantity ...
func (repository ShoppingCartRepository) EditQuantity(c context.Context, model *models.ShoppingCart) (res *string, err error) {
	statement := `UPDATE cart SET 
	qty = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.Qty, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// EditByCartID ...
func (repository ShoppingCartRepository) EditByCartID(c context.Context, cartID string, model *models.ShoppingCart) (res *string, err error) {
	statement := `UPDATE cart SET 
	item_id = $1, uom_id = $2, price = $3, modified_date = $4, 
	modified_by = $5, qty = $6 ,stock_qty = $7 ,total_price = $8 WHERE id = $9 RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.ItemID, model.UomID,
		model.Price, model.ModifiedAt, model.ModifiedBy, model.Qty, model.StockQty, model.TotalPrice, cartID).Scan(&res)
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

func (repository ShoppingCartRepository) SelectAllForGroup(c context.Context, parameter models.ShoppingCartParameter) (data []models.GroupedShoppingCart, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` and def.customer_id = ` + parameter.CustomerID
	}

	statement := models.GroupedShoppingCartSelectStatement + ` ` + models.ShoppingCartWhereStatement +
		` AND (LOWER(it."_name") LIKE $1 ) ` + conditionString + ` group by it.item_category_id,ic._name  `
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanGroupedRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

// FindByID ...
func (repository ShoppingCartRepository) GetTotal(c context.Context, parameter models.ShoppingCartParameter) (data models.ShoppingCheckouAble, err error) {
	statement := ` 
	select 
		((case when ( (select sum(price*qty) from cart where id in(select  unnest(ARRAY(select string_to_array($1,',')))::integer))<( select coalesce(min_omzet_amount,0) from branch where id = (select branch_id from customer where id = $2) )) then 0 else 1 end)) as total_amount,
		(select coalesce(min_omzet_amount,0) from branch where id = (select branch_id from customer where id = $2) )::integer as min_amount,
		(case when((select coalesce(sum(price*qty),0) from cart where id in(select  unnest(ARRAY(select string_to_array($1,',')))::integer)) < 
		coalesce(
		(	select def.min_order from customer_type_branch_min_omzet def where def.branch_id = (select branch_id from customer where id = $2) and def.min_order>0 and def.customer_type_id = (select c1.customer_type_id from customer c1 where id = $2) 
			limit 1
		),( coalesce( (select min_omzet_amount from branch where id= (select branch_id from customer where id = $2)  ),100000)))
		
		) then 0 else 1 end),
		
		coalesce(
			(	select def.min_order from customer_type_branch_min_omzet def where def.branch_id = (select branch_id from customer where id = $2) and def.min_order>0 and def.customer_type_id = (select c1.customer_type_id from customer c1 where id = $2) 
				limit 1
			),( coalesce( (select min_omzet_amount from branch where id= (select branch_id from customer where id = $2)  ),100000)))
			
	`
	row := repository.DB.QueryRowContext(c, statement, parameter.ListLine, parameter.CustomerID)
	data, err = repository.scanIsAbleRow(row)

	return
}

// SelectAll ...
func (repository ShoppingCartRepository) SelectAllBonus(c context.Context, parameter models.ShoppingCartParameter) (data []models.ShoppingCartItemBonus, err error) {
	statement := models.ShoppingCartBonusSelectStatement

	fmt.Println(statement, parameter.ListID)

	rows, err := repository.DB.QueryContext(c, statement, parameter.ListID)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanBonusRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}
