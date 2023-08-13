package controllers

import (
	"irg1008/next-go/internal/services"
	"irg1008/next-go/pkg/crypt"
	"irg1008/next-go/pkg/db"
	"irg1008/next-go/pkg/tokens"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service *services.AuthService
}

func AuthRoutes(e *echo.Group, db *db.DB) {
	u := &AuthController{&services.AuthService{DB: db}}

	g := e.Group("/auth")
	g.POST("/signup", u.signUp)
	g.GET("/login", u.logIn)
}

type CreateUserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5,max=100,password"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

func (u *AuthController) signUp(c echo.Context) error {
	var data CreateUserRequest

	if err := c.Bind(&data); err != nil {
		return err
	}

	if err := c.Validate(data); err != nil {
		return err
	}

	if _, err := u.service.CreateUser(data.Email, data.Password); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "User already exists")
	}

	return c.NoContent(http.StatusCreated)
}

type LogInRequest struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (u *AuthController) logIn(c echo.Context) error {
	var data LogInRequest

	if err := c.Bind(&data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	user, err := u.service.GetUser(data.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if pwdMatch := crypt.CompareHashAndPwd(user.Password, data.Password); !pwdMatch {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect password")
	}

	token, err := tokens.SignUserToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.String(http.StatusOK, token)
}
