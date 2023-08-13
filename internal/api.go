package internal

import (
	"irg1008/next-go/internal/controllers"
	"irg1008/next-go/pkg/config"
	"irg1008/next-go/pkg/db"
	"irg1008/next-go/pkg/log"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getLoggerMiddleware() echo.MiddlewareFunc {
	debug := config.Env.IsDev
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

func applyAPIMiddlewares(e *echo.Group) {
	e.Use(getLoggerMiddleware())
	// TODO: Move rate limit to fast key-value store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
}

func APIRoute(e *echo.Echo, db *db.DB) *echo.Group {
	api := e.Group("/api")
	applyAPIMiddlewares(api)

	// Controllers
	controllers.AuthRoutes(api, db)
	controllers.ProtectedRoutes(api)

	return api
}
