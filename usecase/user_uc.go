package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// UserUC ...
type UserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserUC) BuildBody(res *models.User) {
}

// SelectAll ...
func (uc UserUC) SelectAll(c context.Context, parameter models.UserParameter) (res []models.User, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.UserOrderBy, models.UserOrderByrByString)

	repo := repository.NewUserRepository(uc.DB)
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
func (uc UserUC) FindAll(c context.Context, parameter models.UserParameter) (res []models.User, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.UserOrderBy, models.UserOrderByrByString)

	var count int
	repo := repository.NewUserRepository(uc.DB)
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
func (uc UserUC) FindByID(c context.Context, parameter models.UserParameter) (res models.User, err error) {
	repo := repository.NewUserRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByEmail ...
func (uc UserUC) FindByEmail(c context.Context, parameter models.UserParameter) (res models.User, err error) {
	repo := repository.NewUserRepository(uc.DB)
	res, err = repo.FindByEmail(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByPhone ...
func (uc UserUC) FindByPhone(c context.Context, parameter models.UserParameter) (res models.User, err error) {
	repo := repository.NewUserRepository(uc.DB)
	res, err = repo.FindByPhone(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
