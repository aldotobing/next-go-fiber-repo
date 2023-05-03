package viewmodel

type TicketDocterChatVM struct {
	ID             *string `json:"id"`
	TicketDokterID *int64  `json:"ticket_dokter_id"`
	CreatedDate    *string `json:"created_date"`
	ChatBy         *string `json:"chat_by"`
	Description    *string `json:"description"`
}
