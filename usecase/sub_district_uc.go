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

// SubDistrictUC ...
type SubDistrictUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SubDistrictUC) BuildBody(res *models.SubDistrict) {
}

// SelectAll ...
func (uc SubDistrictUC) SelectAll(c context.Context, parameter models.SubDistrictParameter) (res []models.SubDistrict, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.SubDistrictOrderBy, models.SubDistrictOrderByrByString)

	repo := repository.NewSubDistrictRepository(uc.DB)
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
func (uc SubDistrictUC) FindAll(c context.Context, parameter models.SubDistrictParameter) (res []models.SubDistrict, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.SubDistrictOrderBy, models.SubDistrictOrderByrByString)

	var count int
	repo := repository.NewSubDistrictRepository(uc.DB)
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
func (uc SubDistrictUC) FindByID(c context.Context, parameter models.SubDistrictParameter) (res models.SubDistrict, err error) {
	repo := repository.NewSubDistrictRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByCode ...
func (uc SubDistrictUC) FindByCode(c context.Context, parameter models.SubDistrictParameter) (res models.SubDistrict, err error) {
	repo := repository.NewSubDistrictRepository(uc.DB)
	res, err = repo.FindByCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc SubDistrictUC) Add(c context.Context, data *requests.SubDistrictRequest) (res models.SubDistrict, err error) {
	rt, _ := uc.FindByCode(c, models.SubDistrictParameter{Code: data.Code})
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "mapping_name", c.Value("requestid"))
		return res, errors.New(helper.NameAlreadyExist)
	}

	repo := repository.NewSubDistrictRepository(uc.DB)
	// now := time.Now().UTC()
	res = models.SubDistrict{
		// DistrictID: data.DistrictID,
		Code: data.Code,
		Name: data.Name,
		// PostalCode: data.PostalCode,
		// Status:     data.Status,
		// CreatedAt:  now.Format(time.RFC3339),
		// UpdatedAt:  now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc SubDistrictUC) Edit(c context.Context, id string, data *requests.SubDistrictRequest) (res models.SubDistrict, err error) {
	repo := repository.NewSubDistrictRepository(uc.DB)
	res = models.SubDistrict{
		ID:   id,
		Code: data.Code,
		Name: data.Name,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc SubDistrictUC) Delete(c context.Context, id string) (res viewmodel.SubDistrictVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewSubDistrictRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
