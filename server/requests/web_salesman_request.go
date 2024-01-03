package requests

// CustomerRequest ...
type WebSalesmanRequest struct {
	ID             string `json:"salesman_id"`
	Code           string `json:"salesman_code"`
	PartnerName    string `json:"salesman_name"`
	PartnerAddress string `json:"salesman_address"`
	PartnerPhone   string `json:"salesman_phone"`
	PartnerUserID  string `json:"salesman_user_id"`
	PartnerEmail   string `json:"salesman_email"`
}
