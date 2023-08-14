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

func SetRefreshTokenCookie(c echo.Context, refreshToken string) {
	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Value:    refreshToken,
		Path:     refreshTokenPath,
		MaxAge:   int(config.RefreshTokenDuration.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func GetRefreshTokenCookie(c echo.Context) (value string, err error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return
	}
	return cookie.Value, nil
}
