package usecase

import (
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// SmsUC ...
type SmsUC struct {
	*ContractUC
}

// // SendSMS ...
func (uc SmsUC) SendSMS(sender, message, receiver string) (err error) {
	err = uc.TwilioClient.SendSMS(sender, receiver, message)
	if err != nil {
		fmt.Println("error send sms")
		// logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), helper.SendSms, uc.ReqID)
		// return err
	}

	return err
}

// // SendOtpSMS ...
func (uc SmsUC) SendOtpSMS(from, phone, otp, action string) (err error) {
	action = strings.ReplaceAll(action, "_", " ")
	message := "Your Saham Rakyat " + action + " OTP is " + otp
	err = uc.SendSMS(from, message, phone)
	if err != nil {
		fmt.Println("error send sms")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), helper.SendSms, uc.ReqID)
		return err
	}

	return err
}
