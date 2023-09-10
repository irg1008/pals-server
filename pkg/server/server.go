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
	db := db.NewDB(config.DBUrl)

	return &Server{
		Config: config,
		Client: client.NewClient(config.ClientUrl),
		DB:     db,
		Auth:   NewAuthService(config, db),
		Mailer: NewMailer(config),
	}
}

func NewConfiguredServer(e *echo.Echo) *Server {
	server := NewServer()
	server.SetMiddlewares(e)
	server.SetAuthProviders(server.Config)
	return server
}
