package repository

import (
	"context"
	"database/sql"
	"strconv"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IVoucherRepository ...
type IVoucherRepository interface {
	SelectAll(c context.Context, parameter models.VoucherParameter) ([]models.Voucher, error)
	FindAll(ctx context.Context, parameter models.VoucherParameter) ([]models.Voucher, int, error)
	FindByID(c context.Context, parameter models.VoucherParameter) (models.Voucher, error)
	Add(c context.Context, model viewmodel.VoucherVM) (string, error)
	Update(c context.Context, model viewmodel.VoucherVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// VoucherRepository ...
type VoucherRepository struct {
	DB *sql.DB
}

// NewVoucherRepository ...
func NewVoucherRepository(DB *sql.DB) IVoucherRepository {
	return &VoucherRepository{DB: DB}
}

// Scan rows
func (repository VoucherRepository) scanRows(rows *sql.Rows) (res models.Voucher, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.StartDate,
		&res.EndDate,
		&res.ImageURL,
		&res.VoucherCategoryID,
		&res.CashValue,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Description,
	)

	return
}

// Scan row
func (repository VoucherRepository) scanRow(row *sql.Row) (res models.Voucher, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.Name,
		&res.StartDate,
		&res.EndDate,
		&res.ImageURL,
		&res.VoucherCategoryID,
		&res.CashValue,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Description,
	)

	return
}

// SelectAll ...
func (repository VoucherRepository) SelectAll(c context.Context, parameter models.VoucherParameter) (data []models.Voucher, err error) {
	var conditionString string

	statement := models.VoucherSelectStatement + models.VoucherWhereStatement +
		` AND (LOWER(def."_name") LIKE '%` + parameter.Search + `%') ` + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement)

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
func (repository VoucherRepository) FindAll(ctx context.Context, parameter models.VoucherParameter) (data []models.Voucher, count int, err error) {
	var conditionString string

	statement := models.VoucherSelectStatement + models.VoucherWhereStatement +
		` AND (LOWER(def."_name") LIKE '%` + parameter.Search + `%') ` + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.QueryContext(ctx, statement, parameter.Offset, parameter.Limit)
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

	countQuery := `SELECT COUNT(*) FROM VOUCHER def ` + models.VoucherWhereStatement +
		` AND (LOWER(def."_name") LIKE '%` + parameter.Search + `%') ` + conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository VoucherRepository) FindByID(c context.Context, parameter models.VoucherParameter) (data models.Voucher, err error) {
	statement := models.VoucherSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository VoucherRepository) Add(c context.Context, in viewmodel.VoucherVM) (res string, err error) {
	statement := `INSERT INTO VOUCHER (
			CODE, 
			_NAME,
			START_DATE,
			END_DATE,
			IMAGE_URL,
			VOUCHER_CATEGORY_ID,
			CASH_VALUE,
			CREATED_AT,
			UPDATED_AT,
			DESCRIPTION
		)
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), $8) RETURNING id`

	voucherCategoryID, _ := strconv.Atoi(in.VoucherCategoryID)
	err = repository.DB.QueryRowContext(c, statement,
		in.Code,
		in.Name,
		in.StartDate,
		in.EndDate,
		in.ImageURL,
		voucherCategoryID,
		in.CashValue,
		in.Description,
	).Scan(&res)

	return
}

// Update ...
func (repository VoucherRepository) Update(c context.Context, in viewmodel.VoucherVM) (res string, err error) {
	statement := `UPDATE VOUCHER SET 
		CODE = $1, 
		_NAME = $2, 
		START_DATE = $3, 
		END_DATE = $4,
		IMAGE_URL = $5,
		VOUCHER_CATEGORY_ID = $6,
		CASH_VALUE = $7,
		UPDATED_AT = now(),
		DESCRIPTION = $8
	WHERE id = $9
	RETURNING id`

	voucherCategoryID, _ := strconv.Atoi(in.VoucherCategoryID)
	err = repository.DB.QueryRowContext(c, statement,
		in.Code,
		in.Name,
		in.StartDate,
		in.EndDate,
		in.ImageURL,
		voucherCategoryID,
		in.CashValue,
		in.Description,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository VoucherRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE VOUCHER SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
