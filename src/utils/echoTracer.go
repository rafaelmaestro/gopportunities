package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

// Middleware to automatically capture trace for every HTTP request
func EchoTracer(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Start of the request
			start := time.Now()
			req := c.Request()
			res := c.Response()

			logger.Trace().
				Str("method", req.Method).
				Str("path", req.URL.Path).
				Str("host", req.Host).
				Str("remote_addr", c.RealIP()).
				Msg("Starting request")

			// Process the request
			err := next(c)

			// Check the final status after the request execution
			status := res.Status
			if err != nil && status == http.StatusOK {
				// If an error occurred and the status is 200, modify the status to the correct error status
				if he, ok := err.(*echo.HTTPError); ok {
					status = he.Code
				} else {
					status = http.StatusInternalServerError
				}
			}

			// Use the appropriate log level based on the final status
			switch {
			case status >= 500:
				logger.Error(). // Red by default for errors (500+)
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Error processing the request")

			case status == 404:
				logger.Warn(). // Yellow/Orange by default for 404
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Route not found")

			default:
				logger.Info(). // Green/Blue by default for status 200 and other successful responses
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Request processed successfully")
			}

			return err
		}
	}
}
