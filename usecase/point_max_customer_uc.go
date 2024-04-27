package usecase

import (
	"context"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PointMaxCustomerUC ...
type PointMaxCustomerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PointMaxCustomerUC) BuildBody(data *models.PointMaxCustomer, res *viewmodel.PointMaxCustomerVM) {
	res.ID = data.ID
	startDate, _ := time.Parse("2006-01-02T15:04:05.999999999Z", data.StartDate)
	res.StartDate = startDate.Format("2006-01-02")
	endDate, _ := time.Parse("2006-01-02T15:04:05.999999999Z", data.EndDate)
	res.EndDate = endDate.Format("2006-01-02")
	res.CustomerCode = data.CustomerCode
	res.CustomerName = data.CustomerName.String
	res.MonthlyMaxPoint = data.MonthlyMaxPoint
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	res.BranchCode = data.BranchCode.String
	res.BranchName = data.BranchName.String
}

// FindAll ...
func (uc PointMaxCustomerUC) FindAll(c context.Context, parameter models.PointMaxCustomerParameter) (out []viewmodel.PointMaxCustomerVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointMaxCustomerOrderBy, models.PointMaxCustomerOrderByrByString)

	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	for _, datum := range data {
		var temp viewmodel.PointMaxCustomerVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointMaxCustomerVM, 0)
	}

	return
}

// SelectAll ...
func (uc PointMaxCustomerUC) SelectAll(c context.Context, parameter models.PointMaxCustomerParameter) (out []viewmodel.PointMaxCustomerVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointMaxCustomerOrderBy, models.PointMaxCustomerOrderByrByString)

	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.PointMaxCustomerVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointMaxCustomerVM, 0)
	}

	return
}

// FindByID ...
func (uc PointMaxCustomerUC) FindByID(c context.Context, parameter models.PointMaxCustomerParameter) (out viewmodel.PointMaxCustomerVM, err error) {
	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// FindByCustomerCode ...
func (uc PointMaxCustomerUC) FindByCustomerCode(c context.Context, customerCode string) (out viewmodel.PointMaxCustomerVM, err error) {
	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	data, err := repo.FindByCustomerCode(c, customerCode)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// FindByCustomerCodeWithDateInvoice ...
func (uc PointMaxCustomerUC) FindByCustomerCodeWithDateInvoice(c context.Context, customerCode, date string) (out viewmodel.PointMaxCustomerVM, err error) {
	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	data, err := repo.FindByCustomerCodeWithDateInvoice(c, customerCode, date)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc PointMaxCustomerUC) Add(c context.Context, in requests.PointMaxCustomerRequestHeader) (out []viewmodel.PointMaxCustomerVM, err error) {

	for _, datum := range in.Detail {
		out = append(out, viewmodel.PointMaxCustomerVM{
			StartDate:       datum.StartDate,
			EndDate:         datum.EndDate,
			CustomerCode:    datum.CustomerCode,
			MonthlyMaxPoint: datum.MonthlyMaxPoint,
		})
	}

	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Update ...
func (uc PointMaxCustomerUC) Update(c context.Context, id string, in requests.PointMaxCustomerRequest) (out viewmodel.PointMaxCustomerVM, err error) {
	out = viewmodel.PointMaxCustomerVM{
		ID:              id,
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		CustomerCode:    in.CustomerCode,
		MonthlyMaxPoint: in.MonthlyMaxPoint,
	}

	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc PointMaxCustomerUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewPointMaxCustomerRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
