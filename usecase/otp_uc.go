package usecase

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
)

var (
	// OtpTypeRegister ...
	OtpTypeRegister = "register"
	OtpTypeLogin    = "login"

	// OtpTypeWhiteList ...
	OtpTypeWhiteList = []string{
		OtpTypeRegister,
		OtpTypeLogin,
	}
)

// OtpUC ...
type OtpUC struct {
	*ContractUC
}

// OtpRequest ...
func (uc OtpUC) OtpRequest(c context.Context, id string, data *requests.UserOtpRequest) (res string, err error) {
	if !str.Contains(OtpTypeWhiteList, data.Type) {
		logruslogger.Log(logruslogger.WarnLevel, data.Type, functioncaller.PrintFuncName(), "check_otp_type", uc.ContractUC.ReqID)
		return res, errors.New(helper.InvalidTypeOtp)
	}

	// Generate OTP and save to redis
	rand.Seed(time.Now().UTC().UnixNano())
	res = str.RandomNumberString(4)

	// err = uc.ContractUC.StoreToRedisExp("otp"+id+data.Type, res, OtpLifetime)

	// Check OTP display setting
	if !str.StringToBool(uc.ContractUC.EnvConfig["APP_OTP_DISPLAY"]) {
		res = ""
	}

	return res, err
}

// VerifyOtp ...
func (uc OtpUC) VerifyOtp(c context.Context, id, types, otp string) (res bool, err error) {
	var otpRedis string
	err = uc.ContractUC.GetFromRedis("otp"+id+types, &otpRedis)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_otp_redis", uc.ContractUC.ReqID)
		return res, errors.New(helper.ExpKey)
	}

	if otp != otpRedis {
		// Check wrong otp counter
		err = uc.ContractUC.LimitRetryByKey("invalidOtp"+id, MaxOtpSubmitRetry)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_counter", uc.ContractUC.ReqID)
			return res, err
		}
		res = false
		return res, errors.New(helper.WrongOTP)
	}

	res = true

	return res, err
}
