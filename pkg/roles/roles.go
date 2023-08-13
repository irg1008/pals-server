package roles

import (
	"irg1008/next-go/pkg/config"
	"irg1008/next-go/pkg/tokens"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func userExtractionPanic() {
	slog.Error("Trying to extract user from context, but no auth middleware was assigned to route.")
	panic("Provide a valid middleware before calling this function.")
}

func GetUser(c echo.Context) *tokens.SignedUser {
	token := c.Get("user")
	if token == nil {
		userExtractionPanic()
	}
	return tokens.UnsignUser(token.(*jwt.Token))
}

var IsLogged = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(config.Env.JWTSecret),
})
