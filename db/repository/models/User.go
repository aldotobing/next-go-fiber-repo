package models

// User ...
type User struct {
	ID             string  `json:"id"`
	Code           string  `json:"code"`
	RoleID         string  `json:"role_id"`
	Role           Role    `json:"role"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	EmailValidAt   string  `json:"email_valid_at"`
	Phone          string  `json:"phone"`
	PhoneValidAt   string  `json:"phone_valid_at"`
	ProfilePhotoID string  `json:"profile_photo_id"`
	File           File    `json:"file"`
	Password       string  `json:"password"`
	Status         bool    `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      *string `json:"deleted_at"`
}

// UserParameter ...
type UserParameter struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	RoleID         string `json:"role_id"`
	ProfilePhotoID string `json:"profile_photo_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Status         string `json:"status"`
	Search         string `json:"search"`
	Page           int    `json:"page"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	By             string `json:"by"`
	Sort           string `json:"sort"`
}

var (
	// UserOrderBy ...
	UserOrderBy = []string{"def.id", "def.name", "def.email", "def.created_at", "def.updated_at"}
	// UserOrderByrByString ...
	UserOrderByrByString = []string{
		"def.name", "def.email",
	}

	// UserSelectStatement ...
	UserSelectStatement = `SELECT def.id, def.code, def.role_id, def.name, def.email, def.email_valid_at, 
	def.phone, def.phone_valid_at, def.profile_photo_id, def.password, def.status, def.created_at, 
	def.updated_at, def.deleted_at, r.code, r.name, f.url
	FROM users def
	LEFT JOIN roles r ON r.id = def.role_id
	LEFT JOIN files f ON f.id = def.profile_photo_id`

	// UserWhereStatement ...
	UserWhereStatement = `WHERE def.deleted_at IS NULL`
)
