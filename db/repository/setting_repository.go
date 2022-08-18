package repository

import (
	"database/sql"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/interfacepkg"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ISettingRepository ...
type ISettingRepository interface {
	SelectAll() ([]models.Setting, error)
	FindByID(ID string) (models.Setting, error)
	FindByCode(code string) (models.Setting, error)
	Add(data *viewmodel.SettingVM, now time.Time) (string, error)
	Edit(ID string, model *viewmodel.SettingVM, now time.Time) (string, error)
	Delete(ID string, now time.Time) (string, error)
}

// SettingRepository ...
type SettingRepository struct {
	DB *sql.DB
}

// NewSettingRepository ...
func NewSettingRepository(DB *sql.DB) ISettingRepository {
	return &SettingRepository{DB: DB}
}

// Scan rows
func (repository SettingRepository) scanRows(rows *sql.Rows) (res models.Setting, err error) {
	err = rows.Scan(
		&res.ID, &res.Code, &res.Details, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository SettingRepository) scanRow(row *sql.Row) (res models.Setting, err error) {
	err = row.Scan(
		&res.ID, &res.Code, &res.Details, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SelectAll ...
func (repository SettingRepository) SelectAll() (data []models.Setting, err error) {
	statement := models.SettingSelectStatement + ` ` + models.SettingWhereStatement
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, err
}

// FindByID ...
func (repository SettingRepository) FindByID(ID string) (data models.Setting, err error) {
	statement := models.SettingSelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`
	row := repository.DB.QueryRow(statement, ID)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// FindByCode ...
func (repository SettingRepository) FindByCode(code string) (data models.Setting, err error) {
	statement := models.SettingSelectStatement + ` WHERE def.deleted_at IS NULL AND def.code = $1`
	row := repository.DB.QueryRow(statement, code)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Add ...
func (repository SettingRepository) Add(model *viewmodel.SettingVM, now time.Time) (res string, err error) {
	statement := `INSERT INTO settings (code, details, created_at, updated_at)
	VALUES ($1, $2, $3, $3 ) RETURNING id`

	err = repository.DB.QueryRow(statement,
		model.Code, interfacepkg.Marshal(model.Details), now,
	).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

// Edit ...
func (repository SettingRepository) Edit(ID string, model *viewmodel.SettingVM, now time.Time) (res string, err error) {
	statement := `UPDATE settings SET details = $1, updated_at = $2 WHERE id = $3 RETURNING id`

	err = repository.DB.QueryRow(statement,
		interfacepkg.Marshal(model.Details), now, ID,
	).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete ...
func (repository SettingRepository) Delete(ID string, now time.Time) (res string, err error) {
	statement := `UPDATE settings SET updated_at = $1, deleted_at = $1 WHERE id = $2 RETURNING id`

	err = repository.DB.QueryRow(statement, now, ID).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}
