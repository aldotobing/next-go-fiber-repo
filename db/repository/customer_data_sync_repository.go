package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ICustomerDataSyncRepository ...
type ICustomerDataSyncRepository interface {
	Add(c context.Context, model *models.CustomerDataSync) (*string, error)
	Edit(c context.Context, model *models.CustomerDataSync) (*string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
	FindByCode(c context.Context, parameter models.CustomerDataSyncParameter) (models.CustomerDataSync, error)
}

// CustomerDataSyncRepository ...
type CustomerDataSyncRepository struct {
	DB *sql.DB
}

// NewCustomerDataSyncRepository ...
func NewCustomerDataSyncRepository(DB *sql.DB) ICustomerDataSyncRepository {
	return &CustomerDataSyncRepository{DB: DB}
}

// Scan rows
func (repository CustomerDataSyncRepository) scanRows(rows *sql.Rows) (res models.CustomerDataSync, err error) {
	err = rows.Scan(
		&res.ID, &res.PartnerID, &res.Code, &res.Name,
		&res.Address, &res.PhoneNo, &res.CustomerType, &res.PriceListCode,
		&res.CountryCode, &res.CityCode, &res.DistrictCode, &res.SubDistrictCode,
		&res.ProvinceCode, &res.TermOfPaymentCode, &res.SalesmanCode, &res.BranchID,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CustomerDataSyncRepository) scanRow(row *sql.Row) (res models.CustomerDataSync, err error) {
	err = row.Scan(
		&res.ID, &res.PartnerID, &res.Code, &res.Name,
		&res.Address, &res.PhoneNo, &res.CustomerType, &res.PriceListCode,
		&res.CountryCode, &res.CityCode, &res.DistrictCode, &res.SubDistrictCode,
		&res.ProvinceCode, &res.TermOfPaymentCode, &res.SalesmanCode, &res.BranchID,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository CustomerDataSyncRepository) FindByCode(c context.Context, parameter models.CustomerDataSyncParameter) (data models.CustomerDataSync, err error) {
	statement := models.CustomerDataSyncSelectStatement + ` WHERE c.created_date IS NOT NULL AND p.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository CustomerDataSyncRepository) Add(c context.Context, model *models.CustomerDataSync) (res *string, err error) {

	fmt.Println("insert data post")
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO partner (device_id,code,_name, phone_no, address,
		created_date, modified_date,
		country_id,city_id,district_id,subdistrict_id,company_id)
	VALUES (433, $1, $2, $3, $4, 
		now(), now() ,
		(select id from country where code =$5), 
		(select id from city where code = $6), 
		( select id from district where code = $7), 
		( select id from subdistrict where code = $8),
		2
		) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Code, model.Name, model.PhoneNo, model.Address,
		model.CountryCode, model.CityCode, model.DistrictCode, model.SubDistrictCode,
	).Scan(&res)
	if err != nil {
		fmt.Println("error update partener")
		return res, err
	}

	customerstatement := ` insert into customer(device_id,partner_id,customer_name,customer_address,customer_phone,
		created_date, modified_date,
		payment_terms_id,price_list_id,salesman_id,customer_level_id,
		branch_id,customer_code)
	values
	(433,$1,$2,$3,$4,now(),now(),
		(select id from term_of_payment where code =$5),
		(select id from price_list where code=$6),
		(select id from salesman where partner_id =(select id from partner where code =$7)),
		(select id from customer_level where code = $8),
		$9,$10

		) RETURNING id `

	var rescus string
	err = transaction.QueryRowContext(c, customerstatement,
		&res, model.Name, model.Address, model.PhoneNo,
		model.TermOfPaymentCode, model.PriceListCode, model.SalesmanCode,
		model.CustomerLevelCode, model.BranchID, model.Code,
	).Scan(&rescus)

	if err != nil {
		return res, err
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Edit ...
func (repository CustomerDataSyncRepository) Edit(c context.Context, model *models.CustomerDataSync) (res *string, err error) {
	fmt.Println("update data post")
	statement := `update partner set _name = $1,phone_no =  $2,address=$3 where code = $4 returning id `

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	err = transaction.QueryRowContext(c, statement, model.Name, model.PhoneNo,
		model.Address, model.Code).Scan(&res)
	if err != nil {
		fmt.Println("error update partener")
		return res, err
	}
	model.PartnerID = res
	var rescus string
	customerstatement := ` update customer set customer_name = $1,customer_address=$2,
	payment_terms_id = (select id from term_of_payment where code =$3) ,
	price_list_id= (select id from price_list where code=$4),
	salesman_id =(select id from salesman where partner_id =(select id from partner where code =$5)),
	customer_level_id =(select id from customer_level where code = $6),
	branch_id = $7
	where partner_id = (select id from partner where code = $8)
	returning id `

	err = transaction.QueryRowContext(c, customerstatement,
		model.Name, model.Address,
		model.TermOfPaymentCode, model.PriceListCode, model.SalesmanCode,
		model.CustomerLevelCode, model.BranchID,
		model.Code,
	).Scan(&rescus)

	if err != nil {
		return res, err
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (repository CustomerDataSyncRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE mp_CustomerDataSync SET updated_at_CustomerDataSync = $1, deleted_at_CustomerDataSync = $2 WHERE id_CustomerDataSync = $3 RETURNING id_CustomerDataSync`

	err = repository.DB.QueryRowContext(c, statement, now, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
