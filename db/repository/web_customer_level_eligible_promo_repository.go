package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebCustomerLevelEligiblePromo ...
type IWebCustomerLevelEligiblePromo interface {
	SelectAll(c context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) ([]models.WebCustomerLevelEligiblePromo, error)
	FindAll(ctx context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) ([]models.WebCustomerLevelEligiblePromo, int, error)
}

// WebCustomerLevelEligiblePromo ...
type WebCustomerLevelEligiblePromo struct {
	DB *sql.DB
}

// NewWebCustomerLevelEligiblePromo ...
func NewWebCustomerLevelEligiblePromoRepository(DB *sql.DB) IWebCustomerLevelEligiblePromo {
	return &WebCustomerLevelEligiblePromo{DB: DB}
}

// Scan rows
func (repository WebCustomerLevelEligiblePromo) scanRows(rows *sql.Rows) (res models.WebCustomerLevelEligiblePromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.CustomerLevelId,
		&res.CustomerLevelName,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository WebCustomerLevelEligiblePromo) scanRow(row *sql.Row) (res models.WebCustomerLevelEligiblePromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.CustomerLevelId,
		&res.CustomerLevelName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebCustomerLevelEligiblePromo) SelectAll(c context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) (data []models.WebCustomerLevelEligiblePromo, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	statement := models.WebCustomerLevelEligiblePromoSelectStatement + ` ` + models.WebCustomerLevelEligiblePromoWhereStatement +
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
func (repository WebCustomerLevelEligiblePromo) FindAll(ctx context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) (data []models.WebCustomerLevelEligiblePromo, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	query := models.WebCustomerLevelEligiblePromoSelectStatement + ` ` + models.WebCustomerLevelEligiblePromoWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT (CTEP.ID)
			FROM customer_type_eligible_promo CTEP
			LEFT JOIN customer_type ct ON ct.id = CTEP.customer_type_id
			LEFT JOIN promo PR ON PR.ID = CTEP.promo_id ` + models.WebCustomerLevelEligiblePromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(pr."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}
