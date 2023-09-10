package server

import (
	"irg1008/pals/pkg/auth"
	"irg1008/pals/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware(origins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowCredentials: true,
	})
}

func (s *Server) SetMiddlewares(e *echo.Echo) {
	e.Use(log.LoggerMiddleware(s.Config.IsDev))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(CORSMiddleware([]string{s.Config.ClientUrl}))
}

func (s *Server) SetAuthProviders() {
	s.Auth.AddGoogleProvider(&auth.GoogleProviderOpts{
		ClientID:     s.Config.GoogleClientID,
		ClientSecret: s.Config.GoogleClientSecret,
	})
}
