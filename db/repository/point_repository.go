package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointRepository ...
type IPointRepository interface {
	SelectAll(c context.Context, parameter models.PointParameter) ([]models.Point, error)
	FindAll(ctx context.Context, parameter models.PointParameter) ([]models.Point, int, error)
	FindByID(c context.Context, parameter models.PointParameter) (models.Point, error)
	GetBalance(c context.Context, parameter models.PointParameter) (models.PointGetBalance, error)
	Add(c context.Context, model viewmodel.PointVM) (string, error)
	Update(c context.Context, model viewmodel.PointVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// PointRepository ...
type PointRepository struct {
	DB *sql.DB
}

// NewPointRepository ...
func NewPointRepository(DB *sql.DB) IPointRepository {
	return &PointRepository{DB: DB}
}

// Scan rows
func (repository PointRepository) scanRows(rows *sql.Rows) (res models.Point, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PointType,
		&res.PointTypeName,
		&res.InvoiceID,
		&res.Point,
		&res.CustomerID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// Scan row
func (repository PointRepository) scanRow(row *sql.Row) (res models.Point, err error) {
	err = row.Scan(
		&res.ID,
		&res.PointType,
		&res.PointTypeName,
		&res.InvoiceID,
		&res.Point,
		&res.CustomerID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// SelectAll ...
func (repository PointRepository) SelectAll(c context.Context, parameter models.PointParameter) (data []models.Point, err error) {
	var conditionString string

	statement := models.PointSelectStatement + models.PointWhereStatement +
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
func (repository PointRepository) FindAll(ctx context.Context, parameter models.PointParameter) (data []models.Point, count int, err error) {
	var conditionString string

	if parameter.CustomerID != "" {
		conditionString += `AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += `AND DEF.CREATED_AT BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	}

	statement := models.PointSelectStatement + models.PointWhereStatement +
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

	countQuery := `SELECT COUNT(*) FROM POINTS def ` + models.PointWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository PointRepository) FindByID(c context.Context, parameter models.PointParameter) (data models.Point, err error) {
	statement := models.PointSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// GetBalance ...
func (repository PointRepository) GetBalance(c context.Context, parameter models.PointParameter) (data models.PointGetBalance, err error) {
	statement := `select coalesce(sum(case when pt."_name" = '` + models.PointTypeWithdraw + `' then DEF.point else 0 end),0) as withdraw,
		coalesce(sum(case when pt."_name" = '` + models.PointTypeCashback + `' then DEF.point else 0 end),0) as cashback,
		coalesce(sum(case when pt."_name" = '` + models.PointTypeLoyalty + `' then DEF.point else 0 end),0) as loyalty,
		coalesce(sum(case when pt."_name" = '` + models.PointTypePromo + `' then DEF.point else 0 end),0) as promo
		from points DEF
		left join point_type pt on pt.id = def.point_type ` +
		`WHERE DEF.CUSTOMER_ID = ` + parameter.CustomerID
	row := repository.DB.QueryRowContext(c, statement)

	err = row.Scan(
		&data.Withdraw,
		&data.Cashback,
		&data.Loyalty,
		&data.Promo,
	)

	return
}

// Add ...
func (repository PointRepository) Add(c context.Context, in viewmodel.PointVM) (res string, err error) {
	statement := `INSERT INTO POINTS (
			POINT_TYPE, 
			INVOICE_ID,
			POINT,
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES (` + in.PointType + `, ` + in.InvoiceID + `, '` + in.Point + `', ` + in.CustomerID + `, NOW(), NOW()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}

// Update ...
func (repository PointRepository) Update(c context.Context, in viewmodel.PointVM) (res string, err error) {
	statement := `UPDATE POINTS SET 
		POINT_TYPE = $1, 
		INVOICE_ID = $2, 
		POINT = $3, 
		CUSTOMER_ID = $4,
		UPDATED_AT = now()
	WHERE id = $5
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.PointType,
		in.InvoiceID,
		in.Point,
		in.CustomerID,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository PointRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINTS SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
