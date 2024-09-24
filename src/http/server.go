package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/config"
	"go.uber.org/fx"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(lc fx.Lifecycle, config *config.Config) (*HttpServer, error) {
	echoInstance := echo.New()

	echoInstance.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server := &http.Server{
		Addr:    config.Http.AppPort,
		Handler: echoInstance,
	}


	lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
        	go func() {
				slog.Info("Starting server on port", "port", config.Http.AppPort)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					slog.Error("Failed to start server", "error", err)
				}
            }()
            return nil
        },
        OnStop: func(ctx context.Context) error {
            return server.Shutdown(ctx)
        },
    })

	return &HttpServer{
		server: server,
	}, nil
}
