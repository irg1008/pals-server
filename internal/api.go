package internal

import (
	"irg1008/next-go/internal/config"
	"irg1008/next-go/internal/routes/example"
	"irg1008/next-go/pkg/log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func applyAPIMiddlewares(e *echo.Group) {
	// TODO: Move rate limit to fast key-value store

	logger := log.GetLogger(config.GetConfig().IsDev)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogLatency:  true,
		LogRemoteIP: true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("uri", v.URI).
				Int("status", v.Status).
				Str("latency", v.Latency.String()).
				Str("from", v.RemoteIP).
				Timestamp().
				Msg(v.Method)

			return nil
		},
	}))

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
}

func APIRoute(e *echo.Echo) *echo.Group {
	api := e.Group("/api")
	applyAPIMiddlewares(api)

	// Controllers
	example.Routes(api)

	return api
}
