package repository

import (
	"context"
	"database/sql"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
)

// IDashboardWebRepository ...
type IDashboardWebRepository interface {
	GetData(c context.Context, parameter models.DashboardWebParameter) ([]models.DashboardWeb, error)
	GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) ([]models.DashboardWebRegionDetail, error)
}

// DashboardWebRepository ...
type DashboardWebRepository struct {
	DB *sql.DB
}

// NewDashboardWebRepository ...
func NewDashboardWebRepository(DB *sql.DB) IDashboardWebRepository {
	return &DashboardWebRepository{DB: DB}
}

// Scan row
func (repository DashboardWebRepository) scanRow(row *sql.Row) (res models.DashboardWeb, err error) {
	err = row.Scan(
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalActiveUser,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRows(rows *sql.Rows) (res models.DashboardWeb, err error) {
	err = rows.Scan(
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalActiveUser,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRegionDetailRows(rows *sql.Rows) (res models.DashboardWebRegionDetail, err error) {
	err = rows.Scan(
		&res.BranchID, &res.BranchName, &res.RegionID, &res.RegionGroupName,
		&res.RegionGroupID, &res.RegionGroupName,
		&res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalActiveUser,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) (data []models.DashboardWeb, err error) {
	statement := models.DashboardWebSelectStatement
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

func (repository DashboardWebRepository) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (data []models.DashboardWebRegionDetail, err error) {
	statement := models.DashboardWebRegionDetailSelectStatement
	statement += ` where def.region_id is not null `
	if &parameter.GroupID != nil && strings.Trim(parameter.GroupID, " ") != "" && parameter.GroupID != "0" {
		statement += ` and r.group_id = ` + parameter.GroupID
	}
	statement += ` order by r.sequence,def.region_id `

	rows, err := repository.DB.QueryContext(c, statement)

	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {

		temp, err := repository.scanRegionDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}
