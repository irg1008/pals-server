package mailer

import "fmt"

type smtpServer struct {
	host string
	port int
}

func (s *smtpServer) Host() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

type emailAddress struct {
	sender SenderData
	domain string
}

func (e *emailAddress) Address() string {
	return fmt.Sprintf("%s <%s@%s>", e.sender.Name, e.sender.Subdomain, e.domain)
}
