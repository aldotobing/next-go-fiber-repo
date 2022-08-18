package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// RoleUC ...
type RoleUC struct {
	*ContractUC
}

// BuildBody ...
func (uc RoleUC) BuildBody(res *models.Role) {
}

// SelectAll ...
func (uc RoleUC) SelectAll(c context.Context, parameter models.RoleParameter) (res []models.Role, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.RoleOrderBy, models.RoleOrderByrByString)

	repo := repository.NewRoleRepository(uc.DB)
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

// FindByID ...
func (uc RoleUC) FindByID(c context.Context, parameter models.RoleParameter) (res models.Role, err error) {
	repo := repository.NewRoleRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
