package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// ICustomerOrderHeaderRepository ...
type ICustomerOrderHeaderRepository interface {
	SelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) ([]models.CustomerOrderHeader, error)
	FindAll(ctx context.Context, parameter models.CustomerOrderHeaderParameter) ([]models.CustomerOrderHeader, int, error)
	FindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (models.CustomerOrderHeader, error)
	FindByCode(c context.Context, parameter models.CustomerOrderHeaderParameter) (models.CustomerOrderHeader, error)
	CheckOut(c context.Context, model *models.CustomerOrderHeader) (*string, error)
	SyncVoid(c context.Context, model *models.CustomerOrderHeader) (*string, error)

	//apps
	AppsSelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) ([]models.CustomerOrderHeader, error)
	AppsFindAll(ctx context.Context, parameter models.CustomerOrderHeaderParameter) ([]models.CustomerOrderHeader, int, error)
	AppsFindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (models.CustomerOrderHeader, error)

	ReUpdateModifiedDate(c context.Context) (*string, error)
}

// CustomerOrderHeaderRepository ...
type CustomerOrderHeaderRepository struct {
	DB *sql.DB
}

// NewCustomerOrderHeaderRepository ...
func NewCustomerOrderHeaderRepository(DB *sql.DB) ICustomerOrderHeaderRepository {
	return &CustomerOrderHeaderRepository{DB: DB}
}

// Scan rows
func (repository CustomerOrderHeaderRepository) scanRows(rows *sql.Rows) (res models.CustomerOrderHeader, err error) {
	err = rows.Scan(
		&res.ID, &res.DocumentNo, &res.TransactionDate, &res.TransactionTime,
		&res.CustomerID, &res.CustomerName, &res.TaxCalcMethod,
		&res.SalesmanID, &res.SalesmanName, &res.PaymentTermsID, &res.PaymentTermsName,
		&res.ExpectedDeliveryDate, &res.BranchID, &res.BranchName,
		&res.PriceLIstID, &res.PriceLIstName, &res.PriceLIstVersionID, &res.PriceLIstVersionName,
		&res.Status, &res.GrossAmount, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.DiscAmount,
		&res.CustomerCode, &res.SalesmanCode, &res.CustomerAddress, &res.ModifiedDate,
		&res.VoidReasonText, &res.OrderSource,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository CustomerOrderHeaderRepository) scanRow(row *sql.Row) (res models.CustomerOrderHeader, err error) {
	err = row.Scan(
		&res.ID, &res.DocumentNo, &res.TransactionDate, &res.TransactionTime,
		&res.CustomerID, &res.CustomerName, &res.TaxCalcMethod,
		&res.SalesmanID, &res.SalesmanName, &res.PaymentTermsID, &res.PaymentTermsName,
		&res.ExpectedDeliveryDate, &res.BranchID, &res.BranchName,
		&res.PriceLIstID, &res.PriceLIstName, &res.PriceLIstVersionID, &res.PriceLIstVersionName,
		&res.Status, &res.GrossAmount, &res.TaxableAmount, &res.TaxAmount,
		&res.RoundingAmount, &res.NetAmount, &res.DiscAmount,
		&res.CustomerCode, &res.SalesmanCode, &res.CustomerAddress, &res.ModifiedDate,
		&res.VoidReasonText, &res.OrderSource,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository CustomerOrderHeaderRepository) SelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (data []models.CustomerOrderHeader, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.DateParam != "" {
		conditionString += ` AND def.modified_date > '` + parameter.DateParam + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	statement := models.CustomerOrderHeaderSelectStatement + ` ` + models.CustomerOrderHeaderWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
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
func (repository CustomerOrderHeaderRepository) FindAll(ctx context.Context, parameter models.CustomerOrderHeaderParameter) (data []models.CustomerOrderHeader, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	query := models.CustomerOrderHeaderSelectStatement + ` ` + models.CustomerOrderHeaderWhereStatement + ` ` + conditionString + `
		AND (LOWER(cus."customer_name") LIKE $1 or LOWER(def."document_no") LIKE $1) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `select 
			count(*)
			from customer_order_header def
			join customer cus on cus.id = def.cust_ship_to_id
			left join salesman s on s.id = def.salesman_id
			left join term_of_payment top on top.id = def.payment_terms_id
			left join branch b on b.id = def.branch_id
			left join price_list pl on pl.id = def.price_list_id
			left join price_list_version plv on plv.id = def.price_list_version_id ` + models.CustomerOrderHeaderWhereStatement + ` ` +
		conditionString + ` AND (LOWER(cus."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerOrderHeaderRepository) FindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (data models.CustomerOrderHeader, err error) {
	statement := models.CustomerOrderHeaderSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByID ...
func (repository CustomerOrderHeaderRepository) FindByCode(c context.Context, parameter models.CustomerOrderHeaderParameter) (data models.CustomerOrderHeader, err error) {
	statement := models.CustomerOrderHeaderSelectStatement + ` WHERE def.created_date IS not NULL AND def.document_no = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.DocumentNo)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository CustomerOrderHeaderRepository) CheckOut(c context.Context, model *models.CustomerOrderHeader) (res *string, err error) {
	statement := `
		insert into customer_order_header(
		transaction_date, transaction_time,  cust_bill_to_id, cust_ship_to_id,
		payment_terms_id, expected_delivery_date, gross_amount,disc_amount,
		taxable_amount, tax_amount, rounding_amount, net_amount,
		tax_calc_method, salesman_id,
		branch_id,price_list_id,status
		)values(
			$1,$2,$3,$4,
			$5,$6,$7,$8,
			$9,$10,$11,$12,
			$13,(select salesman_id from customer where id = $14),
			(select branch_id from customer where id = $15),$16,'draft'
		)
	RETURNING id`

	transaction, err := repository.DB.BeginTx(c, nil)
	if err != nil {
		return res, err
	}
	defer transaction.Rollback()

	fmt.Println(statement)
	err = transaction.QueryRowContext(c, statement, model.TransactionDate, model.TransactionTime, model.CustomerID, model.CustomerID,
		model.PaymentTermsID, model.ExpectedDeliveryDate, model.GrossAmount, model.DiscAmount,
		model.TaxableAmount, model.TaxAmount, model.RoundingAmount, model.NetAmount,
		model.TaxCalcMethod, model.CustomerID,
		model.CustomerID, model.PriceLIstID,
	).Scan(&res)

	if err != nil {
		return res, err
	}

	model.ID = res

	line_statement := ` select download_customer_order_line_from_cart as result_text from download_customer_order_line_from_cart($1,$2,$3)`

	err = transaction.QueryRowContext(c, line_statement,
		str.StringToInt(*model.CustomerID), str.StringToInt(*res),
		model.LineList).Scan(&res)

	if err != nil {
		return res, err
	}

	if err = transaction.Commit(); err != nil {
		return res, err
	}
	return model.ID, err
}

func (repository CustomerOrderHeaderRepository) SyncVoid(c context.Context, model *models.CustomerOrderHeader) (res *string, err error) {
	statement := `UPDATE customer_order_header SET 
	status = $1 ,void_reason_id =( select id from master_type where code = $2 )
	WHERE document_no = $3 
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement,
		model.Status, model.VoidReasonCode, model.DocumentNo).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (repository CustomerOrderHeaderRepository) ReUpdateModifiedDate(c context.Context) (res *string, err error) {
	statement := `select update_order_mod_date from update_order_mod_date(10) `
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

//apps select

// SelectAll ...
func (repository CustomerOrderHeaderRepository) AppsSelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (data []models.CustomerOrderHeader, err error) {
	conditionString := ``

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND def.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		conditionString += ` AND def.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.DateParam != "" {
		conditionString += ` AND def.modified_date > '` + parameter.DateParam + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	queryBuilder := ` select * from ( `

	queryBuilder += models.CustomerOrderHeaderSelectStatement + ` ` + models.CustomerOrderHeaderWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1 ) ` + conditionString

	queryBuilder += ` union all `

	queryBuilder += models.CustomerOrderHeaderSFASelectStatement + ` ` + models.CustomerOrderHeaderWhereStatement +
		` AND (LOWER(cus."customer_name") LIKE $1 ) ` + conditionString

	queryBuilder += ` )x ORDER BY x.transaction_date`

	statement := queryBuilder

	fmt.Println(statement)
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
func (repository CustomerOrderHeaderRepository) AppsFindAll(ctx context.Context, parameter models.CustomerOrderHeaderParameter) (data []models.CustomerOrderHeader, count int, err error) {
	conditionString := ``

	if parameter.CustomerID != "" {
		conditionString += ` AND def.cust_ship_to_id = '` + parameter.CustomerID + `'`
	}

	if parameter.UserID != "" {
		conditionString += ` AND def.branch_id in ( select branch_id from user_branch where user_id = ` + parameter.UserID + `)`
	}

	query := models.CustomerOrderHeaderSelectStatement + ` ` + models.CustomerOrderHeaderWhereStatement + ` ` + conditionString + `
		AND (LOWER(cus."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
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

	query = `select 
			count(*)
			from customer_order_header def
			join customer cus on cus.id = def.cust_ship_to_id
			left join salesman s on s.id = def.salesman_id
			left join term_of_payment top on top.id = def.payment_terms_id
			left join branch b on b.id = def.branch_id
			left join price_list pl on pl.id = def.price_list_id
			left join price_list_version plv on plv.id = def.price_list_version_id ` + models.CustomerOrderHeaderWhereStatement + ` ` +
		conditionString + ` AND (LOWER(cus."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}

// FindByID ...
func (repository CustomerOrderHeaderRepository) AppsFindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (data models.CustomerOrderHeader, err error) {
	statement := models.CustomerOrderHeaderSelectStatement + ` WHERE def.created_date IS not NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
