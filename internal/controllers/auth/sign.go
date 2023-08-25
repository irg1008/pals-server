package auth

import (
	"irg1008/pals/ent"
	"irg1008/pals/pkg/cookies"
	"irg1008/pals/pkg/crypt"
	"irg1008/pals/pkg/request"
	"irg1008/pals/pkg/roles"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewPasswordData struct {
	Password        string `json:"password" validate:"required,min=5,max=100,password"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type SignUpRequest struct {
	NewPasswordData
	Email string `json:"email" validate:"required,email"`
}

func (u *AuthController) SignUp(c echo.Context) error {
	data, err := request.BindAndValidate[SignUpRequest](c)
	if err != nil {
		return err
	}

	if _, err := u.service.CreateUser(data.Email, data.Password); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "User already exists")
	}

	if err := u.sendConfirmEmail(data.Email); err != nil {
		return err
	}

	return c.NoContent(http.StatusAccepted)
}

type LogInRequest struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (u *AuthController) LogIn(c echo.Context) error {
	data, err := request.BindAndValidate[LogInRequest](c)
	if err != nil {
		return err
	}

	user, err := u.userExistForEmail(data.Email)
	if err != nil {
		return err
	}

	if pwdMatch := crypt.CompareHashAndPwd(user.Password, data.Password); !pwdMatch {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}

	if err := isNotConfirmed(user); err != nil {
		return err
	}

	return u.renewOrSetTokens(c, user)
}

type RefreshPayload struct {
	Token string `json:"accessToken"`
}

func (u *AuthController) Refresh(c echo.Context) error {
	refreshToken, err := cookies.GetRefreshTokenCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token not found")
	}

	claims, err := u.signing.ParseRefreshToken(refreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := u.userExists(claims.Id)
	if err != nil {
		return err
	}

	return u.renewOrSetTokens(c, user)
}

func (u *AuthController) renewOrSetTokens(c echo.Context, user *ent.User) error {
	tokenPair, err := u.signing.CerateUserTokenPair(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while creating tokens", err)
	}

	cookies.SetRefreshTokenCookie(c, tokenPair.RefreshToken)
	return c.JSON(http.StatusOK, &RefreshPayload{tokenPair.Token})
}

func (u *AuthController) LogOut(c echo.Context) error {
	cookies.DeleteRefreshTokenCookie(c)
	return c.NoContent(http.StatusOK)
}

func (u *AuthController) Me(c echo.Context) error {
	payload := roles.GetUser(c)

	user, err := u.service.GetUserByEmail(payload.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find user")
	}

	return c.JSON(http.StatusOK, user)
}