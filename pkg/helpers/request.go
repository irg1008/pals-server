package helpers

import "github.com/labstack/echo/v4"

func BindAndValidate[T any](c echo.Context) (data T, err error) {
	if err = c.Bind(&data); err != nil {
		return
	}
	if err = c.Validate(data); err != nil {
		return
	}
	return
}

func BindHeaders[T any](c echo.Context) (data T, err error) {
	if err = (&echo.DefaultBinder{}).BindHeaders(c, &data); err != nil {
		return
	}
	return
}
