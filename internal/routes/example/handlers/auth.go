package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func basicAuthCheck(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(basicAuthCheck)
}
