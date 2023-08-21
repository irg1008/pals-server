package cookies

import (
	"irg1008/pals/pkg/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	cookieName       = "refresh_token"
	refreshTokenPath = "/api/auth/refresh"
	loggedCookieName = "is_logged"
	loggedCookiePath = "/"
)

func newCookie(name string, path string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Path:     path,
		MaxAge:   int(config.RefreshTokenDuration.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func deleteCookie(c echo.Context, cookie *http.Cookie) {
	cookie.MaxAge = -1
	c.SetCookie(cookie)
}

func loggedCookie() *http.Cookie {
	return newCookie(loggedCookieName, loggedCookiePath)
}

func refreshCookie() *http.Cookie {
	return newCookie(cookieName, refreshTokenPath)
}

func SetRefreshTokenCookie(c echo.Context, refreshToken string) {
	cookie := refreshCookie()
	cookie.Value = refreshToken
	c.SetCookie(cookie)

	loggedCookie := loggedCookie()
	loggedCookie.Value = "true"
	c.SetCookie(loggedCookie)
}

func DeleteRefreshTokenCookie(c echo.Context) {
	cookie := refreshCookie()
	deleteCookie(c, cookie)
	loggedCookie := loggedCookie()
	deleteCookie(c, loggedCookie)
}

func GetRefreshTokenCookie(c echo.Context) (value string, err error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return
	}
	return cookie.Value, nil
}
