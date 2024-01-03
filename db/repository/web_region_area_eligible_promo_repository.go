package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebRegionAreaEligiblePromo ...
type IWebRegionAreaEligiblePromo interface {
	SelectAll(c context.Context, parameter models.WebRegionAreaEligiblePromoParameter) ([]models.WebRegionAreaEligiblePromo, error)
	FindAll(ctx context.Context, parameter models.WebRegionAreaEligiblePromoParameter) ([]models.WebRegionAreaEligiblePromo, int, error)
}

// WebRegionAreaEligiblePromo ...
type WebRegionAreaEligiblePromo struct {
	DB *sql.DB
}

// NewWebRegionAreaEligiblePromo ...
func NewWebRegionAreaEligiblePromoRepository(DB *sql.DB) IWebRegionAreaEligiblePromo {
	return &WebRegionAreaEligiblePromo{DB: DB}
}

// Scan rows
func (repository WebRegionAreaEligiblePromo) scanRows(rows *sql.Rows) (res models.WebRegionAreaEligiblePromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.RegionID,
		&res.RegionName,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository WebRegionAreaEligiblePromo) scanRow(row *sql.Row) (res models.WebRegionAreaEligiblePromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.RegionID,
		&res.RegionName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebRegionAreaEligiblePromo) SelectAll(c context.Context, parameter models.WebRegionAreaEligiblePromoParameter) (data []models.WebRegionAreaEligiblePromo, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	statement := models.WebRegionAreaEligiblePromoSelectStatement + ` ` + models.WebRegionAreaEligiblePromoWhereStatement +
		` AND (LOWER(pr._name) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")

	//print
	fmt.Println(statement)

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
func (repository WebRegionAreaEligiblePromo) FindAll(ctx context.Context, parameter models.WebRegionAreaEligiblePromoParameter) (data []models.WebRegionAreaEligiblePromo, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	query := models.WebRegionAreaEligiblePromoSelectStatement + ` ` + models.WebRegionAreaEligiblePromoWhereStatement + ` ` + conditionString + `
		AND (LOWER(pr."_name") LIKE $1  )` + `GROUP BY ` + `ORDER BY` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
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

	query = `SELECT 
				count(RAEP.ID) 
				FROM region_area_eligible_promo RAEP
				LEFT JOIN region r ON r.id = RAEP.region_id
				LEFT JOIN promo PR ON PR.ID = RAEP.promo_id ` + models.WebRegionAreaEligiblePromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(pr."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}
