package server

import (
	"irg1008/pals/pkg/auth"
	"irg1008/pals/pkg/client"
	"irg1008/pals/pkg/config"
	"irg1008/pals/pkg/db"
	"irg1008/pals/pkg/mailer"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Config *config.Config
	DB     *db.DB
	Mailer *mailer.Mailer
	Client *client.Client
	Auth   *auth.AuthService
}

func NewServer() *Server {
	config := config.NewConfig()

	auth := auth.NewAuthService(&auth.BaseOptions{
		AppName:   config.AppName,
		JWTSecret: config.JWTSecret,
		URL:       config.APIUrl,
	})

	mailHostInfo := &mailer.HostInfo{
		Username: config.EmailUser,
		Address:  config.EmailHost,
		Port:     config.EmailPort,
		Password: config.EmailPass,
	}

	return &Server{
		Config: config,
		Client: client.NewClient(config.ClientUrl),
		DB:     db.NewDB(config.DBUrl),
		Mailer: mailer.NewMailer(mailHostInfo, config.Domain),
		Auth:   auth,
	}
}

func NewConfiguredServer(e *echo.Echo) *Server {
	server := NewServer()
	server.SetMiddlewares(e)
	server.SetAuthProviders()
	return server
}
