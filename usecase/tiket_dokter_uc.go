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

// TicketDokterUC ...
type TicketDokterUC struct {
	*ContractUC
}

// BuildBody ...
func (uc TicketDokterUC) BuildBody(res *models.TicketDokter) {
}

// SelectAll ...
func (uc TicketDokterUC) SelectAll(c context.Context, parameter models.TicketDokterParameter) (res []models.TicketDokter, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.TicketDokterOrderBy, models.TicketDokterOrderByrByString)

	repo := repository.NewTicketDokterRepository(uc.DB)
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
func (uc TicketDokterUC) FindAll(c context.Context, parameter models.TicketDokterParameter) (res []models.TicketDokter, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.TicketDokterOrderBy, models.TicketDokterOrderByrByString)

	var count int
	repo := repository.NewTicketDokterRepository(uc.DB)
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
func (uc TicketDokterUC) FindByID(c context.Context, parameter models.TicketDokterParameter) (res models.TicketDokter, err error) {
	repo := repository.NewTicketDokterRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc TicketDokterUC) Add(c context.Context, data *requests.TicketDokterRequest) (res models.TicketDokter, err error) {

	repo := repository.NewTicketDokterRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.TicketDokter{
		CustomerID:       &data.CustomerID,
		CustomerName:     &data.CustomerName,
		CustomerHeight:   &data.CustomerHeight,
		CustomerWeight:   &data.CustomerWeight,
		CustomerAge:      &data.CustomerAge,
		CustomerPhone:    &data.CustomerPhone,
		CustomerAltPhone: &data.CustomerAltPhone,
		CustomerProblem:  &data.CustomerProblem,
		Solution:         &data.Solution,
		Allergy:          &data.Allergy,
		Status:           &data.Status,
		CreatedDate:      &data.CreatedDate,
		ModifiedDate:     &data.ModifiedDate,
		CloseDate:        &data.CloseDate,
		DoctorID:         &data.DoctorID,
		DoctorName:       &data.DoctorName,
		Description:      &data.Description,
		Hide:             &data.Hide,
		TicketCode:       &data.TicketCode,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ...
func (uc TicketDokterUC) Edit(c context.Context, id string, data *requests.TicketDokterRequest) (res models.TicketDokter, err error) {
	repo := repository.NewTicketDokterRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.TicketDokter{
		ID:       &id,
		Status:   &data.Status,
		Solution: &data.Solution,
		DoctorID: &data.DoctorID,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
