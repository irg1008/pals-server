package controllers

import (
	"irg1008/next-go/pkg/roles"
	"irg1008/next-go/pkg/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProtectedRoutes(e *echo.Group, s *server.Server) {
	g := e.Group("/protected", s.IsLogged)
	g.GET("", handleExmaple)
}

func handleExmaple(c echo.Context) error {
	user := roles.GetUser(c)
	return c.String(http.StatusOK, "Hello, "+user.Email+"!")
}
