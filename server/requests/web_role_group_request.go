package requests

// WebRoleGroupRequest ...
type WebRoleGroupRequest struct {
	Name       string `json:"name" validate:"required"`
	RoleListID string `json:"role_list_id"`
}
