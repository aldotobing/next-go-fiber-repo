package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IPromoContent ...
type IPromoContent interface {
	SelectAll(c context.Context, parameter models.PromoContentParameter) ([]models.PromoContent, error)
	FindAll(ctx context.Context, parameter models.PromoContentParameter) ([]models.PromoContent, int, error)
	Add(c context.Context, parameter *models.PromoContent) (*string, error)
	Delete(c context.Context, id string) (string, error)
	// 	Edit(c context.Context, model *models.PromoContent) (*string, error)
	// 	EditAddress(c context.Context, model *models.PromoContent) (*string, error)
}

// PromoContent ...
type PromoContent struct {
	DB *sql.DB
}

// NewPromoContent ...
func NewPromoContentRepository(DB *sql.DB) IPromoContent {
	return &PromoContent{DB: DB}
}

// Scan rows
func (repository PromoContent) scanRows(rows *sql.Rows) (res models.PromoContent, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository PromoContent) scanRow(row *sql.Row) (res models.PromoContent, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository PromoContent) SelectAll(c context.Context, parameter models.PromoContentParameter) (data []models.PromoContent, err error) {
	conditionString := ``

	if (parameter.StartDate != "") && (parameter.EndDate != "") {
		conditionString += ` AND start_date >= ` + `'` +
			parameter.StartDate + `'::date` + ` AND end_date <= ` + `'` + parameter.EndDate + `'::date` +
			` + INTERVAL ` + `'1 MONTH' `
	} else {
		conditionString += ` AND now()::date BETWEEN PC.START_DATE AND END_DATE `
	}

	if parameter.CustomerTypeId != "" {
		conditionString += ` AND (PC.ID IN (SELECT promo_id FROM customer_type_eligible_promo ctep WHERE customer_type_id = ` + parameter.CustomerTypeId + `) ` +
			` OR PC.ID NOT IN (SELECT promo_id FROM customer_type_eligible_promo)) `
	}

	if parameter.CustomerLevelID != "" {
		conditionString += ` AND PC.ID IN (SELECT pc2.id 
			FROM promo pc2
			left join customer_level_eligible_promo ctep on ctep.promo_id = pc2.id
			WHERE ctep.customer_level_id = ` + parameter.CustomerLevelID + ` or ctep.id is null)`
	}

	if parameter.BranchID != "" {
		conditionString += ` AND PC.ID IN (SELECT pc2.id 
			FROM promo pc2
			left join branch_eligible_promo BEP on BEP.promo_id = pc2.id
			WHERE BEP.branch_id = ` + parameter.BranchID + ` or BEP.id is null)`
	}

	if parameter.RegionID != "" {
		conditionString += ` AND PC.ID IN (SELECT pc2.id 
			FROM promo pc2
			left join region_area_eligible_promo REP on REP.promo_id = pc2.id
			WHERE REP.region_id = ` + parameter.RegionID + ` or REP.id is null)`
	}

	statement := models.PromoContentSelectStatement + ` ` + models.PromoContentWhereStatement +
		` AND (LOWER(pc._name) LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

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
func (repository PromoContent) FindAll(ctx context.Context, parameter models.PromoContentParameter) (data []models.PromoContent, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	if parameter.CustomerTypeId != "" {
		conditionString += ` AND (PC.ID IN (SELECT promo_id FROM customer_type_eligible_promo ctep WHERE customer_type_id = ` + parameter.CustomerTypeId + `) ` +
			` OR PC.ID NOT IN (SELECT promo_id FROM customer_type_eligible_promo)) `
	}

	query := models.PromoContentSelectStatement + ` ` + models.PromoContentWhereStatement + ` ` + conditionString + `
		AND (LOWER(cus."customer_name") LIKE $1  )` + `GROUP BY ` + `ORDER BY` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.PromoContentWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository PromoContent) FindByID(c context.Context, parameter models.PromoContentParameter) (data models.PromoContent, err error) {
	statement := models.PromoContentSelectStatement + ` WHERE pc.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// // Edit ...
// func (repository PromoContent) Edit(c context.Context, model *models.PromoContent) (res *string, err error) {
// 	statement := `UPDATE customer SET
// 	customer_name = $1,
// 	customer_address = $2,
// 	customer_phone = $3,
// 	customer_email = $4,
// 	customer_cp_name = $5,
// 	customer_profile_picture = $6
// 	WHERE id = $7
// 	RETURNING id`
// 	err = repository.DB.QueryRowContext(c, statement,
// 		model.PromoContentName,
// 		model.PromoContentAddress,
// 		model.PromoContentPhone,
// 		model.PromoContentEmail,
// 		model.PromoContentCpName,
// 		model.PromoContentProfilePicture,
// 		model.ID).Scan(&res)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

func (repository PromoContent) Add(c context.Context, model *models.PromoContent) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO promo (code, _name, description, url_banner,
		start_date, end_date, active)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Code, model.PromoName, model.PromoDescription, model.PromoUrlBanner,
		model.StartDate, model.EndDate, 1).Scan(&res)

	fmt.Println("PROMO INSERT : " + statement)

	if err != nil {
		fmt.Println("INSERT PROMO BERHASIL! :)")
		return res, err
	}

	PromoId := &res

	parts := strings.Split(*model.CustomerTypeIdList, ",")
	if len(parts) >= 1 {
		for pi, _ := range parts {
			linestatement := `INSERT INTO customer_type_eligible_promo (
				customer_type_id, promo_id, created_date, modified_date)
					VALUES ($1, $2, now(), now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, parts[pi], PromoId).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository PromoContent) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE promo set active = 0 where id= $1 RETURNING id `

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
