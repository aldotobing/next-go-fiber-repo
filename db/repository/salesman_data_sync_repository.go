package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISalesmanDataSyncRepository ...
type ISalesmanDataSyncRepository interface {
	Add(c context.Context, model *models.SalesmanDataSync) (*string, error)
	Edit(c context.Context, model *models.SalesmanDataSync) (*string, error)
	Delete(c context.Context, id string, now time.Time) (string, error)
	FindByCode(c context.Context, parameter models.SalesmanDataSyncParameter) (models.SalesmanDataSync, error)
}

// SalesmanDataSyncRepository ...
type SalesmanDataSyncRepository struct {
	DB *sql.DB
}

// NewSalesmanDataSyncRepository ...
func NewSalesmanDataSyncRepository(DB *sql.DB) ISalesmanDataSyncRepository {
	return &SalesmanDataSyncRepository{DB: DB}
}

// Scan rows
func (repository SalesmanDataSyncRepository) scanRows(rows *sql.Rows) (res models.SalesmanDataSync, err error) {
	err = rows.Scan(
		&res.ID, &res.PartnerID, &res.Code, &res.Name,
		&res.Address, &res.PhoneNo, &res.SalesmanType, &res.EffectiveSalesman,
		&res.BranchID,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository SalesmanDataSyncRepository) scanRow(row *sql.Row) (res models.SalesmanDataSync, err error) {
	err = row.Scan(
		&res.ID, &res.PartnerID, &res.Code, &res.Name,
		&res.Address, &res.PhoneNo, &res.SalesmanType, &res.EffectiveSalesman,
		&res.BranchID,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByCode ...
func (repository SalesmanDataSyncRepository) FindByCode(c context.Context, parameter models.SalesmanDataSyncParameter) (data models.SalesmanDataSync, err error) {
	statement := models.SalesmanDataSyncSelectStatement + ` WHERE s.created_date IS NOT NULL AND p.code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository SalesmanDataSyncRepository) Add(c context.Context, model *models.SalesmanDataSync) (res *string, err error) {

	fmt.Println("insert data post")
	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	statement := `INSERT INTO partner (device_id,code,_name, phone_no, address,
		created_date, modified_date,company_id)
	VALUES (433, $1, $2, $3, $4, 
		now(), now() ,
		2
		) RETURNING id`

	err = transaction.QueryRowContext(c, statement, model.Code, model.Name, model.PhoneNo, model.Address).Scan(&res)
	if err != nil {
		fmt.Println("error update partener")
		return res, err
	}

	customerstatement := ` insert into salesman(partner_id,salesman_code,salesman_name,salesman_phone_no,
		created_date, modified_date,
		effective_salesman,salesman_type_id,days_to_generate,
		branch_id)
	values
	($1,$2,$3,$4,now(),now(),
		$5, $6, 7, $7

		) RETURNING id `

	var rescus string
	err = transaction.QueryRowContext(c, customerstatement,
		&res, model.Code, model.Name, model.PhoneNo,
		model.EffectiveSalesman, model.SalesmanType, model.BranchID,
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
func (repository SalesmanDataSyncRepository) Edit(c context.Context, model *models.SalesmanDataSync) (res *string, err error) {
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
	customerstatement := ` update salesman set salesman_name = $1,salesman_phone_no=$2,
	modified_date =now(),effective_salesman = $3,salesman_type_id= $4,
	branch_id = $5
	where partner_id = (select id from partner where code = $6)
	returning id `

	err = transaction.QueryRowContext(c, customerstatement,
		model.Name, model.PhoneNo,
		model.EffectiveSalesman, model.SalesmanType, model.BranchID,
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
func (repository SalesmanDataSyncRepository) Delete(c context.Context, id string, now time.Time) (res string, err error) {
	statement := `UPDATE mp_SalesmanDataSync SET updated_at_SalesmanDataSync = $1, deleted_at_SalesmanDataSync = $2 WHERE id_SalesmanDataSync = $3 RETURNING id_SalesmanDataSync`

	err = repository.DB.QueryRowContext(c, statement, now, now, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
