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

// WeebRoleUC ...
type WeebRoleUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WeebRoleUC) BuildBody(res *models.WeebRole) {
}

// SelectAll ...
func (uc WeebRoleUC) SelectAll(c context.Context, parameter models.WeebRoleParameter) (res []models.WeebRole, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WeebRoleOrderBy, models.WeebRoleOrderByrByString)

	repo := repository.NewWeebRoleRepository(uc.DB)
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
func (uc WeebRoleUC) FindAll(c context.Context, parameter models.WeebRoleParameter) (res []models.WeebRole, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WeebRoleOrderBy, models.WeebRoleOrderByrByString)

	var count int
	repo := repository.NewWeebRoleRepository(uc.DB)
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
func (uc WeebRoleUC) FindByID(c context.Context, parameter models.WeebRoleParameter) (res models.WeebRole, err error) {
	repo := repository.NewWeebRoleRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc WeebRoleUC) Add(c context.Context, data *requests.WeebRoleRequest) (res models.WeebRole, err error) {

	repo := repository.NewWeebRoleRepository(uc.DB)

	res = models.WeebRole{
		ID:     &data.ID,
		Name:   &data.Name,
		Header: &data.Header,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc WeebRoleUC) Edit(c context.Context, id string, data *requests.WeebRoleRequest) (res models.WeebRole, err error) {
	repo := repository.NewWeebRoleRepository(uc.DB)

	res = models.WeebRole{
		ID:     &id,
		Name:   &data.Name,
		Header: &data.Header,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc WeebRoleUC) Delete(c context.Context, id string) (res models.WeebRole, err error) {
	now := time.Now().UTC()
	repo := repository.NewWeebRoleRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
