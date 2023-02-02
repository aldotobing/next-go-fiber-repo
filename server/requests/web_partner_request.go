package requests

// CustomerRequest ...
type WebPartnerRequest struct {
	ID             string `json:"partner_id"`
	Code           string `json:"partner_code"`
	PartnerName    string `json:"partner_name"`
	PartnerAddress string `json:"partner_address"`
	PartnerPhone   string `json:"partner_phone"`
	PartnerUserID  string `json:"partner_user_id"`
	PartnerEmail   string `json:"partner_email"`
}
