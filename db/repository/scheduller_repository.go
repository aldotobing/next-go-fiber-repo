package repository

import (
	"context"
	"database/sql"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
)

// ISchedullerRepository ...
type ISchedullerRepository interface {
	ProcessExpiredPackage(c context.Context, parameter models.SchedullerExpiredPackageParameter) (models.SchedullerExpiredPackage, error)
	ProcessExpiredPackageNoContex()
}

// SchedullerRepository ...
type SchedullerRepository struct {
	DB *sql.DB
}

// NewSchedullerRepository ...
func NewSchedullerRepository(DB *sql.DB) ISchedullerRepository {
	return &SchedullerRepository{DB: DB}
}

// Scan row
func (repository SchedullerRepository) scanRow(row *sql.Row) (res models.SchedullerExpiredPackage, err error) {
	err = row.Scan(
		&res.TotalCount,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository SchedullerRepository) ProcessExpiredPackage(c context.Context, parameter models.SchedullerExpiredPackageParameter) (data models.SchedullerExpiredPackage, err error) {
	statement := models.SchedullerExpiredPackageSelectStatement
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository SchedullerRepository) ProcessExpiredPackageNoContex() {
	fmt.Println("kesini")
	statement := models.SchedullerExpiredPackageSelectStatement
	// ctx := context.Background()
	ress := 1
	row := repository.DB.QueryRow(statement).Scan(&ress)

	if row != nil {

	}

	// data, err = repository.scanRow(row)

	// result, _ := repository.scanRow(row)
	// if result.TotalCount != nil {

	// }
	// if err != nil {
	// 	return data, err
	// }

}
