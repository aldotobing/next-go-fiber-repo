package requests

// PromoContentRequest ...
type NewsRequest struct {
	ID          string `json:"id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      string `json:"active"`
	ImageUrl    string `json:"image_url"`
	Priority    string `json:"priority"`
}

type NewsBulkRequest struct {
	News []NewsRequest `json:"news"`
}
