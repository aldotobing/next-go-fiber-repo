package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// DashboardWebUC ...
type DashboardWebUC struct {
	*ContractUC
}

// BuildBody ...
func (uc DashboardWebUC) BuildBody(res *models.DashboardWeb) {
}

func (uc DashboardWebUC) BuildRegionDetailBody(res *models.DashboardWebRegionDetail) {
}

func (uc DashboardWebUC) BuildBranchDetailCustomerBody(res *models.DashboardWebBranchDetail) {
}

// FindByID ...
func (uc DashboardWebUC) GetData(c context.Context, parameter models.DashboardWebParameter) (res []models.DashboardWeb, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}

// FindByID ...
func (uc DashboardWebUC) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (res []models.DashboardWebRegionDetail, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetRegionDetailData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildRegionDetailBody(&res[i])
	}

	return res, err
}

func (uc DashboardWebUC) GetBranchDetailCustomerData(c context.Context, parameter models.DashboardWebBranchParameter) (res []models.DashboardWebBranchDetail, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DashboardWebBranchDetailOrderBy, models.DashboardWebBranchDetailOrderByrByString)

	var count int
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, count, err = repo.GetBranchDetailCustomerData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range res {
		uc.BuildBranchDetailCustomerBody(&res[i])
	}

	return res, p, err
}
