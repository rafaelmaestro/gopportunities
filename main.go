package main

import (
	"os"
	"strings"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/providers/db"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	// Initialize logger
	loggerEnv := logger.Environment("production")
	logger.Init(loggerEnv)

	appEnv := os.Getenv("APP_ENV")

    var loggerOption fx.Option
	if strings.TrimSpace(appEnv) == "production" {
        loggerOption = fx.WithLogger(func() fxevent.Logger {
            return fxevent.NopLogger
        })
    } else {
        loggerOption = fx.Logger(logger.Get()) // Use custom logger
    }

	// Initialize dd-tracer (Datadog)
	tracer.Start()

	defer tracer.Stop()

	app := fx.New(
		loggerOption,
		fx.Provide(config.Init),
		db.Module(),
		httpServer.Module(),
		preco.Module(),
	)

	app.Run()
}
