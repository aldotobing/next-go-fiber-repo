package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

var (
	// OrderMailExchange ...
	OrderMailExchange = "order_mail.exchange"
	// OrderMail ...
	OrderMail = "order_mail.incoming.queue"
	// OrderMailDeadLetter ...
	OrderMailDeadLetter = "order_mail.deadletter.queue"

	// CancelOrderMailExchange ...
	CancelOrderMailExchange = "cancel_order_mail.exchange"
	// CancelOrderMail ...
	CancelOrderMail = "cancel_order_mail.incoming.queue"
	// CancelOrderMailDeadLetter ...
	CancelOrderMailDeadLetter = "cancel_order_mail.deadletter.queue"

	// Otp ...
	Otp = "saham_rakyat.otp.incoming.queue"
	// OtpExchange ...
	OtpExchange = "saham_rakyat.otp.exchange"
	// OtpDeadLetter ...
	OtpDeadLetter = "saham_rakyat.otp.deadletter.queue"

	// VerifyMailExchange ...
	VerifyMailExchange = "verify_mail.exchange"
	// VerifyMail ...
	VerifyMail = "verify_mail.incoming.queue"
	// VerifyMailDeadLetter ...
	VerifyMailDeadLetter = "verify_mail.deadletter.queue"
)

// IQueue ...
type IQueue interface {
	PushQueue(data map[string]interface{}, types string) error
	PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error)
}

// queue ...
type queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewQueue ...
func NewQueue(conn *amqp.Connection, channel *amqp.Channel) IQueue {
	return &queue{
		Connection: conn,
		Channel:    channel,
	}
}

// PushQueue ...
func (m queue) PushQueue(data map[string]interface{}, types string) error {
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return err
}

// PushQueueReconnect ...
func (m queue) PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": deadLetterKey,
	}
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, args)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return m.Connection, m.Channel, err
}
