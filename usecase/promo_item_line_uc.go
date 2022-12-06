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

// PromoItemLineUC ...
type PromoItemLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PromoItemLineUC) BuildBody(res *models.PromoItemLine) {
}

// SelectAll ...
func (uc PromoItemLineUC) SelectAll(c context.Context, parameter models.PromoItemLineParameter) (res []models.PromoItemLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.PromoItemLineOrderBy, models.PromoItemLineOrderByrByString)

	repo := repository.NewPromoItemLineRepository(uc.DB)
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
func (uc PromoItemLineUC) FindAll(c context.Context, parameter models.PromoItemLineParameter) (res []models.PromoItemLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PromoItemLineOrderBy, models.PromoItemLineOrderByrByString)

	var count int
	repo := repository.NewPromoItemLineRepository(uc.DB)
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
func (uc PromoItemLineUC) FindByID(c context.Context, parameter models.PromoItemLineParameter) (res models.PromoItemLine, err error) {
	repo := repository.NewPromoItemLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc PromoItemLineUC) Add(c context.Context, data *requests.PromoItemLineRequest) (res models.PromoItemLineBreakDown, err error) {

	repo := repository.NewPromoItemLineRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.PromoItemLineBreakDown{
		PromoLineID: &data.PromoLineID,
		ItemID:      &data.ItemID,
		UomID:       &data.UomID,
		Qty:         &data.Qty,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
