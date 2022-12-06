package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerTargetYearUC ...
type CustomerTargetYearUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerTargetYearUC) BuildBody(res *models.CustomerTargetYear) {
}

// SelectAll ...
func (uc CustomerTargetYearUC) SelectAll(c context.Context, parameter models.CustomerTargetYearParameter) (res []models.CustomerTargetYear, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerTargetYearOrderBy, models.CustomerTargetYearOrderByrByString)

	repo := repository.NewCustomerTargetYearRepository(uc.DB)
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
func (uc CustomerTargetYearUC) FindAll(c context.Context, parameter models.CustomerTargetYearParameter) (res []models.CustomerTargetYear, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerTargetYearOrderBy, models.CustomerTargetYearOrderByrByString)

	var count int
	repo := repository.NewCustomerTargetYearRepository(uc.DB)
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
