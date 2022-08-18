package usecase

import (
	"context"
	"strconv"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// SchedullerUC ...
type SchedullerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SchedullerUC) BuildExpiredPackageBody(res *models.SchedullerExpiredPackage) {
}

// FindByID ...
func (uc SchedullerUC) ProccessExpiredPackage(c context.Context, parameter models.SchedullerExpiredPackageParameter) (res models.SchedullerExpiredPackage, err error) {
	repo := repository.NewSchedullerRepository(uc.DB)
	res, err = repo.ProcessExpiredPackage(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildExpiredPackageBody(&res)

	// mail := EmailUc.SendUserOtp(c, res.UserID, res.Otp)
	EmailUc := MailUC{ContractUC: uc.ContractUC}
	strEmail := strconv.Itoa(*res.TotalCount)
	strsubjects := ` Is Your One-Time Password`
	mail := EmailUc.Send(c, "moch.yuliadipurwanto@gmail.com", strsubjects, strEmail)
	if mail != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "mail_sender", uc.ContractUC.ReqID)
		return res, mail
	}

	return res, err
}
