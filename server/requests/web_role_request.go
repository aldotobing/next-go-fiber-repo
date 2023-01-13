package requests

// WeebRoleRequest ...
type WeebRoleRequest struct {
	ID     string `json:"id_role" validate:"required" `
	Name   string `json:"name" validate:"required"`
	Header string `json:"header"`
}
