package mailer

import (
	"irg1008/next-go/pkg/mailer/templates"

	"github.com/jordan-wright/email"
)

type SenderData struct {
	Name      string
	Subdomain string
}

type Sender struct {
	data   *SenderData
	Mailer *Mailer
}

func (s *Sender) newEmail(to string, subject string) *email.Email {
	e := email.NewEmail()
	e.To = []string{to}
	e.Subject = subject
	from := emailAddress{*s.data, s.Mailer.domain}
	e.From = from.Address()
	return e
}

func (s *Sender) SendHTML(to string, subject string, html string) error {
	e := s.newEmail(to, subject)
	e.HTML = []byte(html)
	return s.Mailer.Send(e)
}

type TemplateFunc func(string, string) (string, error)

func (s *Sender) sendTemplateEmail(to, subject, url string, getTemplate TemplateFunc) error {
	html, err := getTemplate(to, url)

	if err != nil {
		return err
	}

	return s.SendHTML(to, subject, html)
}

func (s *Sender) SendConfirmEmail(to, subject, url string) error {
	return s.sendTemplateEmail(to, subject, url, templates.GetConfirm)
}

func (s *Sender) SendResetPassword(to, subject, url string) error {
	return s.sendTemplateEmail(to, subject, url, templates.GetReset)
}
