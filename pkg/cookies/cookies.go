package cookies

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	cookieName       = "refresh_token"
	refreshTokenPath = "/api/auth/refresh"
	loggedCookieName = "is_logged"
	loggedCookiePath = "/"
)

func newCookie(name string, path string, age time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Path:     path,
		MaxAge:   int(age.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func deleteCookie(c echo.Context, cookie *http.Cookie) {
	cookie.MaxAge = -1
	c.SetCookie(cookie)
}
