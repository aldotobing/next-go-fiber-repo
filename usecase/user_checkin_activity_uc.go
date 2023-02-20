package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// UserCheckinActivityUC ...
type UserCheckinActivityUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserCheckinActivityUC) BuildBody(res *models.UserCheckinActivity) {
}

// SelectAll ...
func (uc UserCheckinActivityUC) SelectAll(c context.Context, parameter models.UserCheckinActivityParameter) (res []models.UserCheckinActivity, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.UserCheckinActivityOrderBy, models.UserCheckinActivityOrderByrByString)

	repo := repository.NewUserCheckinActivityRepository(uc.DB)
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
func (uc UserCheckinActivityUC) FindAll(c context.Context, parameter models.UserCheckinActivityParameter) (res []models.UserCheckinActivity, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.UserCheckinActivityOrderBy, models.UserCheckinActivityOrderByrByString)

	var count int
	repo := repository.NewUserCheckinActivityRepository(uc.DB)
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
func (uc UserCheckinActivityUC) FindByID(c context.Context, parameter models.UserCheckinActivityParameter) (res models.UserCheckinActivity, err error) {

	repo := repository.NewUserCheckinActivityRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc UserCheckinActivityUC) Add(c context.Context, data *requests.UserCheckinActivityRequest) (res models.UserCheckinActivity, err error) {

	repo := repository.NewUserCheckinActivityRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.UserCheckinActivity{
		UserID: &data.UserID,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc UserCheckinActivityUC) FindActiveCheckin(c context.Context, parameter models.UserCheckinActivityParameter) (res models.UserCheckinActivity, err error) {

	repo := repository.NewUserCheckinActivityRepository(uc.DB)
	res, err = repo.FindActiveCheckin(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
