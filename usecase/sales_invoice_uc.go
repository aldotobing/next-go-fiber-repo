package usecase

import (
	"context"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// SalesInvoiceUC ...
type SalesInvoiceUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SalesInvoiceUC) BuildBody(res *models.SalesInvoice) {
}

// SelectAll ...
func (uc SalesInvoiceUC) SelectAll(c context.Context, parameter models.SalesInvoiceParameter) (res []models.SalesInvoice, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.SalesInvoiceOrderBy, models.SalesInvoiceOrderByrByString)

	repo := repository.NewSalesInvoiceRepository(uc.DB)
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
func (uc SalesInvoiceUC) FindAll(c context.Context, parameter models.SalesInvoiceParameter) (res []models.SalesInvoice, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.SalesInvoiceOrderBy, models.SalesInvoiceOrderByrByString)

	var count int
	repo := repository.NewSalesInvoiceRepository(uc.DB)
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
func (uc SalesInvoiceUC) FindByID(c context.Context, parameter models.SalesInvoiceParameter) (res models.SalesInvoice, err error) {
	repo := repository.NewSalesInvoiceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByDocumentNo ...
func (uc SalesInvoiceUC) FindByDocumentNo(c context.Context, parameter models.SalesInvoiceParameter) (res models.SalesInvoice, err error) {
	repo := repository.NewSalesInvoiceRepository(uc.DB)
	res, err = repo.FindByDocumentNo(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByCustomerId ...
func (uc SalesInvoiceUC) FindByCustomerId(c context.Context, parameter models.SalesInvoiceParameter) (res models.SalesInvoice, err error) {
	repo := repository.NewSalesInvoiceRepository(uc.DB)
	res, err = repo.FindByCustomerId(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Edit ...
func (uc SalesInvoiceUC) Edit(c context.Context, id string, input *requests.SalesInvoiceRequest) (res models.SalesInvoice, err error) {

	now := time.Now().UTC()
	strnow := now.Format(time.RFC3339)

	res = models.SalesInvoice{
		ID:            &id,
		ModifiedDate:  &strnow,
		TotalPaid:     &input.TotalPaid,
		PaymentMethod: &input.PaymentMethod,
	}
	repo := repository.NewSalesInvoiceRepository(uc.DB)
	_, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
