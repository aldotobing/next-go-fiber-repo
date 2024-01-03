package viewmodel

type TicketDocterChatVM struct {
	ListTicketChatDoctor []TicketDocterChatDetailVM `json:"list_ticket_chat_doctor"`
}

type TicketDocterChatDetailVM struct {
	ID             *string `json:"id"`
	TicketDokterID *int64  `json:"ticket_dokter_id"`
	CreatedDate    *string `json:"created_date"`
	ChatBy         *string `json:"chat_by"`
	Description    *string `json:"description"`
}
