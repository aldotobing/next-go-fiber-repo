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
	SelectReport(c context.Context, parameter models.CouponRedeemParameter) ([]models.CouponRedeemReport, error)
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
		&res.CouponPhotoURL,
		&res.CustomerName,
		&res.CouponCode,
		&res.InvoiceNo,
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
		&res.CouponPhotoURL,
		&res.CustomerName,
		&res.CouponCode,
		&res.InvoiceNo,
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
			EXPIRED_AT,
            COUPON_CODE
		)
	VALUES ($1, $2, NOW(), NOW(), $3, $4) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.CouponID,
		in.CustomerID,
		in.ExpiredAt,
		in.CouponCode,
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

// SelectReport ...
func (repository CouponRedeemRepository) SelectReport(c context.Context, parameter models.CouponRedeemParameter) (data []models.CouponRedeemReport, err error) {
	var conditionString string

	if parameter.ShowAll == "" {
		conditionString += ` AND DEF.REDEEMED_AT IS NULL AND NOW()::DATE<DEF.EXPIRED_AT`
	}
	if parameter.CustomerID != "" {
		conditionString += ` AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND DEF.CREATED_AT::DATE BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		conditionString += ` AND DEF.CREATED_AT::DATE BETWEEN date_trunc('MONTH',now())::DATE AND now()::date`
	}

	if parameter.BranchID != "" {
		conditionString += ` AND B.ID in (` + parameter.BranchID + `)`
	}
	if parameter.RegionGroupID != "" {
		conditionString += ` AND R.GROUP_ID in (` + parameter.RegionGroupID + `)`
	}
	if parameter.RegionID != "" {
		conditionString += ` AND R.ID in (` + parameter.RegionID + `)`
	}

	if parameter.CustomerLevelID != "" {
		conditionString += ` AND C.CUSTOMER_LEVEL_ID IN (` + parameter.CustomerLevelID + `)`
	}

	if parameter.CouponStatus != "" {
		conditionString += ` AND DEF.REDEEMED = '` + parameter.CouponStatus + `'`
	}

	statement := `SELECT 
			DEF.ID, 
			DEF.COUPON_ID,
			DEF.CUSTOMER_ID,
			DEF.REDEEMED,
			DEF.REDEEMED_AT,
			DEF.REDEEM_TO_DOC_NO,
			DEF.CREATED_AT,
			DEF.UPDATED_AT,
			DEF.DELETED_AT,
			DEF.EXPIRED_AT,
			CP._NAME,
			CP.DESCRIPTION,
			CP.POINT_CONVERSION,
			C.CUSTOMER_NAME,
			C.CUSTOMER_CODE,
			B._NAME,
			B.BRANCH_CODE,
			R._NAME,
			R.GROUP_NAME,
			CL._NAME,
			DEF.COUPON_CODE,
			SIH.DOCUMENT_NO,
			SOH.DOCUMENT_NO
		FROM COUPON_REDEEM DEF
		LEFT JOIN SALES_INVOICE_HEADER SIH ON left(SIH.transaction_source_document_no,15) = left(DEF.REDEEM_TO_DOC_NO, 15)
		LEFT JOIN SALES_ORDER_HEADER SOH ON left(SOH.request_document_no,15) = left(DEF.REDEEM_TO_DOC_NO, 15)
		LEFT JOIN COUPONS CP ON CP.ID = DEF.COUPON_ID
		LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
		LEFT JOIN CUSTOMER_LEVEL CL ON CL.ID = C.CUSTOMER_LEVEL_ID
		LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
		LEFT JOIN REGION R ON R.ID = B.REGION_ID
		WHERE DEF.DELETED_AT IS NULL ` + conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.CouponRedeemReport

		err = rows.Scan(
			&temp.ID,
			&temp.CouponID,
			&temp.CustomerID,
			&temp.Redeem,
			&temp.RedeemAt,
			&temp.RedeemedToDocumentNo,
			&temp.CreatedAt,
			&temp.UpdatedAt,
			&temp.DeletedAt,
			&temp.ExpiredAt,
			&temp.CouponName,
			&temp.CouponDescription,
			&temp.CouponPointConversion,
			&temp.CustomerName,
			&temp.CustomerCode,
			&temp.BranchName,
			&temp.BranchCode,
			&temp.RegionName,
			&temp.RegionGroupName,
			&temp.CustomerLevelName,
			&temp.CouponCode,
			&temp.InvoiceNo,
			&temp.SalesOrderDocumentNo,
		)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}
