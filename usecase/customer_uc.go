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

// CustomerUC ...
type CustomerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerUC) BuildBody(res *models.Customer) {
}

// SelectAll ...
func (uc CustomerUC) SelectAll(c context.Context, parameter models.CustomerParameter) (res []models.Customer, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderBy, models.CustomerOrderByrByString)

	repo := repository.NewCustomerRepository(uc.DB)
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
func (uc CustomerUC) FindAll(c context.Context, parameter models.CustomerParameter) (res []models.Customer, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderBy, models.CustomerOrderByrByString)

	var count int
	repo := repository.NewCustomerRepository(uc.DB)
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
func (uc CustomerUC) FindByID(c context.Context, parameter models.CustomerParameter) (res models.Customer, err error) {

	repo := repository.NewCustomerRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Edit ,...
func (uc CustomerUC) Edit(c context.Context, id string, data *requests.CustomerRequest) (res models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.Customer{
		ID:              &id,
		Code:            &data.Code,
		CustomerName:    &data.CustomerName,
		CustomerAddress: &data.CustomerAddress,
		CustomerPhone:   &data.CustomerPhone,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc CustomerUC) EditAddress(c context.Context, id string, data *requests.CustomerRequest) (res models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.Customer{
		ID:                    &id,
		CustomerName:          &data.CustomerName,
		CustomerAddress:       &data.CustomerAddress,
		CustomerProvinceID:    &data.CustomerProvinceID,
		CustomerCityID:        &data.CustomerCityID,
		CustomerDistrictID:    &data.CustomerDistrictID,
		CustomerSubdistrictID: &data.CustomerSubdistrictID,
		CustomerPostalCode:    &data.CustomerPostalCode,
	}

	res.ID, err = repo.EditAddress(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
