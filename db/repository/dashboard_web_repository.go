package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

// IDashboardWebRepository ...
type IDashboardWebRepository interface {
	GetData(c context.Context, parameter models.DashboardWebParameter) ([]models.DashboardWeb, error)
	GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) ([]models.DashboardWebRegionDetail, error)
	GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, int, error)
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
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalVisitUser,
		&res.CustomerCountRepeatOrder, &res.TotalActiveOutlet,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRows(rows *sql.Rows) (res models.DashboardWeb, err error) {
	err = rows.Scan(
		&res.RegionGroupID, &res.RegionGroupName, &res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser, &res.TotalInvoice, &res.TotalVisitUser,
		&res.CustomerCountRepeatOrder, &res.TotalActiveOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRegionDetailRows(rows *sql.Rows) (res models.DashboardWebRegionDetail, err error) {
	err = rows.Scan(
		&res.BranchID, &res.BranchName, &res.RegionID, &res.RegionName,
		&res.RegionGroupID, &res.RegionGroupName,
		&res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalVisitUser, &res.CustomerCountRepeatOrder,
		&res.TotalActiveOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanBranchCustomerDetailRows(rows *sql.Rows) (res models.DashboardWebBranchDetail, err error) {
	err = rows.Scan(
		&res.CustomerID, &res.CustomerName, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// FindByID ...
func (repository DashboardWebRepository) GetData(c context.Context, parameter models.DashboardWebParameter) (data []models.DashboardWeb, err error) {
	statement := models.DashboardWebSelectStatement
	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))

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

	rows, err := repository.DB.QueryContext(c, statement, str.NullOrEmtyString(&parameter.GroupID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))

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

func (repository DashboardWebRepository) GetBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, count int, err error) {

	query := models.DashboardWebBranchDetailSelectStatement + ` OFFSET $4 LIMIT $5`
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate),
		parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = ` select count(*) from os_fetch_dashborad_branchcustomerdata($1::integer,$2,$3,null,null,null) `
	err = repository.DB.QueryRow(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate)).Scan(&count)
	return data, count, err
}
