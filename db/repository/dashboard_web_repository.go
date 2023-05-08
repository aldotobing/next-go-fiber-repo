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
	GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)
	GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.DashboardWebBranchDetail, error)
	GetOmzetValue(ctx context.Context, parameter models.DashboardWebBranchParameter) ([]models.OmzetValueModel, error)
	GetOmzetValueByGroupID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) ([]models.OmzetValueModel, error)
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
		&res.TotalOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanRegionDetailRows(rows *sql.Rows) (res models.DashboardWebRegionDetail, err error) {
	err = rows.Scan(
		&res.BranchID, &res.BranchCode, &res.BranchName, &res.RegionID, &res.RegionName,
		&res.RegionGroupID, &res.RegionGroupName,
		&res.TotalRegisteredUser, &res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalVisitUser, &res.CustomerCountRepeatOrder,
		&res.TotalActiveOutlet,
		&res.TotalOutlet,
	)
	if err != nil {

		return res, err
	}

	return res, nil
}

// Scan rows
func (repository DashboardWebRepository) scanBranchCustomerDetailRows(rows *sql.Rows) (res models.DashboardWebBranchDetail, err error) {
	err = rows.Scan(
		&res.CustomerID, &res.CustomerName, &res.CustomerCode,
		&res.CustomerBranchName, &res.CustomerBranchCode,
		&res.CustomerRegionName, &res.CustomerRegionGroupName,
		&res.CustomerTypeName,
		&res.TotalRepeatUser, &res.TotalOrderUser,
		&res.TotalInvoice, &res.TotalCheckin, &res.TotalAktifOutlet, &res.CustomerClassName, &res.CustomerCityName,
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

func (repository DashboardWebRepository) GetAllBranchDetailCustomerData(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebBranchDetailSelectStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.BarnchID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository DashboardWebRepository) GetAllDetailCustomerDataWithUserID(ctx context.Context, parameter models.DashboardWebBranchParameter) (data []models.DashboardWebBranchDetail, err error) {

	query := models.DashboardWebBranchDetailSelectWithUserIDStatement
	rows, err := repository.DB.Query(query, str.NullOrEmtyString(&parameter.UserID), str.NullOrEmtyString(&parameter.StartDate), str.NullOrEmtyString(&parameter.EndDate))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanBranchCustomerDetailRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, err
}

func (repo DashboardWebRepository) GetOmzetValue(ctx context.Context, parameter models.DashboardWebBranchParameter) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}

	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}

	query := `select r.group_id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null` + whereStatement + `
			group by r.group_id
			order by r.group_id asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.RegionID, &temp.TotalGrossAmount, &temp.TotalNettAmount, &temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}

func (repo DashboardWebRepository) GetOmzetValueByGroupID(ctx context.Context, parameter models.DashboardWebBranchParameter, groupID string) (res []models.OmzetValueModel, err error) {
	var whereStatement string
	if parameter.StartDate != "" && parameter.EndDate != "" {
		whereStatement += ` AND sih.transaction_date BETWEEN '` + parameter.StartDate + `' AND '` + parameter.EndDate + `'`
	} else {
		whereStatement += ` AND sih.transaction_date BETWEEN date_trunc('MONTH',now())::DATE AND now()`
	}
	if parameter.ItemID != "" {
		whereStatement += ` AND sil.item_id = ` + parameter.ItemID
	}
	if parameter.ItemCategoryID != "" {
		whereStatement += ` AND sil.category_id = ` + parameter.ItemCategoryID
	}
	if groupID != "" && groupID != "0" {
		whereStatement += ` AND r.group_id = '` + groupID + `'`
	}

	query := `select r.id,
			coalesce(sum(sil.gross_amount),0) as total_gross_amount, 
			coalesce(sum(sil.net_amount),0) as total_nett_amount, 
			coalesce(sum(sil.qty),0) as total_volume
		from sales_invoice_header sih 
			left join sales_invoice_line sil on sil.header_id = sih.id 
			left join branch b on b.id = sih.branch_id  
			left join region r on r.id = b.region_id
		WHERE sih.transaction_date is not null` + whereStatement + `
			group by r.id
			order by r.id asc`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.OmzetValueModel
		err = rows.Scan(&temp.ID, &temp.TotalGrossAmount, &temp.TotalNettAmount, &temp.TotalQuantity)
		if err != nil {
			return
		}

		res = append(res, temp)
	}

	return
}
