package viewmodel

type CustomerVM struct {
	ID                       *string `json:"customer_id"`
	Code                     *string `json:"customer_code"`
	CustomerName             *string `json:"customer_name"`
	CustomerProfilePicture   *string `json:"customer_profile_picture"`
	CustomerActiveStatus     *string `json:"customer_active_status"`
	CustomerBirthDate        *string `json:"customer_birthdate"`
	CustomerReligion         *string `json:"customer_religion"`
	CustomerLatitude         *string `json:"customer_latitude"`
	CustomerLongitude        *string `json:"customer_longitude"`
	CustomerBranchCode       *string `json:"customer_branch_code"`
	CustomerBranchName       *string `json:"customer_branch_name"`
	CustomerBranchArea       *string `json:"customer_branch_area"`
	CustomerBranchAddress    *string `json:"customer_branch_address"`
	CustomerBranchLat        *string `json:"customer_branch_lat"`
	CustomerBranchLng        *string `json:"customer_branch_lng"`
	CustomerBranchPicPhoneNo *string `json:"customer_branch_pic_phone_no"`
	CustomerRegionCode       *string `json:"customer_region_code"`
	CustomerRegionName       *string `json:"customer_region_name"`
	CustomerRegionGroup      *string `json:"customer_region_group"`
	CustomerEmail            *string `json:"customer_email"`
	CustomerCpName           *string `json:"customer_cp_name"`
	CustomerAddress          *string `json:"customer_address"`
	CustomerPostalCode       *string `json:"customer_postal_code"`
	CustomerProvinceID       *string `json:"customer_province_id"`
	CustomerProvinceName     *string `json:"customer_province_name"`
	CustomerCityID           *string `json:"customer_city_id"`
	CustomerCityName         *string `json:"customer_city_name"`
	CustomerDistrictID       *string `json:"customer_district_id"`
	CustomerDistrictName     *string `json:"customer_district_name"`
	CustomerSubdistrictID    *string `json:"customer_subdistrict_id"`
	CustomerSubdistrictName  *string `json:"customer_subdistrict_name"`
	CustomerSalesmanCode     *string `json:"customer_salesman_code"`
	CustomerSalesmanName     *string `json:"customer_salesman_name"`
	CustomerSalesmanPhone    *string `json:"customer_salesman_phone"`
	CustomerSalesCycle       *string `json:"customer_sales_cycle"`
	CustomerTypeId           *string `json:"customer_type_id"`
	CustomerTypeName         *string `json:"customer_type_name"`
	CustomerPhone            *string `json:"customer_phone"`
	CustomerPoint            *string `json:"customer_point"`
	GiftName                 *string `json:"customer_gift_name"`
	Loyalty                  *string `json:"customer_loyalty"`
	VisitDay                 *string `json:"visit_day"`
	CustomerTaxCalcMethod    *string `json:"customer_tax_calc_method"`
	CustomerBranchID         *string `json:"customer_branch_id"`
	CustomerSalesmanID       *string `json:"customer_salesman_id"`
	CustomerNik              *string `json:"customer_nik"`
	CustomerPhotoKtp         *string `json:"customer_photo_ktp"`
	CustomerLevelID          *int    `json:"customer_level_id"`
	CustomerLevel            *string `json:"customer_level_name"`
	CustomerUserID           *string `json:"customer_user_id"`
	CustomerUserName         *string `json:"customer_user_name"`
	CustomerGender           *string `json:"customer_gender"`
	CustomerProfileStatus    *string `json:"customer_profile_status"`
}