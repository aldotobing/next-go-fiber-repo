package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
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
	res.InvoiceDocumentNo = data.InvoiceDocumentNo.String
	res.Point = data.Point
	res.CustomerID = data.CustomerID
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
	res.ExpiredAt = data.ExpiredAt.String

	res.DetailCustomer = viewmodel.CustomerVM{
		CustomerName:       data.Customer.CustomerName.String,
		Code:               data.Customer.Code.String,
		CustomerBranchCode: data.Customer.CustomerBranchCode.String,
		CustomerBranchName: data.Customer.CustomerBranchName.String,
		CustomerRegionName: data.Customer.CustomerRegionName.String,
	}
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

	if parameter.Renewal != "1" {
		// Try getting data from cache
		cachedData, err := uc.RedisClient.Get(cacheKey)
		if err == nil && string(cachedData) != "" {
			err := json.Unmarshal(cachedData, &out)

			return out, err
		}
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
	point, _ = strconv.ParseFloat(data.Promo, 64)
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

// GetBalanceAll ...
func (uc PointUC) GetBalanceAll(c context.Context, parameter models.PointParameter) (out viewmodel.PointBalanceVM, err error) {

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
	point, _ = strconv.ParseFloat(data.Promo, 64)
	totalPoint += point

	out = viewmodel.PointBalanceVM{
		Balance: strconv.FormatFloat(totalPoint, 'f', 2, 64),
	}

	return
}

// GetPointThisMonth ...
func (uc PointUC) GetPointThisMonth(c context.Context, customerID string) (out viewmodel.PointBalanceVM, err error) {

	repo := repository.NewPointRepository(uc.DB)
	data, err := repo.GetBalance(c, models.PointParameter{
		CustomerID: customerID,
		Month:      strconv.Itoa(int(time.Now().Month())),
		Year:       strconv.Itoa(time.Now().Year()),
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var totalPoint float64
	point, _ := strconv.ParseFloat(data.Cashback, 64)
	totalPoint += point

	out = viewmodel.PointBalanceVM{
		Balance: strconv.FormatFloat(totalPoint, 'f', 2, 64),
	}

	return
}

// Add ...
func (uc PointUC) Add(c context.Context, in requests.PointRequest) (out viewmodel.PointVM, err error) {
	now := time.Now()
	expiredAt := helper.GetExpiredPoint(now)

	var customerCodes string
	for _, datum := range in.CustomerCodes {
		if customerCodes != "" {
			customerCodes += ", '" + datum.CustomerCode + "'"
		} else {
			customerCodes += "'" + datum.CustomerCode + "'"
		}
	}

	customerData, err := WebCustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.WebCustomerParameter{
		Code: customerCodes,
		By:   "c.id",
		Sort: "asc",
	})

	if len(customerData) < 1 {
		err = errors.New(customerCodes + "not found / show_in_app = 0 ")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var customerIDs []string
	for _, datum := range customerData {
		customerIDs = append(customerIDs, datum.ID)
	}

	out = viewmodel.PointVM{
		PointType:         in.PointType,
		InvoiceDocumentNo: in.InvoiceDocumentNo,
		Point:             in.Point,
		CustomerID:        in.CustomerID,
		ExpiredAt:         expiredAt,
		CustomerIDs:       customerIDs,
	}

	repo := repository.NewPointRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddInject ...
func (uc PointUC) AddInject(c context.Context, in requests.PointRequest) (out []viewmodel.PointVM, err error) {
	now := time.Now()
	expiredAt := helper.GetExpiredPoint(now)

	var customerCodes string
	for _, datum := range in.CustomerCodes {
		if customerCodes != "" {
			customerCodes += ", '" + datum.CustomerCode + "'"
		} else {
			customerCodes += "'" + datum.CustomerCode + "'"
		}
	}

	customerData, err := WebCustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.WebCustomerParameter{
		Code: customerCodes,
		By:   "c.id",
		Sort: "asc",
	})

	if len(customerData) < 1 {
		err = errors.New(customerCodes + "not found / show_in_app = 0 ")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range customerData {
		for _, y := range in.CustomerCodes {
			if y.CustomerCode == datum.Code {
				out = append(out, viewmodel.PointVM{
					PointType:         in.PointType,
					InvoiceDocumentNo: "INJECT-" + in.UserID + "-" + y.CustomerCode + "-" + now.Format("2006-01-02 15:04:05"),
					Point:             y.Point,
					CustomerID:        datum.ID,
					ExpiredAt:         expiredAt,
				})
			}
		}
	}

	repo := repository.NewPointRepository(uc.DB)
	_, err = repo.AddInject(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddWithdraw ...
func (uc PointUC) AddWithdraw(c context.Context, in requests.PointRequest) (out viewmodel.PointVM, err error) {
	out = viewmodel.PointVM{
		PointType:  "3",
		Point:      in.Point,
		CustomerID: in.CustomerID,
	}

	repo := repository.NewPointRepository(uc.DB)
	out.ID, err = repo.AddWithdraw(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.GetBalance(c, models.PointParameter{
		CustomerID: in.CustomerID,
		Renewal:    "1",
	})

	return
}

// Update ...
func (uc PointUC) Update(c context.Context, id string, in requests.PointRequest) (out viewmodel.PointVM, err error) {
	out = viewmodel.PointVM{
		ID:                id,
		PointType:         in.PointType,
		InvoiceDocumentNo: in.InvoiceDocumentNo,
		Point:             in.Point,
		CustomerID:        in.CustomerID,
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

// Report ...
func (uc PointUC) Report(c context.Context, parameter models.PointParameter) (out []viewmodel.PointReportVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VoucherOrderBy, models.VoucherOrderByrByString)

	repo := repository.NewPointRepository(uc.DB)
	data, err := repo.Report(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		temp := viewmodel.PointReportVM{
			BranchCode:        *datum.Branch.Code,
			BranchName:        *datum.Branch.Name,
			RegionName:        *datum.Region.Name,
			RegionGroupName:   *datum.Region.GroupName,
			PartnerCode:       *datum.Partner.Code,
			PartnerName:       *datum.Partner.PartnerName,
			InvoiceDocumentNo: datum.InvoiceDocumentNo.String,
			NetAmount:         *datum.SalesInvoice.NetAmount,
			Point:             datum.Point,
			TrasactionDate:    *datum.SalesInvoice.TrasactionDate,
		}

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointReportVM, 0)
	}

	return
}
