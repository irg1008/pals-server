package internal

import (
	"irg1008/next-go/pkg/server"
	"irg1008/next-go/pkg/validation"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = validation.NewCustomValidator()

	server := server.NewServer()
	APIRoutes(e, server)

	e.Static("/", "public")

	err := server.Start(e)
	e.Logger.Fatal(err)
}