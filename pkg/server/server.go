package server

import (
	"irg1008/pals/pkg/client"
	"irg1008/pals/pkg/config"
	"irg1008/pals/pkg/db"
	"irg1008/pals/pkg/mailer"
	"irg1008/pals/pkg/roles"
	"irg1008/pals/pkg/tls"
	"irg1008/pals/pkg/tokens"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setUpTLS(e *echo.Echo) {
	tls.CreateNewCertificate(e, e.Server.Addr)
	e.Use(middleware.HTTPSRedirect())
}

type Server struct {
	Config  *config.Config
	DB      *db.DB
	Signing *tokens.Signing
	Roles   *roles.Roles
	Mailer  *mailer.Mailer
	Client  *client.Client
}

func (s *Server) setMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{s.Config.ClientUrl},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	// TODO: Implement CSRF
	// e.Use(middleware.CSRF())
}

func startDevServer(e *echo.Echo, port string) (err error) {
	return e.Start(port)
}

func startProdServer(e *echo.Echo, port string) (err error) {
	setUpTLS(e)
	return e.StartAutoTLS(port)
}

func (s *Server) Start(e *echo.Echo) (err error) {
	defer s.DB.Close()
	s.setMiddlewares(e)

	port := s.Config.Port
	if s.Config.IsDev {
		return startDevServer(e, port)
	}
	return startProdServer(e, port)
}

func NewServer() *Server {
	config := config.NewConfig()
	signing := tokens.NewSigning(config.JWTSecret)
	return &Server{
		Config:  config,
		Client:  client.NewClient(config.ClientUrl),
		DB:      db.NewDB(config.DBUrl),
		Signing: signing,
		Roles:   roles.NewRoles(signing),
		Mailer:  mailer.NewMailer(config.Domain, config.ResendKey),
	}
}
