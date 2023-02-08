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

// TransactionVAUC ...
type TransactionVAUC struct {
	*ContractUC
}

// BuildBody ...
func (uc TransactionVAUC) BuildBody(res *models.TransactionVA) {
}

// SelectAll ...
func (uc TransactionVAUC) SelectAll(c context.Context, parameter models.TransactionVAParameter) (res []models.TransactionVA, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.TransactionVAOrderBy, models.TransactionVAOrderByrByString)

	repo := repository.NewTransactionVARepository(uc.DB)
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
func (uc TransactionVAUC) FindAll(c context.Context, parameter models.TransactionVAParameter) (res []models.TransactionVA, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.TransactionVAOrderBy, models.TransactionVAOrderByrByString)

	var count int
	repo := repository.NewTransactionVARepository(uc.DB)
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
func (uc TransactionVAUC) FindByID(c context.Context, parameter models.TransactionVAParameter) (res models.TransactionVA, err error) {

	repo := repository.NewTransactionVARepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc TransactionVAUC) Edit(c context.Context, id string, data *requests.TransactionVARequest) (res models.TransactionVA, err error) {

	repo := repository.NewTransactionVARepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.TransactionVA{
		ID: &id,
		// Code:           &data.Code,
		// PartnerName:    &data.PartnerName,
		// PartnerAddress: &data.PartnerAddress,
		// PartnerPhone:   &data.PartnerPhone,
		// PartnerUserID:  &data.PartnerUserID,
		// PartnerEmail:   &data.PartnerEmail,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc TransactionVAUC) Add(c context.Context, data *requests.TransactionVARequest) (res models.TransactionVA, err error) {

	findrepo := repository.NewTransactionVARepository(uc.DB)
	LastTransaction, errfind := findrepo.FindLastActiveVa(c, models.TransactionVAParameter{InvoiceCode: *&data.InvoiceCode})
	if errfind != nil {
		repo := repository.NewTransactionVARepository(uc.DB)
		// now := time.Now().UTC()
		// strnow := now.Format(time.RFC3339)
		res = models.TransactionVA{
			InvoiceCode:   &data.InvoiceCode,
			VACode:        &data.VACode,
			VaPairID:      &data.VaPairID,
			VaRef1:        &data.VaRef1,
			VaRef2:        &data.VaRef2,
			Amount:        &data.Amount,
			StartDate:     &data.StartDate,
			EndDate:       &data.EndDate,
			VAPartnerCode: &data.VAPartnerCode,
		}

		res.ID, err = repo.Add(c, &res)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}

		currenttrans, errcurrent := findrepo.FindByID(c, models.TransactionVAParameter{ID: *res.ID})
		if errcurrent == nil {
			res = currenttrans
		}

	} else {
		res = LastTransaction
	}

	return res, err
}
