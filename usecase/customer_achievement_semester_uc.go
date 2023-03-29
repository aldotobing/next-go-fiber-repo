package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerAchievementSemesterUC ...
type CustomerAchievementSemesterUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerAchievementSemesterUC) BuildBody(res *models.CustomerAchievementSemester) {
}

// SelectAll ...
func (uc CustomerAchievementSemesterUC) SelectAll(c context.Context, parameter models.CustomerAchievementSemesterParameter) (res []models.CustomerAchievementSemester, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerAchievementSemesterOrderBy, models.CustomerAchievementSemesterOrderByrByString)

	repo := repository.NewCustomerAchievementSemesterRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	return res, err
}

// FindAll ...
func (uc CustomerAchievementSemesterUC) FindAll(c context.Context, parameter models.CustomerAchievementSemesterParameter) (res []models.CustomerAchievementSemester, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerAchievementSemesterOrderBy, models.CustomerAchievementSemesterOrderByrByString)

	var count int
	repo := repository.NewCustomerAchievementSemesterRepository(uc.DB)
	res, count, err = repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	return res, p, err
}
