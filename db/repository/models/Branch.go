package models

// Branch ...
type Branch struct {
	ID   *string `json:"branch_id"`
	Code *string `json:"branch_code"`
	Name *string `json:"branch_name"`
}

// BranchParameter ...
type BranchParameter struct {
	ID       string `json:"branch_id"`
	UserID   string `json:"user_id"`
	RegionID string `json:"region_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
	ExceptId string `json:"except_id"`
}

var (
	// BranchOrderBy ...
	BranchOrderBy = []string{"def.id", "def._name", "def.created_date"}
	// BranchOrderByrByString ...
	BranchOrderByrByString = []string{
		"def._name",
	}

	// BranchSelectStatement ...

	BranchSelectStatement = `
	select def.id,def._name,def.branch_code 
	from branch def
	left join region r on r.id = def.region_id
		`

	// BranchWhereStatement ...
	BranchWhereStatement = ` WHERE def.created_date IS not NULL `
)
