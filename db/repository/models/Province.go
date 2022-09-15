package models

// Province ...
type Province struct {
	ID   string  `json:"id"`
	Code *string `json:"code_province"`
	Name *string `json:"name_province"`
}

type MpProvinceDataBreakDown struct {
	ID       *string `json:"id_province"`
	Name     *string `json:"name_province"`
	Code     *string `json:"code_province"`
	OldID    *int    `json:"old_id"`
	NationID *int    `json:"id_nation"`
}

// ProvinceParameter ...
type ProvinceParameter struct {
	ID       string `json:"id_province"`
	Code     string `json:"code_province"`
	Name     string `json:"name_province"`
	IdNation string `json:"id_nation"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
}

var (
	// ProvinceOrderBy ...
	ProvinceOrderBy = []string{"def.id", "def.code", "def._name", "def.created_date"}
	// ProvinceOrderByrByString ...
	ProvinceOrderByrByString = []string{
		"def.id", "def._name",
	}

	// ProvinceSelectStatement ...
	ProvinceSelectStatement = `select def.id,def.code,def._name as name 
	from province def`

	// ProvinceWhereStatement ...
	ProvinceWhereStatement = `WHERE def._name IS NOT NULL`
)
