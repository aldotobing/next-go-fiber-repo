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

// PromoLineUC ...
type PromoLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PromoLineUC) BuildBody(res *models.PromoLine) {
}

// SelectAll ...
func (uc PromoLineUC) SelectAll(c context.Context, parameter models.PromoLineParameter) (res []models.PromoLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.PromoLineOrderBy, models.PromoLineOrderByrByString)

	repo := repository.NewPromoLineRepository(uc.DB)
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
func (uc PromoLineUC) FindAll(c context.Context, parameter models.PromoLineParameter) (res []models.PromoLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PromoLineOrderBy, models.PromoLineOrderByrByString)

	var count int
	repo := repository.NewPromoLineRepository(uc.DB)
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

// Add ...
func (uc PromoLineUC) Add(c context.Context, data *requests.PromoLineRequest) (res models.PromoLine, err error) {

	repo := repository.NewPromoLineRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.PromoLine{
		PromoID:         &data.PromoID,
		GlobalMaxQty:    &data.GlobalMaxQty,
		CustomerMaxQty:  &data.CustomerMaxQty,
		DiscPercent:     &data.DiscPercent,
		DiscAmount:      &data.DiscAmount,
		MinimumValue:    &data.MinimumValue,
		Multiply:        &data.Multiply,
		Description:     &data.Description,
		MinimumQty:      &data.MinimumQty,
		MinimumQtyUomID: &data.MinimumQtyUomID,
		PromoType:       &data.PromoType,
		Strata:          &data.Strata,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
