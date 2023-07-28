package requests

// BroadcastRequest ...
type BroadcastRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}
