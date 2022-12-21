package models

// DashboardWeb ...
type DashboardWeb struct {
	RegionName      *string `json:"region_name"`
	TotalActiveUser *string `json:"total_active_user"`
	TotalRepeatUser *string `json:"total_repeat_order_user"`
	TotalOrderUser  *string `json:"total_order_user"`
	TotalInvoice    *string `json:"total_invoice_user"`
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

var (

	// DashboardWebSelectStatement ...

	DashboardWebSelectStatement = ` 
	select 'Nasional',(select count(*) from _user us join customer c on c.id=us.partner_id where us.fcm_token is not null and length(trim(us.fcm_token))>0) as total_active_user,
	(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP))group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
	(select count(*) from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP))) as total_transaction,
	(select count(*) from sales_invoice_header where cust_bill_to_id in(select cust_bill_to_id from customer_order_header) and (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP))) as total_invoice
	union all
	select * from(
	select x.group_name,
		(select count(*) from _user us join customer c on c.id=us.partner_id where us.fcm_token is not null and length(trim(us.fcm_token))>0 and c.branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) ) ) as total_active_user,
		(select count(*) from (select count(*) as total_transaksi,cust_bill_to_id from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id) )  group by cust_bill_to_id) x where x.total_transaksi>1) as total_repeat_order,
		(select count(*) from customer_order_header where (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP))  and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_transaction,
		(select count(*) from sales_invoice_header where cust_bill_to_id in(select cust_bill_to_id from customer_order_header) and (date_part('month', now()::TIMESTAMP) = date_part('month', now()::TIMESTAMP) and date_part('year', now()::TIMESTAMP)=date_part('year', now()::TIMESTAMP)) and branch_id in (select br.id from branch br where br.region_id in(select rg.id from region rg where rg.group_id = x.group_id)) ) as total_invoice 
	from (
	select r.group_id,r.group_name 
	from region r group by r.group_id,r.group_name
	)x order by x.group_id
		)y
 	`
)
