package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IVideoPromoteRepository ...
type IVideoPromoteRepository interface {
	SelectAll(c context.Context, parameter models.VideoPromoteParameter) ([]models.VideoPromote, error)
	FindAll(ctx context.Context, parameter models.VideoPromoteParameter) ([]models.VideoPromote, int, error)
	Add(c context.Context, model *models.VideoPromote) (*string, error)
	// FindByID(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// FindByDocumentNo(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// FindByCustomerId(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// Add(c context.Context, model *models.SalesInvoice) (*string, error)
	// Edit(c context.Context, model *models.SalesInvoice) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// VideoPromoteRepository ...
type VideoPromoteRepository struct {
	DB *sql.DB
}

// NewVideoPromoteRepository ...
func NewVideoPromoteRepository(DB *sql.DB) IVideoPromoteRepository {
	return &VideoPromoteRepository{DB: DB}
}

// Scan rows
func (repository VideoPromoteRepository) scanRows(rows *sql.Rows) (res models.VideoPromote, err error) {
	err = rows.Scan(
		&res.ID, &res.Title, &res.Description, &res.StartDate, &res.EndDate, &res.Active, &res.Url,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository VideoPromoteRepository) scanRow(row *sql.Row) (res models.VideoPromote, err error) {
	err = row.Scan(
		&res.ID, &res.Title, &res.Description, &res.StartDate, &res.StartDate, &res.Active, &res.Url,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository VideoPromoteRepository) SelectAll(c context.Context, parameter models.VideoPromoteParameter) (data []models.VideoPromote, err error) {
	conditionString := ``

	// if parameter.StartDate != "" && parameter.EndDate != "" {
	// 	conditionString += ` AND def.start_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	// }

	conditionString += ` AND now()::date between def.start_date and def.end_date `
	statement := models.VideoPromoteSelectStatement + ` ` + models.VideoPromoteWhereStatement +
		` AND (LOWER(def."title") LIKE $1) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort
	fmt.Println(statement)
	fmt.Println(parameter.StartDate)
	rows, err := repository.DB.QueryContext(c, statement, "%"+strings.ToLower(parameter.Title)+"%")

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

// FindAll ...
func (repository VideoPromoteRepository) FindAll(ctx context.Context, parameter models.VideoPromoteParameter) (data []models.VideoPromote, count int, err error) {
	conditionString := ``

	// if parameter.StartDate != "" && parameter.EndDate != "" {
	// 	conditionString += ` AND def.start_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	// }
	query := models.VideoPromoteSelectStatement + ` ` + models.VideoPromoteWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."title") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = `SELECT COUNT(*) FROM "video_promote" def ` + models.SalesInvoiceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."title") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err

}

func (repository VideoPromoteRepository) Add(c context.Context, model *models.VideoPromote) (res *string, err error) {
	statement := `INSERT INTO video_promote (start_date, end_date, title, description, active,url)
	VALUES ($1, $2, $3, $4, $5,$6) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.StartDate, model.EndDate, model.Title, model.Description, model.Active, model.Url).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
