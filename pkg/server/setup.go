package server

import (
	"irg1008/pals/ent"
	"irg1008/pals/ent/userdata"
	"irg1008/pals/internal/services/user"
	"irg1008/pals/pkg/auth"
	"irg1008/pals/pkg/config"
	"irg1008/pals/pkg/db"
	"irg1008/pals/pkg/log"
	"irg1008/pals/pkg/mailer"

	"github.com/go-pkgz/auth/token"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware(origins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowCredentials: true,
	})
}

func (s *Server) SetMiddlewares(e *echo.Echo) {
	e.Use(log.LoggerMiddleware(s.Config.IsDev))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(CORSMiddleware([]string{s.Config.ClientUrl}))
}

func (s *Server) SetAuthProviders(config *config.Config) {
	s.Auth.AddGoogleProvider(&auth.GoogleProviderOpts{
		ClientID:     s.Config.GoogleClientID,
		ClientSecret: s.Config.GoogleClientSecret,
	})

	// TODO: This provider won't work until mobile app is ready. This is because enrolling is 99€/year: https://developer.apple.com/enroll/
	// s.Auth.AddCustomAppleProvider(&auth.AppleProviderOpts{
	// 	AppleConfig: &provider.AppleConfig{
	// 		ClientID: s.Config.AppleClientID,
	// 		TeamID:   s.Config.AppleTeamID,
	// 		KeyID:    s.Config.AppleKeyID,
	// 	},
	// 	PrivateKey: s.Config.ApplePrivateKey,
	// })
}

func NewMailer(config *config.Config) *mailer.Mailer {
	mailHostInfo := &mailer.HostInfo{
		Username: config.EmailUser,
		Address:  config.EmailHost,
		Port:     config.EmailPort,
		Password: config.EmailPass,
	}
	return mailer.NewMailer(mailHostInfo, config.Domain)
}

func NewAuthService(config *config.Config, db *db.DB) *auth.AuthService {
	userService := user.UserService{DB: db}
	return auth.NewAuthService(&auth.BaseOptions{
		AppName:      config.AppName,
		JWTSecret:    config.JWTSecret,
		URL:          config.APIUrl,
		Local:        config.IsDev,
		ClaimsUpdate: ClaimsUpdate(userService.GetOrCreteUserData),
	})
}

func ClaimsUpdate(userDataFetcher func(user *token.User) (*ent.UserData, error)) token.ClaimsUpdFunc {
	return func(claims token.Claims) token.Claims {
		if claims.User == nil {
			return claims
		}

		userData, err := userDataFetcher(claims.User)
		if err != nil {
			return claims
		}

		claims.User.Picture = userData.Picture
		claims.User.SetStrAttr("role", userData.Role.String())
		claims.User.SetAdmin(userData.Role == userdata.RoleAdmin)

		return claims
	}
}
