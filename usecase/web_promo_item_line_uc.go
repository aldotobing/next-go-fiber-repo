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

// WebPromoItemLineUC ...
type WebPromoItemLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebPromoItemLineUC) BuildBody(res *models.WebPromoItemLine) {
}

// SelectAll ...
func (uc WebPromoItemLineUC) SelectAll(c context.Context, parameter models.WebPromoItemLineParameter) (res []models.WebPromoItemLine, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebPromoItemLineOrderBy, models.WebPromoItemLineOrderByrByString)

	repo := repository.NewWebPromoItemLineRepository(uc.DB)
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
func (uc WebPromoItemLineUC) FindAll(c context.Context, parameter models.WebPromoItemLineParameter) (res []models.WebPromoItemLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebPromoItemLineOrderBy, models.WebPromoItemLineOrderByrByString)

	var count int
	repo := repository.NewWebPromoItemLineRepository(uc.DB)
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
func (uc WebPromoItemLineUC) FindByID(c context.Context, parameter models.WebPromoItemLineParameter) (res models.WebPromoItemLine, err error) {
	repo := repository.NewWebPromoItemLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc WebPromoItemLineUC) Add(c context.Context, data *requests.WebPromoItemLineRequest) (res models.WebPromoItemLineBreakDown, err error) {

	repo := repository.NewWebPromoItemLineRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.WebPromoItemLineBreakDown{
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

// Delete ...
func (uc WebPromoItemLineUC) Delete(c context.Context, id string) (res viewmodel.CommonDeletedObjectVM, err error) {
	repo := repository.NewWebPromoItemLineRepository(uc.DB)
	res.ID, err = repo.Delete(c, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
