package example

import (
	"irg1008/next-go/internal/routes/example/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {
	g := e.Group("/example")

	g.Group("/admin", handlers.BasicAuth())
	g.GET("", handleExmaple)
}

func handleExmaple(c echo.Context) error {
	params := new(struct {
		Name string `query:"name" validate:"required,startswith=G,min=2"`
	})

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(params); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Hello, "+params.Name+"!")
}
