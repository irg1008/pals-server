package auth

import (
	"irg1008/next-go/internal/services"
	"irg1008/next-go/pkg/client"
	"irg1008/next-go/pkg/mail"
	"irg1008/next-go/pkg/server"
	"irg1008/next-go/pkg/tokens"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	client     *client.Client
	service    *services.AuthService
	signing    *tokens.Signing
	mailSender *mail.Sender
}

func Routes(e *echo.Group, s *server.Server) {
	mailSender := s.Mail.NewSender("Pals", "no-reply")

	u := &AuthController{
		service:    &services.AuthService{DB: s.DB},
		client:     s.Client,
		signing:    s.Signing,
		mailSender: mailSender,
	}

	g := e.Group("/auth")
	g.POST("/signup", u.SignUp)
	g.POST("/login", u.LogIn)
	g.GET("/logout", u.LogOut)
	g.GET("/refresh", u.Refresh)
	g.GET("/request/confirm-user", u.CreateNewConfirmEmailRequest)
	g.GET("/request/reset-password", u.CreateNewResetRequest)
	g.POST("/confirm-user", u.ConfirmEmail)
	g.POST("/reset-password", u.ResetPassword)
}