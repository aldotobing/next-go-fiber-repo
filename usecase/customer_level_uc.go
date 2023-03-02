package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerLevelUC ...
type CustomerLevelUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerLevelUC) BuildBody(data *models.CustomerLevel, res *viewmodel.CustomerLevelVM) {
	res.ID = *data.ID
	res.Code = *data.Code
	res.Name = *data.Name
}

// FindAll ...
func (uc CustomerLevelUC) FindAll(c context.Context, parameter models.CustomerLevelParameter) (res []viewmodel.CustomerLevelVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerLevelOrderBy, models.CustomerLevelOrderByrByString)

	repo := repository.NewCustomerLevelRepository(uc.DB)
	data, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for _, datum := range data {
		var temp viewmodel.CustomerLevelVM
		uc.BuildBody(&datum, &temp)

		res = append(res, temp)
	}

	return res, err
}
