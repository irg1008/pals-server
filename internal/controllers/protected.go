package controllers

import (
	"fmt"
	"irg1008/next-go/pkg/roles"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProtectedRoutes(e *echo.Group) {
	g := e.Group("/protected", roles.IsLogged)
	g.GET("", handleExmaple)
}

func handleExmaple(c echo.Context) error {
	user := roles.GetUser(c)
	fmt.Println(user)

	return c.String(http.StatusOK, "Hello, logged user!")
}
