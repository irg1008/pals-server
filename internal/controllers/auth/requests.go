package auth

import (
	"irg1008/next-go/pkg/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewRequestData struct {
	Email string `query:"email" validate:"required,email"`
}

func (u *AuthController) CreateNewConfirmEmailRequest(c echo.Context) error {
	data, err := helpers.BindAndValidate[NewRequestData](c)
	if err != nil {
		return err
	}

	user, err := u.service.GetUserByEmail(data.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if user.IsConfirmed {
		return echo.NewHTTPError(http.StatusConflict, "User already confirmed")
	}

	if err := u.sendConfirmEmail(data.Email); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (u *AuthController) sendConfirmEmail(email string) error {
	req, err := u.service.CreateConfirmationRequest(email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while creating new confirmation request")
	}

	url := u.client.ConfirmEmailURL(req.Token.String())
	err = u.mailSender.SendConfirmEmail(email, "Confirm email", url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while sending email")
	}

	return nil
}

type ConfirmEmailRequets struct {
	Token string `json:"token" validate:"required,uuid"`
}

func (u *AuthController) ConfirmEmail(c echo.Context) error {
	data, err := helpers.BindAndValidate[ConfirmEmailRequets](c)
	if err != nil {
		return err
	}

	user, err := u.service.ConfirmEmail(data.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return u.renewOrSetTokens(c, user)
}

func (u *AuthController) CreateNewResetRequest(c echo.Context) error {
	data, err := helpers.BindAndValidate[NewRequestData](c)
	if err != nil {
		return err
	}

	user, err := u.service.GetUserByEmail(data.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if !user.IsConfirmed {
		return echo.NewHTTPError(http.StatusConflict, "Confirm your email before attempting a password reset")
	}

	if err := u.createAndSendResetEmail(data.Email); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (u *AuthController) createAndSendResetEmail(email string) error {
	req, err := u.service.CreatePasswordResetRequest(email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while creating new reset request")
	}

	url := u.client.ResetPasswordURL(req.Token.String())
	err = u.mailSender.SendResetPassword(email, "Reset password", url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while sending email")
	}

	return nil
}

type ResetPasswordRequest struct {
	NewPasswordData
	Token string `json:"token" validate:"required,uuid"`
}

func (u *AuthController) ResetPassword(c echo.Context) error {
	data, err := helpers.BindAndValidate[ResetPasswordRequest](c)
	if err != nil {
		return err
	}

	user, err := u.service.ResetPassword(data.Token, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return u.renewOrSetTokens(c, user)
}
