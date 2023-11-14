package usecase

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// BroadcastUC ...
type BroadcastUC struct {
	*ContractUC
}

// BuildBody ...
func (uc BroadcastUC) BuildBody(data *models.Broadcast, res *viewmodel.BroadcastVM) {
	res.ID = data.ID
	res.Title = data.Title
	res.Body = data.Body
	res.BroadcastDate = data.BroadcastDate
	res.BroadcastTime = data.BroadcastTime
	res.RepeatEveryDay = data.RepeatEveryDay
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	var param viewmodel.BroadcastParameterVM
	json.Unmarshal([]byte(data.Parameter.String), &param)
	res.Parameter = param
}

func (uc BroadcastUC) greetingTime(message string) (res string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	envTimeMorning := 5
	envTimeAfternoon := 11
	envTimeEvening := 14
	envTimeNight := 18

	timeMorning := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeMorning, 0, 1, 0, loc)
	timeAfternoon := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeAfternoon, 0, 1, 0, loc)
	timeEvening := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeEvening, 0, 1, 0, loc)
	timeNight := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeNight, 0, 1, 0, loc)

	if now.After(timeMorning) && now.Before(timeAfternoon) {
		res = strings.ReplaceAll(message, "{GREETING_TIME}", "Pagi")
	} else if now.After(timeAfternoon) && now.Before(timeEvening) {
		res = strings.ReplaceAll(message, "{GREETING_TIME}", "Siang")
	} else if now.After(timeEvening) && now.Before(timeNight) {
		res = strings.ReplaceAll(message, "{GREETING_TIME}", "Sore")
	} else if now.After(timeNight) || now.Before(timeMorning) {
		res = strings.ReplaceAll(message, "{GREETING_TIME}", "Malam")
	} else {
		res = message
	}

	return res
}

// FindAll ...
func (uc BroadcastUC) FindAll(c context.Context, parameter models.BroadcastParameter) (out []viewmodel.BroadcastVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.BroadcastOrderBy, models.BroadcastOrderByrByString)

	repo := repository.NewBroadcastRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.BroadcastVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.BroadcastVM, 0)
	}

	return
}

// SelectAll ...
func (uc BroadcastUC) SelectAll(c context.Context, parameter models.BroadcastParameter) (out []viewmodel.BroadcastVM, err error) {
	repo := repository.NewBroadcastRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.BroadcastVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.BroadcastVM, 0)
	}

	return
}

// FindByID ...
func (uc BroadcastUC) FindByID(c context.Context, parameter models.BroadcastParameter) (out viewmodel.BroadcastVM, err error) {
	repo := repository.NewBroadcastRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Broadcast ...
func (uc BroadcastUC) Broadcast(c context.Context, input *requests.BroadcastRequest) (err error) {
	var customerCodes string
	if input.CustomerCode != nil {
		for _, datum := range input.CustomerCode {
			if customerCodes == "" {
				customerCodes += `'` + datum + `'`
			} else {
				customerCodes += `,'` + datum + `'`
			}
		}
	}

	data, err := CustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.CustomerParameter{
		By:             "c.created_date",
		Sort:           "desc",
		FlagToken:      true,
		CustomerTypeId: input.CustomerTypeID,
		CustomerCodes:  customerCodes,
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var requestNotification []requests.UserNotificationRequest

	fcmUC := FCMUC{ContractUC: uc.ContractUC}
	for i := range data {
		if data[i].CustomerFCMToken != nil && *data[i].CustomerFCMToken != "" {
			var body string
			body = input.Body
			body = uc.greetingTime(body)
			body = strings.ReplaceAll(body, "{NAMA_TOKO}", *data[i].CustomerName)
			_, err = fcmUC.SendFCMMessage(c, input.Title, body, *data[i].CustomerFCMToken)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_message", c.Value("requestid"))
				continue
			}
			requestNotification = append(requestNotification, requests.UserNotificationRequest{
				UserID: *data[i].ID,
				RowID:  "0",
				Type:   "4",
				Title:  input.Title,
				Text:   body,
			})
		}
	}

	_ = UserNotificationUC{ContractUC: uc.ContractUC}.AddBulk(c, requestNotification)

	return err
}

// BroadcastWithID ...
func (uc BroadcastUC) BroadcastWithID(c context.Context, id string) (err error) {
	broadcastRepo := repository.BroadcastRepository{uc.DB}
	broadcastData, err := broadcastRepo.FindByID(c, models.BroadcastParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var param viewmodel.BroadcastParameterVM
	json.Unmarshal([]byte(broadcastData.Parameter.String), &param)

	data, err := CustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.CustomerParameter{
		By:              "c.created_date",
		Sort:            "desc",
		FlagToken:       true,
		CustomerTypeId:  param.CustomerTypeID,
		BranchID:        param.BranchID,
		RegionID:        param.RegionID,
		RegionGroupID:   param.RegionGroupID,
		CustomerLevelId: param.CustomerLevelID,
		CustomerCodes:   param.CustomerCodes,
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var requestNotification []requests.UserNotificationRequest
	fcmUC := FCMUC{ContractUC: uc.ContractUC}
	for i := range data {
		if data[i].CustomerFCMToken != nil && *data[i].CustomerFCMToken != "" {
			var body string
			body = broadcastData.Body
			body = uc.greetingTime(body)
			body = strings.ReplaceAll(body, "{NAMA_TOKO}", *data[i].CustomerName)
			_, err = fcmUC.SendFCMMessage(c, broadcastData.Title, body, *data[i].CustomerFCMToken)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_message", c.Value("requestid"))
				continue
			}
			requestNotification = append(requestNotification, requests.UserNotificationRequest{
				UserID: *data[i].ID,
				RowID:  "0",
				Type:   "4",
				Title:  broadcastData.Title,
				Text:   body,
			})
		}
	}

	_ = UserNotificationUC{ContractUC: uc.ContractUC}.AddBulk(c, requestNotification)

	return err
}

// BroadcastWithScheduler ...
func (uc BroadcastUC) BroadcastWithScheduler(c context.Context) (err error) {
	broadcastRepo := repository.BroadcastRepository{uc.DB}
	now := time.Now()
	before := now.Add(-time.Hour * 1)
	broadcastData, err := broadcastRepo.SelectAll(c, models.BroadcastParameter{
		StartAt: before.Format("15:04:05"),
		EndAt:   now.Format("15:04:05"),
		Sort:    "asc",
		By:      "b.id",
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	var requestNotification []requests.UserNotificationRequest
	for _, broadcastDatum := range broadcastData {
		var param viewmodel.BroadcastParameterVM
		json.Unmarshal([]byte(broadcastDatum.Parameter.String), &param)

		data, err := CustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.CustomerParameter{
			By:              "c.created_date",
			Sort:            "desc",
			FlagToken:       true,
			CustomerTypeId:  param.CustomerTypeID,
			BranchID:        param.BranchID,
			RegionID:        param.RegionID,
			RegionGroupID:   param.RegionGroupID,
			CustomerLevelId: param.CustomerLevelID,
			CustomerCodes:   param.CustomerCodes,
		})
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return err
		}

		fcmUC := FCMUC{ContractUC: uc.ContractUC}
		for i := range data {
			if data[i].CustomerFCMToken != nil && *data[i].CustomerFCMToken != "" {
				var body string
				body = broadcastDatum.Body
				body = uc.greetingTime(body)
				body = strings.ReplaceAll(body, "{NAMA_TOKO}", *data[i].CustomerName)
				_, err = fcmUC.SendFCMMessage(c, broadcastDatum.Title, body, *data[i].CustomerFCMToken)
				if err != nil {
					logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_message", c.Value("requestid"))
					continue
				}
				requestNotification = append(requestNotification, requests.UserNotificationRequest{
					UserID: *data[i].ID,
					RowID:  "0",
					Type:   "4",
					Title:  broadcastDatum.Title,
					Text:   body,
				})
			}
		}
	}

	_ = UserNotificationUC{ContractUC: uc.ContractUC}.AddBulk(c, requestNotification)

	return err
}

// Add ...
func (uc BroadcastUC) Add(c context.Context, in requests.BroadcastRequest) (out viewmodel.BroadcastVM, err error) {
	var customerCodes string
	if in.CustomerCode != nil {
		for _, datum := range in.CustomerCode {
			if customerCodes == "" {
				customerCodes += `'` + datum + `'`
			} else {
				customerCodes += `,'` + datum + `'`
			}
		}
	}
	out = viewmodel.BroadcastVM{
		Title:          in.Title,
		Body:           in.Body,
		BroadcastDate:  in.BroadcastDate,
		BroadcastTime:  in.BroadcastTime,
		RepeatEveryDay: in.RepeatEveryDay,
		Parameter: viewmodel.BroadcastParameterVM{
			BranchID:          in.BranchID,
			BranchName:        in.BranchName,
			RegionID:          in.RegionID,
			RegionName:        in.RegionName,
			RegionGroupID:     in.RegionGroupID,
			RegionGroupName:   in.RegionGroupName,
			CustomerTypeID:    in.CustomerTypeID,
			CustomerTypeName:  in.CustomerTypeName,
			CustomerLevelID:   in.CustomerLevelID,
			CustomerLevelName: in.CustomerLevelName,
			CustomerCodes:     customerCodes,
		},
	}

	broadcastRepo := repository.BroadcastRepository{uc.DB}
	out.ID, err = broadcastRepo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Update ...
func (uc BroadcastUC) Update(c context.Context, id string, in requests.BroadcastRequest) (out viewmodel.BroadcastVM, err error) {
	out = viewmodel.BroadcastVM{
		ID:             id,
		Title:          in.Title,
		Body:           in.Body,
		BroadcastDate:  in.BroadcastDate,
		BroadcastTime:  in.BroadcastTime,
		RepeatEveryDay: in.RepeatEveryDay,
		Parameter: viewmodel.BroadcastParameterVM{
			BranchID:         in.BranchID,
			BranchName:       in.BranchName,
			RegionID:         in.RegionID,
			RegionName:       in.RegionName,
			RegionGroupID:    in.RegionGroupID,
			RegionGroupName:  in.RegionGroupName,
			CustomerTypeID:   in.CustomerTypeID,
			CustomerTypeName: in.CustomerTypeName,
		},
	}

	broadcastRepo := repository.BroadcastRepository{uc.DB}
	out.ID, err = broadcastRepo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc BroadcastUC) Delete(c context.Context, in string) (err error) {
	broadcastRepo := repository.BroadcastRepository{uc.DB}
	_, err = broadcastRepo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
