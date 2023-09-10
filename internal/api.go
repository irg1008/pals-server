package internal

import (
	"irg1008/pals/internal/controllers/user"
	"irg1008/pals/pkg/server"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setAPIMiddlewares(e *echo.Group) {
	// TODO: Move rate limit to fast key-value store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
}

func APIRoutes(e *echo.Echo, s *server.Server) *echo.Group {
	api := e.Group("/api")
	setAPIMiddlewares(api)

	user.Routes(api, s)

	return api
}
