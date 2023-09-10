package user

import (
	"irg1008/pals/internal/services/user"
	"irg1008/pals/pkg/server"
	"net/http"

	"github.com/go-pkgz/auth/token"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service *user.UserService
}

func Routes(e *echo.Group, s *server.Server) {
	u := &UserController{
		service: &user.UserService{DB: s.DB},
	}

	g := e.Group("/me", s.Auth.IsLogged)

	g.GET("", u.GetUserData)
}

func (u *UserController) GetUserData(c echo.Context) error {
	user, err := token.GetUserInfo(c.Request())
	if err != nil {
		return err
	}

	data, err := u.service.GetOrCreteUserData(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user data")
	}

	return c.JSON(http.StatusOK, data)
}
