package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"

	"github.com/labstack/echo/v4"
)

func main() {
	opts := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) {
			return "secret", nil
		}),
		TokenDuration:  time.Minute * 15,
		CookieDuration: time.Hour * 24,
		Issuer:         "pals",
		URL:            "http://localhost:8001",
		AvatarStore:    avatar.NewLocalFS("/tmp"),
	}

	service := auth.NewService(opts)

	service.AddDirectProvider("email", provider.CredCheckerFunc(func(user, pass string) (bool, error) {
		valid := user == "pepe" && pass == "pepe"
		return valid, nil
	}))

	service.AddProvider("google", "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com", "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")

	m := service.Middleware()
	authRoutes, _ := service.Handlers()

	router := echo.New()

	router.Any("/auth/*", echo.WrapHandler(authRoutes))
	router.GET("/protected", protectedHandler, echo.WrapMiddleware(m.Auth))

	log.Fatal(router.Start(":8001"))
}

func protectedHandler(c echo.Context) error {
	// New empty struct with name
	user := struct {
		Name string `json:"name"`
	}{
		Name: "John Doe",
	}

	return c.JSON(http.StatusOK, user)
}
