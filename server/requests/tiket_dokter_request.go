package requests

// TicketDokterRequest ...
type TicketDokterRequest struct {
	ID               string `json:"ticket_id"`
	DoctorID         string `json:"doctor_id"`
	DoctorName       string `json:"doctor_name"`
	TicketCode       string `json:"ticket_code"`
	CustomerID       string `json:"customer_id"`
	CustomerName     string `json:"customer_name"`
	CustomerHeight   string `json:"customer_height"`
	CustomerWeight   string `json:"customer_weight"`
	CustomerAge      string `json:"customer_age"`
	CustomerPhone    string `json:"customer_phone"`
	CustomerAltPhone string `json:"customer_alt_phone"`
	CustomerProblem  string `json:"customer_problem"`
	Solution         string `json:"solution"`
	Allergy          string `json:"allergy"`
	Status           string `json:"status"`
	StatusReason     string `json:"status_reason"`
	CreatedDate      string `json:"created_date"`
	ModifiedDate     string `json:"modified_date"`
	CloseDate        string `json:"close_date"`
	Description      string `json:"description"`
	Hide             string `json:"hide"`
}
