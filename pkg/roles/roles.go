package roles

import (
	"irg1008/next-go/pkg/tokens"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const contextKey = "user"

func IsLoggedMiddleware(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(secret),
		ContextKey:    contextKey,
		SigningMethod: tokens.SigningAlgorithm().Alg(),
	})
}
