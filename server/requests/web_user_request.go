package requests

// WebUserRequest ...
type WebUserRequest struct {
	Login               string `json:"login"`
	Password            string `json:"password"`
	CompanyId           string `json:"company_id"`
	Active              string `json:"active"`
	FirestoreUID        string `json:"firestore_uid"`
	FcmToken            string `json:"fcm_token"`
	UserRoleGroupIDList string `json:"role_group_id_list"`
}
