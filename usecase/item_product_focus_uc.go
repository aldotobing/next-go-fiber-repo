package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemProductFocusUC ...
type ItemProductFocusUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemProductFocusUC) BuildBody(res *models.ItemProductFocus) {
}

// SelectAll ...
func (uc ItemProductFocusUC) SelectAll(c context.Context, parameter models.ItemProductFocusParameter) (res []models.ItemProductFocus, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemProductFocusOrderBy, models.ItemProductFocusOrderByrByString)

	repo := repository.NewItemProductFocusRepository(uc.DB)
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
func (uc ItemProductFocusUC) FindAll(c context.Context, parameter models.ItemProductFocusParameter) (res []models.ItemProductFocus, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemProductFocusOrderBy, models.ItemProductFocusOrderByrByString)

	var count int
	repo := repository.NewItemProductFocusRepository(uc.DB)
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
func (uc ItemProductFocusUC) FindByID(c context.Context, parameter models.ItemProductFocusParameter) (res models.ItemProductFocus, err error) {
	repo := repository.NewItemProductFocusRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
