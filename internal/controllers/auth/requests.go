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

	user, err := u.userExistForEmail(data.Email)
	if err != nil {
		return err
	}

	if err := isNotConfirmed(user); err != nil {
		return err
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
	if err := u.mailSender.SendConfirmEmail(email, "Confirm email", url); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while sending confirmation email")
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

	user, err := u.userExistForEmail(data.Email)
	if err != nil {
		return err
	}

	if err := isConfirmed(user); err != nil {
		return err
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
	if err := u.mailSender.SendResetPassword(email, "Reset password", url); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while sending reset email")
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
