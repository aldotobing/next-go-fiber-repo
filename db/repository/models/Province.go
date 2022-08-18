package models

// Province ...
type Province struct {
	ID        string  `json:"id_province"`
	Code      *string `json:"code_province"`
	Name      *string `json:"name_province"`
	IdNation  *string `json:"id_nation"`
	CreatedAt *string `json:"created_at_province"`
	UpdatedAt *string `json:"updated_at_province"`
	DeletedAt *string `json:"deleted_at_province"`
	CreatedBy *int    `json:"created_by_province"`
	UpdatedBy *int    `json:"updated_by_province"`
	DeletedBy *int    `json:"deleted_by_province"`
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
	ProvinceOrderBy = []string{"def.id_province", "def.created_at_province", "def.updated_at_province"}
	// ProvinceOrderByrByString ...
	ProvinceOrderByrByString = []string{
		"def.code_province", "def.name_province",
	}

	// ProvinceSelectStatement ...
	ProvinceSelectStatement = `SELECT def.id_province, def.code_province, def.name_province, def.created_at_province, def.updated_at_province, def.deleted_at_province
	FROM mp_province def`

	// ProvinceWhereStatement ...
	ProvinceWhereStatement = `WHERE def.deleted_at_province IS NULL`
)
