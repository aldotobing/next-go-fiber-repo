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

// BranchUC ...
type BranchUC struct {
	*ContractUC
}

// BuildBody ...
func (uc BranchUC) BuildBody(res *models.Branch) {
}

// SelectAll ...
func (uc BranchUC) SelectAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

	repo := repository.NewBranchRepository(uc.DB)
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
func (uc BranchUC) FindAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

	var count int
	repo := repository.NewBranchRepository(uc.DB)
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
func (uc BranchUC) FindByID(c context.Context, parameter models.BranchParameter) (res models.Branch, err error) {
	repo := repository.NewBranchRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc BranchUC) Update(c context.Context, id string, in *requests.BranchRequest) (res models.Branch, err error) {
	res = models.Branch{
		ID:         &id,
		PICPhoneNo: &in.PICPhoneNo,
		PICName:    &in.PICName,
	}

	repo := repository.NewBranchRepository(uc.DB)
	_, err = repo.Update(c, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return
}
