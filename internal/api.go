package internal

import (
	"irg1008/next-go/internal/controllers"
	"irg1008/next-go/internal/server"
	"irg1008/next-go/pkg/log"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getLoggerMiddleware(s *server.Server) echo.MiddlewareFunc {
	debug := s.Config.IsDev
	log.SetDefaultLogger(debug)

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogLatency:  true,
		LogRemoteIP: true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info(v.Method,
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("duration", v.Latency.String()),
				slog.String("from", v.RemoteIP),
			)

			return nil
		},
	})
}

func applyAPIMiddlewares(e *echo.Group, s *server.Server) {
	e.Use(getLoggerMiddleware(s))
	// TODO: Move rate limit to fast key-value store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
}

func APIRoutes(e *echo.Echo, s *server.Server) *echo.Group {
	api := e.Group("/api")
	applyAPIMiddlewares(api, s)

	// Controllers
	controllers.AuthRoutes(api, s)
	controllers.ProtectedRoutes(api, s)

	return api
}
