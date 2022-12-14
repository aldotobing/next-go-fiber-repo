package requests

// PromoContentRequest ...
type VideoPromoteRequest struct {
	ID          string `json:"id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      string `json:"active"`
	Url         string `json:"url"`
}
