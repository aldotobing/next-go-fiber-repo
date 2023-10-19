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
	GetUserByRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) ([]models.DashboardWebBranchDetail, error)
	GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, int, error)
	GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)

	GetAllReportBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)

	GetAllBranchDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebGetWithUserID, err error)
	GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebGetWithUserID, error)

	GetOmzetValue(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.OmzetValueModel, error)
	GetOmzetValueByGroupID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) ([]models.OmzetValueModel, error)
	GetOmzetValueByRegionID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) ([]models.OmzetValueModel, error)
	GetOmzetValueByBranchID(ctx context.Context, parameter models.DashboardWebBranchParameter, branchID string) ([]models.OmzetValueBranchModel, error)
	GetOmzetValueByCustomerID(ctx context.Context, parameter models.DashboardWebBranchParameter, customerID string) (res []models.OmzetValueModel, err error)

	GetOmzetValueGraph(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.OmzetValueModel, error)

	TrackingInvoice(ctx context.Context, input models.DashboardWebBranchParameter) (res []models.DashboardTrackingInvoice, err error)
	VirtualAccount(ctx context.Context, input models.DashboardWebBranchParameter) (res []models.DashboardVirtualAccount, err error)
}

// DashboardWebRepository ...
type DashboardWebRepository struct {
	DB *sql.DB
}

// NewDashboardWebRepository ...
func NewDashboardWebRepository(DB *sql.DB) IDashboardWebRepository {
	return &DashboardWebRepository{DB: DB}
}

// Scan rows
func (repository DashboardWebRepository) scanRows(rows *sql.Rows) (res models.DashboardWeb, err error) {
	// err = rows.Scan(
	// 	&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalVisitUser,
	// 	&res.CustomerCountRepeatOrder, &res.TotalActiveOutlet,
	// 	&res.TotalOutlet, &res.TotalCompleteCustomer,
	// )

	err = rows.Scan(&res.RegionGroupID,
		&res.RegionGroupName,
		&res.TotalVisitUser,
		&res.TotalRepeatUser,
		&res.TotalOrderUser,
		&res.TotalRegisteredUser,
		&res.CustomerCountRepeatOrder,
		&res.TotalOutlet,
		&res.TotalActiveOutlet,
		&res.TotalInvoice,
		&res.TotalCompleteCustomer)
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
		&res.TotalInvoice, &res.TotalCompleteCustomer,
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
		&res.TotalOutlet, &res.TotalCompleteCustomer,
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
		&res.TotalRepeatUser, &res.CustomerCountRepeatOrder, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.CustomerClassName, &res.CustomerCityName,
		&res.StatusComplete,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// scanGroupDetailWithUserIDRows
func (repository DashboardWebRepository) scanGroupDetailWithUserIDRows(rows *sql.Rows) (res models.DashboardWebGetWithUserID, err error) {
	err = rows.Scan(
		&res.CustomerBranchID, &res.CustomerBranchName, &res.CustomerBranchCode,
		&res.CustomerRegionName, &res.CustomerRegionGroupName,
		&res.TotalRepeatUser, &res.TotalRepeatToko, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.TotalOutlet, &res.TotalOutletAll,
		&res.TotalRegisteredUser, &res.CompleteCustomer,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// scanCustomerDetailWithUserIDRows
func (repository DashboardWebRepository) scanCustomerDetailWithUserIDRows(rows *sql.Rows) (res models.DashboardWebGetWithUserID, err error) {
	err = rows.Scan(
		&res.CustomerID, &res.CustomerName, &res.CustomerCode,
		&res.CustomerBranchCode, &res.CustomerBranchName,
		&res.CustomerRegionName, &res.CustomerRegionGroupName,
		&res.CustomerTypeName, &res.CustomerLevelName, &res.CustomerCityName,
		&res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.CompleteCustomer,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanBranchCustomerDetailReportRows(rows *sql.Rows) (res models.DashboardWebBranchDetail, err error) {
	err = rows.Scan(
		&res.CustomerID, &res.CustomerName, &res.CustomerCode,
		&res.CustomerBranchName, &res.CustomerBranchCode,
		&res.CustomerRegionName, &res.CustomerRegionGroupName,
		&res.CustomerTypeName,
		&res.TotalRepeatUser, &res.CustomerCountRepeatOrder, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.CustomerClassName, &res.CustomerCityName,
		&res.StatusComplete, &res.StatusInstall,
		&res.SalesmanCode, &res.SalesmanName,
		&res.SalesmanTypeCode, &res.SalesmanTypeName,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// FindByID ...
// func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) ([]models.DashboardWeb, error) {

// 	if parameter.StartDate == "" || parameter.EndDate == "" {
// 		return nil, errors.New("invalid dates provided")
// 	}

// 	statement := models.DashboardWebSelectStatement

// 	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var data []models.DashboardWeb
// 	for rows.Next() {
// 		temp, err := repository.scanRows(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, temp)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) (data []models.DashboardWeb, err error) {
	statement := models.DashboardWebSelectStatementNew

	var startDate, endDate string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		startDate = `'` + parameter.StartDate + `'`
		endDate = `'` + parameter.EndDate + `'`
	} else {
		startDate = `date_trunc('MONTH',now())::DATE`
		endDate = `now()::date`
	}
	statement = strings.ReplaceAll(statement, "{START_DATE}", startDate)
	statement = strings.ReplaceAll(statement, "{END_DATE}", endDate)
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

// func (repository DashboardWebRepository) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (data []models.DashboardWebRegionDetail, err error) {
// 	var statement string
// 	var args []interface{}

// 	regionID := str.NullOrEmtyString(&parameter.RegionID)
// 	startDate := str.NullOrEmtyString(&parameter.StartDate)
// 	endDate := str.NullOrEmtyString(&parameter.EndDate)

// 	if parameter.RegionID != "" {
// 		statement = models.DashboardWebRegionDetailByRegionIDSelectStatement
// 		args = append(args, regionID, startDate, endDate)
// 	} else {
// 		groupID := str.NullOrEmtyString(&parameter.GroupID)
// 		statement = models.DashboardWebRegionDetailSelectStatement
// 		args = append(args, groupID, startDate, endDate)
// 	}

// 	rows, err := repository.DB.QueryContext(c, statement, args...)
// 	if err != nil {
// 		return data, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		temp, err := repository.scanRegionDetailRows(rows)
// 		if err != nil {
// 			return data, err
// 		}
// 		data = append(data, temp)
// 	}

// 	return data, err
// }

func (repository DashboardWebRepository) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (data []models.DashboardWebRegionDetail, err error) {
	var rows *sql.Rows

	statement := models.DashboardWebRegionDetailByRegionIDSelectStatement
	rows, err = repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.RegionID), str.NullOrEmtyString(&parameter.GroupID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))

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

func (repository DashboardWebRepository) GetUserByRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (data []models.DashboardWebBranchDetail, err error) {
	statement := models.DashboardWebCustomerDetailByRegionDetailByRegionIDSelectStatement
	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.BranchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
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

	return data, err
}

func (repository DashboardWebRepository) GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, count int, err error) {

	query := models.DashboardWebBranchDetailSelectStatement + ` OFFSET $4 LIMIT $5`
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BranchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate),
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
	err = repository.DB.QueryRow(query, str.NullOrEmtyString(&parameter.BranchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate)).Scan(&count)
	return data, count, err
}

func (repository DashboardWebRepository) GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebBranchDetailSelectStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BranchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
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

func (repository DashboardWebRepository) GetAllReportBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebReportBranchDetailSelectStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BranchID), str.NullOrEmtyString(&parameter.UserID), str.NullOrEmtyString(&parameter.CustomerLevelID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailReportRows(rows)
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

func (repository DashboardWebRepository) GetAllBranchDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebGetWithUserID, err error) {
	var dateStartStatement, dateEndStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		dateStartStatement = `'` + parameter.StartDate + `'`
		dateEndStatement = `'` + parameter.EndDate + `'`
	} else {
		dateStartStatement = `date_trunc('MONTH',now())::DATE`
		dateEndStatement = `now()`
	}

	if parameter.UserID == "" {
		parameter.UserID = "0"
	}

	query := `with customer_repeat_order as(
		select count(*), cust_bill_to_id
		from sales_order_header 
		where transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
			and lower(document_no) like 'oso%' 
			and status='submitted' 
		group by cust_bill_to_id
	), customer_repeat_order_toko as(
		select count(*), sih.cust_bill_to_id
		from sales_invoice_header sih
		where sih.transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + `  
				and sih.transaction_source_document_no like 'CO%' 
				and sih.status = 'submitted'
		group by sih.cust_bill_to_id
	), customer_total_transaction as (
		select count(soh.*), c.branch_id
		from sales_order_header soh
		left join customer c on c.id = soh.cust_bill_to_id
		where soh.transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
			and lower(soh.document_no) like 'oso%'
			and soh.status='submitted'
			and c.show_in_apps = 1
		group by c.branch_id
	), customer_total_invoice as (
		select count(sih.*), sih.cust_bill_to_id
		from sales_invoice_header sih
			left join customer c on c.id = sih.cust_bill_to_id
		where lower(sih.transaction_source_document_no) like 'co%'	
			and sih.transaction_date between  ` + dateStartStatement + ` and ` + dateEndStatement + `
			and c.show_in_apps = 1
		group by sih.cust_bill_to_id
	), customer_total_check_in as (
		select count(uca.*), uca.user_id
		from user_checkin_activity uca
			left join customer cus on cus.user_id = uca.user_id
		where uca.checkin_time::date between  ` + dateStartStatement + ` and ` + dateEndStatement + `
			and cus.show_in_apps = 1
		group by uca.user_id
	), customer_aktif_outlet as (
		select count(*), sih.cust_bill_to_id
		from sales_invoice_header sih
			left join customer c on c.id = sih.cust_bill_to_id
		where c.created_date IS not NULL 
			and c.show_in_apps = 1
			and lower(sih.transaction_source_document_no) like 'co%'
			and sih.transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
		group by sih.cust_bill_to_id
	), registered_user as (
		select count(*), us.id
		from _user us 
		where us.fcm_token is not null 
			and first_login_time::date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
			and length(trim(us.fcm_token))>0 
		group by us.id
	), dataCompleteCustomer as (
		select b.id as branch_id,
		count(c.id) as total_complete_customer
		from branch b
		left join customer c on c.branch_id = b.id
		left join _user us on us.id = c.user_id
		where c.created_date IS not null and c.show_in_apps = 1 
			and coalesce(c.is_data_completed, false) = true
			and c.modified_date::date between ` + dateStartStatement + ` and ` + dateEndStatement + `
			and us.fcm_token is not null and length(trim(us.fcm_token))>0
		group by b.id
	)
	select b.id, b._name, b.branch_code, r._name, r.group_name,
		sum(case when crot.count > 1 then crot.count-1 else 0 end) as repeat_order,
		sum(case when crot.count > 1 then 1 else 0 end) as repeat_order_toko,
		coalesce(cto.count, 0) as total_transaction ,
		coalesce(sum(cti.count), 0) as total_invoice,
		coalesce(sum(ctci.count), 0) as total_check_id,
		sum(case when cao.count >= 1 then 1 else 0 end) as aktif_outlet,
		count(distinct(case when length(trim(us.fcm_token))>0 and us.fcm_token is not null then us.id end)) as total_outlet,
		count(distinct(case when def.show_in_apps = 1 and def.created_date IS not null then def.user_id end )) as total_outlet_all,
		coalesce(sum(ru.count),0) as registered_user,
		coalesce(dcc.total_complete_customer,0) as total_complete_customer
	from customer def
		left join branch b on b.id = def.branch_id
		left join region r on r.id = b.region_id
		left join _user us on us.id = def.user_id
		left join customer_repeat_order cro on cro.cust_bill_to_id = def.id
		left join customer_repeat_order_toko crot on crot.cust_bill_to_id = def.id
		left join customer_total_transaction cto on cto.branch_id = b.id
		left join customer_total_invoice cti on cti.cust_bill_to_id = def.id
		left join customer_total_check_in ctci on ctci.user_id = def.user_id
		left join customer_aktif_outlet cao on cao.cust_bill_to_id = def.id
		left join registered_user ru on ru.id = def.user_id
		left join dataCompleteCustomer dcc on dcc.branch_id = b.id
	WHERE (case when ` + parameter.UserID + ` != 0
		then 
			def.branch_id in(
				select ub.branch_id  
				from user_branch ub
				where ub.user_id = ` + parameter.UserID + `
			)
		else 
			true = true
		end)
		AND def.created_date IS not NULL 
		and def.user_id is not null
		and def.show_in_apps = 1
	GROUP BY b.ID, r.id, dcc.total_complete_customer, cto.count`
	rows, err := repository.DB.Query(query)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanGroupDetailWithUserIDRows(rows)
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

func (repository DashboardWebRepository) GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebGetWithUserID, err error) {
	var dateStartStatement, dateEndStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		dateStartStatement = `'` + parameter.StartDate + `'`
		dateEndStatement = `'` + parameter.EndDate + `'`
	} else {
		dateStartStatement = `date_trunc('MONTH',now())::DATE`
		dateEndStatement = `now()`
	}

	if parameter.BranchID == "" {
		parameter.BranchID = "0"
	}

	query := `with customer_repeat_order as(
		select count(*), cust_bill_to_id
		from sales_order_header 
		where transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
			and lower(document_no) like 'oso%' 
			and status='submitted' 
		group by cust_bill_to_id
	), customer_total_transaction as (
		select count(sih.*), sih.cust_bill_to_id
		from sales_order_header sih
		left join customer c on c.id = sih.cust_bill_to_id
		where sih.transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
			and lower(sih.document_no) like 'oso%' 
			and sih.status='submitted'
			and c.show_in_apps = 1
		group by cust_bill_to_id
	), customer_total_invoice as (
		select count(sih.*), sih.cust_bill_to_id
		from sales_invoice_header sih
		left join customer c on c.id = sih.cust_bill_to_id
		where lower(sih.transaction_source_document_no) like 'co%'	
			and sih.transaction_date between  ` + dateStartStatement + ` and ` + dateEndStatement + `
			and c.show_in_apps = 1
		group by sih.cust_bill_to_id
	), customer_total_check_in as (
		select count(uca.*), uca.user_id
		from user_checkin_activity uca
			left join customer cus on cus.user_id = uca.user_id
		where uca.checkin_time::date between  ` + dateStartStatement + ` and ` + dateEndStatement + `
			and cus.show_in_apps = 1
		group by uca.user_id
	), customer_aktif_outlet as (
		select count(*), sih.cust_bill_to_id
		from sales_invoice_header sih
		left join customer c on c.id = sih.cust_bill_to_id
		where c.created_date IS not NULL 
			and c.show_in_apps = 1
			and lower(sih.transaction_source_document_no) like 'co%'
			and sih.transaction_date between ` + dateStartStatement + ` and ` + dateEndStatement + ` 
		group by sih.cust_bill_to_id
	)
	select def.id, def.customer_name, def.customer_code, b.branch_code, b._name,
		r."_name", r.group_name,
		ctp._name, clv._name, cty."_name",
		case when cro.count > 1 then 1 else 0 end as repeat_order,
		coalesce(cto.count, 0) as total_transaction ,
		coalesce(cti.count, 0) as total_invoice,
		coalesce(ctci.count, 0) as total_check_id,
		case when cao.count >= 1 then 1 else 0 end as aktif_outlet,
		case when def.created_date IS not null and def.show_in_apps = 1
				and coalesce(def.is_data_completed, false) = true
				and def.modified_date::date between  ` + dateStartStatement + ` and ` + dateEndStatement + `
				and us.fcm_token is not null and length(trim(us.fcm_token))>0
			then 1 else 0 end
		as status_complete_customer
	from customer def
		left join _user us on us.id = def.user_id
		left join branch b on b.id = def.branch_id
		left join region r on r.id = b.region_id
		left join customer_type ctp on ctp.id = def.customer_type_id
		left join customer_level clv on clv.id = def.customer_level_id
		left join city cty on cty.id = def.customer_city_id
		left join customer_repeat_order cro on cro.cust_bill_to_id = def.id
		left join customer_total_transaction cto on cto.cust_bill_to_id = def.id
		left join customer_total_invoice cti on cti.cust_bill_to_id = def.id
		left join customer_total_check_in ctci on ctci.user_id = def.user_id
		left join customer_aktif_outlet cao on cao.cust_bill_to_id = def.id
	WHERE def.created_date IS not NULL and def.user_id is not null 
	AND (case when ` + parameter.BranchID + ` != 0
		then
			def.branch_id = ` + parameter.BranchID + `
		else
			true = true
		end)
	and def.show_in_apps = 1
	and us.fcm_token is not null and length(trim(us.fcm_token))>0`

	rows, err := repository.DB.Query(query)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanCustomerDetailWithUserIDRows(rows)
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
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
	}

	query := `select r.group_id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join item i on i.id = sil.item_id
			left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null 
			and coh.id is not null` + whereStatement + `
			group by r.group_id
			order by r.group_id desc`

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
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
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
			left join item i on i.id = sil.item_id
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
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
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
			left join item i on i.id = sil.item_id
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
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
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
		left join item i on i.id = sil.item_id
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
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
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
		and coh.id is not null 
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

func (repo DashboardWebRepository) GetOmzetValueGraph(ctx context.Context, parameter models.DashboardWebBranchParameter) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.Year != "" {
		whereStatement += ` AND TO_CHAR(sih.transaction_date, 'YYYY') = '` + parameter.Year + `'`
	} else {
		whereStatement += ` AND TO_CHAR(sih.transaction_date, 'YYYY') = TO_CHAR(NOW(), 'YYYY')`
	}

	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND i.item_category_id = ` + parameter.ItemCategoryID
	}

	if parameter.ItemIDs != "" {
		whereStatement += ` AND sil.item_id IN (` + parameter.ItemIDs + `)`
	}
	if parameter.ItemCategoryIDs != "" {
		whereStatement += ` AND i.item_category_id IN (` + parameter.ItemCategoryIDs + `)`
	}
	if parameter.UserID != "" {
		whereStatement += ` AND b.branch_id in(
			select ub.branch_id  
			from user_branch ub
			where ub.user_id = ` + parameter.UserID + `
		) `
	}
	if parameter.RegionID != "" {
		whereStatement += `r.ID = ` + parameter.RegionID
	}
	if parameter.RegionGroupID != "" {
		whereStatement += `r.group_id = ` + parameter.RegionGroupID
	}
	if parameter.BranchID != "" {
		whereStatement += `b.id in (` + parameter.BranchID + `)`
	}

	query := `select TO_CHAR(sih.transaction_date, 'YYYY-MM') as transaction_year_month,
			TO_CHAR(sih.transaction_date, 'YYYY') as transaction_year,
			TO_CHAR(sih.transaction_date, 'Month') as transaction_month_name,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join item i on i.id = sil.item_id
			left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null 
			and coh.id is not null` + whereStatement + `
			group by transaction_year_month, transaction_year, transaction_month_name
			order by transaction_year_month asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.TransactionYearMonth, &temp.TransactionYear, &temp.TransactionMonthName,
			&temp.TotalGrossAmount, &temp.TotalNettAmount, &temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) TrackingInvoice(ctx context.Context, input models.DashboardWebBranchParameter) (res []models.DashboardTrackingInvoice, err error) {
	var startDate, endDate string
	if input.StartDate != "" && input.EndDate != "" {
		startDate = "'" + input.StartDate + "'"
		endDate = "'" + input.EndDate + "'"
	} else {
		startDate = "date_trunc('MONTH',now())::DATE"
		endDate = "now()"
	}

	var whereStatement string
	if input.RegionGroupID != "" {
		whereStatement += " AND ct.group_id = " + input.RegionGroupID
	}
	if input.RegionID != "" {
		whereStatement += ` AND ct.region_id = ` + input.RegionID
	}
	if input.BranchID != "" {
		whereStatement += ` AND ct.branch_id in (` + input.BranchID + `)`
	}
	if input.BranchArea != "" {
		whereStatement += ` AND ct.branch_area = '` + input.BranchArea + `'`
	}
	if input.CustomerLevelID != "" {
		whereStatement += ` AND ct.customer_level_id = ` + input.CustomerLevelID
	}
	if input.UserID != "" {
		whereStatement += ` AND ct.branch_id in(
			select ub.branch_id  
			from user_branch ub
			where ub.user_id = ` + input.UserID + `
		) `
	}
	queryStatement := `with customer_temp as (
		select r.group_id as group_id, r.group_name as group_name , 
			r.id as region_id, r."_name" as region_name, 
			b.id as branch_id, b.area as branch_area, b.branch_code as branch_code, b."_name" as branch_name,
			c.id as customer_id, c.customer_name, c.customer_code, 
			c.customer_level_id as customer_level_id, cl."_name" as customer_level_name,
			c.user_id as customer_user_id,
			DIST._NAME AS CUST_DISTRICT_NAME,
			SDIST._NAME AS CUST_SUBDISTRICT_NAME
		from customer c 
		left join customer_level cl on cl.id = c.customer_level_id 
		left join branch b on b.id= c.branch_id 
		left join region r on r.id = b.region_id 
		left JOIN DISTRICT DIST ON DIST.ID = C.CUSTOMER_DISTRICT_ID
		left JOIN SUBDISTRICT SDIST ON SDIST.ID = C.CUSTOMER_SUBDISTRICT_ID
	)
	select ct.group_name, ct.region_name,
		ct.branch_name, ct.branch_area, ct.branch_code,
		ct.customer_name, ct.customer_code, ct.customer_level_name, ct.CUST_DISTRICT_NAME, ct.CUST_SUBDISTRICT_NAME,
		sih.id, sih.document_no as invoice_no,
		coh.document_no as customer_order_document_no, coh.created_date as customer_order_created_date,
		soh.document_no as sales_order_document_no, soh.created_date as sales_order_created_date,
		sih.transaction_date + sih.transaction_time as invoice_created_date, sih.invoice_date, sih.modified_date as invoice_updated_date,
		sih.paid_date,
		top.days
	from sales_invoice_header sih 
		left join term_of_payment top on top.id = sih.payment_terms_id 
		left join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
		left join sales_order_header soh on soh.request_document_no = sih.transaction_source_document_no 
		left join customer_temp ct on ct.customer_id = sih.cust_bill_to_id 
	where sih.transaction_date between ` + startDate + ` and ` + endDate + `
		` + whereStatement + `
	order by sih.transaction_date desc, sih.transaction_time desc`

	rows, err := repo.DB.Query(queryStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.DashboardTrackingInvoice
		err = rows.Scan(&temp.RegionGroupName, &temp.RegionName,
			&temp.BranchName, &temp.BranchArea, &temp.BranchCode,
			&temp.CustomerName, &temp.CustomerCode, &temp.CustomerLevel, &temp.CustomerDistrictName, &temp.CustomerSubDistrictName,
			&temp.InvoiceID, &temp.InvoiceNumber,
			&temp.CustomerOrderDocumentNo, &temp.CustomerOrderCreatedDate,
			&temp.SalesOrderDocumentNo, &temp.SalesOrderCreatedDate,
			&temp.InvoiceCreatedDate, &temp.InvoiceAcceptedDate, &temp.InvoiceUpdatedDate,
			&temp.PaidOffDate,
			&temp.DueDate)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) VirtualAccount(ctx context.Context, input models.DashboardWebBranchParameter) (res []models.DashboardVirtualAccount, err error) {
	var whereStatement string
	if input.StartDate != "" && input.EndDate != "" {
		whereStatement += `AND SIH.transaction_date BETWEEN '` + input.StartDate + `' AND '` + input.EndDate + `'`
	} else {
		whereStatement += `AND SIH.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}

	if input.BranchID != "" {
		whereStatement += `AND b.id in (` + input.BranchID + `)`
	}

	if input.RegionID != "" {
		whereStatement += `AND r.id in (` + input.RegionID + `)`
	}
	if input.RegionGroupID != "" {
		whereStatement += `AND r.group_id in (` + input.RegionGroupID + `)`
	}
	queryStatement := `	select def.va_code, def.invoice_code, sih.transaction_source_document_no, def.start_date, def.end_date,
	c.customer_name, c.customer_code, c.customer_phone,
	b._name, b.branch_code, b.area,
	r._name, r.group_name,
	def.amount, def.va_ref1, def.va_ref2
	from virtual_account_transaction def
		left join sales_invoice_header sih on sih.document_no = def.invoice_code
		left join customer c on c.id = sih.cust_bill_to_id
		left join branch b on b.id = c.branch_id
		left join region r on r.id = b.region_id
	where def.created_date is not null and def.va_code is not null ` + whereStatement + `
	order by def.created_date desc`

	rows, err := repo.DB.Query(queryStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.DashboardVirtualAccount
		err = rows.Scan(&temp.VirtualAccountNumber, &temp.InvoiceNumber, &temp.SourceDocumentNo, &temp.VirtualAccountStartDate, &temp.VirtualAccountEndDate,
			&temp.CustomerName, &temp.CustomerCode, &temp.CustomerPhoneNo,
			&temp.BranchName, &temp.BranchCode, &temp.BranchArea,
			&temp.RegionName, &temp.RegionGroupName,
			&temp.Amount, &temp.VirtualAccountRef1, &temp.VirtualAccountRef2,
		)
		if err != nil {
			return
		}

		res = append(res, temp)
	}
	return
}
