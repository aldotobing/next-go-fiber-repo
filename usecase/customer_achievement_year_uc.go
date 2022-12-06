package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerAchievementYearUC ...
type CustomerAchievementYearUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerAchievementYearUC) BuildBody(res *models.CustomerAchievementYear) {
}

// SelectAll ...
func (uc CustomerAchievementYearUC) SelectAll(c context.Context, parameter models.CustomerAchievementYearParameter) (res []models.CustomerAchievementYear, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerAchievementYearOrderBy, models.CustomerAchievementYearOrderByString)

	repo := repository.NewCustomerAchievementYearRepository(uc.DB)
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
func (uc CustomerAchievementYearUC) FindAll(c context.Context, parameter models.CustomerAchievementYearParameter) (res []models.CustomerAchievementYear, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerAchievementYearOrderBy, models.CustomerAchievementYearOrderByString)

	var count int
	repo := repository.NewCustomerAchievementYearRepository(uc.DB)
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
