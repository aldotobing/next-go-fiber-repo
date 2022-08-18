package repository

import (
	"database/sql"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IFile ...
type IFile interface {
	FindAllUnassignedByUserID(userUpload, types string) (data []models.File, err error)
	FindByID(id string) (models.File, error)
	FindUnassignedByID(id, types, userUpload string) (models.File, error)
	FindByType(id, types string) (models.File, error)
	Store(body *viewmodel.FileVM, changedAt time.Time) (string, error)
	UpdateIsUsed(id string, isUsed bool, changedAt time.Time) (string, error)
	Destroy(id string, changedAt time.Time) (string, error)
}

// fileRepository ...
type fileRepository struct {
	DB *sql.DB
}

// NewFileRepository ...
func NewFileRepository(db *sql.DB) IFile {
	return &fileRepository{DB: db}
}

// FindAllUnassignedByUserID ...
func (model fileRepository) FindAllUnassignedByUserID(userUpload, types string) (data []models.File, err error) {
	query := models.FileSelectString + ` WHERE f."deleted_at" IS NULL AND f."user_upload" = $1 AND f."type" = $2
		` + models.UnassignedQueryString + ` ORDER BY f."created_at"`

	rows, err := model.DB.Query(query, userUpload, types)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		d := models.File{}
		err = rows.Scan(
			&d.ID, &d.Type, &d.URL, &d.UserUpload, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	err = rows.Err()

	return data, err
}

// FindByID ...
func (model fileRepository) FindByID(id string) (data models.File, err error) {
	query := `SELECT "id", "type", "url", "user_upload", "created_at", "updated_at", "deleted_at"
		FROM "files" WHERE "deleted_at" IS NULL AND "id" = $1
		ORDER BY "created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.Type, &data.URL, &data.UserUpload, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
	)

	return data, err
}

// FindUnassignedByID ...
func (model fileRepository) FindUnassignedByID(id, types, userUpload string) (data models.File, err error) {
	query := models.FileSelectString + ` WHERE f."deleted_at" IS NULL AND f."id" = $1 AND f."type" = $2
		AND f."user_upload" = $3 AND (f."is_used" IS NULL OR f."is_used" = false) ` + models.UnassignedQueryString + ` ORDER BY f."created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, id, types, userUpload).Scan(
		&data.ID, &data.Type, &data.URL, &data.UserUpload, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

// FindByType ...
func (model fileRepository) FindByType(id, types string) (data models.File, err error) {
	query := `SELECT "id", "type", "url", "user_upload", "created_at", "updated_at", "deleted_at"
		FROM "files" WHERE "deleted_at" IS NULL AND "id" = $1 AND "type" = $2
		ORDER BY "created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, id, types).Scan(
		&data.ID, &data.Type, &data.URL, &data.UserUpload, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

// Store ...
func (model fileRepository) Store(body *viewmodel.FileVM, changedAt time.Time) (res string, err error) {
	sql := `INSERT INTO "files" ("type", "url", "user_upload", "created_at", "updated_at")
		VALUES($1, $2, $3, $4, $4) RETURNING "id"`
	err = model.DB.QueryRow(sql, body.Type, body.URL, body.UserUpload, changedAt).Scan(&res)

	return res, err
}

// UpdateIsUsed ...
func (model fileRepository) UpdateIsUsed(id string, isUsed bool, changedAt time.Time) (res string, err error) {
	sql := `UPDATE "files" SET is_used = $1, deleted_at = $2 WHERE id = $3 RETURNING "id"`
	err = model.DB.QueryRow(sql, isUsed, changedAt, id).Scan(&res)

	return res, err
}

// Destroy ...
func (model fileRepository) Destroy(id string, changedAt time.Time) (res string, err error) {
	sql := `UPDATE "files" SET deleted_at = $1 WHERE id = $2 RETURNING "id"`
	err = model.DB.QueryRow(sql, changedAt, id).Scan(&res)

	return res, err
}
