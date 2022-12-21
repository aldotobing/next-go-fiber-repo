package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// DashboardWebUC ...
type DashboardWebUC struct {
	*ContractUC
}

// BuildBody ...
func (uc DashboardWebUC) BuildBody(res *models.DashboardWeb) {
}

// FindByID ...
func (uc DashboardWebUC) GetData(c context.Context, parameter models.DashboardWebParameter) (res []models.DashboardWeb, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}
