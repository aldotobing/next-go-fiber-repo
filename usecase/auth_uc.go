package usecase

import (
	"context"
	"errors"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

var (
	// RequestOtp ...
	RequestOtp = "request_otp"
)

// AuthUC ...
type AuthUC struct {
	*ContractUC
}

// GenerateCode randomize code & check uniqueness from DB
func (uc AuthUC) GenerateCode(c context.Context) (res models.UserParameter, err error) {
	repo := repository.NewUserRepository(uc.DB)
	res.Code = str.RandAlphanumericString(8)
	for {
		data, _ := repo.FindByCode(c, res)
		if data.ID == "" {
			break
		}
		res.Code = str.RandAlphanumericString(8)
	}

	return res, err
}

// GenerateToken ...
func (uc AuthUC) GenerateToken(c context.Context, id string) (res viewmodel.JwtVM, err error) {
	payload := map[string]interface{}{
		"id": id,
	}
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(c, payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate_token")
		return res, err
	}

	return res, err
}

// Register ...
func (uc AuthUC) Register(c context.Context, data *requests.RegisterRequest) (res viewmodel.JwtVM, err error) {
	userRequest := requests.AccountOpeningRequest{
		UserID: data.UserID,
		Name:   data.Name,
		Email:  data.Email,
		Phone:  data.Phone,
		Status: models.AccountOpeningStatusPending,
	}
	accountOpeningUC := AccountOpeningUC{ContractUC: uc.ContractUC}
	user, err := accountOpeningUC.Add(c, &userRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query")
		return res, err
	}

	userOtpRequest := requests.UserOtpRequest{
		Type: OtpTypeRegister,
		// CountryCode: data.CountryCode,
		// Phone:       data.Phone,
	}

	otpUc := OtpUC{ContractUC: uc.ContractUC}
	res.Otp, err = otpUc.OtpRequest(c, user.ID, &userOtpRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_request", uc.ContractUC.ReqID)
		return res, err
	}

	err = uc.ContractUC.StoreToRedisExp("latestAction"+user.ID+userOtpRequest.Type, RequestOtp, "1h")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "store_to_redis", uc.ContractUC.ReqID)
		return res, errors.New(helper.InternalServer)
	}
	res.LatestAction = RequestOtp

	return res, err
}

// SubmitOtpRegister ...
func (uc AuthUC) SubmitOtpRegister(c context.Context, id string, data *requests.UserOtpSubmit) (res viewmodel.JwtVM, err error) {
	otpUc := OtpUC{ContractUC: uc.ContractUC}
	verifyOtp, err := otpUc.VerifyOtp(c, id, data.Type, data.Otp)
	if !verifyOtp {
		logruslogger.Log(logruslogger.WarnLevel, "", functioncaller.PrintFuncName(), "otp-not-valid")
		return res, err
	}

	accountOpeningUc := AccountOpeningUC{ContractUC: uc.ContractUC}
	user, err := accountOpeningUc.FindByID(c, models.AccountOpeningParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-find-user")
		return res, err
	}
	_, err = accountOpeningUc.EditPhoneValidAt(c, user.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "phone-not-valid")
		return res, err
	}

	res, err = uc.GenerateToken(c, user.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate-token")
		return res, err
	}

	return res, err
}

func (uc AuthUC) ReqVerifyMail(c context.Context, data *requests.VerifyMailRequest) (res interface{}, err error) {
	accountOpeningUc := AccountOpeningUC{ContractUC: uc.ContractUC}
	user, err := accountOpeningUc.FindByEmail(c, models.AccountOpeningParameter{Email: data.Email})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_user", c.Value("requestid"))
		return res, errors.New(helper.InvalidEmail)
	}

	if data.Email == user.Email {
		// Push data to mqueue
		// mqueue := amqp.NewQueue(AmqpConnection, AmqpChannel)
		// queueBody := map[string]interface{}{
		// 	"qid":  c.Value("requestid"),
		// 	"id":   user.ID,
		// 	"type": "user",
		// }
		// AmqpConnection, AmqpChannel, err = mqueue.PushQueueReconnect(uc.ContractUC.EnvConfig["AMQP_URL"], queueBody, amqp.VerifyMail, amqp.VerifyMailDeadLetter)
		// if err != nil {
		// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "update_cash_request_amqp", c.Value("requestid"))
		// 	return res, errors.New(helper.InternalServer)
		// }
	} else {
		logruslogger.Log(logruslogger.WarnLevel, data.Email, functioncaller.PrintFuncName(), "invalid", c.Value("requestid"))
		return res, errors.New(helper.InvalidParameter)
	}

	return res, err
}

func (uc AuthUC) VerifyMail(c context.Context, key string) (res models.AccountOpening, err error) {
	accountOpeningUc := AccountOpeningUC{ContractUC: uc.ContractUC}

	id := ""
	err = uc.GetFromRedis("verifyMail"+key, &id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_user_id_from_redis", c.Value("requestid"))
		return res, err
	}

	user, err := accountOpeningUc.FindByID(c, models.AccountOpeningParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_user_by_id", c.Value("requestid"))
		return res, err
	}
	keyStored := ""
	err = uc.GetFromRedis("verifyMailKey"+user.ID, &keyStored)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "store_redis", c.Value("requestid"))
		return res, err
	}

	if key != keyStored {
		logruslogger.Log(logruslogger.WarnLevel, helper.InvalidKey, functioncaller.PrintFuncName(), "store_redis", c.Value("requestid"))
		return res, errors.New(helper.InvalidKey)
	}

	res, err = accountOpeningUc.EditEmailValidAt(c, user.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "activated_user", c.Value("requestid"))
		return res, err
	}
	uc.RemoveFromRedis("verifyMail" + key)
	uc.RemoveFromRedis("verifyMailKey" + user.ID)

	return res, err
}
