package viewmodel

// SettingVM ...
type SettingVM struct {
	ID        string      `json:"id"`
	Code      string      `json:"code"`
	Details   interface{} `json:"details"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt string      `json:"deleted_at"`
}
