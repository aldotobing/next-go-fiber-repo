package models

// City ...
type City struct {
	ID   *string `json:"id_city"`
	Code *string `json:"code"`
	Name *string `json:"name_city"`
}

type MpCityDataBreakDown struct {
	ID         *string  `json:"id_city"`
	Name       *string  `json:"name_city"`
	ProvinceID *int     `json:"id_province"`
	OldID      *int     `json:"old_id"`
	NationID   *int     `json:"id_nation"`
	LatCity    *float64 `json:"lat_city"`
	LongCity   *float64 `json:"long_city"`
}

// CityParameter ...
type CityParameter struct {
	ID         string `json:"id_city"`
	ProvinceID string `json:"id_province"`
	Name       string `json:"name_city"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// CityOrderBy ...
	CityOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// CityOrderByrByString ...
	CityOrderByrByString = []string{
		"def._name",
	}

	// CitySelectStatement ...
	CitySelectStatement = `SELECT def.id, def.code,  def._name
	FROM city def
	`

	// CityWhereStatement ...
	CityWhereStatement = `WHERE def.created_date IS not NULL`
)
