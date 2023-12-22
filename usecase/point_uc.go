package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PointUC ...
type PointUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PointUC) BuildBody(data *models.Point, res *viewmodel.PointVM) {
	res.ID = data.ID
	res.PointType = data.PointType
	res.PointTypeName = data.PointTypeName
	res.InvoiceID = data.InvoiceID.String
	res.Point = data.Point
	res.CustomerID = data.CustomerID
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
}

// FindAll ...
func (uc PointUC) FindAll(c context.Context, parameter models.PointParameter) (out []viewmodel.PointVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VoucherOrderBy, models.VoucherOrderByrByString)

	repo := repository.NewPointRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.PointVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointVM, 0)
	}

	return
}

// SelectAll ...
func (uc PointUC) SelectAll(c context.Context, parameter models.PointParameter) (out []viewmodel.PointVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VoucherOrderBy, models.VoucherOrderByrByString)

	repo := repository.NewPointRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.PointVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointVM, 0)
	}

	return
}

// FindByID ...
func (uc PointUC) FindByID(c context.Context, parameter models.PointParameter) (out viewmodel.PointVM, err error) {
	repo := repository.NewPointRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// GetBalance ...
func (uc PointUC) GetBalance(c context.Context, parameter models.PointParameter) (out viewmodel.PointBalanceVM, err error) {
	cacheKey := fmt.Sprintf("balance_customer_id:%s",
		parameter.CustomerID)

	// Try getting data from cache
	cachedData, err := uc.RedisClient.Get(cacheKey)
	if err == nil && string(cachedData) != "" {
		err := json.Unmarshal(cachedData, &out)

		return out, err
	}

	repo := repository.NewPointRepository(uc.DB)
	data, err := repo.GetBalance(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var totalPoint float64
	point, _ := strconv.ParseFloat(data.Cashback, 64)
	totalPoint += point
	point, _ = strconv.ParseFloat(data.Loyalty, 64)
	totalPoint += point
	point, _ = strconv.ParseFloat(data.Withdraw, 64)
	totalPoint -= point

	out = viewmodel.PointBalanceVM{
		Balance: strconv.FormatFloat(totalPoint, 'f', 2, 64),
	}

	dataOut, _ := json.Marshal(out)
	uc.RedisClient.Set(cacheKey, dataOut, time.Hour*1) // Cache for 30 minutes only if there's data

	return
}

// Add ...
func (uc PointUC) Add(c context.Context, in requests.PointRequest) (out viewmodel.PointVM, err error) {
	out = viewmodel.PointVM{
		PointType:  in.PointType,
		InvoiceID:  in.InvoiceID,
		Point:      in.Point,
		CustomerID: in.CustomerID,
	}

	repo := repository.NewPointRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Update ...
func (uc PointUC) Update(c context.Context, id string, in requests.PointRequest) (out viewmodel.PointVM, err error) {
	out = viewmodel.PointVM{
		ID:         id,
		PointType:  in.PointType,
		InvoiceID:  in.InvoiceID,
		Point:      in.Point,
		CustomerID: in.CustomerID,
	}

	repo := repository.NewPointRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc PointUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewPointRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
