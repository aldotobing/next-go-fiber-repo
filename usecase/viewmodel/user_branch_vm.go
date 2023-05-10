package viewmodel

type WebUserBranchVM struct {
	ID         *string `json:"id"`
	UserID     *string `json:"user_id"`
	BranchID   *string `json:"branch_id"`
	BranchName *string `json:"branch_name"`
}
