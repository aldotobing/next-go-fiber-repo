package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IItemOldPriceRepository ...
type IItemOldPriceRepository interface {
	SelectAll(c context.Context, parameter models.ItemOldPriceParameter) ([]models.ItemOldPrice, error)
	FindAll(ctx context.Context, parameter models.ItemOldPriceParameter) ([]models.ItemOldPrice, int, error)
	FindByID(c context.Context, parameter models.ItemOldPriceParameter) (models.ItemOldPrice, error)
	Add(c context.Context, in []viewmodel.ItemOldPriceVM) error
	Update(c context.Context, in viewmodel.ItemOldPriceVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// ItemOldPriceRepository ...
type ItemOldPriceRepository struct {
	DB *sql.DB
}

// NewItemOldPriceRepository ...
func NewItemOldPriceRepository(DB *sql.DB) IItemOldPriceRepository {
	return &ItemOldPriceRepository{DB: DB}
}

// Scan rows
func (repository ItemOldPriceRepository) scanRows(rows *sql.Rows) (res models.ItemOldPrice, err error) {
	err = rows.Scan(
		&res.ID,
		&res.CustomerID,
		&res.CustomerCode,
		&res.CustomerName,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemPicture,
		&res.PriceListID,
		&res.SellPrice,
		&res.Quantity,
		&res.UomID,
		&res.UomName,
		&res.PreservedQty,
		&res.InvoiceQty,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// Scan row
func (repository ItemOldPriceRepository) scanRow(row *sql.Row) (res models.ItemOldPrice, err error) {
	err = row.Scan(
		&res.ID,
		&res.CustomerID,
		&res.CustomerCode,
		&res.CustomerName,
		&res.ItemID,
		&res.ItemCode,
		&res.ItemName,
		&res.ItemPicture,
		&res.PriceListID,
		&res.SellPrice,
		&res.Quantity,
		&res.UomID,
		&res.UomName,
		&res.PreservedQty,
		&res.InvoiceQty,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// SelectAll ...
func (repository ItemOldPriceRepository) SelectAll(c context.Context, parameter models.ItemOldPriceParameter) (data []models.ItemOldPrice, err error) {
	conditionString := ``

	conditionString += ` AND NOW() BETWEEN DEF.START_DATE AND DEF.END_DATE`

	if parameter.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = '` + parameter.CustomerID + `'`
	}

	query := models.ItemOldPriceSelectStatement + models.ItemOldPriceWhereStatement + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return
}

// FindAll ...
func (repository ItemOldPriceRepository) FindAll(ctx context.Context, parameter models.ItemOldPriceParameter) (data []models.ItemOldPrice, count int, err error) {
	var conditionString string

	conditionString += ` AND NOW() BETWEEN DEF.START_DATE AND DEF.END_DATE`

	if parameter.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = '` + parameter.CustomerID + `'`
	}

	query := models.ItemOldPriceSelectStatement + models.ItemOldPriceWhereStatement + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.Query(query, parameter.Offset, parameter.Limit)
	fmt.Println(query)
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

	query = `SELECT COUNT(*) 
		FROM old_item_price def
		LEFT JOIN ITEM I ON I.ID = DEF.ITEM_ID
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
		LEFT JOIN UOM U ON U.ID = DEF.UOM_ID ` +
		models.ItemOldPriceWhereStatement + conditionString
	err = repository.DB.QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByID ...
func (repository ItemOldPriceRepository) FindByID(c context.Context, parameter models.ItemOldPriceParameter) (data models.ItemOldPrice, err error) {
	query := models.ItemOldPriceSelectStatement + models.ItemOldPriceWhereStatement + ` AND DEF.ID = '` + parameter.ID + `'`

	row := repository.DB.QueryRowContext(c, query)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemOldPriceRepository) Add(c context.Context, in []viewmodel.ItemOldPriceVM) (err error) {
	var values string

	for _, datum := range in {
		if values == "" {
			values += ` ('` + datum.CustomerID + `', '` + datum.ItemID + `', '` + datum.PriceListID + `', 
			'` + datum.SellPrice + `', '` + strconv.Itoa(datum.Quantity) + `', '` + datum.PreservedQty + `', 
			'` + datum.InvoiceQty + `', '` + datum.StartDate + `', '` + datum.EndDate + `', '` + datum.UomID + `', now(), now())`
		} else {
			values += `, ('` + datum.CustomerID + `', '` + datum.ItemID + `', '` + datum.PriceListID + `', 
			'` + datum.SellPrice + `', '` + strconv.Itoa(datum.Quantity) + `', '` + datum.PreservedQty + `', 
			'` + datum.InvoiceQty + `', '` + datum.StartDate + `', '` + datum.EndDate + `', '` + datum.UomID + `', now(), now())`
		}
	}

	statement := `INSERT INTO item_old_price (customer_id, item_id, price_list_id,
		sell_price, qty, preserved_qty, 
		invoiced_qty, start_date, end_date, uom_id, created_at, updated_at) VALUES` + values

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

// Update ...
func (repository ItemOldPriceRepository) Update(c context.Context, in viewmodel.ItemOldPriceVM) (res string, err error) {
	statement := `UPDATE ITEM_OLD_PRICE SET 
		START_DATE = $1, 
		END_DATE = $2, 
		QTY = $3,
		UPDATED_AT = now()
	WHERE id = $4
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.Quantity,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository ItemOldPriceRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE ITEM_OLD_PRICE SET 
		DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
