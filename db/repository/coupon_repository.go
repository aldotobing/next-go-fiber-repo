package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ICouponRepository ...
type ICouponRepository interface {
	SelectAll(c context.Context, parameter models.CouponParameter) ([]models.Coupon, error)
	FindAll(ctx context.Context, parameter models.CouponParameter) ([]models.Coupon, int, error)
	FindByID(c context.Context, parameter models.CouponParameter) (models.Coupon, error)
	Add(c context.Context, model viewmodel.CouponVM) (string, error)
	Update(c context.Context, model viewmodel.CouponVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// CouponRepository ...
type CouponRepository struct {
	DB *sql.DB
}

// NewCouponRepository ...
func NewCouponRepository(DB *sql.DB) ICouponRepository {
	return &CouponRepository{DB: DB}
}

// Scan rows
func (repository CouponRepository) scanRows(rows *sql.Rows) (res models.Coupon, err error) {
	err = rows.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.PointConversion,
		&res.Name,
		&res.Description,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// Scan row
func (repository CouponRepository) scanRow(row *sql.Row) (res models.Coupon, err error) {
	err = row.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.PointConversion,
		&res.Name,
		&res.Description,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// SelectAll ...
func (repository CouponRepository) SelectAll(c context.Context, parameter models.CouponParameter) (data []models.Coupon, err error) {
	var conditionString string

	if parameter.Now != "" {
		conditionString += ` AND '` + parameter.Now + `' BETWEEN DEF.START_DATE AND DEF.END_DATE `
	}
	statement := models.CouponSelectStatement + models.CouponWhereStatement +
		conditionString +
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
func (repository CouponRepository) FindAll(ctx context.Context, parameter models.CouponParameter) (data []models.Coupon, count int, err error) {
	var conditionString string

	if parameter.Now != "" {
		conditionString += ` AND '` + parameter.Now + `' BETWEEN DEF.START_DATE AND DEF.END_DATE `
	}
	statement := models.CouponSelectStatement + models.CouponWhereStatement +
		conditionString +
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

	countQuery := `SELECT COUNT(*) FROM POINT_RULES def ` + models.CouponWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository CouponRepository) FindByID(c context.Context, parameter models.CouponParameter) (data models.Coupon, err error) {
	statement := models.CouponSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository CouponRepository) Add(c context.Context, in viewmodel.CouponVM) (res string, err error) {
	statement := `INSERT INTO COUPONS (
			START_DATE, 
			END_DATE,
			POINT_CONVERSION,
			_NAME,
			DESCRIPTION,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.PointConversion,
		in.Name,
		in.Description,
	).Scan(&res)

	return
}

// Update ...
func (repository CouponRepository) Update(c context.Context, in viewmodel.CouponVM) (res string, err error) {
	statement := `UPDATE COUPONS SET 
		START_DATE = $1, 
		END_DATE = $2,
		POINT_CONVERSION = $3,
		_NAME = $4,
		DESCRIPTION = $5,
		UPDATED_AT = now()
	WHERE id = $6
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.PointConversion,
		in.Name,
		in.Description,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository CouponRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE COUPONS SET 
		DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
