package cache

import "go.uber.org/fx"

// var Module = fx.Module("db",
// 	fx.Provide(
// 		NewDatabase,
// 	),
// )

func Module() fx.Option {
	return fx.Module(
		"cache",
		fx.Provide(fx.Annotate(
			NewCacheClient, fx.As(new(ICacheClient)),
		)),
	)
}
