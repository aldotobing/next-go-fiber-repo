package requests

// WebRoleGroupRequest ...
type WebRoleGroupRequest struct {
	Name string `json:"name" validate:"required"`
}
