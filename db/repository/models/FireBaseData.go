package models

type FireStoreUser struct {
	ID     string `json:"user_id" firestore:"user_id"`
	Name   string `json:"nickname" firestore:"nickname"`
	UID    string `json:"uid" firestore:"uid"`
	DBSync string `json:"dbsync" firestore:"dbsync"`
}

// FireStoreUserParameter ...
type FireStoreUserParameter struct {
	ID                  string `json:"firestoreuser_id"`
	Code                string `json:"firestoreuser_code"`
	Name                string `json:"firestoreuser_name"`
	FireStoreUserTypeId string `json:"custome_type_id"`
	UserId              string `json:"admin_user_id"`
	Search              string `json:"search"`
	Page                int    `json:"page"`
	Offset              int    `json:"offset"`
	Limit               int    `json:"limit"`
	By                  string `json:"by"`
	Sort                string `json:"sort"`
}

var (
	// FireStoreUserOrderBy ...
	FireStoreUserOrderBy = []string{"def.id", "def.login", "def.created_date"}
	// FireStoreUserOrderByrByString ...
	FireStoreUserOrderByrByString = []string{
		"def.login",
	}

	// FireStoreUserSelectStatement ...

	FireStoreUserSelectStatement = `
	
		SELECT def.firestoreuid
		from _user def
		
	`

	// FireStoreUserWhereStatement ...
	FireStoreUserWhereStatement = ` where def.firestoreuid is not null and trim(def.firestoreuid) !=''
	`
)
