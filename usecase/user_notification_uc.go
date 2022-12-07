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

type UserNotificationUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserNotificationUC) BuildBody(res *models.UserNotification) {
}

// SelectAll ...
func (uc UserNotificationUC) SelectAll(c context.Context, parameter models.UserNotificationParameter) (res []models.UserNotification, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.UserNotificationOrderBy, models.UserNotificationByrByString)

	repo := repository.NewUserNotificationRepository(uc.DB)
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
func (uc UserNotificationUC) FindAll(c context.Context, parameter models.UserNotificationParameter) (res []models.UserNotification, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.UserNotificationOrderBy, models.UserNotificationByrByString)

	var count int
	repo := repository.NewUserNotificationRepository(uc.DB)
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

func (uc UserNotificationUC) FindByID(c context.Context, parameter models.UserNotificationParameter) (res models.UserNotification, err error) {
	repo := repository.NewUserNotificationRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc UserNotificationUC) Add(c context.Context, data *requests.UserNotificationRequest) (res models.UserNotification, err error) {

	repo := repository.NewUserNotificationRepository(uc.DB)
	now := time.Now().UTC()
	cretaed_date := now.Format(time.RFC3339)
	res = models.UserNotification{
		UserID:    &data.UserID,
		RowID:     &data.RowID,
		Type:      &data.Type,
		Text:      &data.Text,
		CreatedAt: &cretaed_date,
		CreatedBy: &data.CreatedBy,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc UserNotificationUC) UpdateStatus(c context.Context, id string, data *requests.UserNotificationRequest) (res models.UserNotification, err error) {
	repo := repository.NewUserNotificationRepository(uc.DB)
	now := time.Now().Local()
	strNow := now.Format(time.RFC3339)
	res = models.UserNotification{
		ID:        &id,
		UpdatedBy: &data.UpdatedBy,
		UpdatedAt: &strNow,
	}

	res.ID, err = repo.UpdateStatus(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc UserNotificationUC) UpdateAllStatus(c context.Context, user_id string, data *requests.UserNotificationRequest) (res models.UserNotification, err error) {
	repo := repository.NewUserNotificationRepository(uc.DB)
	now := time.Now().Local()
	strNow := now.Format(time.RFC3339)
	res = models.UserNotification{
		UserID:    &user_id,
		UpdatedBy: &data.UpdatedBy,
		UpdatedAt: &strNow,
	}

	res.UserID, err = repo.UpdateAllStatus(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc UserNotificationUC) DeleteStatus(c context.Context, id string, data *requests.UserNotificationRequest) (res models.UserNotification, err error) {
	repo := repository.NewUserNotificationRepository(uc.DB)
	now := time.Now().Local()
	strNow := now.Format(time.RFC3339)
	res = models.UserNotification{
		ID:        &id,
		UpdatedBy: &data.UpdatedBy,
		UpdatedAt: &strNow,
	}

	res.ID, err = repo.DeleteStatus(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc UserNotificationUC) DeleteAllStatus(c context.Context, user_id string, data *requests.UserNotificationRequest) (res models.UserNotification, err error) {
	repo := repository.NewUserNotificationRepository(uc.DB)
	now := time.Now().Local()
	strNow := now.Format(time.RFC3339)
	res = models.UserNotification{
		UserID:    &user_id,
		UpdatedBy: &data.UpdatedBy,
		UpdatedAt: &strNow,
	}

	res.UserID, err = repo.DeleteAllStatus(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
