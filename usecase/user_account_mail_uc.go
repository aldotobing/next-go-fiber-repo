package usecase

import (
	"context"
	"errors"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

type UserAccountMailUC struct {
	*ContractUC
}

func (uc UserAccountMailUC) SendUserOtp(c context.Context, id string, Otp string) (err error) {

	userAccountUc := UserAccountUC{ContractUC: uc.ContractUC}

	user, err := userAccountUc.FindByID(c, models.UserAccountParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_user", c.Value("requestid"))
		return err
	}
	html := `<html>
		<body>
			<p>Hi [name], Your Otp Is [otp]</p>
		</body>
	</html>`

	// Replace value
	html = strings.Replace(html, "[name]", *user.Name, 1)
	html = strings.ReplaceAll(html, "[otp]", Otp)
	EmailUc := MailUC{ContractUC: uc.ContractUC}
	err = EmailUc.Send(c, *user.Name, "Verify Otp", html)
	// mandril
	// err = uc.ContractUC.Mandrill.Send(*user.Email, *user.UserName, "Verify Email Request", html, "")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_mail", c.Value("requestid"))
		return errors.New(helper.SendMail)
	}

	return err
}
