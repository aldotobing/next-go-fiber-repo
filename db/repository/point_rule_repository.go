package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointRuleRepository ...
type IPointRuleRepository interface {
	SelectAll(c context.Context, parameter models.PointRuleParameter) ([]models.PointRule, error)
	FindAll(ctx context.Context, parameter models.PointRuleParameter) ([]models.PointRule, int, error)
	FindByID(c context.Context, parameter models.PointRuleParameter) (models.PointRule, error)
	Add(c context.Context, model viewmodel.PointRuleVM) (string, error)
	Update(c context.Context, model viewmodel.PointRuleVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// PointRuleRepository ...
type PointRuleRepository struct {
	DB *sql.DB
}

// NewPointRuleRepository ...
func NewPointRuleRepository(DB *sql.DB) IPointRuleRepository {
	return &PointRuleRepository{DB: DB}
}

// Scan rows
func (repository PointRuleRepository) scanRows(rows *sql.Rows) (res models.PointRule, err error) {
	err = rows.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.MinOrder,
		&res.PointConversion,
		&res.MonthlyMaxPoint,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// Scan row
func (repository PointRuleRepository) scanRow(row *sql.Row) (res models.PointRule, err error) {
	err = row.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.MinOrder,
		&res.PointConversion,
		&res.MonthlyMaxPoint,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)

	return
}

// SelectAll ...
func (repository PointRuleRepository) SelectAll(c context.Context, parameter models.PointRuleParameter) (data []models.PointRule, err error) {
	var conditionString string

	statement := models.PointRuleSelectStatement + models.PointRuleWhereStatement +
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
func (repository PointRuleRepository) FindAll(ctx context.Context, parameter models.PointRuleParameter) (data []models.PointRule, count int, err error) {
	var conditionString string

	statement := models.PointRuleSelectStatement + models.PointRuleWhereStatement +
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

	countQuery := `SELECT COUNT(*) FROM POINT_RULES def ` + models.PointRuleWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository PointRuleRepository) FindByID(c context.Context, parameter models.PointRuleParameter) (data models.PointRule, err error) {
	statement := models.PointRuleSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository PointRuleRepository) Add(c context.Context, in viewmodel.PointRuleVM) (res string, err error) {
	statement := `INSERT INTO POINT_RULES (
			START_DATE, 
			END_DATE,
			MIN_ORDER,
			POINT_CONVERSION,
			MONTHLY_MAX_POINT,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.MinOrder,
		in.PointConversion,
		in.MonthlyMaxPoint,
	).Scan(&res)

	return
}

// Update ...
func (repository PointRuleRepository) Update(c context.Context, in viewmodel.PointRuleVM) (res string, err error) {
	statement := `UPDATE POINT_RULES SET 
		START_DATE = $1, 
		END_DATE = $2,
		MIN_ORDER = $3,
		POINT_CONVERSION = $4,
		MONTHLY_MAX_POINT = $5,
		UPDATED_AT = now()
	WHERE id = $6
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.MinOrder,
		in.PointConversion,
		in.MonthlyMaxPoint,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository PointRuleRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINT_RULES SET 
		DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
