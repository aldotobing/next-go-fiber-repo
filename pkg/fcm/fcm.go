package fcm

import (
	"nextbasis-service-v-0.1/pkg/interfacepkg"

	"github.com/maddevsio/fcm"
)

// Connection ...
type Connection struct {
	APIKey string
}

// SendAndroid ...
func (cred *Connection) SendAndroid(to []string, title, body string, data map[string]interface{}) (string, error) {
	c := fcm.NewFCM(cred.APIKey)
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  to,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: title,
			Body:  body,
			Sound: "default",
			Badge: "3",
		},
	})
	if err != nil {
		return "", err
	}

	res := interfacepkg.Marshal(response)

	return res, err
}
