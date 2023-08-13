package internal

import (
	"irg1008/next-go/internal/server"
	"irg1008/next-go/pkg/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func applySecurityMiddlewares(e *echo.Echo) {
	// TODO: Add configs
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	// e.Use(middleware.CSRF())
}

func StartServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = validation.NewCustomValidator()
	applySecurityMiddlewares(e)

	// Serve free static files over /public
	e.Static("/", "public")

	// Create all context objects for API
	server := server.NewServer()
	APIRoutes(e, server)

	err := server.Start(e)
	e.Logger.Fatal(err)
}
