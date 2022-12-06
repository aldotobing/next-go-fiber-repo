package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// INewsRepository ...
type INewsRepository interface {
	SelectAll(c context.Context, parameter models.NewsParameter) ([]models.News, error)
	FindAll(ctx context.Context, parameter models.NewsParameter) ([]models.News, int, error)
	Add(c context.Context, model *models.News) (*string, error)
	// FindByID(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// FindByDocumentNo(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// FindByCustomerId(c context.Context, parameter models.SalesInvoiceParameter) (models.SalesInvoice, error)
	// Add(c context.Context, model *models.SalesInvoice) (*string, error)
	// Edit(c context.Context, model *models.SalesInvoice) (*string, error)
	// Delete(c context.Context, id string, now time.Time) (string, error)
}

// NewsRepository ...
type NewsRepository struct {
	DB *sql.DB
}

// NewNewsRepository ...
func NewNewsRepository(DB *sql.DB) INewsRepository {
	return &NewsRepository{DB: DB}
}

// Scan rows
func (repository NewsRepository) scanRows(rows *sql.Rows) (res models.News, err error) {
	err = rows.Scan(
		&res.ID, &res.Title, &res.Description, &res.StartDate, &res.EndDate,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan row
func (repository NewsRepository) scanRow(row *sql.Row) (res models.News, err error) {
	err = row.Scan(
		&res.ID, &res.Title, &res.Description, &res.StartDate, &res.StartDate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository NewsRepository) SelectAll(c context.Context, parameter models.NewsParameter) (data []models.News, err error) {
	conditionString := ``

	// if parameter.StartDate != "" && parameter.EndDate != "" {
	// 	conditionString += ` AND def.start_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	// }

	conditionString += ` AND now()::date between def.start_date and def.end_date `
	statement := models.NewsSelectStatement + ` ` + models.NewsWhereStatement +
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
func (repository NewsRepository) FindAll(ctx context.Context, parameter models.NewsParameter) (data []models.News, count int, err error) {
	conditionString := ``

	// if parameter.StartDate != "" && parameter.EndDate != "" {
	// 	conditionString += ` AND def.start_date between '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	// }
	query := models.NewsSelectStatement + ` ` + models.NewsWhereStatement + ` ` + conditionString + `
		AND (LOWER(def."title") LIKE $4  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $5 LIMIT $6`
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

	query = `SELECT COUNT(*) FROM "news" def ` + models.SalesInvoiceWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."title") LIKE $4)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err

}

func (repository NewsRepository) Add(c context.Context, model *models.News) (res *string, err error) {
	statement := `INSERT INTO news (start_date, end_date, title, description, active)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.StartDate, model.EndDate, model.Title, model.Description, 1).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
