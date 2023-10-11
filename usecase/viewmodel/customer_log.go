package viewmodel

type CustomerLogVM struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	CustomerCode string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
	TypeChanges  string `json:"type_changes"`
	OldData      string `json:"old_data"`
	NewData      string `json:"new_data"`
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	CreatedAt    string `json:"created_at"`
}
