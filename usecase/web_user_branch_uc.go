package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebUserBranchUC ...
type WebUserBranchUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebUserBranchUC) BuildBody(data *models.WebUserBranch, res *viewmodel.WebUserBranchVM) {
	res.ID = data.ID
	res.UserID = data.UserID
	res.BranchID = data.BranchID
	res.BranchName = data.BranchName
	res.BranchCode = data.BranchCode
}

// SelectAll ...
func (uc WebUserBranchUC) SelectAll(c context.Context, parameter models.WebUserBranchParameter) (res []viewmodel.WebUserBranchVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebUserBranchOrderBy, models.WebUserBranchOrderByrByString)

	repo := repository.NewWebUserBranchRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.WebUserBranchVM

		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	return res, err
}

// FindAll ...
func (uc WebUserBranchUC) FindAll(c context.Context, parameter models.WebUserBranchParameter) (res []viewmodel.WebUserBranchVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebUserBranchOrderBy, models.WebUserBranchOrderByrByString)

	var count int
	repo := repository.NewWebUserBranchRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	for i := range data {
		var temp viewmodel.WebUserBranchVM

		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	return res, p, err
}

// FindByID ...
func (uc WebUserBranchUC) FindByID(c context.Context, parameter models.WebUserBranchParameter) (res viewmodel.WebUserBranchVM, err error) {
	repo := repository.NewWebUserBranchRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	uc.BuildBody(&data, &res)

	return res, err
}
