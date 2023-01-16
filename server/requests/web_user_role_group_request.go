package requests

// WebUserRoleGroupRequest ...
type WebUserRoleGroupRequest struct {
	ID     string `json:"id_role" validate:"required" `
	Name   string `json:"name" validate:"required"`
	Header string `json:"header"`
}
