package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IDashboardWebRepository ...
type IDashboardWebRepository interface {
	GetData(c context.Context, parameter models.DashboardWebParameter) ([]models.DashboardWeb, error)
	GetDataByGroupID(c context.Context, parameter models.DashboardWebParameter) ([]models.DashboardWebRegionDetail, error)
	GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) ([]models.DashboardWebRegionDetail, error)
	GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, int, error)
	GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)
	GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)
	GetOmzetValue(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.OmzetValueModel, error)
	GetOmzetValueByGroupID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) ([]models.OmzetValueModel, error)
	GetOmzetValueByRegionID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) ([]models.OmzetValueModel, error)
	GetOmzetValueByBranchID(ctx context.Context, parameter models.DashboardWebBranchParameter, branchID string) ([]models.OmzetValueBranchModel, error)
	GetOmzetValueByCustomerID(ctx context.Context, parameter models.DashboardWebBranchParameter, customerID string) (res []models.OmzetValueModel, err error)
}

// DashboardWebRepository ...
type DashboardWebRepository struct {
	DB *sql.DB
}

// NewDashboardWebRepository ...
func NewDashboardWebRepository(DB *sql.DB) IDashboardWebRepository {
	return &DashboardWebRepository{DB: DB}
}

// Scan row
func (repository DashboardWebRepository) scanRow(row *sql.Row) (res models.DashboardWeb, err error) {
	err = row.Scan(
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalVisitUser,
		&res.CustomerCountRepeatOrder, &res.TotalActiveOutlet,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRows(rows *sql.Rows) (res models.DashboardWeb, err error) {
	err = rows.Scan(
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalVisitUser,
		&res.CustomerCountRepeatOrder, &res.TotalActiveOutlet,
		&res.TotalOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanByRegionIDRows(rows *sql.Rows) (res models.DashboardWebRegionDetail, err error) {
	err = rows.Scan(
		&res.RegionID, &res.RegionName,
		&res.TotalVisitUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalRegisteredUser, &res.CustomerCountRepeatOrder, &res.TotalOutlet,
		&res.TotalActiveOutlet,
		&res.TotalInvoice,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRegionDetailRows(rows *sql.Rows) (res models.DashboardWebRegionDetail, err error) {
	err = rows.Scan(
		&res.BranchID, &res.BranchCode, &res.BranchName, &res.RegionID, &res.RegionName,
		&res.RegionGroupID, &res.RegionGroupName,
		&res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalVisitUser, &res.CustomerCountRepeatOrder,
		&res.TotalActiveOutlet,
		&res.TotalOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanBranchCustomerDetailRows(rows *sql.Rows) (res models.DashboardWebBranchDetail, err error) {
	err = rows.Scan(
		&res.CustomerID, &res.CustomerName, &res.CustomerCode,
		&res.CustomerBranchName, &res.CustomerBranchCode,
		&res.CustomerRegionName, &res.CustomerRegionGroupName,
		&res.CustomerTypeName,
		&res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.CustomerClassName, &res.CustomerCityName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) (data []models.DashboardWeb, err error) {
	statement := models.DashboardWebSelectStatement
	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))

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

// GetDataByGroupID ...
func (repository DashboardWebRepository) GetDataByGroupID(c context.Context, parameter models.DashboardWebParameter) (data []models.DashboardWebRegionDetail, err error) {
	statement := models.DashboardWebSelectByGroupIDStatement

	statement = strings.ReplaceAll(statement, "{START_DATE}", parameter.StartDate)
	statement = strings.ReplaceAll(statement, "{END_DATE}", parameter.EndDate)

	var whereStatement string
	if parameter.GroupID != "" {
		whereStatement = `WHERE r.group_id = ` + parameter.GroupID
	}

	statement = strings.ReplaceAll(statement, "{WHERE_STATEMENT}", whereStatement)

	rows, err := repository.DB.QueryContext(c, statement)

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanByRegionIDRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository DashboardWebRepository) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (data []models.DashboardWebRegionDetail, err error) {
	statement := models.DashboardWebRegionDetailSelectStatement

	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.GroupID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanRegionDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository DashboardWebRepository) GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, count int, err error) {

	query := models.DashboardWebBranchDetailSelectStatement + ` OFFSET $4 LIMIT $5`
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate),
		parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = ` select count(*) from os_fetch_dashborad_branchcustomerdata($1::integer,$2,$3,null,null,null) `
	err = repository.DB.QueryRow(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate)).Scan(&count)
	return data, count, err
}

func (repository DashboardWebRepository) GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebBranchDetailSelectStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository DashboardWebRepository) GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebBranchDetailSelectWithUserIDStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.UserID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, err
}

func (repo DashboardWebRepository) GetOmzetValue(ctx context.Context, parameter models.DashboardWebBranchParameter) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}

	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND sil.category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	query := `select r.group_id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null 
			and coh.id is not null` + whereStatement + `
			group by r.group_id
			order by r.group_id asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.RegionGroupID, &temp.TotalGrossAmount, &temp.TotalNettAmount, &temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) GetOmzetValueByGroupID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND sil.category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	if groupID != "" && groupID != "0" {
		whereStatement += ` AND r.group_id = '` + groupID + `'`
	}

	query := `select r.id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null
			and coh.id is not null` + whereStatement + `
			group by r.id
			order by r.id asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.RegionID, &temp.TotalGrossAmount, &temp.TotalNettAmount, &temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) GetOmzetValueByRegionID(ctx context.Context, parameter models.DashboardWebBranchParameter, regionID string) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND sil.category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	if regionID != "" && regionID != "0" {
		whereStatement += ` AND r.id = '` + regionID + `'`
	}

	query := `select sih.branch_id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume,
			coalesce(count(distinct(sih.cust_bill_to_id)),0) as total_active_customer
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null
			and coh.id is not null` + whereStatement + `
			group by sih.branch_id, r.id
			order by sih.branch_id asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.BranchID,
			&temp.TotalGrossAmount,
			&temp.TotalNettAmount,
			&temp.TotalQuantity,
			&temp.TotalActiveCustomer)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) GetOmzetValueByBranchID(ctx context.Context, parameter models.DashboardWebBranchParameter, branchID string) (res []models.OmzetValueBranchModel, err error) {
	var whereStatement, withWhereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
		withWhereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
		withWhereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND sil.category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	if branchID != "" && branchID != "0" {
		whereStatement += ` AND c.branch_id = '` + branchID + `'`
		withWhereStatement += ` AND c.branch_id = '` + branchID + `'`
	}

	querySelect := `with customerSelected as(
		select 
			reg."_name" as region_name,
			reg.group_name as region_group_name,
			b."_name" as branch_name, 
			b.branch_code as branch_code,
			c.id as customer_id, 
		  	CT._NAME as customer_type_name,
		  	DIST._NAME AS CUSTomer_DISTRICT_NAME, 
			CTY._NAME AS CUSTomer_CITY_NAME,
		  	cl._name as customer_level_name
		from customer c
			left join customer_order_header coh on coh.cust_bill_to_id = c.id
			left join sales_invoice_header sih on sih.transaction_source_document_no = coh.document_no
			left join branch b on b.id = c.branch_id
			LEFT JOIN REGION REG ON REG.ID = B.REGION_ID
			left JOIN CITY CTY ON CTY.ID = C.CUSTOMER_CITY_ID
	  		left JOIN DISTRICT DIST ON DIST.ID = C.CUSTOMER_DISTRICT_ID
	  		left join customer_level cl on cl.id = c.customer_level_id
	  		LEFT JOIN CUSTOMER_TYPE CT ON CT.ID = C.CUSTOMER_TYPE_ID
		where c.show_in_apps = 1  
			and coh.status in ('finish', 'submitted') 
			and sih.id is not null ` + withWhereStatement + `
		group by c.id, b.id, cty.id, dist.id, cl.id, ct.id, reg.id
	)
	select c.id, cs.region_name, cs.region_group_name, cs.branch_name, cs.branch_code, c.customer_name, c.customer_code, cs.customer_type_name, cs.customer_district_name, 
		cs.customer_city_name, customer_level_name, 
		coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
        coalesce(sum(sil.net_amount),0) as total_nett_amount, 
        coalesce(sum(sil.qty),0) as total_volume
    from sales_invoice_header sih 
		left join sales_invoice_line sil on sil.header_id = sih.id 
		left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
		left join customer c on c.id = coh.cust_bill_to_id
		left join branch b on b.id = sih.branch_id  
		left join region r on r.id = b.region_id
		left join customerSelected CS on cs.customer_id = c.id
	WHERE sih.transaction_date is not null 
		and coh.id is not null ` + whereStatement + `
	group by c.id, cs.region_name,cs.region_group_name, cs.branch_name, cs.branch_code, cs.customer_type_name, cs.customer_district_name, 
		cs.customer_city_name, customer_level_name
	order by c.id asc`

	rows, err := repo.DB.Query(querySelect)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueBranchModel
		err = rows.Scan(&temp.CustomerID,
			&temp.RegionName,
			&temp.RegionGroupName,
			&temp.BranchName,
			&temp.BranchCode,
			&temp.CustomerName,
			&temp.CustomerCode,
			&temp.CustomerType,
			&temp.ProvinceName,
			&temp.CityName,
			&temp.CustomerLevel,
			&temp.TotalGrossAmount,
			&temp.TotalNettAmount,
			&temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) GetOmzetValueByCustomerID(ctx context.Context, parameter models.DashboardWebBranchParameter, customerID string) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND sil.category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	if customerID != "" && customerID != "0" {
		whereStatement += ` AND c.id = '` + customerID + `'`
	}

	querySelect := `select i.id, i."_name",
		coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
		coalesce(sum(sil.net_amount),0) as total_nett_amount, 
		coalesce(sum(sil.qty),0) as total_volume
	from sales_invoice_header sih 
		left join sales_invoice_line sil on sil.header_id = sih.id 
		left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
		left join item i on i.id = sil.item_id
		left join customer c on c.id = coh.cust_bill_to_id
		left join branch b on b.id = sih.branch_id  
		left join region r on r.id = b.region_id
	WHERE sih.transaction_date is not null 
		and coh.id is not null AND sih.transaction_date BETWEEN '2023-05-01' AND '2023-05-30'
		` + whereStatement + `
	group by i.id
	order by i.id asc`

	rows, err := repo.DB.Query(querySelect)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.ItemID,
			&temp.ItemName,
			&temp.TotalGrossAmount,
			&temp.TotalNettAmount,
			&temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}
