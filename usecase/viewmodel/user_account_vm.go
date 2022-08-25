package viewmodel

type UserAccountVM struct {
	ID                 string  `json:"user_id"`
	Token              string  `json:"token"`
	ExpiredDate        string  `json:"expired_date"`
	RefreshToken       string  `json:"refresh_token"`
	RefreshExpiredDate string  `json:"refresh_expired_date"`
	LatestAction       string  `json:"latest_action"`
	Otp                string  `json:"otp"`
	RoleList           string  `json:"role_list"`
	RoGroupID          string  `json:"role_group_id"`
	CustomerID         string  `json:"customer_id"`
	CustomerName       string  `json:"customer_name"`
	Code               string  `json:"customer_code"`
	Phone              string  `json:"customer_phone"`
	PriceListID        *string `json:"price_list_id"`
	PriceListVersionID *string `json:"price_list_version_id"`
	CustomerTypeID     *string `json:"customer_type_id"`
	CustomerLevelName  *string `json:"customer_level_name"`
}
