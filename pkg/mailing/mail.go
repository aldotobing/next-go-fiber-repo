package mailing

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Credential struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (cred Credential) SendMail(to []string, subject, message string) error {
	body := "From : " + cred.Username + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", cred.Username, cred.Password, cred.Host)
	smtpAddress := fmt.Sprintf("%s:%s", cred.Host, cred.Port)
	err := smtp.SendMail(smtpAddress, auth, cred.Username, append(to), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
