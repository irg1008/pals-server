package internal

import (
	"irg1008/pals/internal/controllers/auth"
	"irg1008/pals/pkg/server"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, s *server.Server) {
	routes := e.Group("/auth")
	routes.Any("/*", s.Auth.Handler)
	auth.Routes(routes, s)
}
