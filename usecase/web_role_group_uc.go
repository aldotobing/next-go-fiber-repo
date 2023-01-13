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

// WebRoleGroupUC ...
type WebRoleGroupUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebRoleGroupUC) BuildBody(res *models.WebRoleGroup) {
}

// SelectAll ...
func (uc WebRoleGroupUC) SelectAll(c context.Context, parameter models.WebRoleGroupParameter) (res []models.WebRoleGroup, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebRoleGroupOrderBy, models.WebRoleGroupOrderByrByString)

	repo := repository.NewWebRoleGroupRepository(uc.DB)
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
func (uc WebRoleGroupUC) FindAll(c context.Context, parameter models.WebRoleGroupParameter) (res []models.WebRoleGroup, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebRoleGroupOrderBy, models.WebRoleGroupOrderByrByString)

	var count int
	repo := repository.NewWebRoleGroupRepository(uc.DB)
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
func (uc WebRoleGroupUC) FindByID(c context.Context, parameter models.WebRoleGroupParameter) (res models.WebRoleGroup, err error) {
	repo := repository.NewWebRoleGroupRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc WebRoleGroupUC) Add(c context.Context, data *requests.WebRoleGroupRequest) (res models.WebRoleGroup, err error) {

	repo := repository.NewWebRoleGroupRepository(uc.DB)

	res = models.WebRoleGroup{
		Name: &data.Name,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc WebRoleGroupUC) Edit(c context.Context, id string, data *requests.WebRoleGroupRequest) (res models.WebRoleGroup, err error) {
	repo := repository.NewWebRoleGroupRepository(uc.DB)

	res = models.WebRoleGroup{
		ID:   &id,
		Name: &data.Name,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc WebRoleGroupUC) Delete(c context.Context, id string) (res models.WebRoleGroup, err error) {
	now := time.Now().UTC()
	repo := repository.NewWebRoleGroupRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
