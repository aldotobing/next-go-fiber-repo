package models

// Customer ...
type Customer struct {
	ID                      *string `json:"customer_id"`
	Code                    *string `json:"customer_code"`
	CustomerName            *string `json:"customer_name"`
	CustomerCpName          *string `json:"customer_cp_name"`
	CustomerAddress         *string `json:"customer_address"`
	CustomerPostalCode      *string `json:"customer_postal_code"`
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
	Point                   *string `json:"customer_point"`
	GiftName                *string `json:"customer_gift_name"`
	Loyalty                 *string `json:"customer_loyalty"`
}

// CustomerParameter ...
type CustomerParameter struct {
	ID             string `json:"customer_id"`
	Code           string `json:"customer_code"`
	Name           string `json:"customer_name"`
	CustomerTypeId string `json:"custome_type_id"`
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
		PRV._NAME AS CUST_PROVINCE_NAME,
		CTY._NAME AS CUST_CITY_NAME,
		DIST._NAME AS CUST_DISTRICT_NAME,
		SDIST._NAME AS CUST_SUBDISTRICT_NAME,
		PS.CODE as CUST_SALESMAN_CODE,
		PS._NAME AS CUST_SALESMAN_NAME,
		PS.PHONE_NO AS CUST_SALESMAN_PHONE,
		C.SALES_CYCLE CUST_SALES_CYCLE,
		C.CUSTOMER_TYPE_ID AS CUST_TYPE_ID,
		CT._NAME CUST_TYPE_NAME,
		C.CUSTOMER_PHONE AS CUSTOMER_PHONE,
		CP.POINT AS CUST_POINT,
		CG.GIFT_NAME AS CUST_GIFT_NAME,
		LOY.LOYALTY_NAME AS CUST_LOYALTI_NAME
	FROM CUSTOMER C
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
