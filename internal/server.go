package internal

import (
	"irg1008/next-go/internal/config"
	"irg1008/next-go/pkg/tls"
	"irg1008/next-go/pkg/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func applySecurityMiddlewares(e *echo.Echo) {
	// TODO: Add configs
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
}

func setUpTLS(e *echo.Echo) {
	tls.CreateNewCertificate(e, e.Server.Addr)
	e.Use(middleware.HTTPSRedirect())
}

func startServer(e *echo.Echo) error {
	conf := config.GetConfig()

	if conf.IsDev {
		return e.Start(conf.Port)
	}

	setUpTLS(e)
	return e.StartAutoTLS(conf.Port)
}

func StartServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = validation.GetCustomValidator()

	applySecurityMiddlewares(e)

	e.Static("/", "public")
	APIRoute(e)

	e.Logger.Fatal(startServer(e))
}
