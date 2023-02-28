package models

// TicketDoctor ...
type TicketDoctor struct {
	ID           *string `json:"id"`
	CustomerId   *string `json:"customer_id"`
	CustomerName *string `json:"customer_name"`
	Height       *string `json:"height"`
	Weight       *string `json:"weight"`
	PhoneAlt     *string `json:"phone_alt"`
	Problem      *string `json:"problem"`
	Allergy      *string `json:"allergy"`
	Status       *string `json:"status"`
	DoctorName   *string `json:"doctor_name"`
	Description  *string `json:"description"`
}

// TicketDoctorParameter ...
type TicketDoctorParameter struct {
	ID         string `json:"id"`
	CustomerId string `json:"customer_id"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
	ExceptId   string `json:"except_id"`
}

var (
	// TicketDoctorteOrderBy ...
	TicketDoctorteOrderBy = []string{"def.id", "def.customer_name"}
	// TicketDoctorteOrderByrByString ...
	TicketDoctorteOrderByrByString = []string{
		"def.customer_name",
	}

	// TicketDoctorteSelectStatement ...
	TicketDoctorteSelectStatement = `
	`

	// TicketDoctorteWhereStatement ...
	TicketDoctorteWhereStatement = ` where def.id is not null`
)
