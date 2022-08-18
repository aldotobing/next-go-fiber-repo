package twilio

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

type Client struct {
	sid           string
	token         string
	DefaultSender string
	twilio        *gotwilio.Twilio
}

func NewTwilioClient(sid, token, defaultSender string) *Client {
	return &Client{
		sid:           sid,
		token:         token,
		DefaultSender: defaultSender,
		twilio:        gotwilio.NewTwilioClient(sid, token),
	}
}

func (client Client) SendSMS(from, to, message string) (err error) {
	_, _, _ = client.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		fmt.Println("error send sms")
		return err
	}
	return nil
}
