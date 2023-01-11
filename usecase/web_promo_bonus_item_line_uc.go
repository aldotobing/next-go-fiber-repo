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

// WebPromoBonusItemLineUC ...
type WebPromoBonusItemLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebPromoBonusItemLineUC) BuildBody(res *models.WebPromoBonusItemLine) {
}

// SelectAll ...
func (uc WebPromoBonusItemLineUC) SelectAll(c context.Context, parameter models.WebPromoBonusItemLineParameter) (res []models.WebPromoBonusItemLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebPromoBonusItemLineOrderBy, models.WebPromoBonusItemLineOrderByrByString)

	repo := repository.NewWebPromoBonusItemLineRepository(uc.DB)
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
func (uc WebPromoBonusItemLineUC) FindAll(c context.Context, parameter models.WebPromoBonusItemLineParameter) (res []models.WebPromoBonusItemLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebPromoBonusItemLineOrderBy, models.WebPromoBonusItemLineOrderByrByString)

	var count int
	repo := repository.NewWebPromoBonusItemLineRepository(uc.DB)
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
func (uc WebPromoBonusItemLineUC) FindByID(c context.Context, parameter models.WebPromoBonusItemLineParameter) (res models.WebPromoBonusItemLine, err error) {
	repo := repository.NewWebPromoBonusItemLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc WebPromoBonusItemLineUC) Add(c context.Context, data *requests.WebPromoBonusItemLineRequest) (res models.WebPromoBonusItemLineBreakDown, err error) {

	repo := repository.NewWebPromoBonusItemLineRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.WebPromoBonusItemLineBreakDown{
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
