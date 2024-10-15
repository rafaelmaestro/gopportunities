package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/utils"
	"go.uber.org/fx"
)

type HttpServer struct {
	AppServer *echo.Echo
	AppGroup  *echo.Group
}

func NewServer(
	lc fx.Lifecycle,
	config *config.Config,
) *HttpServer {

	logger := utils.ZerologLogger()

	// Configurando o Echo com middleware de trace
	server := echo.New()
	server.Use(utils.EchoTracer(*logger))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			srv := http.Server{
				Addr: fmt.Sprintf(":%d", 3000), // TODO: change to config
			}
			go func() {
				if err := server.Start(srv.Addr); err != nil && err != http.ErrServerClosed {
					logger.Error().Err(err).Msgf("error starting server on port %d: %s", 3000, err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Shutdown(ctx)
			return nil
		},
	})

	appGroup := server.Group(fmt.Sprintf("/%s", os.Getenv("APP_NAME")))

	return &HttpServer{
		server,
		appGroup,
	}

}
