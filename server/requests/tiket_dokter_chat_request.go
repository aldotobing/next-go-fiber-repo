package requests

// TicketDokterChatRequest ...
type TicketDokterChatRequest struct {
	ID             string `json:"ticket_id"`
	TicketDocterID int64  `json:"ticket_dokter_id" validate:"required"`
	ChatBy         string `json:"chat_by"`
	CreatedDate    string `json:"created_date"`
	Description    string `json:"description"`
}
