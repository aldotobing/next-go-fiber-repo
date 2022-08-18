package models

type MuUser struct {
	ID       *string `json:"id_user"`
	BranchID *string `json:"id_branch"`

	BranchCoverageStr    *string `json:"branch_coverage_str"`
	FbId                 *string `json:"id_facebook"`
	GoogleId             *string `json:"id_google"`
	AppleId              *string `json:"id_apple"`
	Name                 *string `json:"name"`
	UserName             *string `json:"username"`
	Password             *string `json:"password"`
	Gender               *string `json:"gender"`
	QrCode               *string `json:"qr_code"`
	Email                *string `json:"email_user"`
	NoTelp               *string `json:"no_telp"`
	Address              *string `json:"address_user"`
	Level                *int    `json:"level"`
	BirthDate            *string `json:"birthdate"`
	BirthDatePlace       *string `json:"birthplace"`
	ReferalCode          *string `json:"referral_code"`
	NIK                  *string `json:"nik_user"`
	ImgKTP               *string `json:"img_ktp"`
	ImgProfile           *string `json:"img_profile"`
	VerifStatusKtp       *string `json:"verif_status_ktp"`
	RoleGroupId          *int    `json:"role_group_id"`
	ReferralCodeLimitUse *string `json:"referral_limit_use"`
	UserActive           *int    `json:"active_user"`
	CreatedAt            *string `json:"created_at_user"`
	UpdatedAt            *string `json:"updated_at_user"`
	DeletedAt            *string `json:"deleted_at_user"`
	CreatedBy            *int    `json:"created_by_user"`
	UpdatedBy            *int    `json:"updated_by_user"`
	DeletedBy            *int    `json:"deleted_by_user"`
	JoinDate             *string `json:"join_date"`
}

type MuUserDataBreakDown struct {
	ID                   *string `json:"id_user"`
	BranchID             *int    `json:"id_branch"`
	BranchCoverageStr    *string `json:"branch_coverage_str"`
	FbId                 *string `json:"id_facebook"`
	GoogleId             *string `json:"id_google"`
	AppleId              *string `json:"id_apple"`
	Name                 *string `json:"name"`
	UserName             *string `json:"username"`
	Password             *string `json:"password"`
	Gender               *string `json:"gender"`
	QrCode               *string `json:"qr_code"`
	Email                *string `json:"email_user"`
	NoTelp               *string `json:"no_telp"`
	Address              *string `json:"address_user"`
	Level                *int    `json:"level"`
	BirthDate            *string `json:"birthdate"`
	ReferalCode          *string `json:"referral_code"`
	NIK                  *string `json:"nik_user"`
	ImgKTP               *string `json:"img_ktp"`
	ImgProfile           *string `json:"img_profile"`
	ReferralCodeLimitUse *string `json:"referral_limit_use"`
	CreatedAt            *string `json:"created_at_user"`
	UpdatedAt            *string `json:"updated_at_user"`
	DeletedAt            *string `json:"deleted_at_user"`
	ActiveStatus         *int    `json:"active_status"`
	OldUserId            *int    `json:"old_user_id"`
	JoinDate             *string `json:"join_date"`
	UserGiveRefferal     *int    `json:"user_give_referral"`
	CityID               *int    `json:"city_id"`
}

type MuUserParameter struct {
	ID            string `json:"id_user"`
	BranchID      string `json:"id_branch"`
	FbId          string `json:"id_facebook"`
	GoogleId      string `json:"id_google"`
	AppleId       string `json:"id_apple"`
	Name          string `json:"name"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Gender        string `json:"gender"`
	QrCode        string `json:"qr_code"`
	Email         string `json:"email_user"`
	NoTelp        string `json:"no_telp"`
	Address       string `json:"address_user"`
	Level         int    `json:"level"`
	BirthDate     string `json:"birthdate"`
	BirtDatePlace string `json:"birthplace"`
	ReferalCode   string `json:"referral_code"`

	RoleGroupID string `json:"role_group_id"`
	Search      string `json:"search"`
	Page        int    `json:"page"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	By          string `json:"by"`
	Sort        string `json:"sort"`
}

var (
	MuUserOrderBy = []string{"def.id_user", "def.name", "def.username", "def.created_at_user", "def.updated_at_user"}

	MuUserOrderByrByString = []string{
		"def.username",
	}

	MuUserSelectStatement = `SELECT def.id_user, def.id_branch, def.id_facebook, def.id_google, 
	def.id_apple, def.name, def.username, def.password, 
	def.gender, def.qr_code,  def.level, def.referral_code, 
	def.email_user,def.no_telp,def.address_user,def.role_group_id,
	 def.birthdate::date, rg.role_group_name, 
	 coalesce((
		select STRING_AGG(name_branch,',') from mp_branch
		where id_branch in (select id_branch from mu_user_branch where id_user =def.id_user)
	 ),''),
	 def.created_at_user,
	 def.updated_at_user,def.deleted_at_user, def.created_by_user, def.updated_by_user, 
	 def.deleted_by_user, mpb.name_branch, def.referral_limit_use
	FROM mu_user def
	left join mp_branch mpb on mpb.id_branch = def.id_branch
	left join mp_role_group rg on rg.id_role_group = def.role_group_id
	`

	MuUserWhereStatement = `WHERE def.deleted_at_user IS NULL`
)
