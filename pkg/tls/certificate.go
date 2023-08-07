package tls

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func CreateNewCertificate(e *echo.Echo, domain string) {
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(domain)
	e.AutoTLSManager.Cache = autocert.DirCache(".cache")
}
