package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerOrderLineUC ...
type CustomerOrderLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerOrderLineUC) BuildBody(res *models.CustomerOrderLine) {
}

// SelectAll ...
func (uc CustomerOrderLineUC) SelectAll(c context.Context, parameter models.CustomerOrderLineParameter) (res []models.CustomerOrderLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderLineOrderBy, models.CustomerOrderLineOrderByrByString)

	repo := repository.NewCustomerOrderLineRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}

// FindAll ...
func (uc CustomerOrderLineUC) FindAll(c context.Context, parameter models.CustomerOrderLineParameter) (res []models.CustomerOrderLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderLineOrderBy, models.CustomerOrderLineOrderByrByString)

	var count int
	repo := repository.NewCustomerOrderLineRepository(uc.DB)
	res, count, err = repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, p, err
}

// FindByID ...
func (uc CustomerOrderLineUC) FindByID(c context.Context, parameter models.CustomerOrderLineParameter) (res models.CustomerOrderLine, err error) {
	repo := repository.NewCustomerOrderLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
