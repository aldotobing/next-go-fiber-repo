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

// ProvinceUC ...
type ProvinceUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ProvinceUC) BuildBody(res *models.Province) {
}

// SelectAll ...
func (uc ProvinceUC) SelectAll(c context.Context, parameter models.ProvinceParameter) (res []models.Province, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ProvinceOrderBy, models.ProvinceOrderByrByString)

	repo := repository.NewProvinceRepository(uc.DB)
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
func (uc ProvinceUC) FindAll(c context.Context, parameter models.ProvinceParameter) (res []models.Province, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ProvinceOrderBy, models.ProvinceOrderByrByString)

	var count int
	repo := repository.NewProvinceRepository(uc.DB)
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
func (uc ProvinceUC) FindByID(c context.Context, parameter models.ProvinceParameter) (res models.Province, err error) {
	repo := repository.NewProvinceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByCode ...
func (uc ProvinceUC) FindByCode(c context.Context, parameter models.ProvinceParameter) (res models.Province, err error) {
	repo := repository.NewProvinceRepository(uc.DB)
	res, err = repo.FindByCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc ProvinceUC) Add(c context.Context, data *requests.ProvinceRequest) (res models.Province, err error) {
	rt, _ := uc.FindByCode(c, models.ProvinceParameter{Code: data.Code})
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "mapping_name", c.Value("requestid"))
		return res, errors.New(helper.NameAlreadyExist)
	}

	repo := repository.NewProvinceRepository(uc.DB)
	// now := time.Now().UTC()
	// nowstr := now.Format(time.RFC3339)
	res = models.Province{
		Code: &data.Code,
		Name: &data.Name,
		// IdNation:  &data.IdNation,
		// CreatedAt: &nowstr,
		// UpdatedAt: &nowstr,
		// CreatedBy: &data.CreatedBy,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc ProvinceUC) Edit(c context.Context, id string, data *requests.ProvinceRequest) (res models.Province, err error) {
	repo := repository.NewProvinceRepository(uc.DB)
	// now := time.Now().UTC()
	// nowstr := now.Format(time.RFC3339)
	res = models.Province{
		ID:   id,
		Code: &data.Code,
		Name: &data.Name,
		// IdNation:  &data.IdNation,
		// UpdatedAt: &nowstr,
		// UpdatedBy: &data.UpdatedBy,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc ProvinceUC) Delete(c context.Context, id string) (res viewmodel.ProvinceVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewProvinceRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}

func (uc ProvinceUC) AddAll(c context.Context, data *[]requests.MpProvinceDataBreakDownRequest) (res []models.MpProvinceDataBreakDown, err error) {

	repo := repository.NewProvinceRepository(uc.DB)

	for _, input := range *data {

		provincesOject := models.MpProvinceDataBreakDown{
			Name:     input.Name,
			Code:     input.Code,
			OldID:    input.OldID,
			NationID: input.NationID,
		}
		provincesOject.ID, err = repo.AddDataBreakDown(c, &provincesOject)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}
		res = append(res, provincesOject)
	}

	return res, err
}
