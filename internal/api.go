package internal

import (
	"irg1008/pals/internal/controllers/auth"
	"irg1008/pals/internal/controllers/protected"
	"irg1008/pals/pkg/log"
	"irg1008/pals/pkg/server"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setAPIMiddlewares(e *echo.Group, s *server.Server) {
	e.Use(log.GetLoggerMiddleware(s))
	// TODO: Move rate limit to fast key-value store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
}

func APIRoutes(e *echo.Echo, s *server.Server) *echo.Group {
	api := e.Group("/api")
	setAPIMiddlewares(api, s)

	// Controllers
	auth.Routes(api, s)
	protected.Routes(api, s)

	return api
}
