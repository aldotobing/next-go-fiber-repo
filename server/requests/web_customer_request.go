package requests

// CustomerRequest ...
type WebCustomerRequest struct {
	ID                     string `json:"customer_id"`
	Code                   string `json:"customer_code"`
	CustomerName           string `json:"customer_name"`
	CustomerAddress        string `json:"customer_address"`
	CustomerPhone          string `json:"customer_phone"`
	CustomerProvinceID     string `json:"customer_province_id"`
	CustomerCityID         string `json:"customer_city_id"`
	CustomerDistrictID     string `json:"customer_district_id"`
	CustomerSubdistrictID  string `json:"customer_subdistrict_id"`
	CustomerPostalCode     string `json:"customer_postal_code"`
	CustomerCpName         string `json:"customer_cp_name"`
	CustomerEmail          string `json:"customer_email"`
	CustomerProfilePicture string `json:"customer_profile_picture"`
	CustomerTaxCalcMethod  string `json:"customer_tax_calc_method"`
	CustomerBranchID       string `json:"customer_branch_id"`
	CustomerSalesmanID     string `json:"customer_salesman_id"`
	CustomerActiveStatus   string `json:"customer_active_status"`
	CustomerNik            string `json:"customer_nik"`
	CustomerBirthDate      string `json:"customer_birthdate"`
	CustomerReligion       string `json:"customer_religion"`
	CustomerUserID         string `json:"customer_user_id"`
	CustomerLevelID        int    `json:"customer_level_id"`
	CustomerGender         string `json:"customer_gender"`
	UserID                 int    `json:"user_id"`
	CustomerShowInApp      string `json:"customer_show_in_app"`
	AdminValidate          bool   `json:"customer_admin_validate"`
	Note                   string `json:"note"`
}

// WebCustomerBulkRequest ...
type WebCustomerBulkRequest struct {
	Active    string               `json:"active"`
	UserID    string               `json:"user_id"`
	ShowInApp string               `json:"show_in_app"`
	Customers []WebCustomerRequest `json:"customers"`
}
