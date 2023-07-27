package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// WebCustomerTypeEligiblePromoUC ...
type WebCustomerTypeEligiblePromoUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebCustomerTypeEligiblePromoUC) BuildBody(res *models.WebCustomerTypeEligiblePromo) {
}

// SelectAll ...
func (uc WebCustomerTypeEligiblePromoUC) SelectAll(c context.Context, parameter models.WebCustomerTypeEligiblePromoParameter) (res []models.WebCustomerTypeEligiblePromo, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerTypeEligiblePromoOrderBy, models.WebCustomerTypeEligiblePromoOrderByrByString)

	repo := repository.NewWebCustomerTypeEligiblePromoRepository(uc.DB)
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
