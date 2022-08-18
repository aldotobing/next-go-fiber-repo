package models

// SubDistrict ...
type SubDistrict struct {
	ID         string   `json:"id"`
	DistrictID string   `json:"district_id"`
	District   District `json:"district"`
	Code       string   `json:"code"`
	Name       string   `json:"name"`
	PostalCode string   `json:"postal_code"`
	Status     bool     `json:"status"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	DeletedAt  *string  `json:"deleted_at"`
}

// SubDistrictParameter ...
type SubDistrictParameter struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	PostalCode string `json:"postal_code"`
	Status     string `json:"status"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// SubDistrictOrderBy ...
	SubDistrictOrderBy = []string{"def.id", "def.code", "def.name", "def.postal_code", "def.created_at", "def.updated_at"}
	// SubDistrictOrderByrByString ...
	SubDistrictOrderByrByString = []string{
		"def.code", "def.name", "def.postal_code",
	}

	// SubDistrictSelectStatement ...
	SubDistrictSelectStatement = `SELECT def.id, def.district_id, def.code, def.name, def.postal_code, def.status, def.created_at, def.updated_at, def.deleted_at, d.name, c.name, p.name
	FROM sub_districts def
	LEFT JOIN districts d ON def.district_id = d.id
	LEFT JOIN cities c ON d.city_id = c.id
	LEFT JOIN provinces p ON c.province_id = p.id`

	// SubDistrictWhereStatement ...
	SubDistrictWhereStatement = `WHERE def.deleted_at IS NULL`
)
