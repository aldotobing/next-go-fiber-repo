package models

// District ...
type District struct {
	ID        string  `json:"id"`
	CityID    string  `json:"city_id"`
	City      City    `json:"city"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

// DistrictParameter ...
type DistrictParameter struct {
	ID     string `json:"id"`
	CityID string `json:"city_id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// DistrictOrderBy ...
	DistrictOrderBy = []string{"def.id", "def.code", "def.name", "def.created_at", "def.updated_at"}
	// DistrictOrderByrByString ...
	DistrictOrderByrByString = []string{
		"def.code", "def.name",
	}

	// DistrictSelectStatement ...
	DistrictSelectStatement = `SELECT def.id, def.city_id, def.code, def.name, def.status, def.created_at, def.updated_at, def.deleted_at, c.name, p.name
	FROM districts def
	LEFT JOIN cities c ON def.city_id = c.id
	LEFT JOIN provinces p ON c.province_id = p.id`

	// DistrictWhereStatement ...
	DistrictWhereStatement = `WHERE def.deleted_at IS NULL`
)
