package models

// WebUser ...
type WebUser struct {
	ID                  *string             `json:"id"`
	Login               *string             `json:"login"`
	Password            *string             `json:"password"`
	CompanyId           *string             `json:"company_id"`
	Active              *string             `json:"active"`
	FirestoreUID        *string             `json:"firestore_uid"`
	FcmToken            *string             `json:"fcm_token"`
	UserRoleGroupList   *[]WebUserRoleGroup `json:"role_group_list"`
	UserRoleGroupIDList *string             `json:"role_group_id_list"`
	BranchIDList        []string            `json:"branch_id_list"`
	BranchList          *[]WebUserBranch    `json:"branch_list"`
}

// WebUserParameter ...
type WebUserParameter struct {
	ID     string `json:"id"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// WebUserOrderBy ...
	WebUserOrderBy = []string{"def.id", "def.login", "def.created_date"}
	// WebUserOrderByrByString ...
	WebUserOrderByrByString = []string{
		"def.login",
	}

	// WebUserSelectStatement ...
	WebUserSelectStatement = `select 
	def.id,def.login,def._password,def.company_id,
	def.active,def.firestoreuid,def.fcm_token
	from _user def
	`

	// WebUserWhereStatement ...
	WebUserWhereStatement = `WHERE def.created_date IS not NULL `
)
