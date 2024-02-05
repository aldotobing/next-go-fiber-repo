package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ICouponRedeemRepository ...
type ICouponRedeemRepository interface {
	SelectAll(c context.Context, parameter models.CouponRedeemParameter) ([]models.CouponRedeem, error)
	FindAll(ctx context.Context, parameter models.CouponRedeemParameter) ([]models.CouponRedeem, int, error)
	FindByID(c context.Context, parameter models.CouponRedeemParameter) (models.CouponRedeem, error)
	Add(c context.Context, model viewmodel.CouponRedeemVM) (string, error)
	Redeem(c context.Context, model viewmodel.CouponRedeemVM) (string, error)
}

// CouponRedeemRepository ...
type CouponRedeemRepository struct {
	DB *sql.DB
}

// NewCouponRedeemRepository ...
func NewCouponRedeemRepository(DB *sql.DB) ICouponRedeemRepository {
	return &CouponRedeemRepository{DB: DB}
}

// Scan rows
func (repository CouponRedeemRepository) scanRows(rows *sql.Rows) (res models.CouponRedeem, err error) {
	err = rows.Scan(
		&res.ID,
		&res.CouponID,
		&res.CustomerID,
		&res.Redeem,
		&res.RedeemAt,
		&res.RedeemedToDocumentNo,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.ExpiredAt,
		&res.CouponName,
		&res.CouponDescription,
		&res.CouponPointConversion,
		&res.CustomerName,
	)

	return
}

// Scan row
func (repository CouponRedeemRepository) scanRow(row *sql.Row) (res models.CouponRedeem, err error) {
	err = row.Scan(
		&res.ID,
		&res.CouponID,
		&res.CustomerID,
		&res.Redeem,
		&res.RedeemAt,
		&res.RedeemedToDocumentNo,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.ExpiredAt,
		&res.CouponName,
		&res.CouponDescription,
		&res.CouponPointConversion,
		&res.CustomerName,
	)

	return
}

// SelectAll ...
func (repository CouponRedeemRepository) SelectAll(c context.Context, parameter models.CouponRedeemParameter) (data []models.CouponRedeem, err error) {
	var conditionString string

	if parameter.ShowAll == "" {
		conditionString += ` AND DEF.REDEEMED_AT IS NULL AND NOW()::DATE<DEF.EXPIRED_AT`
	}
	if parameter.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}
	statement := models.CouponRedeemSelectStatement + models.CouponRedeemWhereStatement +
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
func (repository CouponRedeemRepository) FindAll(ctx context.Context, parameter models.CouponRedeemParameter) (data []models.CouponRedeem, count int, err error) {
	var conditionString string

	conditionString += ` AND DEF.REDEEMED_AT IS NULL AND NOW()::DATE<DEF.EXPIRED_AT`

	if parameter.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}
	statement := models.CouponRedeemSelectStatement + models.CouponRedeemWhereStatement +
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

	countQuery := `SELECT COUNT(*) FROM COUPON_REDEEM def ` + models.CouponRedeemWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository CouponRedeemRepository) FindByID(c context.Context, parameter models.CouponRedeemParameter) (data models.CouponRedeem, err error) {
	statement := models.CouponRedeemSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository CouponRedeemRepository) Add(c context.Context, in viewmodel.CouponRedeemVM) (res string, err error) {
	statement := `INSERT INTO COUPON_REDEEM (
			COUPON_ID, 
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT,
			EXPIRED_AT
		)
	VALUES ($1, $2, NOW(), NOW(), $3) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.CouponID,
		in.CustomerID,
		in.ExpiredAt,
	).Scan(&res)

	return
}

// Redeem ...
func (repository CouponRedeemRepository) Redeem(c context.Context, in viewmodel.CouponRedeemVM) (res string, err error) {
	statement := `UPDATE COUPON_REDEEM SET 
		REDEEMED = $1, 
		REDEEMED_AT = NOW(),
		REDEEM_TO_DOC_NO = $2,
		UPDATED_AT = now()
	WHERE id = $3
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.Redeem,
		in.RedeemedToDocumentNo,
		in.ID).Scan(&res)

	return
}