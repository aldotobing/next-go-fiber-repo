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

// WebRoleGroupRoleLineUC ...
type WebRoleGroupRoleLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebRoleGroupRoleLineUC) BuildBody(res *models.WebRoleGroupRoleLine) {
}

// SelectAll ...
func (uc WebRoleGroupRoleLineUC) SelectAll(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (res []models.WebRoleGroupRoleLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebRoleGroupRoleLineOrderBy, models.WebRoleGroupRoleLineOrderByrByString)

	repo := repository.NewWebRoleGroupRoleLineRepository(uc.DB)
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
func (uc WebRoleGroupRoleLineUC) FindAll(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (res []models.WebRoleGroupRoleLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebRoleGroupRoleLineOrderBy, models.WebRoleGroupRoleLineOrderByrByString)

	var count int
	repo := repository.NewWebRoleGroupRoleLineRepository(uc.DB)
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
func (uc WebRoleGroupRoleLineUC) FindByID(c context.Context, parameter models.WebRoleGroupRoleLineParameter) (res models.WebRoleGroupRoleLine, err error) {
	repo := repository.NewWebRoleGroupRoleLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc WebRoleGroupRoleLineUC) Add(c context.Context, data *requests.WebRoleGroupRoleLineRequest) (res models.WebRoleGroupRoleLine, err error) {

	repo := repository.NewWebRoleGroupRoleLineRepository(uc.DB)

	res = models.WebRoleGroupRoleLine{
		RoleID:      &data.RoleID,
		RoleGroupID: &data.RoleGroupID,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc WebRoleGroupRoleLineUC) Delete(c context.Context, id string) (res models.WebRoleGroupRoleLine, err error) {
	now := time.Now().UTC()
	repo := repository.NewWebRoleGroupRoleLineRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
