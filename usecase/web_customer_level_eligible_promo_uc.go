package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebCustomerLevelEligiblePromoUC ...
type WebCustomerLevelEligiblePromoUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebCustomerLevelEligiblePromoUC) BuildBody(res *models.WebCustomerLevelEligiblePromo) {
}

// SelectAll ...
func (uc WebCustomerLevelEligiblePromoUC) SelectAll(c context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) (res []models.WebCustomerLevelEligiblePromo, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerLevelEligiblePromoOrderBy, models.WebCustomerLevelEligiblePromoOrderByrByString)

	repo := repository.NewWebCustomerLevelEligiblePromoRepository(uc.DB)
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
func (uc WebCustomerLevelEligiblePromoUC) FindAll(c context.Context, parameter models.WebCustomerLevelEligiblePromoParameter) (res []models.WebCustomerLevelEligiblePromo, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerLevelEligiblePromoOrderBy, models.WebCustomerLevelEligiblePromoOrderByrByString)

	var count int
	repo := repository.NewWebCustomerLevelEligiblePromoRepository(uc.DB)
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
