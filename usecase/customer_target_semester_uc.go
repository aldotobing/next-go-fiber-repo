package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerTargetSemesterUC ...
type CustomerTargetSemesterUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerTargetSemesterUC) BuildBody(res *models.CustomerTargetSemester) {
}

// SelectAll ...
func (uc CustomerTargetSemesterUC) SelectAll(c context.Context, parameter models.CustomerTargetSemesterParameter) (res []models.CustomerTargetSemester, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerTargetSemesterOrderBy, models.CustomerTargetSemesterOrderByrByString)

	repo := repository.NewCustomerTargetSemesterRepository(uc.DB)
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
func (uc CustomerTargetSemesterUC) FindAll(c context.Context, parameter models.CustomerTargetSemesterParameter) (res []models.CustomerTargetSemester, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerTargetSemesterOrderBy, models.CustomerTargetSemesterOrderByrByString)

	var count int
	repo := repository.NewCustomerTargetSemesterRepository(uc.DB)
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
