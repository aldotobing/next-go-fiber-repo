package usecase

import (
	"context"
	"errors"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/interfacepkg"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// RecaptchaUC ...
type RecaptchaUC struct {
	*ContractUC
}

// Verify ...
func (uc RecaptchaUC) Verify(c context.Context, response, remoteip string) (res map[string]interface{}, err error) {
	res, err = uc.Recaptcha.Verify(response, remoteip)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "verify", c.Value("requestid"))
		return res, errors.New(helper.Recaptcha)
	}
	if res["success"] == nil {
		logruslogger.Log(logruslogger.WarnLevel, helper.Recaptcha, functioncaller.PrintFuncName(), "null_success", c.Value("requestid"))
		return res, errors.New(helper.Recaptcha)
	}
	if !res["success"].(bool) {
		logruslogger.Log(logruslogger.WarnLevel, interfacepkg.InterfaceArrayToString(res["error-codes"].([]interface{})), functioncaller.PrintFuncName(), "failed", c.Value("requestid"))
		return res, errors.New(helper.Recaptcha)
	}

	return res, err
}
