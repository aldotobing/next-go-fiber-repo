package models

import "database/sql"

// CustomerDataSync ...
type CustomerDataSync struct {
	ID                *string        `json:"customer_id"`
	PartnerID         *string        `json:"partner_id"`
	Code              *string        `json:"customer_code"`
	SalesmanCode      *string        `json:"salesman_code"`
	Name              *string        `json:"customer_name"`
	CustomerType      *string        `json:"customer_type_code"`
	Address           *string        `json:"customer_address"`
	PhoneNo           *string        `json:"customer_phone_no"`
	PriceListCode     *string        `json:"price_list_code"`
	CountryCode       *string        `json:"country_code"`
	CityCode          *string        `json:"city_code"`
	DistrictCode      *string        `json:"district_code"`
	SubDistrictCode   *string        `json:"subdistrict_code"`
	ProvinceCode      *string        `json:"province_code"`
	BranchID          *string        `json:"branch_id"`
	TermOfPaymentCode *string        `json:"top_code"`
	CustomerLevelCode *string        `json:"customer_level_code"`
	UserID            sql.NullString `json:"user_id"`
}

// CustomerDataSyncParameter ...
type CustomerDataSyncParameter struct {
	ID        string `json:"id_customer"`
	Code      string `json:"customer_code"`
	DateParam string `json:"date_param"`
	Name      string `json:"name"`
	Search    string `json:"search"`
	MysmOnly  string `json:"mysm_only"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
}

var (
	// CustomerDataSyncOrderBy ...
	CustomerDataSyncOrderBy = []string{"c.id", "p._name", "c.created_date"}
	// CustomerDataSyncOrderByrByString ...
	CustomerDataSyncOrderByrByString = []string{
		"p._name",
	}

	// CustomerDataSyncSelectStatement ...
	CustomerDataSyncSelectStatement = `select c.id as customer_id,p.id as partner_id,p.code as customer_code, p._name as cusomer_name, 
	p.address as customer_addrss,p.phone_no as customer_phone_no,ctp.code as customer_type_code, 
	p.code as plice_list_code,cntr.code as country_code,cty.code as city_code, dst.code as district_code,
	sdst.code as subdistrict_code,pv.code as province_code, top.code as top_code, 
	ps.code as salesman_code ,b.id::character varying as branch_id,
	c.user_id
	from customer c  
	left join partner p on p.id=c.partner_id 
	left join salesman s on s.id = c.salesman_id 
	left join partner ps on ps.id = s.partner_id 
	left join customer_type ctp on ctp.id = c.customer_type_id 
	left join price_list pl on pl.id =c.price_list_id 
	left join country cntr on cntr.id = p.country_id 
	left join city cty on cty.id = p.city_id  
	left join district dst on dst.id = p.district_id 
	left join province pv on pv.id = p.province_id 
	left join term_of_payment top on top.id = c.payment_terms_id 
	left join subdistrict sdst on sdst.id = p.subdistrict_id  
	left join branch b on b.id = c.branch_id
	`

	// CustomerDataSyncWhereStatement ...
	CustomerDataSyncWhereStatement = `WHERE p._name IS not NULL`
)
