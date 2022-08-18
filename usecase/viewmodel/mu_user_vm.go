package viewmodel

type MuUserVM struct {
	ID             string `json:"id_user"`
	BranchID       string `json:"id_branch"`
	FbId           string `json:"id_facebook"`
	GoogleId       string `json:"id_google"`
	AppleId        string `json:"id_apple"`
	Name           string `json:"name"`
	UserName       string `json:"username"`
	Password       string `json:"password"`
	Gender         string `json:"gender"`
	QrCode         string `json:"qr_code"`
	Email          string `json:"email_user"`
	NoTelp         string `json:"no_telp"`
	Address        string `json:"address_user"`
	Level          int    `json:"level"`
	BirtDate       string `json:"birthdate"`
	BirtDatePlace  string `json:"birthplace"`
	ReferalCode    string `json:"referral_code"`
	NIK            string `json:"nik_user"`
	ImgKTP         string `json:"img_ktp"`
	VerifStatusKtp string `json:"verif_status_ktp"`
	UserActive     int    `json:"active_user"`
	CreatedAt      string `json:"created_at_user"`
	UpdatedAt      string `json:"updated_at_user"`
	DeletedAt      string `json:"deleted_at_user"`
	CreatedBy      int    `json:"created_by_user"`
	UpdatedBy      int    `json:"updated_by_user"`
	DeletedBy      *int   `json:"deleted_by_user"`
}
