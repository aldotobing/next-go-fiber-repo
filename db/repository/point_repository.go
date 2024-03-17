package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointRepository ...
type IPointRepository interface {
	SelectAll(c context.Context, parameter models.PointParameter) ([]models.Point, error)
	FindAll(ctx context.Context, parameter models.PointParameter) ([]models.Point, int, error)
	FindByID(c context.Context, parameter models.PointParameter) (models.Point, error)
	GetBalance(c context.Context, parameter models.PointParameter) (models.PointGetBalance, error)
	GetBalanceUsingInvoiceDate(c context.Context, parameter models.PointParameter) (models.PointGetBalance, error)
	Add(c context.Context, model viewmodel.PointVM) (string, error)
	AddInject(c context.Context, model []viewmodel.PointVM) (string, error)
	AddWithdraw(c context.Context, in viewmodel.PointVM) (res string, err error)
	Update(c context.Context, model viewmodel.PointVM) (string, error)
	Delete(c context.Context, id string) (string, error)
	Report(c context.Context, parameter models.PointParameter) ([]models.Point, error)
	SingleAdd(c context.Context, model viewmodel.PointVM) (string, error)
}

// PointRepository ...
type PointRepository struct {
	DB *sql.DB
}

// NewPointRepository ...
func NewPointRepository(DB *sql.DB) IPointRepository {
	return &PointRepository{DB: DB}
}

// Scan rows
func (repository PointRepository) scanRows(rows *sql.Rows) (res models.Point, err error) {
	err = rows.Scan(
		&res.ID,
		&res.PointType,
		&res.PointTypeName,
		&res.InvoiceDocumentNo,
		&res.Point,
		&res.CustomerID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.ExpiredAt,

		&res.Customer.CustomerName,
		&res.Customer.Code,
		&res.Customer.CustomerBranchCode,
		&res.Customer.CustomerBranchName,
		&res.Customer.CustomerRegionName,

		&res.InvoiceDate,
		&res.Note,
	)

	return
}

// Scan row
func (repository PointRepository) scanRow(row *sql.Row) (res models.Point, err error) {
	err = row.Scan(
		&res.ID,
		&res.PointType,
		&res.PointTypeName,
		&res.InvoiceDocumentNo,
		&res.Point,
		&res.CustomerID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.ExpiredAt,

		&res.Customer.CustomerName,
		&res.Customer.Code,
		&res.Customer.CustomerBranchCode,
		&res.Customer.CustomerBranchName,
		&res.Customer.CustomerRegionName,

		&res.InvoiceDate,
		&res.Note,
	)

	return
}

// SelectAll ...
func (repository PointRepository) SelectAll(c context.Context, parameter models.PointParameter) (data []models.Point, err error) {
	var conditionString string

	if parameter.PointType != "" {
		conditionString += `AND DEF.POINT_TYPE = ` + parameter.PointType
	}

	if parameter.CustomerID != "" {
		conditionString += `AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}

	statement := models.PointSelectStatement + models.PointWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort

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
func (repository PointRepository) FindAll(ctx context.Context, parameter models.PointParameter) (data []models.Point, count int, err error) {
	var conditionString string

	if parameter.CustomerID != "" {
		conditionString += `AND DEF.CUSTOMER_ID = ` + parameter.CustomerID
	}

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += `AND DEF.CREATED_AT BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	}

	if parameter.PointType != "" {
		conditionString += `AND DEF.POINT_TYPE = ` + parameter.PointType
	}

	statement := models.PointSelectStatement + models.PointWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
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

	countQuery := `SELECT COUNT(*) FROM POINTS def ` + models.PointWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository PointRepository) FindByID(c context.Context, parameter models.PointParameter) (data models.Point, err error) {
	statement := models.PointSelectStatement + ` WHERE DEF.ID = ` + parameter.ID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// GetBalance ...
func (repository PointRepository) GetBalance(c context.Context, parameter models.PointParameter) (data models.PointGetBalance, err error) {
	var whereStatement string
	if parameter.Month != "" && parameter.Year != "" {
		whereStatement += ` AND extract(month from def.created_at) = '` + parameter.Month + `' 
		and EXTRACT(YEAR from def.created_at) = '` + parameter.Year + `'`
	}

	statement := `select coalesce(sum(case when pt."_name" = '` + models.PointTypeWithdraw + `' then DEF.point else 0 end),0) as withdraw,
		coalesce(sum(case when pt."_name" = '` + models.PointTypeCashback + `' then DEF.point else 0 end),0) as cashback,
		coalesce(sum(case when pt."_name" = '` + models.PointTypeLoyalty + `' then DEF.point else 0 end),0) as loyalty,
		coalesce(sum(case when pt."_name" = '` + models.PointTypePromo + `' then DEF.point else 0 end),0) as promo
		from points DEF
		left join point_type pt on pt.id = def.point_type 
		LEFT JOIN SALES_INVOICE_HEADER SIH ON SIH.DOCUMENT_NO = DEF.INVOICE_DOCUMENT_NO
		WHERE DEF.DELETED_AT IS NULL AND DEF.CUSTOMER_ID = ` + parameter.CustomerID + whereStatement
	row := repository.DB.QueryRowContext(c, statement)

	err = row.Scan(
		&data.Withdraw,
		&data.Cashback,
		&data.Loyalty,
		&data.Promo,
	)

	return
}

// GetBalanceUsingInvoiceDate ...
func (repository PointRepository) GetBalanceUsingInvoiceDate(c context.Context, parameter models.PointParameter) (data models.PointGetBalance, err error) {
	var whereStatement string
	if parameter.Month != "" && parameter.Year != "" {
		whereStatement += ` AND extract(month from SIH.TRANSACTION_DATE) = '` + parameter.Month + `' 
		and EXTRACT(YEAR from SIH.TRANSACTION_DATE) = '` + parameter.Year + `'`
	}

	statement := `select coalesce(sum(case when pt."_name" = '` + models.PointTypeCashback + `' then DEF.point else 0 end),0) as cashback
		from points DEF
		left join point_type pt on pt.id = def.point_type 
		LEFT JOIN SALES_INVOICE_HEADER SIH ON SIH.DOCUMENT_NO = DEF.INVOICE_DOCUMENT_NO
		WHERE DEF.DELETED_AT IS NULL AND DEF.CUSTOMER_ID = ` + parameter.CustomerID + whereStatement
	row := repository.DB.QueryRowContext(c, statement)

	err = row.Scan(
		&data.Cashback,
	)

	return
}

// Add ...
func (repository PointRepository) Add(c context.Context, in viewmodel.PointVM) (res string, err error) {
	var statementInsert string

	statementInsert += `(` + in.PointType + `, '` + in.InvoiceDocumentNo + `', '` + in.Point + `', ` + in.CustomerID + `, NOW(), NOW(), '` + in.ExpiredAt + `')`

	statement := `INSERT INTO POINTS (
			POINT_TYPE, 
			INVOICE_DOCUMENT_NO,
			POINT,
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT,
			EXPIRED_AT
		)
	VALUES ` + statementInsert

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

func (repository PointRepository) SingleAdd(c context.Context, in viewmodel.PointVM) (res string, err error) {
	var statementInsert string
	statementInsert += `(` + in.PointType + `, '` + in.InvoiceDocumentNo + `', '` + in.Point + `', ` + in.CustomerID + `, NOW(), NOW(), '` + in.ExpiredAt + `')`

	statement := `INSERT INTO POINTS (
			POINT_TYPE, 
			INVOICE_DOCUMENT_NO,
			POINT,
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT,
			EXPIRED_AT
		)
	VALUES ` + statementInsert + ` returning id `

	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return res, err
}

// AddInject ...
func (repository PointRepository) AddInject(c context.Context, in []viewmodel.PointVM) (res string, err error) {
	var statementInsert string
	if len(in) > 0 {
		for _, datum := range in {
			if statementInsert == "" {
				statementInsert += `(` + datum.PointType + `, '` + datum.InvoiceDocumentNo + `', '` + datum.Point + `', ` + datum.CustomerID + `, NOW(), NOW(), '` + datum.ExpiredAt + `', '` + datum.Note + `')`
			} else {
				statementInsert += `, (` + datum.PointType + `, '` + datum.InvoiceDocumentNo + `', '` + datum.Point + `', ` + datum.CustomerID + `, NOW(), NOW(), '` + datum.ExpiredAt + `', '` + datum.Note + `')`
			}
		}
	}
	statement := `INSERT INTO POINTS (
			POINT_TYPE, 
			INVOICE_DOCUMENT_NO,
			POINT,
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT,
			EXPIRED_AT, 
			NOTE
		)
	VALUES ` + statementInsert

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

// AddWithdraw ...
func (repository PointRepository) AddWithdraw(c context.Context, in viewmodel.PointVM) (res string, err error) {
	statement := `INSERT INTO POINTS (
			POINT_TYPE,
			POINT,
			CUSTOMER_ID,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES (
		$1, $2, $3, now(), now()
	)`

	err = repository.DB.QueryRowContext(c, statement, in.PointType, in.Point, in.CustomerID).Err()

	return
}

// Update ...
func (repository PointRepository) Update(c context.Context, in viewmodel.PointVM) (res string, err error) {
	statement := `UPDATE POINTS SET 
		POINT_TYPE = $1, 
		INVOICE_DOCUMENT_NO = $2, 
		POINT = $3, 
		CUSTOMER_ID = $4,
		UPDATED_AT = now()
	WHERE id = $5
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		in.PointType,
		in.InvoiceDocumentNo,
		in.Point,
		in.CustomerID,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository PointRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINTS SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}

func (repository PointRepository) Report(c context.Context, parameter models.PointParameter) (data []models.Point, err error) {
	var conditionString string

	if parameter.StartDate != "" && parameter.EndDate != "" {
		conditionString += ` AND sih.transaction_date between '` + parameter.StartDate + `' and '` + parameter.EndDate + `'`
	}

	if parameter.BranchID != "" {
		conditionString += ` AND B.ID = ` + parameter.BranchID
	}

	if parameter.RegionGroupID != "" {
		conditionString += ` AND R.GROUP_ID = ` + parameter.RegionGroupID
	}
	if parameter.RegionID != "" {
		conditionString += ` AND R.ID = ` + parameter.RegionID
	}

	statement := `select b.branch_code, b."_name", r._name, r.group_name, 
		pt.code, pt."_name", 
		p.invoice_document_no, sih.net_amount, p.point, sih.transaction_date
		from points p
		left join sales_invoice_header sih on sih.document_no = p.invoice_document_no 
		left join customer c on c.id = p.customer_id 
		left join branch b on b.id = c.branch_id 
		left join region r on r.id = b.region_id
		left join partner pt on pt.id = c.partner_id
		WHERE P.DELETED_AT IS NULL AND P.POINT_TYPE = 2 ` + conditionString + `
		order by branch_code asc;`

	rows, err := repository.DB.QueryContext(c, statement)

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Point
		err = rows.Scan(
			&temp.Branch.Code, &temp.Branch.Name, &temp.Region.Name, &temp.Region.GroupName,
			&temp.Partner.Code, &temp.Partner.PartnerName,
			&temp.InvoiceDocumentNo, &temp.SalesInvoice.NetAmount, &temp.Point, &temp.SalesInvoice.TrasactionDate,
		)

		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}
