package models

// WebUserBranch ...
type WebUserBranch struct {
	ID         *string `json:"id"`
	UserID     *string `json:"user_id"`
	BranchID   *string `json:"branch_id"`
	BranchName *string `json:"branch_name"`
	BranchCode *string `json:"branch_code"`
}

// WebUserBranchParameter ...
type WebUserBranchParameter struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	By       string `json:"by"`
	Sort     string `json:"sort"`
	RegionID string `json:"region_id"`
}

var (
	// WebUserBranchOrderBy ...
	WebUserBranchOrderBy = []string{"def.id", "br._name", "def.created_date"}
	// WebUserBranchOrderByrByString ...
	WebUserBranchOrderByrByString = []string{
		"br._name",
	}

	// WebUserBranchSelectStatement ...
	WebUserBranchSelectStatement = `SELECT def.id, def.user_id, br.id, br._name, br.branch_code
	FROM user_branch def
	join branch br on br.id = def.branch_id
	`

	// WebUserBranchWhereStatement ...
	WebUserBranchWhereStatement = `WHERE br._name IS not NULL `
)
