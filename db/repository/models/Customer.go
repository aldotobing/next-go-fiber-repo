package models

// Customer ...
type Customer struct {
	ID                    *string `json:"customer_id"`
	Code                  *string `json:"customer_code"`
	Name                  *string `json:"customer_name"`
	CustomerCpName        *string `json:"customer_cp_name"`
	Address               *string `json:"customer_address"`
	CustomerProvince      *string `json:"customer_province"`
	CustomerCity          *string `json:"customer_city"`
	CustomerDistrict      *string `json:"customer_district"`
	CustomerSubdistrict   *string `json:"customer_subdistrict"`
	CustomerSalesmanCode  *string `json:"customer_salesman_code"`
	CustomerSalesmanName  *string `json:"customer_salesman_name"`
	CustomerSalesmanPhone *string `json:"customer_salesman_phone"`
	CustomerSalesCycle    *string `json:"customer_sales_cycle"`
	CustomerTypeId        *string `json:"customer_type_id"`
	CustomerTypeName      *string `json:"customer_type_name"`
	Phone                 *string `json:"customer_phone"`
	Point                 *string `json:"customer_point"`
	GiftName              *string `json:"customer_gift_name"`
	Loyalty               *string `json:"customer_loyalty"`
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
	CustomerOrderBy = []string{"def.id", "def.customer_name", "def.created_date"}
	// CustomerOrderByrByString ...
	CustomerOrderByrByString = []string{
		"def.customer_name",
	}

	// CustomerSelectStatement ...
	/*
	 */
	CustomerSelectStatement = `
	SELECT DEF.ID,
		DEF.CUSTOMER_CODE,
		DEF.CUSTOMER_NAME,
		CT.ID,
		CT._NAME,
		DEF.CUSTOMER_ADDRESS,
		DEF.CUSTOMER_PHONE,
		CP.POINT,
		CG.GIFT_NAME,
		CG.GIFT_DESCRIPTION,
		LOY.LOYALTY_NAME
	FROM CUSTOMER DEF
	JOIN CUSTOMER_TYPE CT ON CT.ID = DEF.CUSTOMER_TYPE_ID
	LEFT JOIN CUSTOMER_GIFT CG ON CG.CUSTOMER_ID = DEF.ID
	LEFT JOIN CUSTOMER_POINT CP ON CP.CUSTOMER_ID = DEF.ID
	LEFT JOIN LOYALTY LOY ON LOY.CUSTOMER_ID = DEF.ID
	`

	// CustomerWhereStatement ...
	CustomerWhereStatement = ` WHERE def.created_date IS not NULL `
)
