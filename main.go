package main

import (
	"os"
	"strings"

	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco"
	"github.com/rafaelmaestro/gopportunities/src/providers/cache"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/providers/db"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
	"github.com/rafaelmaestro/gopportunities/src/utils"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	// Initialize logger
	loggerEnv := logger.Environment(os.Getenv("APP_ENV"))
	logger.Init(loggerEnv)

	appEnv := os.Getenv("APP_ENV")

	var loggerOption fx.Option
	if strings.TrimSpace(appEnv) == "production" {
		loggerOption = fx.Options(
			fx.Provide(func() zerolog.Logger {
				return *utils.ZerologLogger()
			}),
			fx.WithLogger(func() fxevent.Logger {
				return fxevent.NopLogger
			}),
		)
	} else {
		loggerOption = fx.Options(
			fx.Provide(func() zerolog.Logger {
				return *utils.ZerologLogger()
			}),
			fx.WithLogger(fxzerolog.Init()),
		)
	}

	// Initialize dd-tracer (Datadog)
	// Commented because we dont have the datadog agent running on the new o2b cluster
	// tracer.Start()

	// defer tracer.Stop()

	app := fx.New(
		loggerOption,
		fx.Provide(config.Init),
		db.Module(),
		httpServer.Module(),
		preco.Module(),
		cache.Module(),
	)

	app.Run()
}
