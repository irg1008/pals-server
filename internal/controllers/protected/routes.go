package protected

import (
	"irg1008/pals/pkg/roles"
	"irg1008/pals/pkg/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, s *server.Server) {
	g := e.Group("/protected", s.Roles.IsLogged)
	g.GET("", HandleExmaple)
}

func HandleExmaple(c echo.Context) error {
	user := roles.GetUser(c)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, " + user.Email,
	})
}
