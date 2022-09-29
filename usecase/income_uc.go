package usecase

import (
	"context"
	"errors"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IncomeUC ...
type IncomeUC struct {
	*ContractUC
}

// BuildBody ...
func (uc IncomeUC) BuildBody(res *models.Income) {
}

// SelectAll ...
func (uc IncomeUC) SelectAll(c context.Context, parameter models.IncomeParameter) (res []models.Income, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.IncomeOrderBy, models.IncomeOrderByrByString)

	repo := repository.NewIncomeRepository(uc.DB)
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
func (uc IncomeUC) FindAll(c context.Context, parameter models.IncomeParameter) (res []models.Income, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.IncomeOrderBy, models.IncomeOrderByrByString)

	var count int
	repo := repository.NewIncomeRepository(uc.DB)
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
func (uc IncomeUC) FindByID(c context.Context, parameter models.IncomeParameter) (res models.Income, err error) {
	repo := repository.NewIncomeRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByMappingName ...
func (uc IncomeUC) FindByMappingName(c context.Context, parameter models.IncomeParameter) (res models.Income, err error) {
	repo := repository.NewIncomeRepository(uc.DB)
	res, err = repo.FindByMappingName(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc IncomeUC) Add(c context.Context, data *requests.IncomeRequest) (res models.Income, err error) {
	rt, _ := uc.FindByMappingName(c, models.IncomeParameter{MappingName: data.MappingName})
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "mapping_name", c.Value("requestid"))
		return res, errors.New(helper.NameAlreadyExist)
	}

	repo := repository.NewIncomeRepository(uc.DB)
	now := time.Now().UTC()
	res = models.Income{
		Name:        data.Name,
		MappingName: data.MappingName,
		MinValue:    data.MinValue,
		MaxValue:    data.MaxValue,
		Status:      data.Status,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc IncomeUC) Edit(c context.Context, id string, data *requests.IncomeRequest) (res models.Income, err error) {
	repo := repository.NewIncomeRepository(uc.DB)
	now := time.Now().UTC()
	res = models.Income{
		ID:          id,
		Name:        data.Name,
		MappingName: data.MappingName,
		MinValue:    data.MinValue,
		MaxValue:    data.MaxValue,
		Status:      data.Status,
		UpdatedAt:   now.Format(time.RFC3339),
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc IncomeUC) Delete(c context.Context, id string) (res viewmodel.IncomeVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewIncomeRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}

func (uc IncomeUC) Send(c context.Context) (res string, err error) {
	fcmUc := FCMUC{ContractUC: uc.ContractUC}
	fireBaseMessage := models.FireBaseCloudMessage{
		Body:  "Fcm Structure",
		Title: "Fcm Structure",
		Token: "cI5dKxAGSSyGPC5qDNoacn:APA91bEP3GokGS6nFxITK16TgeKoBxQStlkkmY-A1ZqRdJVWbSVBrBJGUy5jPmYrTCuAGFgUhKUoutEanCygm-IVbtletnlR39F9r31suL0D9-_D7hmroCHcmCJ-3hlMcqiPAZJnYqEk",
		Type:  "Pesan",
	}

	fcmmesage, _ := fcmUc.SendMessage(c, &fireBaseMessage)

	return fcmmesage, err
}
