package models

// Customer ...
type Customer struct {
	ID                      *string `json:"customer_id"`
	Code                    *string `json:"customer_code"`
	CustomerName            *string `json:"customer_name"`
	CustomerProfilePicture  *string `json:"customer_profile_picture"`
	CustomerActiveStatus    *string `json:"customer_active_status"`
	CustomerLatitude        *string `json:"customer_latitude"`
	CustomerLongitude       *string `json:"customer_longitude"`
	CustomerBranchCode      *string `json:"customer_branch_code"`
	CustomerBranchName      *string `json:"customer_branch_name"`
	CustomerRegionCode      *string `json:"customer_region_code"`
	CustomerRegionName      *string `json:"customer_region_name"`
	CustomerEmail           *string `json:"customer_email"`
	CustomerCpName          *string `json:"customer_cp_name"`
	CustomerAddress         *string `json:"customer_address"`
	CustomerPostalCode      *string `json:"customer_postal_code"`
	CustomerProvinceID      *string `json:"customer_province_id"`
	CustomerProvinceName    *string `json:"customer_province_name"`
	CustomerCityID          *string `json:"customer_city_id"`
	CustomerCityName        *string `json:"customer_city_name"`
	CustomerDistrictID      *string `json:"customer_district_id"`
	CustomerDistrictName    *string `json:"customer_district_name"`
	CustomerSubdistrictID   *string `json:"customer_subdistrict_id"`
	CustomerSubdistrictName *string `json:"customer_subdistrict_name"`
	CustomerSalesmanCode    *string `json:"customer_salesman_code"`
	CustomerSalesmanName    *string `json:"customer_salesman_name"`
	CustomerSalesmanPhone   *string `json:"customer_salesman_phone"`
	CustomerSalesCycle      *string `json:"customer_sales_cycle"`
	CustomerTypeId          *string `json:"customer_type_id"`
	CustomerTypeName        *string `json:"customer_type_name"`
	CustomerPhone           *string `json:"customer_phone"`
	CustomerPoint           *string `json:"customer_point"`
	GiftName                *string `json:"customer_gift_name"`
	Loyalty                 *string `json:"customer_loyalty"`
	VisitDay                *string `json:"visit_day"`
}

// CustomerParameter ...
type CustomerParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	Name           string `json:"customer_name"`
	CustomerTypeId string `json:"custome_type_id"`
	UserId         string `json:"admin_user_id"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// CustomerOrderBy ...
	CustomerOrderBy = []string{"c.id", "c.customer_name", "c.created_date"}
	// CustomerOrderByrByString ...
	CustomerOrderByrByString = []string{
		"c.customer_name",
	}

	// CustomerSelectStatement ...

	CustomerSelectStatement = `
	SELECT 
		C.ID AS CUST_ID,
		C.CUSTOMER_CODE AS CUST_CODE,
		C.CUSTOMER_NAME AS CUST_NAME,
		C.CUSTOMER_CP_NAME AS CUST_CP_NAME,
		C.CUSTOMER_ADDRESS AS CUST_ADDRESS,
		C.CUSTOMER_PROFILE_PICTURE AS CUST_PROFILE_PICTURE,
		C.CUSTOMER_EMAIL AS CUST_EMAIL,
		CASE C.ACTIVE
				WHEN 1 THEN 'Active'
				WHEN 0 THEN 'Inactive'
		END AS CUST_ACTIVE_STATUS,
		C.LATITUDE AS CUST_LATITUDE,
		C.LONGITUDE AS CUST_LONGITUDE,
		B.BRANCH_CODE AS BRANCH_CODE,
		B._NAME AS BRANCH_NAME,
		REG.CODE AS REGION_CODE,
		REG._NAME AS REGION_NAME,
		PRV.ID AS CUST_PROVINCE_ID,
		PRV._NAME AS CUST_PROVINCE_NAME,
		CTY.ID AS CUST_CITY_ID,
		CTY._NAME AS CUST_CITY_NAME,
		DIST.ID AS CUST_DISTRICT_ID,
		DIST._NAME AS CUST_DISTRICT_NAME,
		SDIST.ID AS CUST_SUBDISTRICT_ID,
		SDIST._NAME AS CUST_SUBDISTRICT_NAME,
		PS.CODE as CUST_SALESMAN_CODE,
		PS._NAME AS CUST_SALESMAN_NAME,
		PS.PHONE_NO AS CUST_SALESMAN_PHONE,
		C.SALES_CYCLE CUST_SALES_CYCLE,
		C.CUSTOMER_TYPE_ID AS CUST_TYPE_ID,
		CT._NAME CUST_TYPE_NAME,
		C.CUSTOMER_PHONE AS CUSTOMER_PHONE,
		(SELECT SUM(SIH.TRANSACTION_POINT) FROM SALES_INVOICE_HEADER SIH WHERE SIH.CUST_BILL_TO_ID = C.ID) AS CUSTOMER_POINT,
		CG.GIFT_NAME AS CUST_GIFT_NAME,
		LOY.LOYALTY_NAME AS CUST_LOYALTY_NAME
	FROM CUSTOMER C
	LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
	LEFT JOIN REGION REG ON REG.ID = B.REGION_ID
	LEFT JOIN CUSTOMER_TYPE CT ON CT.ID = C.CUSTOMER_TYPE_ID
	JOIN PROVINCE PRV ON PRV.ID = C.CUSTOMER_PROVINCE_ID
	JOIN CITY CTY ON CTY.ID = C.CUSTOMER_CITY_ID
	JOIN DISTRICT DIST ON DIST.ID = C.CUSTOMER_DISTRICT_ID
	JOIN SUBDISTRICT SDIST ON SDIST.ID = C.CUSTOMER_SUBDISTRICT_ID
	LEFT JOIN SALESMAN S ON S.ID = C.SALESMAN_ID
	LEFT JOIN PARTNER PS ON PS.ID = S.PARTNER_ID
	LEFT JOIN CUSTOMER_GIFT CG ON CG.CUSTOMER_ID = C.ID
	LEFT JOIN CUSTOMER_POINT CP ON CP.CUSTOMER_ID = C.ID
	LEFT JOIN LOYALTY LOY ON LOY.CUSTOMER_ID = C.ID
	`

	// CustomerWhereStatement ...
	CustomerWhereStatement = ` WHERE c.created_date IS not NULL `
)
