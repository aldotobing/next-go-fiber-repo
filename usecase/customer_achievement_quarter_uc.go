package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerAchievementQuarterUC ...
type CustomerAchievementQuarterUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerAchievementQuarterUC) BuildBody(res *models.CustomerAchievementQuarter) {
}

// SelectAll ...
func (uc CustomerAchievementQuarterUC) SelectAll(c context.Context, parameter models.CustomerAchievementQuarterParameter) (res []models.CustomerAchievementQuarter, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerAchievementQuarterOrderBy, models.CustomerAchievementQuarterOrderByrByString)

	repo := repository.NewCustomerAchievementQuarterRepository(uc.DB)
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
func (uc CustomerAchievementQuarterUC) FindAll(c context.Context, parameter models.CustomerAchievementQuarterParameter) (res []models.CustomerAchievementQuarter, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerAchievementQuarterOrderBy, models.CustomerAchievementQuarterOrderByrByString)

	var count int
	repo := repository.NewCustomerAchievementQuarterRepository(uc.DB)
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
