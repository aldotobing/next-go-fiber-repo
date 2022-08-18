package usecase

import (
	"context"
	"errors"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/interfacepkg"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// SettingUC ...
type SettingUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SettingUC) BuildBody(data *models.Setting, res *viewmodel.SettingVM) {
	res.ID = data.ID
	res.Code = data.Code.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt
	res.DeletedAt = data.DeletedAt.String

	interfacepkg.UnmarshallCb(data.Details.String, &res.Details)
}

// SelectAll ...
func (uc SettingUC) SelectAll(c context.Context) (res []viewmodel.SettingVM, err error) {
	repo := repository.NewSettingRepository(uc.DB)
	data, err := repo.SelectAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for _, r := range data {
		tempRes := viewmodel.SettingVM{}
		uc.BuildBody(&r, &tempRes)
		res = append(res, tempRes)
	}

	return res, err
}

// FindByID ...
func (uc SettingUC) FindByID(c context.Context, id string) (res viewmodel.SettingVM, err error) {
	repo := repository.NewSettingRepository(uc.DB)
	data, err := repo.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&data, &res)

	return res, err
}

// FindByCode ...
func (uc SettingUC) FindByCode(c context.Context, code string) (res viewmodel.SettingVM, err error) {
	repo := repository.NewSettingRepository(uc.DB)
	data, err := repo.FindByCode(code)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&data, &res)

	return res, err
}

// Add ...
func (uc SettingUC) Add(c context.Context, data *requests.SettingRequest) (res viewmodel.SettingVM, err error) {
	rt, _ := uc.FindByCode(c, data.Code)
	if rt.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.CodeAlredyUsed, functioncaller.PrintFuncName(), "code", c.Value("requestid"))
		return res, errors.New(helper.CodeAlredyUsed)
	}

	repo := repository.NewSettingRepository(uc.DB)
	now := time.Now().UTC()
	res = viewmodel.SettingVM{
		Code:      data.Code,
		Details:   data.Details,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	res.ID, err = repo.Add(&res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc SettingUC) Edit(c context.Context, ID string, data *requests.SettingRequest) (res viewmodel.SettingVM, err error) {
	repo := repository.NewSettingRepository(uc.DB)
	now := time.Now().UTC()
	res = viewmodel.SettingVM{
		Code:      data.Code,
		Details:   data.Details,
		UpdatedAt: now.Format(time.RFC3339),
	}

	res.ID, err = repo.Edit(ID, &res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc SettingUC) Delete(c context.Context, id string) (res viewmodel.SettingVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewSettingRepository(uc.DB)
	res.ID, err = repo.Delete(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
