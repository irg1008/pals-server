package log

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(debug bool) echo.MiddlewareFunc {
	SetDefaultLogger(debug)

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
				slog.Int64("content-length", c.Response().Size),
			)

			return nil
		},
	})
}
