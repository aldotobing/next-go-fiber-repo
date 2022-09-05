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

// CustomerOrderHeaderUC ...
type CustomerOrderHeaderUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerOrderHeaderUC) BuildBody(res *models.CustomerOrderHeader) {
}

// SelectAll ...
func (uc CustomerOrderHeaderUC) SelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
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
func (uc CustomerOrderHeaderUC) FindAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	var count int
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
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
func (uc CustomerOrderHeaderUC) FindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (res models.CustomerOrderHeader, err error) {
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc CustomerOrderHeaderUC) CheckOut(c context.Context, data *requests.CustomerOrderHeaderRequest) (res models.CustomerOrderHeader, err error) {

	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	round_amount := "0"
	gross_amount := "0"
	taxable_amount := "0"
	tax_amount := "0"
	net_amount := "0"
	disc_amount := "0"
	res = models.CustomerOrderHeader{
		TransactionDate:      &data.TransactionDate,
		TransactionTime:      &data.TransactionTime,
		CustomerID:           &data.CustomerID,
		PaymentTermsID:       &data.PaymentTermsID,
		ExpectedDeliveryDate: &data.ExpectedDeliveryDate,
		GrossAmount:          &gross_amount,
		DiscAmount:           &disc_amount,
		TaxableAmount:        &taxable_amount,
		TaxAmount:            &tax_amount,
		RoundingAmount:       &round_amount,
		NetAmount:            &net_amount,
		TaxCalcMethod:        &data.TaxCalcMethod,
		SalesmanID:           &data.SalesmanID,
		BranchID:             &data.BranchID,
		PriceLIstID:          &data.PriceLIstID,
		LineList:             &data.LineList,
	}
	res.ID, err = repo.CheckOut(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
