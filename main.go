package main

import (
	"github.com/rafaelmaestro/gopportunities/src/config"
	"github.com/rafaelmaestro/gopportunities/src/db"
	"github.com/rafaelmaestro/gopportunities/src/http"
	"github.com/rafaelmaestro/gopportunities/src/preco"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		http.Module,
		preco.Module,
	)
	app.Run()
}
