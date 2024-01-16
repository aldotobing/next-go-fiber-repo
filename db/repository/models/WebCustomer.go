package models

import "database/sql"

// Customer ...
type WebCustomer struct {
	ID                         sql.NullString `json:"customer_id"`
	Code                       sql.NullString `json:"customer_code"`
	CustomerName               sql.NullString `json:"customer_name"`
	CustomerProfilePicture     sql.NullString `json:"customer_profile_picture"`
	CustomerActiveStatus       sql.NullString `json:"customer_active_status"`
	CustomerBirthDate          sql.NullString `json:"customer_birthdate"`
	CustomerReligion           sql.NullString `json:"customer_religion"`
	CustomerLatitude           sql.NullString `json:"customer_latitude"`
	CustomerLongitude          sql.NullString `json:"customer_longitude"`
	CustomerBranchCode         sql.NullString `json:"customer_branch_code"`
	CustomerBranchName         sql.NullString `json:"customer_branch_name"`
	CustomerBranchArea         sql.NullString `json:"customer_branch_area"`
	CustomerBranchAddress      sql.NullString `json:"customer_branch_address"`
	CustomerBranchLat          sql.NullString `json:"customer_branch_lat"`
	CustomerBranchLng          sql.NullString `json:"customer_branch_lng"`
	CustomerBranchPicPhoneNo   sql.NullString `json:"customer_branch_pic_phone_no"`
	CustomerBranchPicName      sql.NullString `json:"customer_branch_pic_name"`
	CustomerRegionCode         sql.NullString `json:"customer_region_code"`
	CustomerRegionName         sql.NullString `json:"customer_region_name"`
	CustomerRegionGroup        sql.NullString `json:"customer_region_group"`
	CustomerEmail              sql.NullString `json:"customer_email"`
	CustomerCpName             sql.NullString `json:"customer_cp_name"`
	CustomerAddress            sql.NullString `json:"customer_address"`
	CustomerPostalCode         sql.NullString `json:"customer_postal_code"`
	CustomerProvinceID         sql.NullString `json:"customer_province_id"`
	CustomerProvinceName       sql.NullString `json:"customer_province_name"`
	CustomerCityID             sql.NullString `json:"customer_city_id"`
	CustomerCityName           sql.NullString `json:"customer_city_name"`
	CustomerDistrictID         sql.NullString `json:"customer_district_id"`
	CustomerDistrictName       sql.NullString `json:"customer_district_name"`
	CustomerSubdistrictID      sql.NullString `json:"customer_subdistrict_id"`
	CustomerSubdistrictName    sql.NullString `json:"customer_subdistrict_name"`
	CustomerSalesmanCode       sql.NullString `json:"customer_salesman_code"`
	CustomerSalesmanName       sql.NullString `json:"customer_salesman_name"`
	CustomerSalesmanPhone      sql.NullString `json:"customer_salesman_phone"`
	CustomerSalesCycle         sql.NullString `json:"customer_sales_cycle"`
	CustomerTypeId             sql.NullString `json:"customer_type_id"`
	CustomerTypeName           sql.NullString `json:"customer_type_name"`
	CustomerPhone              sql.NullString `json:"customer_phone"`
	CustomerPoint              sql.NullString `json:"customer_point"`
	GiftName                   sql.NullString `json:"customer_gift_name"`
	Loyalty                    sql.NullString `json:"customer_loyalty"`
	VisitDay                   sql.NullString `json:"visit_day"`
	CustomerTaxCalcMethod      sql.NullString `json:"customer_tax_calc_method"`
	CustomerBranchID           sql.NullString `json:"customer_branch_id"`
	CustomerSalesmanID         sql.NullString `json:"customer_salesman_id"`
	CustomerNik                sql.NullString `json:"customer_nik"`
	CustomerPhotoKtp           sql.NullString `json:"customer_photo_ktp"`
	CustomerPhotoKtpDashboard  sql.NullString `json:"customer_photo_ktp_dashboard"`
	CustomerPhotoNpwp          sql.NullString `json:"customer_photo_npwp"`
	CustomerPhotoNpwpDashboard sql.NullString `json:"customer_photo_npwp_dashboard"`
	CustomerLevelID            sql.NullInt64  `json:"customer_level_id"`
	CustomerLevel              sql.NullString `json:"customer_level_name"`
	CustomerUserID             sql.NullString `json:"customer_user_id"`
	CustomerUserName           sql.NullString `json:"customer_user_name"`
	CustomerUserToken          sql.NullString `json:"customer_user_token"`
	CustomerUserFirstLoginTime sql.NullString `json:"customer_user_first_login_time"`
	CustomerGender             sql.NullString `json:"customer_gender"`
	CustomerProfileStatus      sql.NullString `json:"customer_profile_status"`
	ModifiedDate               sql.NullString `json:"modified_date"`
	ModifiedBy                 sql.NullString `json:"modified_by"`
	UserID                     sql.NullInt64  `json:"user_id"`
	RegionID                   sql.NullString `json:"region_id"`
	RegionGroupID              sql.NullString `json:"region_group_id"`
	CreatedDate                sql.NullString `json:"created_date"`
	CustomerPriceListID        sql.NullString `json:"customer_price_list_id"`
	CustomerPriceListName      sql.NullString `json:"customer_price_list_name"`
	ShowInApp                  sql.NullString `json:"show_in_app"`
	IsDataComplete             bool           `json:"is_data_complete"`
	SalesmanTypeCode           sql.NullString `json:"salesman_type_code"`
	SalesmanTypeName           sql.NullString `json:"salesman_type_name"`
	CustomerAdminValidate      bool           `json:"customer_admin_validate"`
	IndexPoint                 int            `json:"index_point"`
}

// CustomerParameter ...
type WebCustomerParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	Name           string `json:"customer_name"`
	CustomerTypeId string `json:"custome_type_id"`
	SalesmanTypeID string `json:"salesman_type_id"`
	UserId         string `json:"admin_user_id"`
	BranchId       string `json:"branch_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
	PhoneNumber    string `json:"phone_number"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	ShowInApp      string `json:"show_in_app"`
	Active         string `json:"active"`
	IsDataComplete string `json:"is_data_complete"`
	AdminValidate  string `json:"admin_validate"`
}

// WebCustomerReportParameter ...
type WebCustomerReportParameter struct {
	RegionGroupID         string `json:"region_group_id"`
	RegionID              string `json:"region_id"`
	BranchArea            string `json:"branch_area"`
	CustomerTypeID        string `json:"customer_type_id"`
	BranchIDs             string `json:"branch_ids"`
	CustomerLevelID       string `json:"customer_level_id"`
	CustomerProfileStatus string `json:"customer_profile_status"`
	AdminUserID           string `json:"admin_user_id"`
}

var (
	// CustomerOrderBy ...
	WebCustomerOrderBy = []string{"c.id", "c.customer_name", "c.show_in_apps", "c.active", "c.customer_phone", "c.created_date"}
	// CustomerOrderByrByString ...
	WebCustomerOrderByrByString = []string{
		"c.customer_name",
	}

	//CustomerGenderList ...
	CustomerGenderList = []string{"male", "female"}

	//CustomerProfileStatusComplete ...
	CustomerProfileStatusComplete   = "Lengkap"
	CustomerProfileStatusIncomplete = "Tidak Lengkap"

	// CustomerSelectStatement ...

	WebCustomerSelectStatement = `
	SELECT 
		C.ID AS CUST_ID,
		C.CUSTOMER_CODE AS CUST_CODE,
		C.CUSTOMER_NAME AS CUST_NAME,
		C.CUSTOMER_CP_NAME AS CUST_CP_NAME,
		C.CUSTOMER_ADDRESS AS CUST_ADDRESS,
		C.CUSTOMER_PROFILE_PICTURE AS CUST_PROFILE_PICTURE,
		C.CUSTOMER_EMAIL AS CUST_EMAIL,
		C.CUSTOMER_BIRTHDATE AS BIRTHDATE,
		C.CUSTOMER_RELIGION AS RELIGION,
		C.CUSTOMER_GENDER AS GENDER,
		CASE C.ACTIVE
				WHEN 1 THEN 'Active'
				WHEN 0 THEN 'Inactive'
		END AS CUST_ACTIVE_STATUS,
		C.LATITUDE AS CUST_LATITUDE,
		C.LONGITUDE AS CUST_LONGITUDE,
		B.BRANCH_CODE AS BRANCH_CODE,
		B._NAME AS BRANCH_NAME,
		B.AREA AS BRANCH_AREA,
		B.ADDRESS AS BRANCH_ADRESS,
		B.LATITUDE AS BRANCH_LATITUDE,
		B.LONGITUDE AS BRANCH_LONGITUDE,
		B.PIC_PHONE_NO AS PIC_PHONE_NO,
		B.PIC_NAME AS PIC_NAME,
		REG.CODE AS REGION_CODE,
		REG._NAME AS REGION_NAME,
		REG.GROUP_NAME AS REGION_GROUP,
		PRV.ID AS CUST_PROVINCE_ID,
		PRV._NAME AS CUST_PROVINCE_NAME,
		CTY.ID AS CUST_CITY_ID,
		CTY._NAME AS CUST_CITY_NAME,
		DIST.ID AS CUST_DISTRICT_ID,
		DIST._NAME AS CUST_DISTRICT_NAME,
		SDIST.ID AS CUST_SUBDISTRICT_ID,
		SDIST._NAME AS CUST_SUBDISTRICT_NAME,
		S.SALESMAN_CODE as CUST_SALESMAN_CODE,
		S.SALESMAN_NAME AS CUST_SALESMAN_NAME,
		S.SALESMAN_PHONE_NO AS CUST_SALESMAN_PHONE,
		C.SALES_CYCLE CUST_SALES_CYCLE,
		C.CUSTOMER_TYPE_ID AS CUST_TYPE_ID,
		CT._NAME CUST_TYPE_NAME,
		customer_phone  AS CUSTOMER_PHONE,
		(SELECT SUM(SIH.TRANSACTION_POINT) FROM SALES_INVOICE_HEADER SIH WHERE SIH.CUST_BILL_TO_ID = C.ID) AS CUSTOMER_POINT,
		CG.GIFT_NAME AS CUST_GIFT_NAME,
		LOY.LOYALTY_NAME AS CUST_LOYALTY_NAME,
		C.tax_calc_method as tax_calc_method,
		C.branch_id as c_branch_id,
		C.salesman_id as c_salesman_id,
		C.customer_photo_ktp AS CUST_KTP_PICTURE,
		c.customer_nik,
		cl._name as cus_level_name,
		c.customer_level_id,
		usr.id as user_id,
		usr.login as user_name,
		usr.fcm_token as user_token,
		usr.first_login_time as first_login_time,
		usr_edited.login,
		c.modified_date,
		C.price_list_id,
		PL._name,
		coalesce(c.show_in_apps,0),
		coalesce(C.is_data_completed,false),
		ST.CODE,
		ST._NAME,
		c.admin_validate,
		c.index_point,
		C.customer_photo_ktp_dashboard,
		C.customer_photo_npwp,
		C.customer_photo_npwp_dashboard
	FROM CUSTOMER C
	LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
	LEFT JOIN REGION REG ON REG.ID = B.REGION_ID
	LEFT JOIN CUSTOMER_TYPE CT ON CT.ID = C.CUSTOMER_TYPE_ID
	left JOIN PROVINCE PRV ON PRV.ID = C.CUSTOMER_PROVINCE_ID
	left JOIN CITY CTY ON CTY.ID = C.CUSTOMER_CITY_ID
	left JOIN DISTRICT DIST ON DIST.ID = C.CUSTOMER_DISTRICT_ID
	left JOIN SUBDISTRICT SDIST ON SDIST.ID = C.CUSTOMER_SUBDISTRICT_ID
	LEFT JOIN SALESMAN S ON S.ID = C.SALESMAN_ID
	LEFT JOIN SALESMAN_TYPE ST ON ST.ID = S.SALESMAN_TYPE_ID
	LEFT JOIN PARTNER PS ON PS.ID = S.PARTNER_ID
	LEFT JOIN CUSTOMER_GIFT CG ON CG.CUSTOMER_ID = C.ID
	LEFT JOIN CUSTOMER_POINT CP ON CP.CUSTOMER_ID = C.ID
	LEFT JOIN LOYALTY LOY ON LOY.CUSTOMER_ID = C.ID
	left join customer_level cl on cl.id = c.customer_level_id
	left join _user usr on usr.id = C.user_id
	left join _user usr_edited on usr_edited.id = c.modified_by
	LEFT JOIN PRICE_LIST PL ON PL.ID = C.PRICE_LIST_ID
	`

	// CustomerWhereStatement ...
	WebCustomerWhereStatement = ` WHERE c.show_in_apps = 1 and c.created_date IS not NULL `

	// CustomerWhereStatement ...
	WebCustomerWhereStatementAll = ` WHERE c.created_date IS not NULL `

	WebCustomerWithInvoiceSelectStatement = `with invoice_count as(
		select c.id, count(coh.id) as invoices
		from customer c
		left join customer_order_header coh on coh.cust_bill_to_id = c.id
		left join sales_invoice_header sih on sih.transaction_source_document_no = coh.document_no
		where c.show_in_apps = 1 and sih.id is not null
		{WHERE_CONDITION} 
		group by c.id
	) 
	select C.ID AS CUST_ID,
		  C.CUSTOMER_CODE AS CUST_CODE,
		  C.CUSTOMER_NAME AS CUST_NAME,
		  C.CUSTOMER_CP_NAME AS CUST_CP_NAME,
		  C.CUSTOMER_ADDRESS AS CUST_ADDRESS,
		  C.CUSTOMER_PROFILE_PICTURE AS CUST_PROFILE_PICTURE,
		  C.CUSTOMER_EMAIL AS CUST_EMAIL,
		  C.CUSTOMER_BIRTHDATE AS BIRTHDATE,
		  C.CUSTOMER_RELIGION AS RELIGION,
		  C.CUSTOMER_GENDER AS GENDER,
		  CASE C.ACTIVE
						  WHEN 1 THEN 'Active'
						  WHEN 0 THEN 'Inactive'
		  END AS CUST_ACTIVE_STATUS,
		  C.LATITUDE AS CUST_LATITUDE,
		  C.LONGITUDE AS CUST_LONGITUDE,
		  B.BRANCH_CODE AS BRANCH_CODE,
		  B._NAME AS BRANCH_NAME,
		  B.AREA AS BRANCH_AREA,
		  B.ADDRESS AS BRANCH_ADRESS,
		  B.LATITUDE AS BRANCH_LATITUDE,
		  B.LONGITUDE AS BRANCH_LONGITUDE,
		  B.PIC_PHONE_NO AS PIC_PHONE_NO,
		  REG.CODE AS REGION_CODE,
		  REG._NAME AS REGION_NAME,
		  REG.GROUP_NAME AS REGION_GROUP,
		  PRV.ID AS CUST_PROVINCE_ID,
		  PRV._NAME AS CUST_PROVINCE_NAME,
		  CTY.ID AS CUST_CITY_ID,
		  CTY._NAME AS CUST_CITY_NAME,
		  DIST.ID AS CUST_DISTRICT_ID,
		  DIST._NAME AS CUST_DISTRICT_NAME,
		  SDIST.ID AS CUST_SUBDISTRICT_ID,
		  SDIST._NAME AS CUST_SUBDISTRICT_NAME,
		  S.SALESMAN_CODE as CUST_SALESMAN_CODE,
		  S.SALESMAN_NAME AS CUST_SALESMAN_NAME,
		  S.SALESMAN_PHONE_NO AS CUST_SALESMAN_PHONE,
		  C.SALES_CYCLE CUST_SALES_CYCLE,
		  C.CUSTOMER_TYPE_ID AS CUST_TYPE_ID,
		  CT._NAME CUST_TYPE_NAME,
		  customer_phone  AS CUSTOMER_PHONE,
		  (SELECT SUM(SIH.TRANSACTION_POINT) FROM SALES_INVOICE_HEADER SIH WHERE SIH.CUST_BILL_TO_ID = C.ID) AS CUSTOMER_POINT,
		  CG.GIFT_NAME AS CUST_GIFT_NAME,
		  LOY.LOYALTY_NAME AS CUST_LOYALTY_NAME,
		  C.tax_calc_method as tax_calc_method,
		  C.branch_id as c_branch_id,
		  C.salesman_id as c_salesman_id,
		  C.customer_photo_ktp AS CUST_KTP_PICTURE,
		  c.customer_nik,
		  cl._name as cus_level_name,
		  c.customer_level_id,
		  usr.id as user_id,
		  usr.login as user_name
	FROM CUSTOMER C
	  LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
	  LEFT JOIN REGION REG ON REG.ID = B.REGION_ID
	  LEFT JOIN CUSTOMER_TYPE CT ON CT.ID = C.CUSTOMER_TYPE_ID
	  left JOIN PROVINCE PRV ON PRV.ID = C.CUSTOMER_PROVINCE_ID
	  left JOIN CITY CTY ON CTY.ID = C.CUSTOMER_CITY_ID
	  left JOIN DISTRICT DIST ON DIST.ID = C.CUSTOMER_DISTRICT_ID
	  left JOIN SUBDISTRICT SDIST ON SDIST.ID = C.CUSTOMER_SUBDISTRICT_ID
	  LEFT JOIN SALESMAN S ON S.ID = C.SALESMAN_ID
	  LEFT JOIN PARTNER PS ON PS.ID = S.PARTNER_ID
	  LEFT JOIN CUSTOMER_GIFT CG ON CG.CUSTOMER_ID = C.ID
	  LEFT JOIN CUSTOMER_POINT CP ON CP.CUSTOMER_ID = C.ID
	  LEFT JOIN LOYALTY LOY ON LOY.CUSTOMER_ID = C.ID
	  left join customer_level cl on cl.id = c.customer_level_id
	  left join _user usr on usr.id = C.user_id
	  left join invoice_count ic on ic.id = c.id `

	// WebCustomerWithInvoiceWhereStatement ...
	WebCustomerWithInvoiceWhereStatement = ` WHERE c.created_date IS not NULL and c.show_in_apps = 1 and ic.invoices > 0`

	WebCustomerWithInvoiceGroupByStatement = `group by c.id, b.id, reg.id, prv.id,
		cty.id, dist.id, sdist.id, ps.id, 
		s.id, ct.id, cg.id, loy.id, 
		cl.id,usr.id`
)

// '+62' || regexp_replace(SUBSTRING (customer_phone, 2, length(customer_phone)),'[^\w]+','','g' )  AS CUSTOMER_PHONE,
