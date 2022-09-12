package models

// District ...
type District struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
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
	DistrictOrderBy = []string{"def.id", "def.code", "def._name", "def.created_at", "def.created_date"}
	// DistrictOrderByrByString ...
	DistrictOrderByrByString = []string{
		"def.code", "def._name",
	}

	// DistrictSelectStatement ...
	DistrictSelectStatement = `SELECT def.id, def.code, def._name
	FROM district def
	LEFT JOIN city c ON def.city_id = c.id`

	// DistrictWhereStatement ...
	DistrictWhereStatement = `WHERE def._name IS NOT NULL`
)
