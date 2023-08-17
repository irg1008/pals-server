package cookies

import (
	"irg1008/next-go/pkg/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	cookieName       = "refresh_token"
	refreshTokenPath = "/api/auth/refresh"
)

func refreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     refreshTokenPath,
		MaxAge:   int(config.RefreshTokenDuration.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func SetRefreshTokenCookie(c echo.Context, refreshToken string) {
	cookie := refreshCookie()
	cookie.Value = refreshToken
	c.SetCookie(cookie)
}

func GetRefreshTokenCookie(c echo.Context) (value string, err error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return
	}
	return cookie.Value, nil
}

func DeleteRefreshTokenCookie(c echo.Context) {
	cookie := refreshCookie()
	cookie.MaxAge = -1
	c.SetCookie(cookie)
}
