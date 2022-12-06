package usecase

import (
	"context"
	"errors"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// DistrictUC ...
type DistrictUC struct {
	*ContractUC
}

// BuildBody ...
func (uc DistrictUC) BuildBody(res *models.District) {
}

// SelectAll ...
func (uc DistrictUC) SelectAll(c context.Context, parameter models.DistrictParameter) (res []models.District, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.DistrictOrderBy, models.DistrictOrderByrByString)

	repo := repository.NewDistrictRepository(uc.DB)
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
func (uc DistrictUC) FindAll(c context.Context, parameter models.DistrictParameter) (res []models.District, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DistrictOrderBy, models.DistrictOrderByrByString)

	var count int
	repo := repository.NewDistrictRepository(uc.DB)
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
func (uc DistrictUC) FindByID(c context.Context, parameter models.DistrictParameter) (res models.District, err error) {
	repo := repository.NewDistrictRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByCode ...
func (uc DistrictUC) FindByCode(c context.Context, parameter models.DistrictParameter) (res models.District, err error) {
	repo := repository.NewDistrictRepository(uc.DB)
	res, err = repo.FindByCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc DistrictUC) Add(c context.Context, data *requests.DistrictRequest) (res models.District, err error) {
	rt, _ := uc.FindByCode(c, models.DistrictParameter{Code: data.Code})
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "mapping_name", c.Value("requestid"))
		return res, errors.New(helper.NameAlreadyExist)
	}

	repo := repository.NewDistrictRepository(uc.DB)
	// now := time.Now().UTC()
	res = models.District{
		// CityID:    data.CityID,
		Code: data.Code,
		Name: data.Name,
		// Status:    data.Status,
		// CreatedAt: now.Format(time.RFC3339),
		// UpdatedAt: now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc DistrictUC) Edit(c context.Context, id string, data *requests.DistrictRequest) (res models.District, err error) {
	repo := repository.NewDistrictRepository(uc.DB)
	// now := time.Now().UTC()
	res = models.District{
		ID: id,
		// CityID:    data.CityID,
		Code: data.Code,
		Name: data.Name,
		// Status:    data.Status,
		// UpdatedAt: now.Format(time.RFC3339),
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc DistrictUC) Delete(c context.Context, id string) (res viewmodel.DistrictVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewDistrictRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
