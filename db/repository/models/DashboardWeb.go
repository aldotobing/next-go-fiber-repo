package models

// DashboardWeb ...
type DashboardWeb struct {
	RegionGroupID       *string                    `json:"region_id"`
	RegionGroupName     *string                    `json:"region_name"`
	TotalActiveUser     *string                    `json:"total_active_user"`
	TotalRepeatUser     *string                    `json:"total_repeat_order_user"`
	TotalOrderUser      *string                    `json:"total_order_user"`
	TotalInvoice        *string                    `json:"total_invoice_user"`
	TotalRegisteredUser *string                    `json:"total_registered_user"`
	DetailData          []DashboardWebRegionDetail `json:"detailed_data"`
}

type DashboardWebRegionDetail struct {
	RegionID            *string `json:"region_id_detail"`
	RegionName          *string `json:"region_name_detail"`
	RegionGroupID       *string `json:"region_group_id_detail"`
	RegionGroupName     *string `json:"region_group_name_detail"`
	BranchID            *string `json:"branch_id_detail"`
	BranchName          *string `json:"branch_name_detail"`
	TotalActiveUser     *string `json:"total_active_user_detail"`
	TotalRepeatUser     *string `json:"total_repeat_order_user_detail"`
	TotalOrderUser      *string `json:"total_order_user_detail"`
	TotalInvoice        *string `json:"total_invoice_user_detail"`
	TotalRegisteredUser *string `json:"total_registered_user_detail"`
}

type DashboardWebBranchDetail struct {
	CustomerID      *string `json:"customer_id_detail"`
	CustomerName    *string `json:"customer_name_detail"`
	TotalRepeatUser *string `json:"total_repeat_order_user_customer_detail"`
	TotalOrderUser  *string `json:"total_order_user_customer_detail"`
	TotalInvoice    *string `json:"total_invoice_user_customer_detail"`
}

// DashboardWebParameter ...
type DashboardWebParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

type DashboardWebRegionParameter struct {
	GroupID string `json:"group_id"`
	Search  string `json:"search"`
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	By      string `json:"by"`
	Sort    string `json:"sort"`
}

type DashboardWebBranchParameter struct {
	BarnchID string `json:"branch_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
}

var (

	// DashboardWebSelectStatement ...

	DashboardWebSelectStatement = ` 
	select 0 as group_id,'Nasional' as group_name,(select count(*) from _user us join customer c on c.id=us.partner_id where us.fcm_token is not null and length(trim(us.fcm_token))>0) as total_register_user,
	(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
	(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))) as total_transaction,
	(select count(*) from sales_invoice_header where cust_bill_to_id in(select cust_bill_to_id from customer_order_header) 
	 and (date_part('month',transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))
	) as total_invoice,
	(select count(*) from (select count(distinct(cust_bill_to_id)) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP))group by cust_bill_to_id) x) as total_active_user

	union all
	select * from(
	select x.group_id as group_id,x.group_name as group_name, 
		(select count(*) from _user us join customer c on c.id=us.partner_id where us.fcm_token is not null and length(trim(us.fcm_token))>0 and c.branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) ) ) as total_register_user,
		(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) )  group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
		(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', transaction_date::TIMESTAMP))  and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_transaction,
		(select count(*) from sales_invoice_header where cust_bill_to_id in(select distinct(cust_bill_to_id) from customer_order_header) 
		 and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) 
		 and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_invoice,
		(select count(*) from (select distinct(cust_bill_to_id) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) )  group by cust_bill_to_id) x) as total_active_user

	from (
	select r.group_id,r.group_name 
	from region r group by r.group_id,r.group_name
	)x order by x.group_id
		)y
 	`

	DashboardWebRegionDetailSelectStatement = `
	 
	 select 
		def.id as b_id,def._name as b_name,
		r.id as region_id, r._name as region_name, r.group_id as region_group_id, r.group_name as region_group_name,
		(select count(*) from _user us join customer c on c.id=us.partner_id where us.fcm_token is not null and length(trim(us.fcm_token))>0 and c.branch_id = def.id  ) as total_register_user,
		(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id =def.id )  group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
		(select count(*) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', transaction_date::TIMESTAMP))  and branch_id =def.id ) as total_transaction,
		(select count(*) from sales_invoice_header where cust_bill_to_id in(select distinct(cust_bill_to_id) from customer_order_header) 
				and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) 
				and branch_id =def.id ) as total_invoice,
		(select count(*) from (select distinct(cust_bill_to_id) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id =def.id  group by cust_bill_to_id) x) as total_active_user


		from branch def
		left join region r on r.id = def.region_id
		
	 `

	DashboardWebBranchDetailOrderBy = []string{"def.id", "def.customer_name"}
	// CustomerOrderLineOrderByrByString ...
	DashboardWebBranchDetailOrderByrByString = []string{
		"def.id",
	}

	DashboardWebBranchDetailSelectStatement = ` select 
	def.id as cus_id,def.customer_name as cus_name,
	(case when ( select count(*) from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and cust_bill_to_id = def.id)>0 then
	( select count(*) from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and cust_bill_to_id = def.id) else 0 end
	)as total_repeat_order,
	(select count(*) from customer_order_header where cust_bill_to_id = def.id  and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', transaction_date::TIMESTAMP)) ) as total_transaction,
	(select count(*) from sales_invoice_header where  cust_bill_to_id in (select distinct(cust_bill_to_id) from customer_order_header where (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) )  and cust_bill_to_id = def.id 
			and (date_part('month', transaction_date::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', transaction_date::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) 
	) as total_invoice
	from customer def
	left join branch b on b.id = def.branch_id
	left join region r on r.id = b.region_id
	   `
)
