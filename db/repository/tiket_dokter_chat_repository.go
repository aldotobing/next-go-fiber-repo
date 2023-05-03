package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ITicketDokterChat ...
type ITicketDokterChat interface {
	FindByTicketDocterID(c context.Context, parameter models.TicketDokterChatParameter, ticketDokterID string) ([]models.TicketDokterChat, error)
	Add(c context.Context, parameter *models.TicketDokterChat) (*string, *string, error)
	Delete(c context.Context, id string) error
}

// TicketDokterChat ...
type TicketDokterChat struct {
	DB *sql.DB
}

// NewTicketDokterChat ...
func NewTicketDokterChatRepository(DB *sql.DB) ITicketDokterChat {
	return &TicketDokterChat{DB: DB}
}

// Scan rows
func (repository TicketDokterChat) scanRows(rows *sql.Rows) (res models.TicketDokterChat, err error) {
	err = rows.Scan(
		&res.ID,
		&res.TicketDokterID,
		&res.Description,
		&res.ChatBy,
		&res.CreatedDate,
	)

	return
}

// Scan row
func (repository TicketDokterChat) scanRow(row *sql.Row) (res models.TicketDokterChat, err error) {
	err = row.Scan(
		&res.ID,
		&res.TicketDokterID,
		&res.Description,
		&res.ChatBy,
		&res.CreatedDate,
	)

	return
}

// FindByTicketDocterID ...
func (repository TicketDokterChat) FindByTicketDocterID(c context.Context, parameter models.TicketDokterChatParameter, ticketDokterID string) (data []models.TicketDokterChat, err error) {
	whereStatement := `AND def.ticket_dokter_id = ` + ticketDokterID
	statement := models.TicketDokterChatSelectStatement +
		models.TicketDokterChatWhereStatement + whereStatement +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	rows, err := repository.DB.QueryContext(c, statement)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository TicketDokterChat) Add(c context.Context, model *models.TicketDokterChat) (res *string, createdDate *string, err error) {
	statement := `INSERT INTO ticket_dokter_chat (
		ticket_dokter_id, 
		chat_by,
		description,
		created_date
	)
	VALUES ($1, $2, $3, now()) RETURNING id, created_date`
	err = repository.DB.QueryRowContext(c, statement,
		model.TicketDokterID,
		model.ChatBy,
		model.Description).Scan(&res, &createdDate)

	return
}

// Delete ...
func (repository TicketDokterChat) Delete(c context.Context, id string) (err error) {
	statement := `UPDATE ticket_dokter_chat SET
		deleted_at = now()
		WHERE id = $1 AND deleted_at IS NULL`
	err = repository.DB.QueryRowContext(c, statement,
		id).Err()

	return
}
