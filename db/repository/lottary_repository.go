package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ILottaryRepository ...
type ILottaryRepository interface {
	SelectAll(c context.Context, parameter models.LottaryParameter) ([]models.Lottary, error)
	FindAll(ctx context.Context, parameter models.LottaryParameter) ([]models.Lottary, int, error)
	FindByID(c context.Context, parameter models.LottaryParameter) (models.Lottary, error)
	FindExsistingLottaryCupon(c context.Context, parameter models.LottaryParameter) (models.Lottary, error)
	FindByCustomerCode(c context.Context, customerCode string) (models.Lottary, error)
	Add(c context.Context, model viewmodel.LottaryVM) error
	Update(c context.Context, model viewmodel.LottaryVM) (string, error)
	Delete(c context.Context, id string) (string, error)
}

// LottaryRepository ...
type LottaryRepository struct {
	DB *sql.DB
}

// NewLottaryRepository ...
func NewLottaryRepository(DB *sql.DB) ILottaryRepository {
	return &LottaryRepository{DB: DB}
}

// Scan rows
func (repository LottaryRepository) scanRows(rows *sql.Rows) (res models.Lottary, err error) {
	err = rows.Scan(
		&res.ID,
		&res.SerailNo,
		&res.Status,
		&res.CustomerCode,
		&res.CustomerName,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Year,
		&res.Quartal,
		&res.Sequence,
		&res.BranchName,
		&res.RegionCode,
		&res.RegionName,
		&res.RegionGroup,
		&res.CustomerCpName,
		&res.CustomerLevel,
		&res.CustomerType,
	)

	return
}

// Scan row
func (repository LottaryRepository) scanRow(row *sql.Row) (res models.Lottary, err error) {
	err = row.Scan(
		&res.ID,
		&res.SerailNo,
		&res.Status,
		&res.CustomerCode,
		&res.CustomerName,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
		&res.Year,
		&res.Quartal,
		&res.Sequence,
		&res.BranchName,
		&res.RegionCode,
		&res.RegionName,
		&res.RegionGroup,
		&res.CustomerCpName,
		&res.CustomerLevel,
		&res.CustomerType,
	)

	return
}

// SelectAll ...
func (repository LottaryRepository) SelectAll(c context.Context, parameter models.LottaryParameter) (data []models.Lottary, err error) {
	var conditionString string

	if parameter.Quartal != "" {
		conditionString += ` AND def._quartal = ` + parameter.Quartal
	}
	if parameter.Year != "" {
		conditionString += ` AND def._year = ` + parameter.Year
	}
	if parameter.BranchID != "" {
		conditionString += ` AND c.branch_id in (` + parameter.BranchID + `)`
	}
	if parameter.RegionID != "" {
		conditionString += ` AND b.region_id in (` + parameter.RegionID + `)`
	}

	if parameter.Search != "" {
		conditionString += ` AND (LOWER(DEF.CUSTOMER_CODE) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(C.CUSTOMER_NAME) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(B.BRANCH_CODE) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(B._NAME) LIKE LOWER('%` + parameter.Search + `%'))`
	}

	statement := models.LottarySelectStatement + models.LottaryWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort

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
func (repository LottaryRepository) FindAll(ctx context.Context, parameter models.LottaryParameter) (data []models.Lottary, count int, err error) {
	var conditionString string

	if parameter.Quartal != "" {
		conditionString += ` AND def._quartal = ` + parameter.Quartal
	}
	if parameter.Year != "" {
		conditionString += ` AND def._year = ` + parameter.Year
	}
	if parameter.BranchID != "" {
		conditionString += ` AND c.branch_id in (` + parameter.BranchID + `)`
	}
	if parameter.RegionID != "" {
		conditionString += ` AND b.region_id in (` + parameter.RegionID + `)`
	}

	if parameter.Search != "" {
		conditionString += ` AND (LOWER(C.CUSTOMER_CODE) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(C.CUSTOMER_NAME) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(B.BRANCH_CODE) LIKE LOWER('%` + parameter.Search + `%') OR
		LOWER(B._NAME) LIKE LOWER('%` + parameter.Search + `%'))`
	}

	statement := models.LottarySelectStatement + models.LottaryWhereStatement +
		conditionString +
		` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $1 LIMIT $2`
	rows, err := repository.DB.QueryContext(ctx, statement, parameter.Offset, parameter.Limit)
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

	countQuery := `SELECT COUNT(*) 
	FROM lottary DEF
	LEFT JOIN CUSTOMER C ON C.ID = DEF.CUSTOMER_ID
	LEFT JOIN BRANCH B ON B.ID = C.BRANCH_ID
	left join region r on r.id = b.region_id
	left join customer_level cl on cl.id = c.customer_level_id
	left join customer_type ctp on ctp.id= c.customer_type_id  ` +
		models.LottaryWhereStatement +
		conditionString
	err = repository.DB.QueryRow(countQuery).Scan(&count)

	return
}

// FindByID ...
func (repository LottaryRepository) FindByID(c context.Context, parameter models.LottaryParameter) (data models.Lottary, err error) {
	statement := models.LottarySelectStatement + ` WHERE DEF.ID = ` + parameter.ID

	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// FindExsistingLottaryCupon ...
func (repository LottaryRepository) FindExsistingLottaryCupon(c context.Context, parameter models.LottaryParameter) (data models.Lottary, err error) {
	statement := models.LottarySelectStatement + ` WHERE DEF.serial_no = '` + parameter.SerialNo + `' and def._year = '` + parameter.Year + `' and def._quartal = '` + parameter.Quartal + `'` + ` and def.customer_id = ` + parameter.CustomerID
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// FindByCustomerCode ...
func (repository LottaryRepository) FindByCustomerCode(c context.Context, customerCode string) (data models.Lottary, err error) {
	var conditionString string

	conditionString += ` AND NOW() BETWEEN DEF.START_DATE AND DEF.END_DATE`

	statement := models.LottarySelectStatement + ` WHERE DEF.CUSTOMER_CODE = '` + customerCode + `'` + conditionString +
		` ORDER BY DEF.CREATED_AT DESC LIMIT 1`
	row := repository.DB.QueryRowContext(c, statement)

	data, err = repository.scanRow(row)

	return
}

// Add ...
func (repository LottaryRepository) Add(c context.Context, model viewmodel.LottaryVM) (err error) {

	statement := `INSERT INTO lottary (
		serial_no, 
		status,
			CUSTOMER_ID,
			_year,
			_quartal,
			CREATED_DATE,
			MODIFIED_DATE,
			_sequence
		)
	VALUES ($1, 1, $2, $3, $4, now(),now(), $5 )`
	// + statementInsert

	err = repository.DB.QueryRowContext(c, statement, model.SerialNo, model.CustomerID, model.Year, model.Quartal, model.Sequence).Err()

	return
}

// Update ...
func (repository LottaryRepository) Update(c context.Context, in viewmodel.LottaryVM) (res string, err error) {
	statement := `UPDATE POINT_MAX_CUSTOMER SET 
		START_DATE = $1, 
		END_DATE = $2, 
		CUSTOMER_CODE = $3, 
		MONTHLY_MAX_POINT = $4,
		UPDATED_AT = now()
	WHERE id = $5
	RETURNING id`

	err = repository.DB.QueryRowContext(c, statement,
		// in.StartDate,
		// in.EndDate,
		// in.CustomerCode,
		// in.MonthlyMaxPoint,
		in.ID).Scan(&res)

	return
}

// Delete ...
func (repository LottaryRepository) Delete(c context.Context, id string) (res string, err error) {
	statement := `UPDATE POINT_MAX_CUSTOMER SET 
	DELETED_AT = NOW()
	WHERE id = ` + id + `
	RETURNING id`
	err = repository.DB.QueryRowContext(c, statement).Scan(&res)

	return
}
