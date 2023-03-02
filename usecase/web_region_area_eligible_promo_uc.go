package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebRegionAreaEligiblePromoUC ...
type WebRegionAreaEligiblePromoUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebRegionAreaEligiblePromoUC) BuildBody(res *models.WebRegionAreaEligiblePromo) {
}

// SelectAll ...
func (uc WebRegionAreaEligiblePromoUC) SelectAll(c context.Context, parameter models.WebRegionAreaEligiblePromoParameter) (res []models.WebRegionAreaEligiblePromo, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebRegionAreaEligiblePromoOrderBy, models.WebRegionAreaEligiblePromoOrderByrByString)

	repo := repository.NewWebRegionAreaEligiblePromoRepository(uc.DB)
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
func (uc WebRegionAreaEligiblePromoUC) FindAll(c context.Context, parameter models.WebRegionAreaEligiblePromoParameter) (res []models.WebRegionAreaEligiblePromo, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebRegionAreaEligiblePromoOrderBy, models.WebRegionAreaEligiblePromoOrderByrByString)

	var count int
	repo := repository.NewWebRegionAreaEligiblePromoRepository(uc.DB)
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
