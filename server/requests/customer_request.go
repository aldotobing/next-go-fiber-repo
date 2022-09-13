package requests

// CustomerRequest ...
type CustomerRequest struct {
	ID                    string `json:"customer_id"`
	Code                  string `json:"customer_code"`
	CustomerName          string `json:"customer_name"`
	CustomerAddress       string `json:"customer_address"`
	CustomerPhone         string `json:"customer_phone"`
	CustomerCityID        string `json:"customer_city_id"`
	CustomerDistrictID    string `json:"customer_district_id"`
	CustomerSubdistrictID string `json:"customer_subdistrict_id"`
	CustomerPostalCode    string `json:"customer_postal_code"`
}

type MpCustomerDataBreakDownRequest struct {
	Name         string  `json:"name"`
	ProvinceID   int     `json:"provinsi_id"`
	OldID        int     `json:"id"`
	NationID     int     `json:"id_nation"`
	LatCustomer  float64 `json:"latitude"`
	LongCustomer float64 `json:"longitude"`
}
