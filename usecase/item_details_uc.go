package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemDetailsUC ...
type ItemDetailsUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemDetailsUC) BuildBody(res *models.ItemDetails) {
}

// SelectAll ...
func (uc ItemDetailsUC) SelectAll(c context.Context, parameter models.ItemDetailsParameter) (res []models.ItemDetails, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemDetailsOrderBy, models.ItemDetailsOrderByrByString)

	repo := repository.NewItemDetailsRepository(uc.DB)
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
func (uc ItemDetailsUC) FindAll(c context.Context, parameter models.ItemDetailsParameter) (res []models.ItemDetails, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemDetailsOrderBy, models.ItemDetailsOrderByrByString)

	var count int
	repo := repository.NewItemDetailsRepository(uc.DB)
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
func (uc ItemDetailsUC) FindByID(c context.Context, parameter models.ItemDetailsParameter) (res models.ItemDetails, err error) {
	repo := repository.NewItemDetailsRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
