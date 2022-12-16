package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerTypeUC ...
type CustomerTypeUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerTypeUC) BuildBody(res *models.CustomerType) {
}

// SelectAll ...
func (uc CustomerTypeUC) SelectAll(c context.Context, parameter models.CustomerTypeParameter) (res []models.CustomerType, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerTypeOrderBy, models.CustomerTypeOrderByrByString)

	repo := repository.NewCustomerTypeRepository(uc.DB)
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
func (uc CustomerTypeUC) FindAll(c context.Context, parameter models.CustomerTypeParameter) (res []models.CustomerType, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerTypeOrderBy, models.CustomerTypeOrderByrByString)

	var count int
	repo := repository.NewCustomerTypeRepository(uc.DB)
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
func (uc CustomerTypeUC) FindByID(c context.Context, parameter models.CustomerTypeParameter) (res models.CustomerType, err error) {
	repo := repository.NewCustomerTypeRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
