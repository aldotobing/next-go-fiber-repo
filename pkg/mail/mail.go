package mail

import (
	"errors"
	"net/smtp"
	"strconv"
	"strings"
)

// Connection ...
type Connection struct {
	Host     string
	Port     int
	Username string
	Password string
}

type plainAuthOverTLSConn struct {
	smtp.Auth
}

// PlainAuthOverTLSConn ...
func PlainAuthOverTLSConn(identity, username, password, host string) smtp.Auth {
	return &plainAuthOverTLSConn{smtp.PlainAuth(identity, username, password, host)}
}

func (a *plainAuthOverTLSConn) Start(server *smtp.ServerInfo) (string, []byte, error) {
	server.TLS = true
	return a.Auth.Start(server)
}

func (a *plainAuthOverTLSConn) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}

// Send ...
func (conn *Connection) Send(from string, to []string, subject, body string) error {
	emailBody := "To: " + strings.Join(to, ", ") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" + body + "\r\n"

	// auth := PlainAuthOverTLSConn("", conn.Username, conn.Password, conn.Host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		conn.Host+":"+strconv.Itoa(conn.Port),
		nil,
		conn.Username,
		to,
		[]byte(emailBody),
	)
	if err != nil {
		return err
	}

	return nil
}
