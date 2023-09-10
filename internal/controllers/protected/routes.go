package protected

import (
	"irg1008/pals/pkg/server"
	"net/http"

	"github.com/go-pkgz/auth/token"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, s *server.Server) {
	// m := s.Auth.Service.Middleware()
	// mid := echo.WrapMiddleware(m.Auth)

	g := e.Group("/protected", s.Auth.IsLogged)
	g.GET("", HandleExmaple)
}

func HandleExmaple(c echo.Context) error {
	user, err := token.GetUserInfo(c.Request())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + user.Name + "!",
	})
}
