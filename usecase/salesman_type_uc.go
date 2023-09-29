package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// SalesmanTypeUC ...
type SalesmanTypeUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SalesmanTypeUC) BuildBody(in *models.SalesmanType, out *viewmodel.SalesmanTypeVM) {
	out.ID = in.ID
	out.Code = in.Code
	out.Name = in.Name
}

// SelectAll ...
func (uc SalesmanTypeUC) SelectAll(c context.Context, parameter models.SalesmanParameter) (res []viewmodel.SalesmanTypeVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.SalesmanTypeOrderBy, models.SalesmanTypeOrderByrByString)

	repo := repository.NewSalesmanTypeRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for _, datum := range data {
		var temp viewmodel.SalesmanTypeVM
		uc.BuildBody(&datum, &temp)

		res = append(res, temp)
	}

	if res == nil {
		res = make([]viewmodel.SalesmanTypeVM, 0)
	}

	return res, err
}
