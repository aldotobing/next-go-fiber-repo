package viewmodel

// UserInfoVM ...
type UserInfoVM struct {
	Birthdate  string `json:"birthdate"`
	Email      string `json:"email"`
	ID         string `json:"id"`
	MotherName string `json:"mother_name"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

// CashInfoVM ...
type CashInfoVM struct {
	Body []CashInfoBodyVM `json:"body"`
}

// CashInfoBodyVM ...
type CashInfoBodyVM struct {
	Accountno string  `json:"accountno"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}
