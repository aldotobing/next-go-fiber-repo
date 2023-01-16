package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebUserRoleGroupUC ...
type WebUserRoleGroupUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebUserRoleGroupUC) BuildBody(res *models.WebUserRoleGroup) {
}

// SelectAll ...
func (uc WebUserRoleGroupUC) SelectAll(c context.Context, parameter models.WebUserRoleGroupParameter) (res []models.WebUserRoleGroup, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebUserRoleGroupOrderBy, models.WebUserRoleGroupOrderByrByString)

	repo := repository.NewWebUserRoleGroupRepository(uc.DB)
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
func (uc WebUserRoleGroupUC) FindAll(c context.Context, parameter models.WebUserRoleGroupParameter) (res []models.WebUserRoleGroup, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebUserRoleGroupOrderBy, models.WebUserRoleGroupOrderByrByString)

	var count int
	repo := repository.NewWebUserRoleGroupRepository(uc.DB)
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
func (uc WebUserRoleGroupUC) FindByID(c context.Context, parameter models.WebUserRoleGroupParameter) (res models.WebUserRoleGroup, err error) {
	repo := repository.NewWebUserRoleGroupRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
