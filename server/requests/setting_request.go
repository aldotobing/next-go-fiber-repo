package requests

// SettingRequest ...
type SettingRequest struct {
	Code    string      `json:"code" validate:"required"`
	Details interface{} `json:"details"`
}
