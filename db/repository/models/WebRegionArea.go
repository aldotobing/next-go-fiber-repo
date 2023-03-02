package models

// WebRegionArea ...
type WebRegionArea struct {
	ID        *string `json:"id"`
	Name      *string `json:"name"`
	Code      *string `json:"code"`
	GroupName *string `json:"group_name"`
}

// WebRegionAreaParameter ...
type WebRegionAreaParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WebRegionAreaOrderBy ...
	WebRegionAreaOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// WebRegionAreaOrderByrByString ...
	WebRegionAreaOrderByrByString = []string{
		"def._name",
	}

	// WebRegionAreaSelectStatement ...
	WebRegionAreaSelectStatement = `SELECT def.id, def._name, def.code, def.group_name
	FROM region def
	`

	// WebRegionAreaWhereStatement ...
	WebRegionAreaWhereStatement = `WHERE def._name IS not NULL and def.created_date is not null `
)
