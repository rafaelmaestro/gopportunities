package main

import (
	"os"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/providers/db"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {

	loggerEnv := logger.Environment("production")
	logger.Init(loggerEnv)

	appEnv := os.Getenv("APP_ENV")

    var loggerOption fx.Option
    if appEnv == "production" {
        loggerOption = fx.WithLogger(func() fxevent.Logger {
            return fxevent.NopLogger
        })
    } else {
        loggerOption = fx.Logger(logger.Get()) // Use custom logger
    }

	app := fx.New(
		loggerOption, // Disable fx logger
		fx.Logger(logger.Get()), // Use custom logger
		fx.Provide(config.Init),
		db.Module(),
		httpServer.Module(),
		preco.Module(),
	)

	app.Run()
}
