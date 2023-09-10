package auth

import (
	"irg1008/pals/ent/userdata"

	"github.com/go-pkgz/auth/middleware"
	"github.com/labstack/echo/v4"
)

type Roles struct {
	IsLogged echo.MiddlewareFunc
	IsAdmin  echo.MiddlewareFunc
	IsRole   func(roles ...userdata.Role) echo.MiddlewareFunc
}

func NewRoles(m middleware.Authenticator) *Roles {
	mid := echo.WrapMiddleware

	return &Roles{
		IsLogged: mid(m.Auth),
		IsAdmin:  mid(m.AdminOnly),
		IsRole: func(roles ...userdata.Role) echo.MiddlewareFunc {
			values := make([]string, len(roles))
			for i, role := range roles {
				values[i] = role.String()
			}
			return mid(m.RBAC(values...))
		},
	}
}
