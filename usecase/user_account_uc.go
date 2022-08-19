package usecase

import (
	"context"
	"errors"
	"fmt"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

type UserAccountUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserAccountUC) BuildBody(res *models.UserAccount) {
}

func (uc UserAccountUC) GenerateReferalCode(c context.Context, userName string) (res models.UserAccountParameter, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res.ReferalCode = userName + str.RandNumericString(3)
	for {
		data, _ := repo.FindByRefferalCode(c, res)
		if data.ID == "" {
			break
		}
		res.ReferalCode = userName + str.RandAlphanumericString(3)
	}

	return res, err
}

func (uc UserAccountUC) GenerateToken(c context.Context, id string) (res viewmodel.JwtVM, err error) {

	payload := map[string]interface{}{
		"user_id": id,
	}
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(c, payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate_token")
		return res, err
	}
	return res, err
}

func (uc UserAccountUC) FindByRefferalCode(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByRefferalCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc UserAccountUC) FindByPhoneNo(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByPhoneNo(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc UserAccountUC) FindByEmail(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByEmail(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc UserAccountUC) FindByEmailAndPassword(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByEmailAndPassword(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc UserAccountUC) FindByID(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)
	return res, err
}

func (uc UserAccountUC) Login(c context.Context, data *requests.UserAccountLoginRequest) (res viewmodel.UserAccountVM, err error) {
	chkuser, _ := uc.FindByPhoneNo(c, models.UserAccountParameter{PhoneNo: data.PhoneNo, Code: data.Code})
	if chkuser.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "email", c.Value("requestid"))
		return res, errors.New(helper.InvalidEmail)
	}
	userOtpRequest := requests.UserOtpRequest{
		Type:  OtpTypeLogin,
		Phone: data.PhoneNo,
	}
	res.UserID = chkuser.ID
	otpUc := OtpUC{ContractUC: uc.ContractUC}
	res.Otp, err = otpUc.OtpRequest(c, res.UserID, &userOtpRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_request", uc.ContractUC.ReqID)
		return res, err
	}

	tokens, err := uc.GenerateToken(c, res.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "token_request", uc.ContractUC.ReqID)
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate
	res.UserID = chkuser.ID
	senDwaMessage := uc.ContractUC.WhatsApp.SendWA("081329998633", res.Otp)
	if senDwaMessage != nil {
		fmt.Println("sukses")
	}

	// res.RoleList = *chkuser.RoleList
	// res.RoGroupID = *chkuser.RoleGroupID
	// res.QrCode = *&chkuser.QrCode
	// EmailUc := UserAccountMailUC{ContractUC: uc.ContractUC}
	// // mail := EmailUc.SendUserOtp(c, res.UserID, res.Otp)
	// EmailUc := MailUC{ContractUC: uc.ContractUC}
	// strEmail := helper.MailOTPTemplate(res.Otp)
	// strsubjects := res.Otp + ` Is Your One-Time Password`
	// mail := EmailUc.Send(c, data.Email, strsubjects, strEmail)
	// if mail != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "mail_sender", uc.ContractUC.ReqID)
	// 	return res, mail
	// }

	// SendSmsUc := SmsUC{ContractUC: uc.ContractUC}
	// sendsms := SendSmsUc.SendOtpSMS("+15398003148", "+6281329998633", res.Otp, "sender")
	// if sendsms != nil {
	// 	// logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "sms_sender", uc.ContractUC.ReqID)
	// 	// return res, sendsms
	// }

	return res, nil
}

func (uc UserAccountUC) ResendOtp(c context.Context, id string, data *requests.UserOtpRequest) (res viewmodel.UserAccountVM, err error) {

	chkuser, _ := uc.FindByID(c, models.UserAccountParameter{ID: id})
	if chkuser.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.ReferralNotFound, functioncaller.PrintFuncName(), "referral_code", c.Value("requestid"))
		return res, errors.New(helper.ReferralNotFound)
	}

	res.UserID = id
	userOtpRequest := requests.UserOtpRequest{
		Type:  data.Type,
		Phone: data.Phone,
	}

	otpUc := OtpUC{ContractUC: uc.ContractUC}
	res.Otp, err = otpUc.OtpRequest(c, res.UserID, &userOtpRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_request", uc.ContractUC.ReqID)
		return res, err
	}

	tokens, err := uc.GenerateToken(c, res.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "token_request", uc.ContractUC.ReqID)
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate

	senDwaMessage := uc.ContractUC.WhatsApp.SendWA("081329998633", res.Otp)
	if senDwaMessage != nil {
		fmt.Println("sukses")
	}

	// res.RoleList = *chkuser.RoleList
	// res.RoGroupID = *chkuser.RoleGroupID
	// res.QrCode = *&chkuser.QrCode

	// // EmailUc := UserAccountMailUC{ContractUC: uc.ContractUC}
	// // mail := EmailUc.SendUserOtp(c, res.UserID, res.Otp)
	// EmailUc := MailUC{ContractUC: uc.ContractUC}
	// strEmail := helper.MailOTPTemplate(res.Otp)
	// strsubjects := res.Otp + ` Is Your One-Time Password`
	// mail := EmailUc.Send(c, chkuser.Email, strsubjects, strEmail)

	// if mail != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "mail_sender", uc.ContractUC.ReqID)
	// 	return res, mail
	// }
	return res, err
}

func (uc UserAccountUC) SubmitOtpUser(c context.Context, id string, data *requests.UserOtpSubmit) (res viewmodel.UserAccountVM, err error) {
	otpUc := OtpUC{ContractUC: uc.ContractUC}
	verifyOtp, err := otpUc.VerifyOtp(c, id, data.Type, data.Otp)
	if !verifyOtp {
		logruslogger.Log(logruslogger.WarnLevel, "", functioncaller.PrintFuncName(), "otp-not-valid")
		return res, err
	}
	Useracountuc := UserAccountUC{ContractUC: uc.ContractUC}
	user, err := Useracountuc.FindByID(c, models.UserAccountParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-find-user")
		return res, err
	}

	// if data.Type == OtpTypeRegister {
	// 	_, err = uc.SetActiveUser(c, user.ID)
	// 	if err != nil {
	// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "set-active-fail")
	// 		return res, err
	// 	}
	// }

	tokens, err := uc.GenerateToken(c, user.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate-token")
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate
	res.UserID = user.ID
	// res.RoleList = *user.RoleList
	// res.RoGroupID = *user.RoleGroupID
	// res.QrCode = *&user.QrCode
	return res, err
}
