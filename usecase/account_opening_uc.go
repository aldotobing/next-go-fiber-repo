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

// AccountOpeningUC ...
type AccountOpeningUC struct {
	*ContractUC
}

// BuildBody ...
func (uc AccountOpeningUC) BuildBody(res *models.AccountOpening) {
}

// SelectAll ...
func (uc AccountOpeningUC) SelectAll(c context.Context, parameter models.AccountOpeningParameter) (res []models.AccountOpening, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.AccountOpeningOrderBy, models.AccountOpeningOrderByrByString)

	repo := repository.NewAccountOpeningRepository(uc.DB)
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
func (uc AccountOpeningUC) FindAll(c context.Context, parameter models.AccountOpeningParameter) (res []models.AccountOpening, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.AccountOpeningOrderBy, models.AccountOpeningOrderByrByString)

	var count int
	repo := repository.NewAccountOpeningRepository(uc.DB)
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
func (uc AccountOpeningUC) FindByID(c context.Context, parameter models.AccountOpeningParameter) (res models.AccountOpening, err error) {
	repo := repository.NewAccountOpeningRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByEmail ...
func (uc AccountOpeningUC) FindByEmail(c context.Context, parameter models.AccountOpeningParameter) (res models.AccountOpening, err error) {
	repo := repository.NewAccountOpeningRepository(uc.DB)
	res, err = repo.FindByEmail(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByPhone ...
func (uc AccountOpeningUC) FindByPhone(c context.Context, parameter models.AccountOpeningParameter) (res models.AccountOpening, err error) {
	repo := repository.NewAccountOpeningRepository(uc.DB)
	res, err = repo.FindByPhone(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc AccountOpeningUC) Add(c context.Context, data *requests.AccountOpeningRequest) (res models.AccountOpening, err error) {
	// Check Duplicate Email
	AccountOpeningEmail, _ := uc.FindByEmail(c, models.AccountOpeningParameter{Email: data.Email})
	if AccountOpeningEmail.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.DuplicateEmail, functioncaller.PrintFuncName(), "duplicate_email", c.Value("requestid"))
		return res, errors.New(helper.DuplicateEmail)
	}

	// Check Duplicate Phone Number
	AccountOpeningPhone, _ := uc.FindByPhone(c, models.AccountOpeningParameter{Phone: data.Phone})
	if AccountOpeningPhone.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.DuplicatePhone, functioncaller.PrintFuncName(), "duplicate_phone", c.Value("requestid"))
		return res, errors.New(helper.DuplicatePhone)
	}

	repo := repository.NewAccountOpeningRepository(uc.DB)
	now := time.Now().UTC()
	res = models.AccountOpening{
		UserID:           data.UserID,
		Name:             data.Name,
		Email:            data.Email,
		MaritalStatusID:  data.MaritalStatusID,
		GenderID:         data.GenderID,
		BirthPlace:       data.BirthPlace,
		BirthPlaceCityID: data.BirthPlaceCityID,
		BirthDate:        data.BirthDate,
		MotherName:       data.MotherName,
		Phone:            data.Phone,
		Status:           data.Status,
		CreatedAt:        now.Format(time.RFC3339),
		UpdatedAt:        now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc AccountOpeningUC) Edit(c context.Context, id string, data *requests.AccountOpeningRequest) (res models.AccountOpening, err error) {
	// Check Duplicate Email
	AccountOpeningEmail, _ := uc.FindByEmail(c, models.AccountOpeningParameter{Email: data.Email})
	if AccountOpeningEmail.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.DuplicateEmail, functioncaller.PrintFuncName(), "duplicate_email", c.Value("requestid"))
		return res, errors.New(helper.DuplicateEmail)
	}

	// Check Duplicate Phone Number
	AccountOpeningPhone, _ := uc.FindByPhone(c, models.AccountOpeningParameter{Phone: data.Phone})
	if AccountOpeningPhone.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.DuplicatePhone, functioncaller.PrintFuncName(), "duplicate_phone", c.Value("requestid"))
		return res, errors.New(helper.DuplicatePhone)
	}

	repo := repository.NewAccountOpeningRepository(uc.DB)
	now := time.Now().UTC()
	res = models.AccountOpening{
		ID:        id,
		Name:      data.Name,
		UpdatedAt: now.Format(time.RFC3339),
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc AccountOpeningUC) Delete(c context.Context, id string) (res viewmodel.AccountOpeningVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewAccountOpeningRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}

// EditPhoneValidAt ...
func (uc AccountOpeningUC) EditPhoneValidAt(c context.Context, id string) (res models.AccountOpening, err error) {
	repo := repository.NewAccountOpeningRepository(uc.DB)
	now := time.Now().UTC()
	res = models.AccountOpening{
		ID:           id,
		Status:       models.AccountOpeningStatusActive,
		PhoneValidAt: now.Format(time.RFC3339),
	}
	res.ID, err = repo.EditPhoneValidAt(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-AccountOpening-edit")
		return res, err
	}

	return res, err
}

// EditEmailValidAt ...
func (uc AccountOpeningUC) EditEmailValidAt(c context.Context, id string) (res models.AccountOpening, err error) {
	repo := repository.NewAccountOpeningRepository(uc.DB)
	now := time.Now().UTC()
	res = models.AccountOpening{
		ID:           id,
		EmailValidAt: now.Format(time.RFC3339),
	}
	res.ID, err = repo.EditEmailValidAt(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-AccountOpening-edit")
		return res, err
	}

	return res, err
}
