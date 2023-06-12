package models

// SubDistrict ...
type SubDistrict struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// SubDistrictParameter ...
type SubDistrictParameter struct {
	ID         string `json:"id"`
	IDs        string `json:"ids"`
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
	SubDistrictOrderBy = []string{"def.id", "def.code", "def._name", "def.created_date"}
	// SubDistrictOrderByrByString ...
	SubDistrictOrderByrByString = []string{
		"def.code", "def._name",
	}

	// SubDistrictSelectStatement ...
	SubDistrictSelectStatement = `SELECT def.id, def.code, def._name
	FROM subdistrict def
	LEFT JOIN district d ON def.district_id = d.id`

	// SubDistrictWhereStatement ...
	SubDistrictWhereStatement = `WHERE def._name IS NOT NULL`
)
