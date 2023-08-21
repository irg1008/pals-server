package roles

import (
	"irg1008/pals/pkg/tokens"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const contextKey = "user"

type Roles struct {
	signing      *tokens.Signing
	HasAuthToken echo.MiddlewareFunc
}

func NewRoles(signing *tokens.Signing) *Roles {
	return &Roles{
		signing:      signing,
		HasAuthToken: withAuthToken(signing.Secret),
	}
}

func withAuthToken(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(secret),
		ContextKey:    contextKey,
		SigningMethod: tokens.SigningAlgorithm().Alg(),
	})
}

func (r *Roles) IsLogged(next echo.HandlerFunc) echo.HandlerFunc {
	isLogged := func(c echo.Context) error {
		token := c.Get(contextKey).(*jwt.Token)
		err := r.signing.IsValidAccessToken(token)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		return next(c)
	}

	return r.HasAuthToken(isLogged)
}
