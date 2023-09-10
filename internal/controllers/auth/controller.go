package auth

import (
	"irg1008/pals/ent"
	service "irg1008/pals/internal/services/auth"
	"irg1008/pals/pkg/auth"
	"irg1008/pals/pkg/client"
	"irg1008/pals/pkg/crypt"
	"irg1008/pals/pkg/mailer"
	"irg1008/pals/pkg/request"
	"irg1008/pals/pkg/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	client     *client.Client
	service    *service.AuthService
	mailSender *mailer.Sender
}

func Routes(e *echo.Group, s *server.Server) {
	mailSender := s.Mailer.NewSender(s.Config.AppName, "no-reply")

	u := &AuthController{
		service:    &service.AuthService{DB: s.DB},
		client:     s.Client,
		mailSender: mailSender,
	}

	emailEndpoint := "email"
	s.Auth.AddEmailProvider(&auth.EmailProviderOpts{
		Name:  emailEndpoint,
		Check: u.ValidLogin,
	})

	g := e.Group("/" + emailEndpoint)

	g.POST("/signup", u.SignUp)
	g.POST("/confirm", u.ConfirmEmail)
	g.GET("/request/confirm", u.CreateNewConfirmEmailRequest)
	g.POST("/passwd-reset", u.ResetPassword)
	g.GET("/request/passwd-reset", u.CreateNewResetRequest)
}

type NewRequestData struct {
	Email string `query:"email" validate:"required,email"`
}

type ConfirmEmailRequets struct {
	Token string `json:"token" validate:"required,uuid"`
}

type NewPasswordData struct {
	Password string `json:"password" validate:"required,min=5,max=72,password"`
}

type SignUpRequest struct {
	NewPasswordData
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	NewPasswordData
	Token string `json:"token" validate:"required,uuid"`
}

func (u *AuthController) CreateNewConfirmEmailRequest(c echo.Context) error {
	data, err := request.BindAndValidate[NewRequestData](c)
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
		return echo.NewHTTPError(http.StatusInternalServerError, err, "Error while sending confirmation email")
	}

	return nil
}

func (u *AuthController) ConfirmEmail(c echo.Context) error {
	data, err := request.BindAndValidate[ConfirmEmailRequets](c)
	if err != nil {
		return err
	}

	_, err = u.service.ConfirmEmail(data.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}

func (u *AuthController) CreateNewResetRequest(c echo.Context) error {
	data, err := request.BindAndValidate[NewRequestData](c)
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

func (u *AuthController) ResetPassword(c echo.Context) error {
	data, err := request.BindAndValidate[ResetPasswordRequest](c)
	if err != nil {
		return err
	}

	_, err = u.service.ResetPassword(data.Token, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
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

func (u *AuthController) ValidLogin(email string, pass string) (bool, error) {
	user, err := u.userExistForEmail(email)
	if err != nil {
		return false, nil
	}

	if err = isConfirmed(user); err != nil {
		return false, err
	}

	return crypt.CompareHashAndPwd(user.Password, pass), nil
}

func (u *AuthController) userExistForEmail(email string) (*ent.User, error) {
	user, err := u.service.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "No user found with this email")
	}
	return user, nil
}

func isConfirmed(user *ent.User) error {
	if !user.IsConfirmed {
		return echo.NewHTTPError(http.StatusForbidden, "Confirm your email first")
	}
	return nil

}

func isNotConfirmed(user *ent.User) error {
	if user.IsConfirmed {
		return echo.NewHTTPError(http.StatusConflict, "User already confirmed")
	}
	return nil
}
