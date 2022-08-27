package requests

// CustomerRequest ...
type CustomerRequest struct {
	ID      string `json:"customer_id"`
	Code    string `json:"customer_code"`
	Name    string `json:"customer_name"`
	Address string `json:"customer_address"`
	Phone   string `json:"customer_phone"`
}

type MpCustomerDataBreakDownRequest struct {
	Name         string  `json:"name"`
	ProvinceID   int     `json:"provinsi_id"`
	OldID        int     `json:"id"`
	NationID     int     `json:"id_nation"`
	LatCustomer  float64 `json:"latitude"`
	LongCustomer float64 `json:"longitude"`
}
