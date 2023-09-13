package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IVoucherRedeemRepository ...
type IVoucherRedeemRepository interface {
	SelectAll(c context.Context, in models.VoucherRedeemParameter) ([]models.VoucherRedeem, error)
	FindAll(ctx context.Context, in models.VoucherRedeemParameter) ([]models.VoucherRedeem, int, error)
	FindByID(c context.Context, in models.VoucherRedeemParameter) (models.VoucherRedeem, error)
	FindByDocumentNo(c context.Context, in models.VoucherRedeemParameter) (models.VoucherRedeem, error)
	Add(c context.Context, model viewmodel.VoucherRedeemVM) (string, error)
	AddBulk(c context.Context, model []viewmodel.VoucherRedeemVM) error
	Update(c context.Context, model viewmodel.VoucherRedeemVM) (string, error)
	Redeem(c context.Context, model viewmodel.VoucherRedeemVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// VoucherRedeemRepository ...
type VoucherRedeemRepository struct {
	DB *sql.DB
}

// NewVoucherRedeemRepository ...
func NewVoucherRedeemRepository(DB *sql.DB) IVoucherRedeemRepository {
	return &VoucherRedeemRepository{DB: DB}
}

// Scan rows
func (repository VoucherRedeemRepository) scanRows(rows *sql.Rows) (res models.VoucherRedeem, err error) {
	err = rows.Scan(
		&res.ID,
		&res.CustomerID,
		&res.Redeemed,
		&res.RedeemedAt,
		&res.RedeemedToDocNo,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.VoucherID,
		&res.VoucherName,
		&res.VoucherCashValue,
		&res.VoucherDescription,
		&res.VoucherImageURL,
	)

	return
}

// Scan row
func (repository VoucherRedeemRepository) scanRow(row *sql.Row) (res models.VoucherRedeem, err error) {
	err = row.Scan(
		&res.ID,
		&res.CustomerID,
		&res.Redeemed,
		&res.RedeemedAt,
		&res.RedeemedToDocNo,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.VoucherID,
		&res.VoucherName,
		&res.VoucherCashValue,
		&res.VoucherDescription,
		&res.VoucherImageURL,
	)

	return
}

// SelectAll ...
func (repository VoucherRedeemRepository) SelectAll(c context.Context, in models.VoucherRedeemParameter) (data []models.VoucherRedeem, err error) {
	var conditionString string

	if in.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = ` + in.CustomerID
	}

	statement := models.VoucherRedeemSelectStatement + models.VoucherRedeemWhereStatement +
		` AND DEF.REDEEMED_AT IS NULL AND DEF.REDEEMED_TO_DOC_NO IS NULL ` + conditionString +
		` ORDER BY ` + in.By + ` ` + in.Sort

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
func (repository VoucherRedeemRepository) FindAll(ctx context.Context, in models.VoucherRedeemParameter) (data []models.VoucherRedeem, count int, err error) {
	var conditionString string

	conditionString += ` AND DEF.REDEEMED_AT IS NULL `

	statement := models.VoucherRedeemSelectStatement + models.VoucherRedeemWhereStatement +
		conditionString +
		` ORDER BY ` + in.By + ` ` + in.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.QueryContext(ctx, statement, in.Offset, in.Limit)
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

	countQuery := `SELECT COUNT(*) FROM VOUCHER_REDEEM DEF ` + models.VoucherRedeemWhereStatement + conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository VoucherRedeemRepository) FindByID(c context.Context, in models.VoucherRedeemParameter) (data models.VoucherRedeem, err error) {
	statement := models.VoucherRedeemSelectStatement + ` WHERE DEF.ID = ` + in.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// FindByDocumentNo ...
func (repository VoucherRedeemRepository) FindByDocumentNo(c context.Context, in models.VoucherRedeemParameter) (data models.VoucherRedeem, err error) {
	statement := models.VoucherRedeemSelectStatement + ` WHERE DEF.REDEEMED_TO_DOC_NO = '` + in.DocumentNo + `'`
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository VoucherRedeemRepository) Add(c context.Context, in viewmodel.VoucherRedeemVM) (res string, err error) {
	statement := `INSERT INTO VOUCHER_REDEEM (
			CUSTOMER_ID, 
			VOUCHER_ID,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		in.CustomerID,
		in.VoucherID,
	).Scan(&res)

	return
}

// AddBulk ...
func (repository VoucherRedeemRepository) AddBulk(c context.Context, in []viewmodel.VoucherRedeemVM) (err error) {
	var valueStatement string
	for _, datum := range in {
		if valueStatement == "" {
			valueStatement += `(` + datum.CustomerID + `, ` + datum.VoucherID + `, NOW(), NOW())`
		} else {
			valueStatement += `, (` + datum.CustomerID + `, ` + datum.VoucherID + `, NOW(), NOW())`
		}
	}
	statement := `INSERT INTO VOUCHER_REDEEM (
			CUSTOMER_ID, 
			VOUCHER_ID,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ` + valueStatement
	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

// Update ...
func (repository VoucherRedeemRepository) Update(c context.Context, in viewmodel.VoucherRedeemVM) (res string, err error) {
	statement := `UPDATE VOUCHER_REDEEM SET 
		CUSTOMER_ID = $1, 
		VOUCHER_ID = $2, 
		UPDATED_AT = now()
	WHERE id = $3
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		in.CustomerID,
		in.VoucherID,
		in.ID).Scan(&res)

	return
}

// Redeem ...
func (repository VoucherRedeemRepository) Redeem(c context.Context, in viewmodel.VoucherRedeemVM) (res string, err error) {
	statement := `UPDATE VOUCHER_REDEEM SET 
		REDEEMED_TO_DOC_NO = $1,
		REDEEMED_AT = now(),
		UPDATED_AT = NOW()
	WHERE id = $2
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		in.RedeemedToDocumentNo,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository VoucherRedeemRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE VOUCHER_REDEEM SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
