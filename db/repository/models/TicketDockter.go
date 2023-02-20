package models

// TicketDokter ...
type TicketDokter struct {
	ID               *string `json:"ticket_id"`
	DoctorID         *string `json:"doctor_id"`
	DoctorName       *string `json:"doctor_name"`
	TicketCode       *string `json:"ticket_code"`
	CustomerID       *string `json:"customer_id"`
	CustomerName     *string `json:"customer_name"`
	CustomerHeight   *string `json:"customer_height"`
	CustomerWeight   *string `json:"customer_weight"`
	CustomerAge      *string `json:"customer_age"`
	CustomerPhone    *string `json:"customer_phone"`
	CustomerAltPhone *string `json:"customer_alt_phone"`
	CustomerProblem  *string `json:"customer_problem"`
	Solution         *string `json:"solution"`
	Allergy          *string `json:"allergy"`
	Status           *string `json:"status"`
	StatusReason     *string `json:"status_reason"`
	CreatedDate      *string `json:"created_date"`
	ModifiedDate     *string `json:"modified_date"`
	CloseDate        *string `json:"close_date"`
	Description      *string `json:"description"`
	Hide             *string `json:"hide"`
}

type TicketDokterBreakDown struct {
	ID           *string `json:"id"`
	TicketCode   *string `json:"ticket_code"`
	CustomerID   *string `json:"customer_id"`
	CustomerName *string `json:"customer_name"`
	DoctorID     *string `json:"doctor_id"`
	DoctorName   *string `json:"doctor_name"`
}

// TicketDokterParameter ...
type TicketDokterParameter struct {
	ID         string `json:"ticket_id"`
	TicketCode string `json:"ticket_code"`
	CustomerID string `json:"customer_id"`
	DoctorID   string `json:"doctor_id"`
	Status     string `json:"status"`
	Search     string `json:"search"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	By         string `json:"by"`
	Sort       string `json:"sort"`
}

var (
	// TicketDokterOrderBy ...
	TicketDokterOrderBy = []string{"td.id", "td.customer_name", "td.created_date"}
	// TicketDokterOrderByrByString ...
	TicketDokterOrderByrByString = []string{
		"td.customer_name",
	}

	// TicketDokterSelectStatement ...
	TicketDokterSelectStatement = `
	SELECT 
		TD.ID AS TICKET_ID,
		TD.DOCTOR_ID AS DOCTOR_ID,
		TD.DOCTOR_NAME AS DOCTOR_NAME,
		TD.TICKET_CODE AS TICKET_CODE,
		TD.CUSTOMER_ID AS CUSTOMER_ID,
		TD.CUSTOMER_NAME AS CUSTOMER_NAME,
		TD.HEIGHT AS CUSTOMER_HEIGHT,
		TD.WEIGHT AS CUSTOMER_WEIGHT,
		TD.AGE AS CUSTOMER_AGE,
		TD.PHONE AS CUSTOMER_PHONE,
		TD.PHONE AS CUSTOMER_ALT_PHONE,
		TD.PROBLEM AS CUSTOMER_PROBLEM,
		TD.SOLUTION AS SOLUTION,
		TD.ALLERGY AS ALLERGY,
		TD.STATUS AS STATUS,
		TD.STATUS_REASON AS STATUS_REASON,
		to_char(TD.CREATED_DATE, 'DD-MM-YYYY HH24:MI:SS') AS CREATED_DATE,
		to_char(TD.MODIFIED_DATE, 'DD-MM-YYYY HH24:MI:SS') AS MODIFIED_DATE,
		to_char(TD.CLOSE_DATE, 'DD-MM-YYYY HH24:MI:SS') AS CLOSE_DATE,
		TD.DESCRIPTION AS DESCRIPTION,
		TD.HIDE AS HIDE
	FROM ticket_dokter TD
	`
	// TicketDokterWhereStatement ...
	TicketDokterWhereStatement = ` 
	WHERE TD.ID IS NOT NULL
	`
)
