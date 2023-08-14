package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IItemRepository ...
type IItemRepository interface {
	SelectAll(c context.Context, parameter models.ItemParameter) ([]models.Item, error)
	FindAll(ctx context.Context, parameter models.ItemParameter) ([]models.Item, int, error)
	FindByID(c context.Context, parameter models.ItemParameter) (models.Item, error)
	SelectAllV2(c context.Context, parameter models.ItemParameter) (data []models.ItemV2, err error)
	// Add(c context.Context, model *models.Item) (*string, error)
	// Edit(c context.Context, model *models.Item) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// ItemRepository ...
type ItemRepository struct {
	DB *sql.DB
}

// NewItemRepository ...
func NewItemRepository(DB *sql.DB) IItemRepository {
	return &ItemRepository{DB: DB}
}

// Scan rows
func (repository ItemRepository) scanRows(rows *sql.Rows) (res models.Item, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionId,
		&res.ItemPicture,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository ItemRepository) scanRow(row *sql.Row) (res models.Item, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.Description,
		&res.ItemCategoryId,
		&res.ItemCategoryName,
		&res.UomID,
		&res.UomName,
		&res.UomLineConversion,
		&res.ItemPrice,
		&res.PriceListVersionId,
		&res.ItemPicture,
	)

	fmt.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository ItemRepository) SelectAll(c context.Context, parameter models.ItemParameter) (data []models.Item, err error) {
	conditionString := ``

	var args []interface{}
	var index int = 1 // Starting from 1 since PostgreSQL uses 1-indexed placeholders.

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = $` + strconv.Itoa(index)
		args = append(args, parameter.ID)
		index++
	}

	if parameter.ItemCategoryId != "" {
		if parameter.ItemCategoryId == "2" {
			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower(_name) LIKE $` + strconv.Itoa(index) + `) `
			args = append(args, "%tac%")
			index++
		} else {
			conditionString += ` AND def.item_category_id = $` + strconv.Itoa(index)
			args = append(args, parameter.ItemCategoryId)
			index++
		}
	}

	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id NOT IN (SELECT item_id FROM item_exception) `
	}

	if parameter.UomID != "" {
		conditionString += ` AND IUL.UOM_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.UomID)
		index++
	}

	if parameter.ExceptId != "" {
		conditionString += ` AND def.id <> $` + strconv.Itoa(index)
		args = append(args, parameter.ExceptId)
		index++
	}

	args = append([]interface{}{parameter.PriceListVersionId}, "%"+strings.ToLower(parameter.Search)+"%")

	statement := models.ItemSelectStatement + ` ` + models.ItemWhereStatement +
		` AND (LOWER(def."_name") LIKE $` + strconv.Itoa(index+1) + `) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, args...)

	fmt.Println(statement) // consider logging this instead of printing, or removing it entirely after testing

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

// func (repository ItemRepository) SelectAll(c context.Context, parameter models.ItemParameter) (data []models.Item, err error) {
// 	conditionString := ``

// 	if parameter.ID != "" {
// 		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
// 	}

// 	if parameter.ItemCategoryId != "" {
// 		if parameter.ItemCategoryId == "2" {
// 			//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
// 			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
// 		} else {
// 			conditionString += ` AND def.item_category_id = ` + parameter.ItemCategoryId + ``
// 		}
// 	}

// 	/*
// 		customerType 7 = Apotek Lokal
// 		customerType 15 = MT LOKAL INDEPENDEN
// 		defId 83 = TOLAK ANGIN CAIR /D5
// 		Tampilkan TAC D5 hanya pada kedua customerType di atas
// 	*/
// 	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
// 		conditionString += ` AND def.id NOT IN (SELECT item_id FROM item_exception) `
// 	}

// 	if parameter.UomID != "" {
// 		conditionString += ` AND IUL.UOM_ID = '` + parameter.UomID + `'`
// 	}

// 	if parameter.ExceptId != "" {
// 		conditionString += ` AND def.id <> '` + parameter.ExceptId + `'`
// 	}

// 	statement := models.ItemSelectStatement + ` ` + models.ItemWhereStatement +
// 		` AND (LOWER(def."_name") LIKE $2) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
// 	rows, err := repository.DB.QueryContext(c, statement, parameter.PriceListVersionId, "%"+strings.ToLower(parameter.Search)+"%")

// 	fmt.Println(statement)

// 	if err != nil {
// 		return data, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {

// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, temp)
// 	}

// 	return data, err
// }

// FindAll ...
func (repository ItemRepository) FindAll(ctx context.Context, parameter models.ItemParameter) (data []models.Item, count int, err error) {
	conditionString := ``

	var args []interface{}
	var index int = 1 // Starting from 1 since PostgreSQL uses 1-indexed placeholders.

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = $` + strconv.Itoa(index)
		args = append(args, parameter.ID)
		index++
	}

	if parameter.ItemCategoryId != "" {
		if parameter.ItemCategoryId == "2" {
			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower(_name) LIKE $` + strconv.Itoa(index) + `) `
			args = append(args, "%tac%")
			index++
		} else {
			conditionString += ` AND def.item_category_id = $` + strconv.Itoa(index)
			args = append(args, parameter.ItemCategoryId)
			index++
		}
	}

	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND def.id NOT IN (SELECT item_id FROM item_exception) `
	}

	if parameter.PriceListVersionId != "" {
		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.PriceListVersionId)
		index++
	}

	if parameter.UomID != "" {
		conditionString += ` AND IUL.UOM_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.UomID)
		index++
	}

	if parameter.ExceptId != "" {
		conditionString += ` AND def.id <> $` + strconv.Itoa(index)
		args = append(args, parameter.ExceptId)
		index++
	}

	args = append(args, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)

	query := models.ItemSelectStatement + ` ` + models.ItemWhereStatement + ` ` + conditionString +
		` AND (LOWER(def."_name") LIKE $` + strconv.Itoa(index) + `) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $` + strconv.Itoa(index+1) + ` LIMIT $` + strconv.Itoa(index+2)
	rows, err := repository.DB.QueryContext(ctx, query, args...)

	fmt.Println(query) // consider logging this instead of printing, or removing it entirely after testing

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

	// Now the count query
	queryCount := `SELECT COUNT(*) FROM "item" DEF ` +
		`JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID ` +
		`LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID ` +
		`LEFT JOIN UOM UOM ON UOM.ID = IUL.UOM_ID ` +
		`JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID AND IP.ITEM_ID = IUL.ITEM_ID` +
		models.ItemWhereStatement + ` ` + conditionString + ` AND (LOWER(def."_name") LIKE $` + strconv.Itoa(index) + `)`
	err = repository.DB.QueryRowContext(ctx, queryCount, args[:index]...).Scan(&count)
	return data, count, err
}

// func (repository ItemRepository) FindAll(ctx context.Context, parameter models.ItemParameter) (data []models.Item, count int, err error) {
// 	conditionString := ``

// 	if parameter.ID != "" {
// 		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
// 	}

// 	if parameter.ItemCategoryId != "" {
// 		if parameter.ItemCategoryId == "2" {
// 			//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
// 			conditionString += ` AND def.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
// 		} else {
// 			conditionString += ` AND def.item_category_id = ` + parameter.ItemCategoryId + ``
// 		}
// 	}

// 	/*
// 		customerType 7 = Apotek Lokal
// 		customerType 15 = MT LOKAL INDEPENDEN
// 		defId 83 = TOLAK ANGIN CAIR /D5
// 		Tampilkan TAC D5 hanya pada kedua customerType di atas
// 	*/
// 	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
// 		conditionString += ` AND def.id NOT IN (SELECT item_id FROM item_exception) `
// 	}

// 	if parameter.PriceListVersionId != "" {
// 		conditionString += ` AND IP.PRICE_LIST_VERSION_ID = '` + parameter.PriceListVersionId + `'`
// 	}

// 	if parameter.UomID != "" {
// 		conditionString += ` AND IUL.UOM_ID = '` + parameter.UomID + `'`
// 	}

// 	if parameter.ExceptId != "" {
// 		conditionString += ` AND def.id <> '` + parameter.ExceptId + `'`
// 	}

// 	query := models.ItemSelectStatement + ` ` + models.ItemWhereStatement + ` ` + conditionString + `
// 		AND (LOWER(def."_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
// 	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
// 	if err != nil {
// 		return data, count, err
// 	}

// 	fmt.Println(query)

// 	defer rows.Close()
// 	for rows.Next() {
// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return data, count, err
// 		}
// 		data = append(data, temp)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return data, count, err
// 	}

// 	query = `SELECT COUNT(*) FROM "item" DEF ` +
// 		`JOIN ITEM_UOM_LINE IUL ON IUL.ITEM_ID = DEF.ID ` +
// 		`LEFT JOIN ITEM_CATEGORY IC ON IC.ID = DEF.ITEM_CATEGORY_ID ` +
// 		`LEFT JOIN UOM UOM ON UOM.ID = IUL.UOM_ID ` +
// 		`JOIN ITEM_PRICE IP ON IP.UOM_ID = UOM.ID AND IP.ITEM_ID = IUL.ITEM_ID` +
// 		models.ItemWhereStatement + ` ` +
// 		conditionString + ` AND (LOWER(def."_name") LIKE $1)`
// 	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
// 	return data, count, err
// }

// FindByID ...
func (repository ItemRepository) FindByID(c context.Context, parameter models.ItemParameter) (data models.Item, err error) {
	statement := models.ItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository ItemRepository) FindByCategoryID(c context.Context, parameter models.ItemParameter) (data models.Item, err error) {
	statement := models.ItemSelectStatement + ` WHERE def.created_date IS NOT NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// SelectAllV2 ...
func (repository ItemRepository) SelectAllV2(c context.Context, parameter models.ItemParameter) (out []models.ItemV2, err error) {
	conditionString := ``

	var args []interface{}
	var index int = 1

	if parameter.ID != "" {
		conditionString += ` AND DEF.ID = $` + strconv.Itoa(index)
		args = append(args, parameter.ID)
		index++
	}

	if parameter.ItemCategoryId != "" {
		if parameter.ItemCategoryId == "2" {
			conditionString += ` AND DEF.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE $` + strconv.Itoa(index) + `) `
			args = append(args, "%tac%")
			index++
		} else {
			conditionString += ` AND DEF.item_category_id = $` + strconv.Itoa(index)
			args = append(args, parameter.ItemCategoryId)
			index++
		}
	}

	if parameter.PriceListId != "" {
		conditionString += ` and ip.price_list_version_id= (
			select plv.id 
			from price_list pl
			left join price_list_version plv on plv.price_list_id = pl.id
			where pl.id = $` + strconv.Itoa(index) + `
			order by plv.created_date desc
			limit 1
		)`
		args = append(args, parameter.PriceListId)
		index++
	} else if parameter.PriceListVersionId != "" {
		conditionString += ` and ip.price_list_version_id = $` + strconv.Itoa(index)
		args = append(args, parameter.PriceListVersionId)
		index++
	}

	if parameter.ExceptId != "" {
		conditionString += ` AND DEF.id <> $` + strconv.Itoa(index)
		args = append(args, parameter.ExceptId)
		index++
	}

	if parameter.UomID != "" {
		conditionString += ` AND IUL.UOM_ID = $` + strconv.Itoa(index)
		args = append(args, parameter.UomID)
		index++
	}

	if parameter.ItemCategoryName != "" {
		conditionString += ` OR LOWER(ic."_name") LIKE $` + strconv.Itoa(index)
		args = append(args, "%"+strings.ToLower(parameter.ItemCategoryName)+"%")
		index++
	}

	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
		conditionString += ` AND DEF.id NOT IN (83, 307, 393) `
	}

	statement := models.ItemV2SelectStatement + conditionString +
		`GROUP by def.id, td.MULTIPLY_DATA ` +
		`ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement, args...)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.ItemV2
		err := rows.Scan(
			&temp.ID,
			&temp.Code,
			&temp.Name,
			&temp.Description,
			&temp.ItemCategoryId,
			&temp.ItemCategoryName,
			&temp.AdditionalData,
			&temp.MultiplyData,
			&temp.ItemPicture,
		)
		if err != nil {
			return out, err
		}
		out = append(out, temp)
	}

	return
}

// func (repository ItemRepository) SelectAllV2(c context.Context, parameter models.ItemParameter) (out []models.ItemV2, err error) {
// 	conditionString := ``

// 	if parameter.ID != "" {
// 		conditionString += ` AND DEF.ID = '` + parameter.ID + `'`
// 	}

// 	if parameter.ItemCategoryId != "" {
// 		if parameter.ItemCategoryId == "2" {
// 			//KHUSUS TAC sendiri, tampilkan semua item dengan category TAC (TAC ANAK, BEBAS GULA, DLL)
// 			conditionString += ` AND DEF.item_category_id IN (SELECT id FROM item_category WHERE lower (_name) LIKE '%tac%') `
// 		} else {
// 			conditionString += ` AND DEF.item_category_id = ` + parameter.ItemCategoryId + ``
// 		}
// 	}

// 	if parameter.PriceListId != "" {
// 		conditionString += ` and ip.price_list_version_id= (
// 		select plv.id
// 		from price_list pl
// 		left join price_list_version plv on plv.price_list_id = pl.id
// 		where pl.id = ` + parameter.PriceListId + `
// 		order by plv.created_date desc
// 		limit 1
// 		)`
// 	} else if parameter.PriceListVersionId != "" {
// 		conditionString += ` and ip.price_list_version_id = ` + parameter.PriceListVersionId + ` `
// 	}

// 	if parameter.ExceptId != "" {
// 		conditionString += ` AND DEF.id <> '` + parameter.ExceptId + `'`
// 	}

// 	if parameter.UomID != "" {
// 		conditionString += ` AND IUL.UOM_ID = '` + parameter.UomID + `'`
// 	}

// 	if parameter.ItemCategoryName != "" {
// 		conditionString += ` or (LOWER (ic."_name") like ` + `'%` + strings.ToLower(parameter.ItemCategoryName) + `%')`
// 	}

// 	if parameter.CustomerTypeId != "" && (parameter.CustomerTypeId != "7" && parameter.CustomerTypeId != "15") {
// 		conditionString += ` AND def.id NOT IN (83, 307, 393) `
// 	}

// 	statement := models.ItemV2SelectStatement + conditionString +
// 		`GROUP by def.id, td.MULTIPLY_DATA ` +
// 		`ORDER BY ` + parameter.By + ` ` + parameter.Sort
// 	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")
// 	if err != nil {
// 		return
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		var temp models.ItemV2
// 		err := rows.Scan(
// 			&temp.ID,
// 			&temp.Code,
// 			&temp.Name,
// 			&temp.Description,
// 			&temp.ItemCategoryId,
// 			&temp.ItemCategoryName,
// 			&temp.AdditionalData,
// 			&temp.MultiplyData,
// 			&temp.ItemPicture,
// 		)
// 		if err != nil {
// 			return out, err
// 		}
// 		out = append(out, temp)
// 	}

// 	return
// }
