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

// WebPartnerUC ...
type WebPartnerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebPartnerUC) BuildBody(res *models.WebPartner) {
}

// SelectAll ...
func (uc WebPartnerUC) SelectAll(c context.Context, parameter models.WebPartnerParameter) (res []models.WebPartner, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebPartnerOrderBy, models.WebPartnerOrderByrByString)

	repo := repository.NewWebPartnerRepository(uc.DB)
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
func (uc WebPartnerUC) FindAll(c context.Context, parameter models.WebPartnerParameter) (res []models.WebPartner, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebPartnerOrderBy, models.WebPartnerOrderByrByString)

	var count int
	repo := repository.NewWebPartnerRepository(uc.DB)
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
func (uc WebPartnerUC) FindByID(c context.Context, parameter models.WebPartnerParameter) (res models.WebPartner, err error) {

	repo := repository.NewWebPartnerRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc WebPartnerUC) Edit(c context.Context, id string, data *requests.WebPartnerRequest) (res models.WebPartner, err error) {

	repo := repository.NewWebPartnerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebPartner{
		ID:             &id,
		Code:           &data.Code,
		PartnerName:    &data.PartnerName,
		PartnerAddress: &data.PartnerAddress,
		PartnerPhone:   &data.PartnerPhone,
		PartnerUserID:  &data.PartnerUserID,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc WebPartnerUC) Add(c context.Context, data *requests.WebPartnerRequest) (res models.WebPartner, err error) {

	repo := repository.NewWebPartnerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebPartner{
		Code:           &data.Code,
		PartnerName:    &data.PartnerName,
		PartnerAddress: &data.PartnerAddress,
		PartnerPhone:   &data.PartnerPhone,
		PartnerUserID:  &data.PartnerUserID,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
