package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	*auth.Service
	*Roles
	Opts    *auth.Opts
	Handler echo.HandlerFunc
}

type BaseOptions struct {
	AppName   string
	JWTSecret string
	URL       string
}

const (
	tokenDuration  = 5 * time.Minute
	cookieDuration = 24 * time.Hour
)

func getAuthOptions(opts *BaseOptions) auth.Opts {
	return auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) {
			return opts.JWTSecret, nil
		}),
		TokenDuration:  tokenDuration,
		CookieDuration: cookieDuration,
		Issuer:         opts.AppName,
		URL:            opts.URL,
		AvatarStore:    avatar.NewNoOp(),
		SecureCookies:  true,
		SameSiteCookie: http.SameSiteStrictMode,
		ClaimsUpd:      token.ClaimsUpdFunc(AddUserRoles),
	}
}

func AddUserRoles(claims token.Claims) token.Claims {
	fmt.Print(claims)
	claims.User.SetSliceAttr("roles", []string{"admin"})
	claims.User.Picture = ""
	return claims
}

func NewAuthService(baseOpts *BaseOptions) *AuthService {
	opts := getAuthOptions(baseOpts)

	service := auth.NewService(opts)
	authHandler, _ := service.Handlers()

	return &AuthService{
		Service: service,
		Opts:    &opts,
		Roles:   NewRoles(service.Middleware()),
		Handler: echo.WrapHandler(authHandler),
	}
}
