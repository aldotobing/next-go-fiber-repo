package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebRegionAreaUC ...
type WebRegionAreaUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebRegionAreaUC) BuildBody(data *models.WebRegionArea, res *viewmodel.RegionAreaVM) {
	if data.ID != nil {
		res.ID = data.ID
	}
	if data.Code != nil {
		res.Code = data.Code
	}
	if data.Name != nil {
		res.Name = data.Name
	}
	if data.GroupID != nil {
		res.GroupID = data.GroupID
	}
	if data.GroupName != nil {
		res.GroupName = data.GroupName
	}
}

// SelectAll ...
func (uc WebRegionAreaUC) SelectAll(c context.Context, parameter models.WebRegionAreaParameter) (res []models.WebRegionArea, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebRegionAreaOrderBy, models.WebRegionAreaOrderByrByString)

	repo := repository.NewWebRegionAreaRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// SelectAllGroupByRegion ...
func (uc WebRegionAreaUC) SelectAllGroupByRegion(c context.Context) (res []viewmodel.RegionAreaVM, err error) {
	repo := repository.NewWebRegionAreaRepository(uc.DB)
	data, err := repo.SelectAllGroupByRegion(c)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.RegionAreaVM
		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	return res, err
}

// SelectByGroupID ...
func (uc WebRegionAreaUC) SelectByGroupID(c context.Context, groupID string) (res []viewmodel.RegionAreaVM, err error) {
	repo := repository.NewWebRegionAreaRepository(uc.DB)
	data, err := repo.SelectByRegionGroupID(c, groupID)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.RegionAreaVM
		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	return res, err
}

// FindAll ...
func (uc WebRegionAreaUC) FindAll(c context.Context, parameter models.WebRegionAreaParameter) (res []models.WebRegionArea, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebRegionAreaOrderBy, models.WebRegionAreaOrderByrByString)

	var count int
	repo := repository.NewWebRegionAreaRepository(uc.DB)
	res, count, err = repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	return res, p, err
}

// FindByID ...
func (uc WebRegionAreaUC) FindByID(c context.Context, parameter models.WebRegionAreaParameter) (res models.WebRegionArea, err error) {
	repo := repository.NewWebRegionAreaRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
