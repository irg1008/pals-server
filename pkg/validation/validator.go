package validation

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func compiles(value string, fl validator.FieldLevel) bool {
	re := regexp.MustCompile(value)
	return re.MatchString(fl.Field().String())
}

var containsLetter = validator.Func(func(fl validator.FieldLevel) bool {
	return compiles(`[a-zA-Z]`, fl)
})

var containsNumber = validator.Func(func(fl validator.FieldLevel) bool {
	return compiles(`[0-9]`, fl)
})

var containsSpecial = validator.Func(func(fl validator.FieldLevel) bool {
	return compiles(`[^a-zA-Z0-9]`, fl)
})

var isValidPassword = validator.Func(func(fl validator.FieldLevel) bool {
	return containsLetter(fl) && containsNumber(fl) && containsSpecial(fl)
})

func NewCustomValidator() *CustomValidator {
	validator := validator.New()
	err := validator.RegisterValidation("password", isValidPassword)

	if err != nil {
		panic("Failed to register custom validator")
	}

	return &CustomValidator{validator}
}
