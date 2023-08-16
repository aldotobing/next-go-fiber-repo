package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebBranchEligiblePromo ...
type IWebBranchEligiblePromo interface {
	SelectAll(c context.Context, parameter models.WebBranchEligiblePromoParameter) ([]models.WebBranchEligiblePromo, error)
	FindAll(ctx context.Context, parameter models.WebBranchEligiblePromoParameter) ([]models.WebBranchEligiblePromo, int, error)
}

// WebBranchEligiblePromo ...
type WebBranchEligiblePromo struct {
	DB *sql.DB
}

// NewWebBranchEligiblePromo ...
func NewWebBranchEligiblePromoRepository(DB *sql.DB) IWebBranchEligiblePromo {
	return &WebBranchEligiblePromo{DB: DB}
}

// Scan rows
func (repository WebBranchEligiblePromo) scanRows(rows *sql.Rows) (res models.WebBranchEligiblePromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.BranchId,
		&res.BranchName,
	)

	return
}

// Scan row
func (repository WebBranchEligiblePromo) scanRow(row *sql.Row) (res models.WebBranchEligiblePromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.PromoID,
		&res.Code,
		&res.PromoName,
		&res.BranchId,
		&res.BranchName,
	)

	return
}

// SelectAll ...
func (repository WebBranchEligiblePromo) SelectAll(c context.Context, parameter models.WebBranchEligiblePromoParameter) (data []models.WebBranchEligiblePromo, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	statement := models.WebBranchEligiblePromoSelectStatement + ` ` + models.WebBranchEligiblePromoWhereStatement +
		` AND (LOWER(pr._name) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Search)+"%")
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
func (repository WebBranchEligiblePromo) FindAll(ctx context.Context, parameter models.WebBranchEligiblePromoParameter) (data []models.WebBranchEligiblePromo, count int, err error) {
	conditionString := ``

	if parameter.PromoID != "" {
		conditionString += ` AND PR.ID = ` + parameter.PromoID + ` `
	}

	query := models.WebBranchEligiblePromoSelectStatement + ` ` + models.WebBranchEligiblePromoWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT (BEP.ID)
			FROM branch_eligible_promo BEP
			LEFT JOIN branch b ON ct.id = BEP.branch_id
			LEFT JOIN promo PR ON PR.ID = BEP.promo_id ` + models.WebBranchEligiblePromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(pr."_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}
