package server

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func handleIndex(c echo.Context) error {
	name := c.QueryParam("name")

	if err := validate.Var(name, "required,startswith=P,min=2"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Hello, "+name+"!")
}

func basicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

func StartServer() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handleIndex)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth((basicAuth)))

	err := e.Start(":8080")
	e.Logger.Fatal(err)
}
