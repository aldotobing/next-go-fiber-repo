package usecase

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/gomodul/envy"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
)

// BroadcastUC ...
type BroadcastUC struct {
	*ContractUC
}

func (uc BroadcastUC) greetingTime(message string) (res string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	envTimeMorning, _ := strconv.Atoi(envy.Get("TIME_MORNING"))
	envTimeAfternoon, _ := strconv.Atoi(envy.Get("TIME_AFTERNOON"))
	envTimeEvening, _ := strconv.Atoi(envy.Get("TIME_EVENING"))
	envTimeNight, _ := strconv.Atoi(envy.Get("TIME_NIGHT"))

	timeMorning := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeMorning, 0, 1, 0, loc)
	timeAfternoon := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeAfternoon, 0, 1, 0, loc)
	timeEvening := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeEvening, 0, 1, 0, loc)
	timeNight := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), envTimeNight, 0, 1, 0, loc)

	if now.After(timeMorning) && now.Before(timeAfternoon) {
		res = strings.Replace(message, "{GREETING_TIME}", "Pagi", 5)
	} else if now.After(timeAfternoon) && now.Before(timeEvening) {
		res = strings.Replace(message, "{GREETING_TIME}", "Siang", 5)
	} else if now.After(timeEvening) && now.Before(timeNight) {
		res = strings.Replace(message, "{GREETING_TIME}", "Sore", 5)
	} else if now.After(timeNight) || now.Before(timeMorning) {
		res = strings.Replace(message, "{GREETING_TIME}", "Malam", 5)
	} else {
		res = message
	}

	return res
}

// FindAll ...
func (uc BroadcastUC) Broadcast(c context.Context, input *requests.BroadcastRequest) (err error) {
	data, err := CustomerUC{ContractUC: uc.ContractUC}.SelectAll(c, models.CustomerParameter{
		By:   "c.created_date",
		Sort: "desc",
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	fcmUC := FCMUC{ContractUC: uc.ContractUC}
	for i := range data {
		if data[i].CustomerFCMToken != nil && *data[i].CustomerFCMToken != "" {
			var body string
			body = uc.greetingTime(input.Body)
			body = strings.ReplaceAll(body, "{NAMA_TOKO}", *data[i].CustomerName)

			_, err = fcmUC.SendFCMMessage(c, input.Title, body, *data[i].CustomerFCMToken)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_message", c.Value("requestid"))
				continue
			}
		}
	}

	return err
}
