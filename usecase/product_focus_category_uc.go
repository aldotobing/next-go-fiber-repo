package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ProductFocusCategoryUC ...
type ProductFocusCategoryUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ProductFocusCategoryUC) BuildBody(res *models.ProductFocusCategory) {
}

// SelectAll ...
func (uc ProductFocusCategoryUC) SelectAll(c context.Context, parameter models.ProductFocusCategoryParameter) (res []models.ProductFocusCategory, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ProductFocusCategoryOrderBy, models.ProductFocusCategoryOrderByrByString)

	repo := repository.NewProductFocusCategoryRepository(uc.DB)
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
func (uc ProductFocusCategoryUC) FindAll(c context.Context, parameter models.ProductFocusCategoryParameter) (res []models.ProductFocusCategory, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ProductFocusCategoryOrderBy, models.ProductFocusCategoryOrderByrByString)

	var count int
	repo := repository.NewProductFocusCategoryRepository(uc.DB)
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
func (uc ProductFocusCategoryUC) FindByID(c context.Context, parameter models.ProductFocusCategoryParameter) (res models.ProductFocusCategory, err error) {
	repo := repository.NewProductFocusCategoryRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByBrancID ...
func (uc ProductFocusCategoryUC) FindByBranchID(c context.Context, parameter models.ProductFocusCategoryParameter) (res models.ProductFocusCategory, err error) {
	repo := repository.NewProductFocusCategoryRepository(uc.DB)
	res, err = repo.FindByBranchID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByCategoryID ...
func (uc ProductFocusCategoryUC) FindByCategoryID(c context.Context, parameter models.ProductFocusCategoryParameter) (res []models.ProductFocusCategory, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ProductFocusCategoryOrderBy, models.ProductFocusCategoryOrderByrByString)

	var count int
	repo := repository.NewProductFocusCategoryRepository(uc.DB)
	res, count, err = repo.FindByCategoryID(c, parameter)
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
