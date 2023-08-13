package server

import (
	"irg1008/next-go/pkg/config"
	"irg1008/next-go/pkg/db"
	"irg1008/next-go/pkg/roles"
	"irg1008/next-go/pkg/tls"
	"irg1008/next-go/pkg/tokens"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setUpTLS(e *echo.Echo) {
	tls.CreateNewCertificate(e, e.Server.Addr)
	e.Use(middleware.HTTPSRedirect())
}

type Server struct {
	Config   *config.Config
	DB       *db.DB
	Signing  *tokens.Signing
	IsLogged echo.MiddlewareFunc
}

func (s *Server) Start(e *echo.Echo) (err error) {
	defer s.DB.Close()

	port := s.Config.Port

	if s.Config.IsDev {
		return e.Start(port)
	}

	setUpTLS(e)
	return e.StartAutoTLS(port)
}

func NewServer() *Server {
	config := config.NewConfig()
	signing := tokens.NewSigning(config.JWTSecret)
	db := db.NewDB(config.DBUrl)
	isLogged := roles.IsLoggedMiddleware(config.JWTSecret)
	return &Server{config, db, signing, isLogged}
}
