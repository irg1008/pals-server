package internal

import (
	"irg1008/next-go/pkg/config"
	"irg1008/next-go/pkg/db"
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
	// e.Use(middleware.CSRF())
}

func setUpTLS(e *echo.Echo) {
	tls.CreateNewCertificate(e, e.Server.Addr)
	e.Use(middleware.HTTPSRedirect())
}

func startServer(e *echo.Echo) (err error) {
	port := config.Env.Port

	if config.Env.IsDev {
		return e.Start(port)
	}

	err = e.StartAutoTLS(port)
	if err != nil {
		return
	}

	setUpTLS(e)
	return
}

func StartServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = validation.NewCustomValidator()

	db := db.New()
	defer db.Close()

	applySecurityMiddlewares(e)

	e.Static("/", "public")
	APIRoute(e, db)

	e.Logger.Fatal(startServer(e))
}
