package mailer

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	hostUsername = "resend"
	hostAddress  = "smtp.resend.com"
	hostPort     = 587
)

type Mailer struct {
	smtpServer smtpServer
	auth       smtp.Auth
	domain     string
}

func NewMailer(domain string, pwd string) *Mailer {
	smtpServer := smtpServer{hostAddress, hostPort}
	auth := smtp.PlainAuth("", hostUsername, pwd, hostAddress)
	return &Mailer{smtpServer, auth, domain}
}

func (m *Mailer) Send(email *email.Email) error {
	return email.Send(m.smtpServer.Host(), m.auth)
}

func (m *Mailer) NewSender(name string, subdomain string) *Sender {
	return &Sender{&SenderData{name, subdomain}, m}
}
