package controllers

import (
	"irg1008/next-go/ent"
	"irg1008/next-go/internal/services"
	"irg1008/next-go/pkg/cookies"
	"irg1008/next-go/pkg/crypt"
	"irg1008/next-go/pkg/helpers"
	"irg1008/next-go/pkg/server"
	"irg1008/next-go/pkg/tokens"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	signing *tokens.Signing
	service *services.AuthService
}

func AuthRoutes(e *echo.Group, s *server.Server) {
	u := &AuthController{s.Signing, &services.AuthService{DB: s.DB}}

	g := e.Group("/auth")
	g.POST("/signup", u.signUp)
	g.POST("/login", u.logIn)
	g.POST("/refresh", u.refresh)
}

type SignUpRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5,max=100,password"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

func (u *AuthController) signUp(c echo.Context) error {
	data, err := helpers.BindAndValidate[SignUpRequest](c)
	if err != nil {
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
	data, err := helpers.BindAndValidate[LogInRequest](c)
	if err != nil {
		return err
	}

	user, err := u.service.GetUserByEmail(data.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if pwdMatch := crypt.CompareHashAndPwd(user.Password, data.Password); !pwdMatch {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect password")
	}

	return u.renewOrSetTokens(c, user)
}

type RefreshPayload struct {
	Token string `json:"accessToken"`
}

func (u *AuthController) refresh(c echo.Context) error {
	refreshToken, err := cookies.GetRefreshTokenCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token not found")
	}

	claims, err := u.signing.ParseRefreshToken(refreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := u.service.GetUserById(claims.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return u.renewOrSetTokens(c, user)
}

func (u *AuthController) renewOrSetTokens(c echo.Context, user *ent.User) error {
	tokenPair, err := u.signing.CerateUserTokenPair(user)
	if err != nil {
		return echo.ErrInternalServerError
	}

	cookies.SetRefreshTokenCookie(c, tokenPair.RefreshToken)
	return c.JSON(http.StatusOK, &RefreshPayload{tokenPair.Token})
}
