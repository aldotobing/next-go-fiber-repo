package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IWebPromo ...
type IWebPromo interface {
	SelectAll(c context.Context, parameter models.WebPromoParameter) ([]models.WebPromo, error)
	FindAll(ctx context.Context, parameter models.WebPromoParameter) ([]models.WebPromo, int, error)
	Add(c context.Context, parameter *models.WebPromo) (*string, error)
	Edit(c context.Context, parameter *models.WebPromo) (*string, error)
	Delete(c context.Context, id string) (string, error)
	FindByID(c context.Context, parameter models.WebPromoParameter) (models.WebPromo, error)
	// 	Edit(c context.Context, model *models.WebPromo) (*string, error)
	// 	EditAddress(c context.Context, model *models.WebPromo) (*string, error)
}

// WebPromo ...
type WebPromo struct {
	DB *sql.DB
}

// NewWebPromo ...
func NewWebPromoRepository(DB *sql.DB) IWebPromo {
	return &WebPromo{DB: DB}
}

// Scan rows
func (repository WebPromo) scanRows(rows *sql.Rows) (res models.WebPromo, err error) {
	err = rows.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
		&res.ShowInApp,
	)
	if err != nil {

		return res, err
	}
	return res, nil
}

// Scan row
func (repository WebPromo) scanRow(row *sql.Row) (res models.WebPromo, err error) {
	err = row.Scan(
		&res.ID,
		&res.Code,
		&res.PromoName,
		&res.PromoDescription,
		&res.PromoUrlBanner,
		&res.StartDate,
		&res.EndDate,
		&res.Active,
		&res.ShowInApp,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository WebPromo) SelectAll(c context.Context, parameter models.WebPromoParameter) (data []models.WebPromo, err error) {
	conditionString := ``

	if (parameter.StartDate != "") && (parameter.EndDate != "") {
		conditionString += ` AND start_date >= ` + `'` +
			parameter.StartDate + `'::date` + ` AND end_date <= ` + `'` + parameter.EndDate + `'::date` +
			` + INTERVAL ` + `'1 MONTH' `
	}

	statement := models.WebPromoSelectStatement + ` ` + models.WebPromoWhereStatement +
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
func (repository WebPromo) FindAll(ctx context.Context, parameter models.WebPromoParameter) (data []models.WebPromo, count int, err error) {
	conditionString := ``

	if parameter.ID != "" {
		conditionString += ` AND cus.id = '` + parameter.ID + `'`
	}

	query := models.WebPromoSelectStatement + ` ` + models.WebPromoWhereStatement + ` ` + conditionString + `
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

	query = `SELECT COUNT(*) FROM "customer" c ` + models.WebPromoWhereStatement + ` ` +
		conditionString + ` AND (LOWER(c."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository WebPromo) FindByID(c context.Context, parameter models.WebPromoParameter) (data models.WebPromo, err error) {
	statement := models.WebPromoSelectStatement + ` WHERE PC.ID = $1 `
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	fmt.Println("Promo find by ID : " + statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Edit ...
func (repository WebPromo) Edit(c context.Context, model *models.WebPromo) (res *string, err error) {
	statement := `UPDATE promo SET
		_name = $1,
		description = $2,
		url_banner = $3,
		show_in_app = $4,
		start_date = $5,
		end_date = $6,
		active = $7,
		code = $8
	WHERE id = $9
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.PromoName,
		model.PromoDescription,
		model.PromoUrlBanner,
		model.ShowInApp,
		model.StartDate,
		model.EndDate,
		model.Active,
		model.Code,
		model.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	if *model.CustomerTypeIdList != "" {
		customerTypeDeleteStatement := `DELETE FROM customer_type_eligible_promo where promo_id = $1`
		err = repository.DB.QueryRowContext(c, customerTypeDeleteStatement, *res).Err()
		if err != nil {
			return
		}

		customerTypeIDArr := strings.Split(*model.CustomerTypeIdList, ",")

		var customerTypeIDValuesStatement string
		for _, datum := range customerTypeIDArr {
			if customerTypeIDValuesStatement == "" {
				customerTypeIDValuesStatement += `(` + datum + `, ` + *res + `, now(), now())`
			} else {
				customerTypeIDValuesStatement += `, (` + datum + `, ` + *res + `, now(), now())`
			}
		}
		customerTypeUpdateStatement := `insert into customer_type_eligible_promo 
		(customer_type_id, promo_id, created_date, modified_date)
		Values ` + customerTypeIDValuesStatement

		err = repository.DB.QueryRowContext(c, customerTypeUpdateStatement).Err()
		if err != nil {
			return
		}
	}

	if *model.RegionAreaIdList != "" {
		regionAreaDeleteStatement := `DELETE FROM region_area_eligible_promo where promo_id = $1`
		err = repository.DB.QueryRowContext(c, regionAreaDeleteStatement, *res).Err()
		if err != nil {
			return
		}

		regionAreaIDArr := strings.Split(*model.RegionAreaIdList, ",")

		var regionAreaValuesStatement string
		for _, datum := range regionAreaIDArr {
			if regionAreaValuesStatement == "" {
				regionAreaValuesStatement += `(` + datum + `, ` + *res + `, now(), now())`
			} else {
				regionAreaValuesStatement += `, (` + datum + `, ` + *res + `, now(), now())`
			}
		}
		regionAreaUpdateStatement := `insert into region_area_eligible_promo 
		(region_id, promo_id, created_date, modified_date)
		Values ` + regionAreaValuesStatement

		err = repository.DB.QueryRowContext(c, regionAreaUpdateStatement).Err()
		if err != nil {
			return
		}
	}

	if *model.CustomerLevelIdList != "" {
		regionAreaDeleteStatement := `DELETE FROM customer_level_eligible_promo where promo_id = $1`
		err = repository.DB.QueryRowContext(c, regionAreaDeleteStatement, *res).Err()
		if err != nil {
			return
		}

		regionAreaIDArr := strings.Split(*model.CustomerLevelIdList, ",")

		var customerLevelValuesStatement string
		for _, datum := range regionAreaIDArr {
			if customerLevelValuesStatement == "" {
				customerLevelValuesStatement += `(` + datum + `, ` + *res + `, now(), now())`
			} else {
				customerLevelValuesStatement += `, (` + datum + `, ` + *res + `, now(), now())`
			}
		}
		customerLevelUpdateStatement := `insert into customer_level_eligible_promo 
		(customer_level_id, promo_id, created_date, modified_date)
		Values ` + customerLevelValuesStatement

		_, err = repository.DB.Query(customerLevelUpdateStatement)
		if err != nil {
			return
		}
	}
	return res, err
}

func (repository WebPromo) Add(c context.Context, model *models.WebPromo) (res *string, err error) {

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO promo (code, _name, description, url_banner,
		start_date, end_date, active, show_in_app)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Code, model.PromoName, model.PromoDescription, model.PromoUrlBanner,
		model.StartDate, model.EndDate, 1, model.ShowInApp).Scan(&res)

	fmt.Println("PROMO INSERT : " + statement)

	if err != nil {
		fmt.Println("INSERT PROMO BERHASIL! :)")
		return res, err
	}

	PromoId := &res

	parts := strings.Split(*model.CustomerTypeIdList, ",")
	if len(parts) >= 1 && parts[0] != "" {
		for pi, _ := range parts {
			linestatement := `INSERT INTO customer_type_eligible_promo 
			(customer_type_id, promo_id, created_date, modified_date)
					VALUES ($1, $2, now(), now()) RETURNING id`
			var resLine string
			err = transaction.QueryRowContext(c, linestatement, parts[pi], PromoId).Scan(&resLine)
			if err != nil {
				return res, err
			}
		}
	}

	regionparts := strings.Split(*model.RegionAreaIdList, ",")
	if len(regionparts) >= 1 && regionparts[0] != "" {
		for pi, _ := range regionparts {
			linestatement := `INSERT INTO region_area_eligible_promo 
			(region_id, promo_id, created_date, modified_date)
					VALUES ($1, $2, now(), now()) RETURNING id`
			var resregionIDLine string
			err = transaction.QueryRowContext(c, linestatement, regionparts[pi], PromoId).Scan(&resregionIDLine)
			if err != nil {
				return res, err
			}
		}
	}

	if *model.CustomerLevelIdList != "" {
		customerLevelIDs := strings.Split(*model.CustomerLevelIdList, ",")
		var customerLevelValuesStatement string
		for _, datum := range customerLevelIDs {
			if customerLevelValuesStatement == "" {
				customerLevelValuesStatement += `(` + datum + `, ` + *res + `, now(), now())`
			} else {
				customerLevelValuesStatement += `, (` + datum + `, ` + *res + `, now(), now())`
			}
		}

		insertStatement := `insert into customer_level_eligible_promo 
		(customer_level_id, promo_id, created_date, modified_date)
		Values ` + customerLevelValuesStatement

		err = repository.DB.QueryRowContext(c, insertStatement).Err()
		if err != nil {
			return
		}
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (repository WebPromo) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE promo set active = 0 where id= $1 RETURNING id `

	err = repository.DB.QueryRowContext(c, statement, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
