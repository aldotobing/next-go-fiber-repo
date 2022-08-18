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

// ResidenceOwnershipUC ...
type ResidenceOwnershipUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ResidenceOwnershipUC) BuildBody(res *models.ResidenceOwnership) {
}

// SelectAll ...
func (uc ResidenceOwnershipUC) SelectAll(c context.Context, parameter models.ResidenceOwnershipParameter) (res []models.ResidenceOwnership, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ResidenceOwnershipOrderBy, models.ResidenceOwnershipOrderByrByString)

	repo := repository.NewResidenceOwnershipRepository(uc.DB)
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
func (uc ResidenceOwnershipUC) FindAll(c context.Context, parameter models.ResidenceOwnershipParameter) (res []models.ResidenceOwnership, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ResidenceOwnershipOrderBy, models.ResidenceOwnershipOrderByrByString)

	var count int
	repo := repository.NewResidenceOwnershipRepository(uc.DB)
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
func (uc ResidenceOwnershipUC) FindByID(c context.Context, parameter models.ResidenceOwnershipParameter) (res models.ResidenceOwnership, err error) {
	repo := repository.NewResidenceOwnershipRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByMappingName ...
func (uc ResidenceOwnershipUC) FindByMappingName(c context.Context, parameter models.ResidenceOwnershipParameter) (res models.ResidenceOwnership, err error) {
	repo := repository.NewResidenceOwnershipRepository(uc.DB)
	res, err = repo.FindByMappingName(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc ResidenceOwnershipUC) Add(c context.Context, data *requests.ResidenceOwnershipRequest) (res models.ResidenceOwnership, err error) {
	rt, _ := uc.FindByMappingName(c, models.ResidenceOwnershipParameter{MappingName: data.MappingName})
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "mapping_name", c.Value("requestid"))
		return res, errors.New(helper.NameAlreadyExist)
	}

	repo := repository.NewResidenceOwnershipRepository(uc.DB)
	now := time.Now().UTC()
	res = models.ResidenceOwnership{
		Name:        data.Name,
		MappingName: data.MappingName,
		Status:      data.Status,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc ResidenceOwnershipUC) Edit(c context.Context, id string, data *requests.ResidenceOwnershipRequest) (res models.ResidenceOwnership, err error) {
	repo := repository.NewResidenceOwnershipRepository(uc.DB)
	now := time.Now().UTC()
	res = models.ResidenceOwnership{
		ID:          id,
		Name:        data.Name,
		MappingName: data.MappingName,
		Status:      data.Status,
		UpdatedAt:   now.Format(time.RFC3339),
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc ResidenceOwnershipUC) Delete(c context.Context, id string) (res viewmodel.ResidenceOwnershipVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewResidenceOwnershipRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
