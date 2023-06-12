package models

// TicketDokterChat ...
type TicketDokterChat struct {
	ID             *string `json:"ticket_id"`
	TicketDokterID *int64  `json:"ticket_dokter_id"`
	CreatedDate    *string `json:"created_date"`
	ChatBy         *string `json:"chat_by"`
	Description    *string `json:"description"`
}

// TicketDokterChatParameter ...
type TicketDokterChatParameter struct {
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	By     string `json:"by"`
	Sort   string `json:"sort"`
}

var (
	// TicketDokterChatOrderBy ...
	TicketDokterChatOrderBy = []string{"def.created_date"}
	// TicketDokterChatOrderByrByString ...
	TicketDokterChatOrderByrByString = []string{
		"def.created_date",
	}

	// TicketDokterChatSelectStatement ...
	TicketDokterChatSelectStatement = `
	SELECT 
		def.id, def.ticket_dokter_id, def.description, def.chat_by, def.created_date
	FROM ticket_dokter_chat def
	`
	// TicketDokterChatWhereStatement ...
	TicketDokterChatWhereStatement = ` 
	WHERE def.deleted_at is NULL AND def.ID IS NOT NULL
	`
)
