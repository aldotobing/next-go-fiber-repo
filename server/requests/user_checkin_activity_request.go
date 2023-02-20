package requests

type UserCheckinActivityRequest struct {
	ID     string `json:"va_id"`
	UserID string `json:"user_id"`
}
