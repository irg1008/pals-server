package server

import (
	"irg1008/pals/pkg/tls"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startDevServer(e *echo.Echo, port string) (err error) {
	return e.Start(port)
}

func setUpTLS(e *echo.Echo) {
	tls.CreateNewCertificate(e, e.Server.Addr)
	e.Use(middleware.HTTPSRedirect())
}

func startProdServer(e *echo.Echo, port string) (err error) {
	setUpTLS(e)
	return e.StartAutoTLS(port)
}

func (s *Server) Start(e *echo.Echo) (err error) {
	defer s.DB.Close()

	port := s.Config.Port
	if s.Config.IsDev {
		return startDevServer(e, port)
	}
	return startProdServer(e, port)
}
