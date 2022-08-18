package usecase

import (
	"context"
	"errors"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/str"
)

var (
	activationURL = "/email-activation"
)

// MailUC ...
type MailUC struct {
	*ContractUC
}

// Send ...
func (uc MailUC) Send(c context.Context, to, subject, body string) (err error) {
	ctx := "MailUC.Send"

	if uc.EnvConfig["SMTP_PROVIDER"] == "mandrill" {
		err = uc.ContractUC.Mailing.Send(to, subject, body)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "mandrill", c.Value("requestid"))
			return err
		}
	} else {
		err = uc.ContractUC.Mail.Send(uc.EnvConfig["SMTP_FROM"], []string{to}, subject, body)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "smtp", c.Value("requestid"))
			return err
		}
	}

	return err
}

func (uc MailUC) GenerateUserVerifyMailKeyURL(c context.Context, user *models.AccountOpening) (verifyEmailURL string, err error) {
	// Generate Key
	rand := str.RandAlphanumericString(10)
	keys, err := uc.ContractUC.Aes.Encrypt(rand)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "aes_key", c.Value("requestid"))
		return verifyEmailURL, errors.New(helper.InternalServer)
	}

	err = uc.StoreToRedisExp("verifyMail"+keys, user.ID, "24h")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "store_redis", c.Value("requestid"))
		return verifyEmailURL, errors.New(helper.InternalServer)
	}

	err = uc.StoreToRedisExp("verifyMailKey"+user.ID, keys, "24h")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "store_key_redis", c.Value("requestid"))
		return verifyEmailURL, errors.New(helper.InternalServer)
	}

	// Activation path using key
	verifyEmailURL = uc.ContractUC.EnvConfig["APP_ENDPOINT_URL"] + activationURL + "?key=" + keys + "&email=" + user.Email
	return verifyEmailURL, err
}

func (uc MailUC) UserVerifyMail(c context.Context, id string) (err error) {
	accountOpeningUc := AccountOpeningUC{ContractUC: uc.ContractUC}
	user, err := accountOpeningUc.FindByID(c, models.AccountOpeningParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_user", c.Value("requestid"))
		return err
	}

	keyURL, err := uc.GenerateUserVerifyMailKeyURL(c, &user)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_key_url", c.Value("requestid"))
		return errors.New(helper.InternalServer)
	}
	html := `<html>
		<body>
			<p>Hi [name], Verify your email by click this link <a href="[link-verify-mail]">link</a></p>
		</body>
	</html>`

	// Replace value
	html = strings.Replace(html, "[name]", user.Name, 1)
	html = strings.ReplaceAll(html, "[link-verify-mail]", keyURL)

	err = uc.ContractUC.Mandrill.Send(user.Email, user.Name, "Verify Email Request", html, "")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_mail", c.Value("requestid"))
		return errors.New(helper.SendMail)
	}

	return err
}
