package usecase

import (
	"context"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebUserUC ...
type WebUserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebUserUC) BuildBody(res *models.WebUser) {
}

// SelectAll ...
func (uc WebUserUC) SelectAll(c context.Context, parameter models.WebUserParameter) (res []models.WebUser, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebUserOrderBy, models.WebUserOrderByrByString)

	repo := repository.NewWebUserRepository(uc.DB)
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
func (uc WebUserUC) FindAll(c context.Context, parameter models.WebUserParameter) (res []models.WebUser, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebUserOrderBy, models.WebUserOrderByrByString)

	var count int
	repo := repository.NewWebUserRepository(uc.DB)
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
func (uc WebUserUC) FindByID(c context.Context, parameter models.WebUserParameter) (res models.WebUser, err error) {
	repo := repository.NewWebUserRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc WebUserUC) Add(c context.Context, data *requests.WebUserRequest) (res models.WebUser, err error) {

	repo := repository.NewWebUserRepository(uc.DB)

	res = models.WebUser{
		Login:               &data.Login,
		Password:            &data.Password,
		UserRoleGroupIDList: &data.UserRoleGroupIDList,
		BranchIDList:        data.BranchIDList,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc WebUserUC) Edit(c context.Context, id string, data *requests.WebUserRequest) (res models.WebUser, err error) {
	repo := repository.NewWebUserRepository(uc.DB)

	res = models.WebUser{
		ID:                  &id,
		Login:               &data.Login,
		Password:            &data.Password,
		UserRoleGroupIDList: &data.UserRoleGroupIDList,
		BranchIDList:        data.BranchIDList,
	}

	resID, err := repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	res.ID = &resID

	return res, err
}

// Delete ...
func (uc WebUserUC) Delete(c context.Context, id string) (res models.WebUser, err error) {
	now := time.Now().UTC()
	repo := repository.NewWebUserRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
