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

// WebSalesmanUC ...
type WebSalesmanUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebSalesmanUC) BuildBody(res *models.WebSalesman) {
}

// SelectAll ...
func (uc WebSalesmanUC) SelectAll(c context.Context, parameter models.WebSalesmanParameter) (res []models.WebSalesman, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebSalesmanOrderBy, models.WebSalesmanOrderByrByString)

	repo := repository.NewWebSalesmanRepository(uc.DB)
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
func (uc WebSalesmanUC) FindAll(c context.Context, parameter models.WebSalesmanParameter) (res []models.WebSalesman, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebSalesmanOrderBy, models.WebSalesmanOrderByrByString)

	var count int
	repo := repository.NewWebSalesmanRepository(uc.DB)
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
func (uc WebSalesmanUC) FindByID(c context.Context, parameter models.WebSalesmanParameter) (res models.WebSalesman, err error) {

	repo := repository.NewWebSalesmanRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc WebSalesmanUC) Edit(c context.Context, id string, data *requests.WebSalesmanRequest) (res models.WebSalesman, err error) {

	repo := repository.NewWebSalesmanRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebSalesman{
		ID:            &id,
		Code:          &data.Code,
		PartnerName:   &data.PartnerName,
		PartnerPhone:  &data.PartnerPhone,
		PartnerUserID: &data.PartnerUserID,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc WebSalesmanUC) Add(c context.Context, data *requests.WebSalesmanRequest) (res models.WebSalesman, err error) {

	repo := repository.NewWebSalesmanRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebSalesman{
		Code:          &data.Code,
		PartnerName:   &data.PartnerName,
		PartnerPhone:  &data.PartnerPhone,
		PartnerUserID: &data.PartnerUserID,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
