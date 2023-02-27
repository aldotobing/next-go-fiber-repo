package repository

import (
	"context"
	"database/sql"
	"strings"

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
		&res.BranchID, &res.BranchName, &res.RegionID, &res.RegionName,
		&res.RegionGroupID, &res.RegionGroupName,
		&res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalActiveUser,
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
	conditionString := ` WHERE def.created_date IS not NULL and def.user_id is not null and def.user_id in(select us.id from _user us join customer cs on cs.user_id = us.id where us.fcm_token is not null and length(trim(us.fcm_token))>0 ) `

	if parameter.BarnchID != "" {
		conditionString += ` AND def.branch_id = '` + parameter.BarnchID + `'`
	}

	query := models.DashboardWebBranchDetailSelectStatement + ` ` + conditionString + `
		AND (LOWER(def."customer_name") LIKE $1  ) ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`
	rows, err := repository.DB.Query(query, "%"+strings.ToLower(parameter.Search)+"%", parameter.Offset, parameter.Limit)
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

	query = `select 
			count(*)
			from customer def 
			left join branch b on b.id = def.branch_id
			left join region r on r.id = b.region_id ` + ` ` +
		conditionString + ` AND (LOWER(def."customer_name") LIKE $1)`
	err = repository.DB.QueryRow(query, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)
	return data, count, err
}
