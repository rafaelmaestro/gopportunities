package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/fnunezzz/go-logger"
	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"go.uber.org/fx"
)

type HttpServer struct {
	AppServer *echo.Echo
	AppGroup *echo.Group
}

// func NewHttpServer(lc fx.Lifecycle, config *config.Config) (*HttpServer, error) {
// 	s := echo.New()

// 	srv := http.Server{
// 		Addr: config.Http.AppPort,
// 	}

// 	lc.Append(fx.Hook{
//         OnStart: func(ctx context.Context) error {
//         	go func() {
// 				if err := s.Start(srv.Addr); err != nil && err != http.ErrServerClosed {
// 					log.Fatal("shutting down the server")
// 				}
// 		}()
//             return nil
//         },
//         OnStop: func(ctx context.Context) error {
//             return s.Shutdown(ctx)
//         },
//     })

// 	return &HttpServer{
// 		Server: s,
// 	}, nil
// }

func NewServer(
	lc fx.Lifecycle,
	config *config.Config,
) *HttpServer {
	server := echo.New()

	sLog := logger.Get()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			srv := http.Server{
				Addr: fmt.Sprintf(":%d", 3000), // TODO: change to config
			}
			go func() {
				if err := server.Start(srv.Addr); err != nil && err != http.ErrServerClosed {
					sLog.Errorf("error starting server on port %d: %s", 3000, err)
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
