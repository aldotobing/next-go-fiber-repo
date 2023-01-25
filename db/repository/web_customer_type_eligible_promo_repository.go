package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebCustomerTypeEligiblePromo ...
type IWebCustomerTypeEligiblePromo interface {
	SelectAll(c context.Context, parameter models.WebCustomerTypeEligiblePromoParameter) ([]models.WebCustomerTypeEligiblePromo, error)
	FindAll(ctx context.Context, parameter models.WebCustomerTypeEligiblePromoParameter) ([]models.WebCustomerTypeEligiblePromo, int, error)
}

// WebCustomerTypeEligiblePromo ...
type WebCustomerTypeEligiblePromo struct {
	DB *sql.DB
}

// NewWebCustomerTypeEligiblePromo ...
func NewWebCustomerTypeEligiblePromoRepository(DB *sql.DB) IWebCustomerTypeEligiblePromo {
	return &WebCustomerTypeEligiblePromo{DB: DB}
}

// Scan rows
func (repository WebCustomerTypeEligiblePromo) scanRows(rows *sql.Rows) (res models.WebCustomerTypeEligiblePromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository WebCustomerTypeEligiblePromo) scanRow(row *sql.Row) (res models.WebCustomerTypeEligiblePromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.CustomerTypeId,
		&res.CustomerTypeName,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebCustomerTypeEligiblePromo) SelectAll(c context.Context, parameter models.WebCustomerTypeEligiblePromoParameter) (data []models.WebCustomerTypeEligiblePromo, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	statement := models.WebCustomerTypeEligiblePromoSelectStatement + ` ` + models.WebCustomerTypeEligiblePromoWhereStatement +
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
func (repository WebCustomerTypeEligiblePromo) FindAll(ctx context.Context, parameter models.WebCustomerTypeEligiblePromoParameter) (data []models.WebCustomerTypeEligiblePromo, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	query := models.WebCustomerTypeEligiblePromoSelectStatement + ` ` + models.WebCustomerTypeEligiblePromoWhereStatement + ` ` + conditionString + `
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
			LEFT JOIN promo PR ON PR.ID = CTEP.promo_id ` + models.WebCustomerTypeEligiblePromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(pr."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}
