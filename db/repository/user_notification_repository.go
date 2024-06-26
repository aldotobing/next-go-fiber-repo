package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
)

type IUserNotificationRepository interface {
	SelectAll(c context.Context, parameter models.UserNotificationParameter) ([]models.UserNotification, error)
	FindAll(ctx context.Context, parameter models.UserNotificationParameter) ([]models.UserNotification, int, error)
	FindByID(c context.Context, parameter models.UserNotificationParameter) (models.UserNotification, error)
	Add(c context.Context, model *models.UserNotification) (*string, error)
	AddBulk(c context.Context, model []models.UserNotification) (err error)
	UpdateStatus(c context.Context, model *models.UserNotification) (*string, error)
	UpdateAllStatus(c context.Context, model *models.UserNotification) (*string, error)
	DeleteStatus(c context.Context, model *models.UserNotification) (*string, error)
	DeleteAllStatus(c context.Context, model *models.UserNotification) (*string, error)
}

type UserNotificationRepository struct {
	DB *sql.DB
}

func NewUserNotificationRepository(DB *sql.DB) IUserNotificationRepository {
	return &UserNotificationRepository{DB: DB}
}

// Scan rows
func (repository UserNotificationRepository) scanRows(rows *sql.Rows) (res models.UserNotification, err error) {
	err = rows.Scan(
		&res.ID, &res.UserID, &res.RowID, &res.Type,
		&res.Text,
		&res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
		&res.CreatedBy, &res.UpdatedBy, &res.DeletedBy,
		&res.Status, &res.Title,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository UserNotificationRepository) scanRow(row *sql.Row) (res models.UserNotification, err error) {
	err = row.Scan(
		&res.ID, &res.UserID, &res.RowID, &res.Type,
		&res.Text,
		&res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
		&res.CreatedBy, &res.UpdatedBy, &res.DeletedBy,
		&res.Status, &res.Title,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository UserNotificationRepository) SelectAll(c context.Context, parameter models.UserNotificationParameter) (data []models.UserNotification, err error) {
	conditionString := ``

	if parameter.UserID != "" {
		conditionString += ` and def.user_id = ` + parameter.UserID
	}

	statement := models.UserNotificationSelectStatement + ` ` + models.UserNotificationWhereStatement +
		`  ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort

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

// FindAll ...
func (repository UserNotificationRepository) FindAll(ctx context.Context, parameter models.UserNotificationParameter) (data []models.UserNotification, count int, err error) {
	conditionString := ``

	query := models.UserNotificationSelectStatement + ` ` + models.UserNotificationWhereStatement + ` ` + conditionString + `
		 ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`

	rows, err := repository.DB.Query(query, parameter.Offset, parameter.Limit)
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

	query = `SELECT COUNT(*) FROM "user_notification" def ` + models.UserNotificationWhereStatement + ` ` +
		conditionString + `  `
	err = repository.DB.QueryRow(query).Scan(&count)

	return data, count, err
}

func (repository UserNotificationRepository) FindByID(c context.Context, parameter models.UserNotificationParameter) (data models.UserNotification, err error) {
	statement := models.UserNotificationSelectStatement + ` WHERE def.deleted_date IS NULL AND def.id = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserNotificationRepository) Add(c context.Context, model *models.UserNotification) (res *string, err error) {
	statement := `INSERT INTO user_notification (user_id, row_id, type_notification, notification_text,
		created_date, created_by,notification_title,notification_status)
	VALUES ($1, $2, $3, $4, $5, $6,$7,'unread') RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.UserID, model.RowID, model.Type,
		model.Text, model.CreatedAt, model.CreatedBy, model.Title).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository UserNotificationRepository) AddBulk(c context.Context, model []models.UserNotification) (err error) {
	var valueStatement string
	for i := range model {
		if valueStatement == "" {
			valueStatement += `(
				'` + *model[i].UserID + `', '` + *model[i].RowID + `', '` + *model[i].Type + `',
				'` + *model[i].Text + `', now(),'` + *model[i].Title + `','unread')`
		} else {
			valueStatement += `, (
				'` + *model[i].UserID + `', '` + *model[i].RowID + `', '` + *model[i].Type + `',
				'` + *model[i].Text + `', now(),'` + *model[i].Title + `','unread')`
		}
	}
	statement := `INSERT INTO user_notification (user_id, row_id, type_notification, notification_text,
		created_date, notification_title, notification_status)
	VALUES ` + valueStatement

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

func (repository UserNotificationRepository) UpdateStatus(c context.Context, model *models.UserNotification) (res *string, err error) {
	statement := `UPDATE user_notification SET  notification_status = 'read', 
	modified_date = $1, modified_by = $2

	WHERE id = $3
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement, model.UpdatedAt, model.UpdatedBy,
		model.ID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository UserNotificationRepository) UpdateAllStatus(c context.Context, model *models.UserNotification) (res *string, err error) {
	statement := `UPDATE user_notification SET  notification_status = 'read', 
	modified_date = $1, modified_by = $2

	WHERE user_id = $3
	RETURNING user_id`

	err = repository.DB.QueryRowContext(c, statement, model.UpdatedAt, model.UpdatedBy,
		model.UserID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository UserNotificationRepository) DeleteStatus(c context.Context, model *models.UserNotification) (res *string, err error) {
	statement := `UPDATE user_notification SET  notification_status = 'read', 
	deleted_date = now()

	WHERE id = $1
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		model.ID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository UserNotificationRepository) DeleteAllStatus(c context.Context, model *models.UserNotification) (res *string, err error) {
	statement := `UPDATE user_notification SET  notification_status = 'read', 
	deleted_date = now()

	WHERE user_id = $1
	RETURNING user_id`

	err = repository.DB.QueryRowContext(c, statement,
		model.UserID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
