package httpServer

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"go.uber.org/fx"
)

type HttpServer struct {
	*echo.Echo
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

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			srv := http.Server{
				Addr: fmt.Sprintf(":%d", 3000),
			}


			fmt.Printf("%s", srv.Addr)
			go func() {
				if err := server.Start(srv.Addr); err != nil && err != http.ErrServerClosed {
					log.Fatal("shutting down the server")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping client")
			server.Shutdown(ctx)
			return nil
		},
	})
	return &HttpServer{
		server,
	}

}
