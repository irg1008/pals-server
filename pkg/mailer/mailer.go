package mailer

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type HostInfo struct {
	Username string
	Address  string
	Port     int
	Password string
}

type Mailer struct {
	smtpServer smtpServer
	auth       smtp.Auth
	domain     string
}

func NewMailer(host *HostInfo, domain string) *Mailer {
	smtpServer := smtpServer{host.Address, host.Port}
	auth := smtp.PlainAuth("", host.Username, host.Password, host.Address)
	return &Mailer{smtpServer, auth, domain}
}

func (m *Mailer) Send(email *email.Email) error {
	return email.Send(m.smtpServer.Host(), m.auth)
}

func (m *Mailer) NewSender(name string, subdomain string) *Sender {
	return &Sender{&SenderData{name, subdomain}, m}
}
