package main

import (
	"github.com/rafaelmaestro/gopportunities/src/modules/preco"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/providers/db"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.Init),
		db.Module(),
		httpServer.Module(),
		preco.Module(),
	)

	app.Run()
}
