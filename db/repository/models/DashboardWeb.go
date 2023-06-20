package models

import "database/sql"

// DashboardWeb ...
type DashboardWeb struct {
	RegionGroupID            *string                    `json:"region_id"`
	RegionGroupName          *string                    `json:"region_name"`
	TotalVisitUser           *string                    `json:"total_visit_user"`
	TotalRepeatUser          *string                    `json:"total_repeat_order_user"`
	TotalOrderUser           *string                    `json:"total_order_user"`
	TotalInvoice             *string                    `json:"total_invoice_user"`
	TotalRegisteredUser      *string                    `json:"total_registered_user"`
	DetailData               []DashboardWebRegionDetail `json:"detailed_data"`
	CustomerCountRepeatOrder *string                    `json:"customer_count_repeat_order"`
	TotalActiveOutlet        *string                    `json:"total_active_outlet"`
	TotalOutlet              *string                    `json:"total_outlet"`
	TotalCompleteCustomer    *string                    `json:"total_complete_customer"`
}

type DashboardWebRegionDetail struct {
	RegionID                 *string `json:"region_id_detail"`
	RegionName               *string `json:"region_name_detail"`
	RegionGroupID            *string `json:"region_group_id_detail"`
	RegionGroupName          *string `json:"region_group_name_detail"`
	BranchID                 *string `json:"branch_id_detail"`
	BranchName               *string `json:"branch_name_detail"`
	BranchCode               *string `json:"branch_code_detail"`
	TotalVisitUser           *string `json:"total_visit_user_detail"`
	TotalRepeatUser          *string `json:"total_repeat_order_user_detail"`
	TotalOrderUser           *string `json:"total_order_user_detail"`
	TotalInvoice             *string `json:"total_invoice_user_detail"`
	TotalRegisteredUser      *string `json:"total_registered_user_detail"`
	CustomerCountRepeatOrder *string `json:"customer_count_repeat_order_detail"`
	TotalActiveOutlet        *string `json:"total_active_outlet_detail"`
	TotalOutlet              *string `json:"total_outlet"`
	TotalCompleteCustomer    *string `json:"total_complete_customer"`
}

type DashboardWebBranchDetail struct {
	CustomerID              *string `json:"customer_id_detail"`
	CustomerName            *string `json:"customer_name_detail"`
	CustomerCode            *string `json:"customer_code_detail"`
	CustomerBranchName      *string `json:"customer_branch_name_detail"`
	CustomerBranchCode      *string `json:"customer_branch_code_detail"`
	CustomerRegionName      *string `json:"customer_region_name_detail"`
	CustomerRegionGroupName *string `json:"customer_region_group_name_detail"`
	CustomerTypeName        *string `json:"customer_type_name_detail"`
	TotalRepeatUser         *string `json:"total_repeat_order_user_customer_detail"`
	TotalOrderUser          *string `json:"total_order_user_customer_detail"`
	TotalInvoice            *string `json:"total_invoice_user_customer_detail"`
	TotalCheckin            *string `json:"total_checkin_user_customer_detail"`
	TotalAktifOutlet        *string `json:"total_aktif_outlet"`
	CustomerClassName       *string `json:"customer_class_name_detail"`
	CustomerCityName        *string `json:"customer_city_name_detail"`
	StatusInstall           *string `json:"status_install"`
}

type DashboardWebGetWithUserID struct {
	CustomerID              *string `json:"customer_id_detail"`
	CustomerName            *string `json:"customer_name_detail"`
	CustomerCode            *string `json:"customer_code_detail"`
	CustomerBranchID        *string `json:"customer_branch_id"`
	CustomerBranchName      *string `json:"customer_branch_name_detail"`
	CustomerBranchCode      *string `json:"customer_branch_code_detail"`
	CustomerRegionID        *string `json:"customer_region_id"`
	CustomerRegionName      *string `json:"customer_region_name_detail"`
	CustomerRegionGroupName *string `json:"customer_region_group_name_detail"`
	CustomerTypeName        *string `json:"customer_type_name_detail"`
	CustomerLevelName       *string `json:"customer_level_name"`
	TotalRepeatUser         *string `json:"total_repeat_order_user_customer_detail"`
	TotalRepeatToko         *string `json:"total_repeat_order_toko"`
	TotalOrderUser          *string `json:"total_order_user_customer_detail"`
	TotalInvoice            *string `json:"total_invoice_user_customer_detail"`
	TotalCheckin            *string `json:"total_checkin_user_customer_detail"`
	TotalAktifOutlet        *string `json:"total_aktif_outlet"`
	TotalOutlet             *string `json:"total_outlet"`
	TotalOutletAll          *string `json:"total_outlet_all"`
	TotalRegisteredUser     *string `json:"total_registered_user_detail"`
	CustomerClassName       *string `json:"customer_class_name_detail"`
	CustomerCityName        *string `json:"customer_city_name_detail"`
	StatusInstall           *string `json:"status_install"`
}
type OmzetValueModel struct {
	RegionID            sql.NullString `json:"region_id"`
	RegionName          sql.NullString `json:"region_name"`
	RegionGroupID       sql.NullString `json:"region_group_id"`
	RegionGroupName     sql.NullString `json:"region_group_name"`
	BranchID            sql.NullString `json:"branch_id"`
	CustomerID          sql.NullString `json:"customer_id"`
	ItemID              sql.NullString `json:"item_id"`
	ItemName            sql.NullString `json:"item_name"`
	TotalGrossAmount    string         `json:"total_gross_amount"`
	TotalNettAmount     string         `json:"total_nett_amount"`
	TotalQuantity       string         `json:"total_quantity"`
	TotalActiveCustomer string         `json:"total_active_customer"`
}

type OmzetValueBranchModel struct {
	RegionGroupName  *string `json:"region_group_name"`
	RegionName       *string `json:"region_name"`
	BranchName       *string `json:"branch_name"`
	BranchCode       *string `json:"branch_code"`
	CustomerID       *string `json:"customer_id"`
	CustomerCode     *string `json:"customer_code"`
	CustomerName     *string `json:"customer_name"`
	CustomerType     *string `json:"customer_type"`
	ProvinceName     *string `json:"customer_province_name"`
	CityName         *string `json:"customer_city_name"`
	CustomerLevel    *string `json:"customer_level"`
	TotalGrossAmount string  `json:"total_gross_amount"`
	TotalNettAmount  string  `json:"total_nett_amount"`
	TotalQuantity    string  `json:"total_quantity"`
}

// DashboardWebParameter ...
type DashboardWebParameter struct {
	ID        string `json:"id"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	GroupID   string `json:"group_id"`
}

type DashboardWebRegionParameter struct {
	GroupID   string `json:"group_id"`
	RegionID  string `json:"region_id"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type DashboardWebBranchParameter struct {
	BranchID        string `json:"branch_id"`
	Search          string `json:"search"`
	Page            int    `json:"page"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	By              string `json:"by"`
	Sort            string `json:"sort"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	UserID          string `json:"user_id"`
	ItemID          string `json:"item_id"`
	ItemCategoryID  string `json:"item_category_id"`
	GroupID         string `json:"group_id"`
	ItemIDs         string `json:"item_ids"`
	ItemCategoryIDs string `json:"item_category_ids"`
	RegionGroupID   string `json:"region_group_id"`
	RegionID        string `json:"region_id"`
	CustomerLevelID string `json:"customer_level_id"`
	CustomerTypeID  string `json:"customer_type_id"`
	BranchArea      string `json:"branch_area"`
}

var (

	// DashboardWebSelectStatement ...

	DashboardWebSelectStatement = ` 
	select * from os_fetch_dashborad_regiongroupdata($1,$2,null,null,null)
	`

	DashboardWebSelectByGroupIDStatement = ` 
	with dataRepeatOrder as (
		select c.id as customer_id, count(soh.id) as total_transaction
		from customer c 
		left join sales_order_header soh ON soh.cust_bill_to_id = c.id 
		where lower(document_no) like '%oso%' 
			and soh.status='submitted'
			and soh.transaction_date between '{START_DATE}' and '{END_DATE}' 
		group by c.id 
	),
	dataInvoice as (
		select r.id region_id, count(sih.id) total_invoice
		from sales_invoice_header sih 
			join customer_order_header coh on coh.document_no = sih.transaction_source_document_no
			left join branch b on b.id = sih.branch_id 
			left join region r on r.id = b.region_id 
		where lower(sih.transaction_source_document_no) like '%co%'
			and coh.status in ('submitted','finish')
			and sih.transaction_date between '{START_DATE}' and '{END_DATE}' 
		group by r.id
	),
	dataVisitedUser as (
		select uca.user_id as user_id, count(uca.id) as visit_user
		from user_checkin_activity uca
		where uca.checkin_time between '{START_DATE}' and '{END_DATE}'
		group by uca.user_id
	),
	dataOutlet as (
		select r.id as region_id, count(distinct c.id) as total_outlet
		from branch b 
		left join customer c on c.branch_id = b.id
		left join sales_invoice_header sih on sih.cust_bill_to_id = c.id
		left join region r on r.id = b.region_id
		where c.created_date IS not NULL 
			and c.show_in_apps = 1 
		group by r.id
	), 
	dataActiveOutlet as (
		select r.id as region_id,
		count(distinct sih.cust_bill_to_id) as active_outlet
		from branch b 
		left join customer c on c.branch_id = b.id
		left join sales_invoice_header sih on sih.cust_bill_to_id = c.id
		join customer_order_header coh on coh.document_no = sih.transaction_source_document_no 
		left join region r on r.id = b.region_id
		where c.created_date IS not NULL 
			and c.show_in_apps = 1 and sih.transaction_date between '{START_DATE}' and '{END_DATE}'
			and lower(sih.transaction_source_document_no) like 'co%' 
			and coh.status in ('submitted','finish')
		group by r.id 
	),
	dataCompleteCustomer as (
		select r.id as region_id,
		count(c.id) as total_complete_customer
		from region r 
		left join branch b on b.region_id = r.id
		left join customer c on c.branch_id = b.id
		where c.modified_date between '{START_DATE}' and '{END_DATE}'
			and(c.customer_nik is not null or c.customer_nik != '')
			and (c.customer_name is not null or c.customer_name != '')
			and (c.customer_birthdate is not null)
			and (c.customer_religion is not null or c.customer_religion != '')
			and (c.customer_photo_ktp is not null or c.customer_photo_ktp != '')
			and (c.customer_profile_picture is not null or c.customer_profile_picture != '')
			and (c.customer_phone is not null or c.customer_phone != '')
			and (c.customer_code is not null or c.customer_code != '')
			and c.created_date IS not null and c.show_in_apps = 1
		group by r.id
	)
	select r.id, r."_name",
	coalesce(sum(dvs.visit_user),0) as total_visit_user,
	sum(case when dro.total_transaction>1 then(dro.total_transaction-1) else 0 end) as total_repeat_order_user,
	coalesce (sum(dro.total_transaction), 0) as total_order_user,
	count(u.id) filter (
		where u.fcm_token is not null 
			and u.first_login_time between '{START_DATE}' and '{END_DATE}' 
			and length(trim(u.fcm_token))>0
	) as total_register_user,
	coalesce(count(dro.customer_id) filter (where dro.total_transaction > 1), 0) as customer_count_repeat_order,
	coalesce (ddo.total_outlet,0) as total_outlet,
	coalesce(dao.active_outlet,0) as total_active_outlet,
	coalesce (di.total_invoice,0) as total_invoice,
	coalesce (dcc.total_complete_customer,0) as total_complete_customer
	from customer c 
		left join branch b on b.id = c.branch_id
		left join region r on r.id = b.region_id
		left join "_user" u on u.id = c.user_id
		left join dataRepeatOrder dro on dro.customer_id = c.id
		left join dataVisitedUser dvs on dvs.user_id = u.id
		left join dataOutlet ddo on ddo.region_id = r.id
		left join dataActiveOutlet dao on dao.region_id= r.id
		left join dataInvoice di on di.region_id = r.id
		left join dataCompleteCustomer dcc on dcc.region_id = r.id
	{WHERE_STATEMENT}
	group by r.id, ddo.total_outlet, dao.active_outlet, di.total_invoice, dcc.total_complete_customer
	`

	DashboardWebRegionDetailSelectStatement = `
	select * from os_fetch_dashborad_regiongroupdetaildata($1::integer,$2,$3,null,null,null)
	 
		
	 `
	DashboardWebRegionDetailByRegionIDSelectStatement = `
	select * from os_fetch_dashborad_regiongroupdetaildata_by_region_id($1::integer,$2,$3,null,null,null) `

	DashboardWebBranchDetailOrderBy = []string{"def.id", "def.customer_name"}
	// CustomerOrderLineOrderByrByString ...
	DashboardWebBranchDetailOrderByrByString = []string{
		"def.id",
	}

	DashboardWebBranchDetailSelectStatement = ` select * from os_fetch_dashborad_branchcustomerdata($1::integer,$2,$3,null,null,null)
	   `

	DashboardWebReportBranchDetailSelectStatement = ` select * from os_fetch_dashborad_branchcustomerdata2($1::varchar,$2,$3,null,null,null)
	   `

	DashboardWebBranchDetailSelectWithUserIDStatement = ` select * from os_fetch_dashborad_customerdata_using_user_id($1::integer,$2,$3,null,null,null)
	   `
)

// old query dashboard region group
// select 0 as group_id,'Nasional' as group_name,(select count(*) from _user us join customer c on c.user_id=us.id where us.fcm_token is not null and length(trim(us.fcm_token))>0) as total_register_user,
// 	(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
// 	(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))) as total_transaction,
// 	(select count(*) from sales_invoice_header where cust_bill_to_id in(select cust_bill_to_id from customer_order_header)
// 	and transaction_source_document_no like '%co%'
// 	 and (date_part('month',transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))
// 	) as total_invoice,
// 	(select count(*) from (select count(distinct(cust_bill_to_id)) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))group by cust_bill_to_id) x) as total_active_user

// 	union all
// 	select * from(
// 	select x.group_id as group_id,x.group_name as group_name,
// 		(select count(*) from _user us join customer c on c.user_id=us.id where us.fcm_token is not null and length(trim(us.fcm_token))>0 and c.branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) ) ) as total_register_user,
// 		(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) )  group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
// 		(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', transaction_date::TIMESTAMP))  and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_transaction,
// 		(select count(*) from sales_invoice_header where cust_bill_to_id in(select distinct(cust_bill_to_id) from customer_order_header)
// 		and transaction_source_document_no like '%co%'
// 		and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))
// 		 and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_invoice,
// 		(select count(*) from (select distinct(cust_bill_to_id) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) )  group by cust_bill_to_id) x) as total_active_user

// 	from (
// 	select r.group_id,r.group_name
// 	from region r group by r.group_id,r.group_name
// 	)x order by x.group_id
// 		)y

// query dashboard detail old

// select
// 		def.id as b_id,def._name as b_name,
// 		r.id as region_id, r._name as region_name, r.group_id as region_group_id, r.group_name as region_group_name,
// 		(select count(*) from _user us join customer c on c.user_id=us.id where us.fcm_token is not null and length(trim(us.fcm_token))>0 and c.branch_id = def.id  ) as total_register_user,
// 		(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id =def.id )  group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
// 		(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', transaction_date::TIMESTAMP))  and branch_id =def.id ) as total_transaction,
// 		(select count(*) from sales_invoice_header where cust_bill_to_id in(select distinct(cust_bill_to_id) from customer_order_header)
// 		and transaction_source_document_no like '%co%'
// 		and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))
// 				and branch_id =def.id ) as total_invoice,
// 		(select count(*) from (select distinct(cust_bill_to_id) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id =def.id  group by cust_bill_to_id) x) as total_active_user

// from branch def
// left join region r on r.id = def.region_id
