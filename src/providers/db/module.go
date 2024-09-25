package db

import "go.uber.org/fx"

// var Module = fx.Module("db",
// 	fx.Provide(
// 		NewDatabase,
// 	),
// )

func Module() fx.Option {
	return fx.Module(
		"db",
		fx.Provide(NewDatabase),
	)
}
