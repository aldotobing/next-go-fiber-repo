package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemSearchUC ...
type ItemSearchUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemSearchUC) BuildBody(res *models.ItemSearch) {
}

// SelectAll ...
func (uc ItemSearchUC) SelectAll(c context.Context, parameter models.ItemSearchParameter) (res []models.ItemSearch, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemSearchOrderBy, models.ItemSearchOrderByrByString)

	repo := repository.NewItemSearchRepository(uc.DB)
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
func (uc ItemSearchUC) FindAll(c context.Context, parameter models.ItemSearchParameter) (res []models.ItemSearch, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemSearchOrderBy, models.ItemSearchOrderByrByString)

	var count int
	repo := repository.NewItemSearchRepository(uc.DB)
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
func (uc ItemSearchUC) FindByID(c context.Context, parameter models.ItemSearchParameter) (res models.ItemSearch, err error) {
	repo := repository.NewItemSearchRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
