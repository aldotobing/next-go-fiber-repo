package requests

// CustomerDataSyncRequest ...
type CustomerDataSyncRequest struct {
	ID                *string `json:"customer_id"`
	Code              *string `json:"customer_code"`
	Name              *string `json:"customer_name"`
	CustomerType      *string `json:"customer_type_code"`
	Address           *string `json:"customer_address"`
	PhoneNo           *string `json:"customer_phone_no"`
	PriceListCode     *string `json:"price_list_code"`
	CountryCode       *string `json:"country_code"`
	CityCode          *string `json:"city_code"`
	DistrictCode      *string `json:"district_code"`
	SubDistrictCode   *string `json:"subdistrict_code"`
	ProvinceCode      *string `json:"province_code"`
	BranchID          *string `json:"branch_id"`
	TermOfPaymentCode *string `json:"top_code"`
	SalesmanCode      *string `json:"salesman_code"`
	CustomerLevelCode *string `json:"customer_level_code"`
}
