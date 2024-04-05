package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointPromoRepository ...
type IPointPromoRepository interface {
	SelectAll(c context.Context, parameter models.PointPromoParameter) ([]models.PointPromo, error)
	FindAll(ctx context.Context, parameter models.PointPromoParameter) ([]models.PointPromo, int, error)
	FindByID(c context.Context, parameter models.PointPromoParameter) (models.PointPromo, error)
	Add(c context.Context, model viewmodel.PointPromoVM) (string, error)
	Update(c context.Context, model viewmodel.PointPromoVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// PointPromoRepository ...
type PointPromoRepository struct {
	DB *sql.DB
}

// NewPointPromoRepository ...
func NewPointPromoRepository(DB *sql.DB) IPointPromoRepository {
	return &PointPromoRepository{DB: DB}
}

// Scan rows
func (repository PointPromoRepository) scanRows(rows *sql.Rows) (res models.PointPromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Multiplicator,
		&res.PointConversion,
		&res.QuantityConversion,
		&res.PromoType,
		&res.Strata,
		&res.Items,
		&res.Image,
		&res.Title,
		&res.Description,
	)

	return
}

// Scan row
func (repository PointPromoRepository) scanRow(row *sql.Row) (res models.PointPromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Multiplicator,
		&res.PointConversion,
		&res.QuantityConversion,
		&res.PromoType,
		&res.Strata,
		&res.Items,
		&res.Image,
		&res.Title,
		&res.Description,
	)

	return
}

// SelectAll ...
func (repository PointPromoRepository) SelectAll(c context.Context, parameter models.PointPromoParameter) (data []models.PointPromo, err error) {
	var conditionString string

	if parameter.Now {
		conditionString += ` AND NOW()::date BETWEEN DEF.START_DATE AND END_DATE`
	}
	statement := models.PointPromoSelectStatement + models.PointPromoWhereStatement +
		conditionString +
		` GROUP BY DEF.ID ORDER BY ` + parameter.By + ` ` + parameter.Sort

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
func (repository PointPromoRepository) FindAll(ctx context.Context, parameter models.PointPromoParameter) (data []models.PointPromo, count int, err error) {
	var conditionString string

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += `AND DEF.CREATED_AT BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	}

	statement := models.PointPromoSelectStatement + models.PointPromoWhereStatement +
		conditionString +
		` GROUP BY DEF.ID ORDER BY ` + parameter.By + ` ` + parameter.Sort +
		` OFFSET $1 LIMIT $2`

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

	countQuery := `SELECT COUNT(*) FROM POINT_PROMO def ` + models.PointPromoWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository PointPromoRepository) FindByID(c context.Context, parameter models.PointPromoParameter) (data models.PointPromo, err error) {
	statement := models.PointPromoSelectStatement + ` WHERE DEF.ID = ` + parameter.ID + `GROUP BY DEF.ID`
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository PointPromoRepository) Add(c context.Context, in viewmodel.PointPromoVM) (res string, err error) {
	var statementInsert string

	j, _ := json.Marshal(in.Strata)
	statementInsert += `('` + in.StartDate + `', '` + in.EndDate + `', NOW(), NOW(), '` +
		strconv.FormatBool(in.Multiplicator) + `', '` + in.PointConversion + `', '` + string(j) + `', '` + in.QuantityConversion + `', '` +
		in.PromoType + `', '` + in.Image + `', '` + in.Title + `', '` + in.Description + `')`

	statement := `INSERT INTO POINT_PROMO (
			START_DATE, 
			END_DATE,
			CREATED_AT,
			UPDATED_AT,
			MULTIPLICATOR,
			POINT_CONVERSION,
			STRATA,
			QUANTITY_CONVERSION,
			PROMO_TYPE,
			IMAGE_URL,
			TITLE,
			DESCRIPTION
		)
	VALUES ` + statementInsert + `RETURNING ID`

	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}

// Update ...
func (repository PointPromoRepository) Update(c context.Context, in viewmodel.PointPromoVM) (res string, err error) {
	statement := `UPDATE POINT_PROMO SET 
		START_DATE = $1,
		END_DATE = $2,
		UPDATED_AT = NOW(),
		MULTIPLICATOR = $3,
		POINT_CONVERSION = $4,
		STRATA = $5,
		QUANTITY_CONVERSION = $6,
		PROMO_TYPE = $7,
		IMAGE_URL = $8,
		TITLE = $9,
		DESCRIPTION = $10
	WHERE ID = $11
	RETURNING ID`

	j, _ := json.Marshal(in.Strata)
	err = repository.DB.QueryRowContext(c, statement,
		in.StartDate,
		in.EndDate,
		in.Multiplicator,
		in.PointConversion,
		string(j),
		in.QuantityConversion,
		in.PromoType,
		in.Image,
		in.Title,
		in.Description,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository PointPromoRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINT_PROMO SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
