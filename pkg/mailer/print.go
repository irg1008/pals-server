package mailer

import "fmt"

type smtpServer struct {
	host string
	port int
}

func (s *smtpServer) Host() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

type EmailAddress struct {
	Sender SenderData
	Domain string
}

func (e *EmailAddress) Address() string {
	return fmt.Sprintf("%s <%s@%s>", e.Sender.Name, e.Sender.Subdomain, e.Domain)
}
