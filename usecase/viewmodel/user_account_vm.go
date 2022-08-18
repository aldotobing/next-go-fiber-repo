package viewmodel

type UserAccountVM struct {
	Token              string `json:"token"`
	ExpiredDate        string `json:"expired_date"`
	RefreshToken       string `json:"refresh_token"`
	RefreshExpiredDate string `json:"refresh_expired_date"`
	LatestAction       string `json:"latest_action"`
	UserID             string `json:"user_id"`
	Otp                string `json:"otp"`
	RoleList           string `json:"role_list"`
	RoGroupID          string `json:"role_group_id"`
	QrCode             string `json:"qr_code"`
}
