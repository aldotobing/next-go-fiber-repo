package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

type UserAccountUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserAccountUC) BuildBody(res *models.UserAccount) {
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

func (uc UserAccountUC) FindByEmailAndPass(c context.Context, parameter models.UserAccountParameter) (res models.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(uc.DB)
	res, err = repo.FindByEmailAndPass(c, parameter)
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
	parts := strings.Split(data.Code, "--*")
	CodeUser := ""
	if len(parts) >= 1 {
		CodeUser = parts[0]
	}
	chkuser, _ := uc.FindByPhoneNo(c, models.UserAccountParameter{PhoneNo: data.PhoneNo, Code: CodeUser})
	if chkuser.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "email", c.Value("requestid"))
		return res, errors.New(helper.InvalidEmail)
	}
	userOtpRequest := requests.UserOtpRequest{
		Type:  OtpTypeLogin,
		Phone: data.PhoneNo,
	}
	res.CustomerID = *chkuser.CustomerID
	otpUc := OtpUC{ContractUC: uc.ContractUC}
	if len(parts) > 1 {
		res.Otp, err = otpUc.OtpAnonumousRequest(c, res.CustomerID, &userOtpRequest)

	} else {
		res.Otp, err = otpUc.OtpRequest(c, res.CustomerID, &userOtpRequest)
	}

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_request", uc.ContractUC.ReqID)
		return res, err
	}

	tokens, err := uc.GenerateToken(c, res.CustomerID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "token_request", uc.ContractUC.ReqID)
		return res, err
	}

	repo := repository.NewUserAccountRepository(uc.DB)
	_, errfcm := repo.FCMUpdate(c, &models.UserAccount{ID: chkuser.ID, FCMToken: &data.FCMToken})
	if errfcm != nil {
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate
	res.ID = chkuser.ID
	res.Code = chkuser.Code
	res.CustomerID = *chkuser.CustomerID
	res.CustomerName = *chkuser.Name
	res.Phone = *chkuser.Phone
	res.PriceListID = chkuser.PriceListID
	res.PriceListVersionID = chkuser.PriceListVersionID
	res.CustomerTypeID = chkuser.CustomerTypeID
	res.CustomerLevelName = chkuser.CustomerLevelName
	res.CustomerAddress = chkuser.CustomerAddress
	res.SalesmanID = chkuser.SalesmanID
	res.SalesmanName = chkuser.SalesmanName
	res.SalesmanCode = chkuser.SalesmanCode

	if len(parts) == 1 {
		senDwaMessage := uc.ContractUC.WhatsApp.SendWA(res.Phone, res.Otp)
		if senDwaMessage != nil {
			fmt.Println("sukses")
		}
	}
	return res, nil
}

func (uc UserAccountUC) ResendOtp(c context.Context, id string, data *requests.UserOtpRequest) (res viewmodel.UserAccountVM, err error) {

	chkuser, _ := uc.FindByID(c, models.UserAccountParameter{CustomerID: id})
	if chkuser.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.ReferralNotFound, functioncaller.PrintFuncName(), "referral_code", c.Value("requestid"))
		return res, errors.New(helper.ReferralNotFound)
	}

	res.CustomerID = id
	userOtpRequest := requests.UserOtpRequest{
		Type:  data.Type,
		Phone: data.Phone,
	}

	otpUc := OtpUC{ContractUC: uc.ContractUC}
	res.Otp, err = otpUc.OtpRequest(c, res.CustomerID, &userOtpRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "otp_request", uc.ContractUC.ReqID)
		return res, err
	}

	tokens, err := uc.GenerateToken(c, res.CustomerID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "token_request", uc.ContractUC.ReqID)
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate

	res.ID = chkuser.ID
	res.Code = chkuser.Code
	res.CustomerID = *chkuser.CustomerID
	res.CustomerName = *chkuser.Name
	res.Phone = *chkuser.Phone
	res.PriceListID = chkuser.PriceListID
	res.PriceListVersionID = chkuser.PriceListVersionID
	res.CustomerTypeID = chkuser.CustomerTypeID
	res.CustomerLevelName = chkuser.CustomerLevelName
	res.SalesmanID = chkuser.SalesmanID
	res.SalesmanName = chkuser.SalesmanName
	res.SalesmanCode = chkuser.SalesmanCode
	senDwaMessage := uc.ContractUC.WhatsApp.SendWA(res.Phone, res.Otp)
	if senDwaMessage != nil {
		fmt.Println("sukses")
	}

	return res, err
}

func (uc UserAccountUC) SubmitOtpUser(c context.Context, id string, data *requests.UserOtpSubmit) (res viewmodel.UserAccountVM, err error) {
	otpUc := OtpUC{ContractUC: uc.ContractUC}
	fmt.Println(id, data.Type)
	verifyOtp, err := otpUc.VerifyOtp(c, id, data.Type, data.Otp)
	if !verifyOtp {
		logruslogger.Log(logruslogger.WarnLevel, "", functioncaller.PrintFuncName(), "otp-not-valid")
		return res, err
	}
	Useracountuc := UserAccountUC{ContractUC: uc.ContractUC}
	user, err := Useracountuc.FindByID(c, models.UserAccountParameter{CustomerID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-find-user")
		return res, err
	}

	tokens, err := uc.GenerateToken(c, user.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate-token")
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate
	res.ID = user.ID
	res.Code = user.Code
	res.CustomerID = *user.CustomerID
	res.CustomerName = *user.Name
	res.Phone = *user.Phone
	res.PriceListID = user.PriceListID
	res.PriceListVersionID = user.PriceListVersionID
	res.CustomerTypeID = user.CustomerTypeID
	res.CustomerLevelName = user.CustomerLevelName
	res.SalesmanID = user.SalesmanID
	res.SalesmanName = user.SalesmanName
	res.SalesmanCode = user.SalesmanCode
	return res, err
}

func (uc UserAccountUC) LoginBackEnd(c context.Context, data *requests.UserAccountBackendLoginRequest) (res viewmodel.UserAccountVM, err error) {
	chkuser, _ := uc.FindByEmailAndPass(c, models.UserAccountParameter{Password: data.Password, Email: data.Email})
	if chkuser.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, helper.NameAlreadyExist, functioncaller.PrintFuncName(), "email", c.Value("requestid"))
		return res, errors.New(helper.InvalidEmail)
	}

	res.CustomerID = *&chkuser.ID

	tokens, err := uc.GenerateToken(c, chkuser.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "token_request", uc.ContractUC.ReqID)
		return res, err
	}
	res.Token = tokens.Token
	res.ExpiredDate = tokens.ExpiredDate
	res.RefreshToken = tokens.RefreshToken
	res.RefreshExpiredDate = tokens.RefreshExpiredDate
	res.ID = chkuser.ID
	res.Code = chkuser.Code
	res.PriceListID = chkuser.PriceListID
	res.PriceListVersionID = chkuser.PriceListVersionID
	res.CustomerTypeID = chkuser.CustomerTypeID
	res.CustomerLevelName = chkuser.CustomerLevelName
	res.CustomerAddress = chkuser.CustomerAddress
	res.SalesmanID = chkuser.SalesmanID
	res.SalesmanName = chkuser.SalesmanName
	res.SalesmanCode = chkuser.SalesmanCode

	return res, nil
}
